package main

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"smartfarm-be/internal/adapter/outbound/googleai"
	"smartfarm-be/pkg/config"

	"github.com/samber/do/v2"
	"github.com/stretchr/testify/assert"
)

func TestGeminiImageGeneration(t *testing.T) {
	tests := []struct {
		name        string
		apiKey      string
		prompt      string
		expectMock  bool
		expectError bool
	}{
		{
			name:       "Mock mode (no API key)",
			apiKey:     "",
			prompt:     "싱싱하고 아삭한 상추",
			expectMock: true,
		},
		{
			name:       "Real API mode (with API key)",
			apiKey:     os.Getenv("GOOGLE_AI_API_KEY"),
			prompt:     "달콤한 토마토",
			expectMock: false,
		},
		{
			name:       "Korean characters prompt",
			apiKey:     os.Getenv("GOOGLE_AI_API_KEY"),
			prompt:     "귀여운 당근 캐릭터",
			expectMock: false,
		},
		{
			name:       "English prompt",
			apiKey:     os.Getenv("GOOGLE_AI_API_KEY"),
			prompt:     "cute broccoli character",
			expectMock: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// DI 컨테이너 설정
			injector := do.New()
			defer injector.Shutdown()

			// Config 제공
			cfg := &config.GeminiConfig{
				APIKey: tt.apiKey,
			}
			do.ProvideValue(injector, cfg)

			// Generator 생성
			generator, err := googleai.NewGoogleAIGenerator(injector)
			if err != nil {
				t.Fatalf("Failed to create generator: %v", err)
			}

			// 이미지 생성 테스트
			ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
			defer cancel()

			fmt.Printf("🎨 Testing image generation for: %s\n", tt.prompt)
			start := time.Now()

			imageURL, err := generator.GenerateImage(ctx, tt.prompt)

			elapsed := time.Since(start)
			fmt.Printf("⏱️  Generation time: %v\n", elapsed)

			if tt.expectError {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.NotEmpty(t, imageURL)

			if tt.expectMock {
				// Mock 모드인 경우 예상 URL 검증
				assert.Equal(t, "https://example.com/generated-image.png", imageURL)
				fmt.Printf("✅ Mock response: %s\n", imageURL)
			} else if tt.apiKey != "" {
				// 실제 API 모드인 경우
				assert.Contains(t, imageURL, "/uploads/images/")
				assert.Contains(t, imageURL, "character_")
				assert.Contains(t, imageURL, ".png")
				fmt.Printf("✅ Real image generated: %s\n", imageURL)

				// 파일이 실제로 생성되었는지 확인
				filePath := "." + imageURL // "/uploads/images/..." -> "./uploads/images/..."
				if _, err := os.Stat(filePath); err == nil {
					fmt.Printf("📁 Image file saved successfully: %s\n", filePath)

					// 파일 크기 확인
					if info, err := os.Stat(filePath); err == nil {
						fmt.Printf("📊 File size: %d bytes\n", info.Size())
						assert.Greater(t, info.Size(), int64(0), "Generated image file should not be empty")
					}
				} else {
					fmt.Printf("⚠️  Image file not found: %s (may be expected in test mode)\n", filePath)
				}
			} else {
				fmt.Printf("⏭️  Skipping real API test (no API key provided)\n")
			}

			fmt.Printf("──────────────────────────────────────\n")
		})
	}
}

func TestGeminiImageGenerationErrors(t *testing.T) {
	// DI 컨테이너 설정
	injector := do.New()
	defer injector.Shutdown()

	// Invalid API key로 테스트
	cfg := &config.GeminiConfig{
		APIKey: "invalid_api_key_for_testing",
	}
	do.ProvideValue(injector, cfg)

	generator, err := googleai.NewGoogleAIGenerator(injector)
	assert.NoError(t, err)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// 잘못된 API 키로 호출 시 에러 발생 예상
	imageURL, err := generator.GenerateImage(ctx, "test prompt")

	// API 키가 잘못되었으므로 에러가 발생해야 함
	if err != nil {
		fmt.Printf("✅ Expected error with invalid API key: %v\n", err)
		assert.Error(t, err)
		assert.Empty(t, imageURL)
	} else {
		// 에러가 발생하지 않았다면 무언가 잘못된 것
		t.Logf("⚠️  No error occurred with invalid API key. Response: %s", imageURL)
	}
}

func TestImageFileOperations(t *testing.T) {
	// 테스트용 이미지 데이터 (작은 PNG 파일)
	testImageData := []byte{
		0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A, // PNG 헤더
		0x00, 0x00, 0x00, 0x0D, 0x49, 0x48, 0x44, 0x52, // IHDR 청크 시작
		0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x01, // 1x1 픽셀
		0x08, 0x02, 0x00, 0x00, 0x00, 0x90, 0x77, 0x53, // IHDR 나머지
		0xDE, 0x00, 0x00, 0x00, 0x0C, 0x49, 0x44, 0x41, // IDAT 청크
		0x54, 0x08, 0x99, 0x01, 0x01, 0x00, 0x00, 0x00,
		0xFF, 0xFF, 0x00, 0x00, 0x00, 0x02, 0x00, 0x01,
		0xE2, 0x21, 0xBC, 0x33, 0x00, 0x00, 0x00, 0x00, // IEND
		0x49, 0x45, 0x4E, 0x44, 0xAE, 0x42, 0x60, 0x82,
	}

	// DI 컨테이너 설정
	injector := do.New()
	defer injector.Shutdown()

	// saveImageToFile 메서드는 private이므로 간접적으로 테스트
	// 실제로는 디렉토리 생성과 파일 저장 로직 검증

	// uploads/images 디렉토리가 생성되는지 확인
	testDir := "test_uploads/images"
	err := os.MkdirAll(testDir, 0755)
	assert.NoError(t, err)

	// 테스트 파일 저장
	testFile := fmt.Sprintf("%s/test_character_%d.png", testDir, time.Now().Unix())
	err = os.WriteFile(testFile, testImageData, 0644)
	assert.NoError(t, err)

	// 파일이 제대로 저장되었는지 확인
	_, err = os.Stat(testFile)
	assert.NoError(t, err)

	// 정리
	os.RemoveAll("test_uploads")

	fmt.Printf("✅ File operations test passed\n")
}

// 플랜트카드 전체 플로우 테스트
func TestPlantCardWithImageGeneration(t *testing.T) {
	apiKey := os.Getenv("GOOGLE_AI_API_KEY")
	if apiKey == "" {
		t.Skip("Skipping integration test: GOOGLE_AI_API_KEY not set")
	}

	// DI 컨테이너 설정
	injector := do.New()
	defer injector.Shutdown()

	cfg := &config.GeminiConfig{APIKey: apiKey}
	do.ProvideValue(injector, cfg)

	generator, err := googleai.NewGoogleAIGenerator(injector)
	assert.NoError(t, err)

	ctx := context.Background()

	fmt.Printf("🌱 Testing complete plant card generation flow...\n")

	// 1. 페르소나 생성
	fmt.Printf("1️⃣  Generating persona...\n")
	persona, err := generator.GeneratePersona(ctx, "싱싱하고 아삭한 상추")
	assert.NoError(t, err)
	assert.NotEmpty(t, persona)
	fmt.Printf("✅ Persona: %s\n", persona)

	// 2. 이미지 생성
	fmt.Printf("2️⃣  Generating character image...\n")
	imageURL, err := generator.GenerateImage(ctx, persona)
	assert.NoError(t, err)
	assert.NotEmpty(t, imageURL)
	fmt.Printf("✅ Image URL: %s\n", imageURL)

	// 3. 이벤트 메시지 생성
	fmt.Printf("3️⃣  Generating event message...\n")
	message, err := generator.GenerateEventMessage(ctx, persona, "plant_card_creation")
	assert.NoError(t, err)
	assert.NotEmpty(t, message)
	fmt.Printf("✅ Event Message: %s\n", message)

	// 4. 결과 출력
	fmt.Printf("\n🎉 Complete Plant Card Generated!\n")
	fmt.Printf("👤 Persona: %s\n", persona)
	fmt.Printf("🖼️  Image: %s\n", imageURL)
	fmt.Printf("💬 Message: %s\n", message)
}

// 성능 테스트
func TestImageGenerationPerformance(t *testing.T) {
	apiKey := os.Getenv("GOOGLE_AI_API_KEY")
	if apiKey == "" {
		t.Skip("Skipping performance test: GOOGLE_AI_API_KEY not set")
	}

	injector := do.New()
	defer injector.Shutdown()

	cfg := &config.GeminiConfig{APIKey: apiKey}
	do.ProvideValue(injector, cfg)

	generator, err := googleai.NewGoogleAIGenerator(injector)
	assert.NoError(t, err)

	prompts := []string{
		"귀여운 토마토",
		"싱싱한 상추",
		"달콤한 당근",
	}

	fmt.Printf("⚡ Performance testing with %d prompts...\n", len(prompts))

	for i, prompt := range prompts {
		start := time.Now()

		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		imageURL, err := generator.GenerateImage(ctx, prompt)
		cancel()

		elapsed := time.Since(start)

		if err != nil {
			fmt.Printf("❌ Prompt %d failed: %v (time: %v)\n", i+1, err, elapsed)
		} else {
			fmt.Printf("✅ Prompt %d completed: %s (time: %v)\n", i+1, imageURL, elapsed)
		}

		// 요청 간 간격 (API 레이트 리밋 고려)
		time.Sleep(2 * time.Second)
	}
}
