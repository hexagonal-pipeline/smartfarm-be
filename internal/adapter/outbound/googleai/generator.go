package googleai

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"smartfarm-be/internal/ports/outbound"
	"smartfarm-be/pkg/config"
	"time"

	"github.com/google/generative-ai-go/genai"
	"github.com/rs/zerolog/log"
	"github.com/samber/do/v2"
	"google.golang.org/api/option"
	veo "google.golang.org/genai"
)

type GoogleAIGenerator struct {
	client    *genai.Client
	veoClient *veo.Client
	config    *config.GeminiConfig
	veoConfig *config.VeoConfig
}

// Veo API 구조체들
type VeoVideoRequest struct {
	Instances  []VeoInstance `json:"instances"`
	Parameters VeoParameters `json:"parameters"`
}

type VeoInstance struct {
	Prompt string   `json:"prompt"`
	Image  *VeoFile `json:"image,omitempty"`
}

type VeoFile struct {
	BytesBase64Encoded string `json:"bytesBase64Encoded,omitempty"`
	GcsUri             string `json:"gcsUri,omitempty"`
	MimeType           string `json:"mimeType"`
}

type VeoParameters struct {
	AspectRatio      string `json:"aspectRatio,omitempty"`
	DurationSeconds  int    `json:"durationSeconds,omitempty"`
	EnhancePrompt    bool   `json:"enhancePrompt,omitempty"`
	NegativePrompt   string `json:"negativePrompt,omitempty"`
	PersonGeneration string `json:"personGeneration,omitempty"`
	SampleCount      int    `json:"sampleCount,omitempty"`
	Seed             uint32 `json:"seed,omitempty"`
	StorageUri       string `json:"storageUri,omitempty"`
}

type VeoVideoResponse struct {
	Name string `json:"name"`
}

type VeoOperationStatusRequest struct {
	OperationName string `json:"operationName"`
}

type VeoOperationResponse struct {
	Name     string       `json:"name"`
	Done     bool         `json:"done"`
	Response *VeoResponse `json:"response,omitempty"`
	Error    *VeoError    `json:"error,omitempty"`
}

type VeoResponse struct {
	Type   string     `json:"@type"`
	Videos []VeoVideo `json:"videos"`
}

type VeoVideo struct {
	GcsUri   string `json:"gcsUri"`
	MimeType string `json:"mimeType"`
}

type VeoError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Gemini REST API 구조체들
type GeminiImageRequest struct {
	Contents         []GeminiContent `json:"contents"`
	GenerationConfig GeminiGenConfig `json:"generationConfig"`
}

type GeminiContent struct {
	Parts []GeminiPart `json:"parts"`
}

type GeminiPart struct {
	Text string `json:"text"`
}

type GeminiGenConfig struct {
	ResponseModalities []string `json:"responseModalities"`
}

type GeminiImageResponse struct {
	Candidates []GeminiCandidate `json:"candidates"`
}

type GeminiCandidate struct {
	Content GeminiResponseContent `json:"content"`
}

type GeminiResponseContent struct {
	Parts []GeminiResponsePart `json:"parts"`
}

type GeminiResponsePart struct {
	Text       string            `json:"text,omitempty"`
	InlineData *GeminiInlineData `json:"inlineData,omitempty"`
}

type GeminiInlineData struct {
	MimeType string `json:"mimeType"`
	Data     string `json:"data"`
}

func NewGoogleAIGenerator(i do.Injector) (outbound.AIGenerator, error) {
	geminiCfg := do.MustInvoke[*config.GeminiConfig](i)

	// VeoConfig는 선택적으로 주입 (없으면 nil로 처리)
	var veoConfig *config.VeoConfig
	defer func() {
		if r := recover(); r != nil {
			// VeoConfig가 등록되지 않은 경우 panic이 발생하므로 nil로 처리
			veoConfig = nil
		}
	}()

	// VeoConfig 주입 시도
	func() {
		veoConfig = do.MustInvoke[*config.VeoConfig](i)
	}()

	if geminiCfg.APIKey == "" {
		log.Warn().Msg("Google AI API key not provided, using mock responses")
		return nil, fmt.Errorf("google AI API key not provided")
	}

	client, err := genai.NewClient(context.Background(), option.WithAPIKey(geminiCfg.APIKey))
	if err != nil {
		return nil, fmt.Errorf("failed to create Gemini client: %w", err)
	}

	veoClient, err := veo.NewClient(context.Background(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create Veo client: %w", err)
	}

	return &GoogleAIGenerator{
		client:    client,
		config:    geminiCfg,
		veoConfig: veoConfig,
		veoClient: veoClient,
	}, nil
}

// GeneratePersona는 작물에 대한 페르소나를 생성합니다.
func (g *GoogleAIGenerator) GeneratePersona(ctx context.Context, prompt string) (string, error) {
	if g.client == nil {
		// API 키가 없을 때 목업 응답
		return fmt.Sprintf("안녕하세요! 저는 %s입니다.", prompt), nil
	}

	model := g.client.GenerativeModel("gemini-2.0-flash")
	model.SetTemperature(0.8)
	model.SetMaxOutputTokens(200)

	fullPrompt := fmt.Sprintf(`
다음 농작물 설명을 바탕으로 친근하고 매력적인 페르소나를 영어로 생성해주세요:
"%s"

요구사항:
- 1인칭 관점으로 작성
- 친근하고 따뜻한 말투
- 농작물의 특성을 살린 성격
- 100자 이내로 간결하게
- SNS에서 사용할 수 있는 톤앤매너

예시: "안녕! 나는 햇살을 받고 자란 싱싱한 상추야. 아삭한 식감으로 여러분의 식탁을 더 건강하게 만들어줄게!"
주의: 영어로 작성하시오.
`, prompt)

	resp, err := model.GenerateContent(ctx, genai.Text(fullPrompt))
	if err != nil {
		log.Error().Err(err).Str("prompt", prompt).Msg("failed to generate persona with Gemini")
		return "", fmt.Errorf("failed to generate persona: %w", err)
	}

	if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
		return "", fmt.Errorf("empty response from Gemini")
	}

	persona := fmt.Sprintf("%v", resp.Candidates[0].Content.Parts[0])
	log.Info().Str("prompt", prompt).Str("persona", persona).Msg("successfully generated persona")
	return persona, nil
}

// GenerateEventMessage는 특정 이벤트에 대한 메시지를 생성합니다.
func (g *GoogleAIGenerator) GenerateEventMessage(ctx context.Context, persona, event string) (string, error) {
	if g.client == nil {
		// API 키가 없을 때 목업 응답
		return fmt.Sprintf("%s가 %s 이벤트를 알려드립니다!", persona, event), nil
	}

	model := g.client.GenerativeModel("gemini-2.0-flash")
	model.SetTemperature(0.9)
	model.SetMaxOutputTokens(150)

	fullPrompt := fmt.Sprintf(`
다음 페르소나로 이벤트 메시지를 작성해주세요:

페르소나: "%s"
이벤트: "%s"

요구사항:
- 위 페르소나의 말투와 성격을 유지
- SNS 공유에 적합한 톤
- 이모지 포함 가능
- 80자 이내로 간결하게
- 친근하고 재미있게
- 이벤트의 성격과 목적에 맞는 메시지 작성
- 반드시 성격을 중요시 할 것

이벤트 별 예시 (밝은 성격):
- 발아: "🌱 휴~! 발아했어요! 내 성장 과정을 영상으로 만나보세요 ✨"
- 성장: "💪 쑥쑥 자라고 있어요! 오늘의 성장 소식을 전해드려요 🌿"
- 수확: "🌾 수확의 계절이 왔어요! 오늘이 제가 가장 맛있는 날이에요 😋"
- 건강: "☔️ 비가 많이 오네요! 걱정 마세요, 저는 잘 지내고 있답니다 🌱💪"

주어진 페르소나의 성격을 파악하여 적절한 메시지를 생성해주세요.
`, persona, event)

	resp, err := model.GenerateContent(ctx, genai.Text(fullPrompt))
	if err != nil {
		log.Error().Err(err).Str("persona", persona).Str("event", event).Msg("failed to generate event message")
		return "", fmt.Errorf("failed to generate event message: %w", err)
	}

	if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
		return "", fmt.Errorf("empty response from Gemini")
	}

	message := fmt.Sprintf("%v", resp.Candidates[0].Content.Parts[0])
	log.Info().Str("persona", persona).Str("event", event).Str("message", message).Msg("successfully generated event message")
	return message, nil
}

// GenerateImage는 주어진 프롬프트로 이미지를 생성합니다.
func (g *GoogleAIGenerator) GenerateImage(ctx context.Context, prompt string) (string, error) {
	if g.config.APIKey == "" {
		// API 키가 없을 때 목업 응답
		log.Info().Str("prompt", prompt).Msg("image generation requested (mock)")
		return "https://example.com/generated-image.png", nil
	}

	// 캐릭터 이미지 생성에 특화된 프롬프트
	fullPrompt := fmt.Sprintf(`
Create a cute and friendly cartoon character image for sns: "%s"
`, prompt)

	// REST API 요청 준비
	requestBody := GeminiImageRequest{
		Contents: []GeminiContent{
			{
				Parts: []GeminiPart{
					{Text: fullPrompt},
				},
			},
		},
		GenerationConfig: GeminiGenConfig{
			ResponseModalities: []string{"TEXT", "IMAGE"},
		},
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal request")
		return "", fmt.Errorf("failed to marshal request: %w", err)
	}

	// HTTP 요청 생성
	url := "https://generativelanguage.googleapis.com/v1beta/models/gemini-2.0-flash-preview-image-generation:generateContent"
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Error().Err(err).Msg("failed to create HTTP request")
		return "", fmt.Errorf("failed to create HTTP request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-goog-api-key", g.config.APIKey)

	// HTTP 요청 실행
	client := &http.Client{Timeout: 60 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		log.Error().Err(err).Msg("failed to call Gemini API")
		return "", fmt.Errorf("failed to call Gemini API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		log.Error().Int("status", resp.StatusCode).Str("body", string(body)).Msg("Gemini API error")
		return "", fmt.Errorf("Gemini API error: %d - %s", resp.StatusCode, string(body))
	}

	// 응답 파싱
	var geminiResp GeminiImageResponse
	if err := json.NewDecoder(resp.Body).Decode(&geminiResp); err != nil {
		log.Error().Err(err).Msg("failed to decode response")
		return "", fmt.Errorf("failed to decode response: %w", err)
	}

	if len(geminiResp.Candidates) == 0 {
		return "", fmt.Errorf("no candidates in Gemini response")
	}

	// 이미지 데이터 추출 및 저장
	for _, part := range geminiResp.Candidates[0].Content.Parts {
		if part.InlineData != nil && part.InlineData.Data != "" {
			// Base64 디코딩
			imageData, err := base64.StdEncoding.DecodeString(part.InlineData.Data)
			if err != nil {
				log.Error().Err(err).Msg("failed to decode base64 image")
				continue
			}

			// 이미지 파일 저장
			fileName := fmt.Sprintf("character_%d.png", time.Now().Unix())
			imageURL, err := g.saveImageToFile(imageData, fileName)
			if err != nil {
				log.Error().Err(err).Str("file", fileName).Msg("failed to save generated image")
				return "", fmt.Errorf("failed to save generated image: %w", err)
			}

			log.Info().Str("prompt", prompt).Str("image_url", imageURL).Msg("successfully generated image with Gemini 2.0")
			return imageURL, nil
		}
	}

	return "", fmt.Errorf("no image data found in Gemini response")
}

// saveImageToFile는 이미지 데이터를 파일로 저장하고 URL을 반환합니다.
func (g *GoogleAIGenerator) saveImageToFile(imageData []byte, fileName string) (string, error) {
	// uploads/images 디렉토리 생성
	dir := "uploads/images"
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", fmt.Errorf("failed to create directory: %w", err)
	}

	// 파일 저장
	filePath := filepath.Join(dir, fileName)
	if err := os.WriteFile(filePath, imageData, 0644); err != nil {
		return "", fmt.Errorf("failed to write image file: %w", err)
	}

	// 정적 파일 URL 생성 (Fiber에서 /uploads 경로로 서빙)
	imageURL := fmt.Sprintf("/uploads/images/%s", fileName)
	return imageURL, nil
}

func (g *GoogleAIGenerator) GenerateVideo(ctx context.Context, prompt, imageURL string) (string, error) {

	imageData, err := os.ReadFile(path.Join(".", imageURL))
	if err != nil {
		return "", fmt.Errorf("failed to read image file: %w", err)
	}

	if g.veoConfig == nil {
		log.Info().Str("prompt", prompt).Str("image_url", imageURL).Msg("Veo video generation requested (mock - no VEO config)")
		return "", nil
	}

	if g.veoClient == nil {
		return "", fmt.Errorf("veo client is not initialized")
	}

	videoConfig := &veo.GenerateVideosConfig{
		AspectRatio:      "9:16",
		PersonGeneration: "dont_allow",
		NumberOfVideos:   1,
	}
	operation, err := g.veoClient.Models.GenerateVideos(ctx, "veo-2.0-generate-001", prompt, &veo.Image{
		ImageBytes: imageData,
		MIMEType:   "image/png",
	}, videoConfig)
	if err != nil {
		return "", fmt.Errorf("failed to generate video: %w", err)
	}
	if operation == nil {
		return "", fmt.Errorf("received nil operation from GenerateVideos")
	}
	for !operation.Done {
		time.Sleep(10 * time.Second)
		operation, err = g.veoClient.Operations.GetVideosOperation(ctx, operation, nil)
		if err != nil {
			return "", fmt.Errorf("failed to get operation status: %w", err)
		}
		if operation == nil {
			return "", fmt.Errorf("received nil operation from GetVideosOperation")
		}
	}
	if operation.Response == nil {
		return "", fmt.Errorf("operation completed but response is nil")
	}
	for n, video := range operation.Response.GeneratedVideos {
		_, err := g.veoClient.Files.Download(ctx, video.Video, nil)
		if err != nil {
			log.Error().Err(err).Msg("failed to download video file")
			continue
		}
		fname := fmt.Sprintf("uploads/video_%d_%d.mp4", time.Now().Unix(), n)
		err = os.WriteFile(fname, video.Video.VideoBytes, 0644)
		if err != nil {
			return "", fmt.Errorf("failed to write video file: %w", err)
		}
		return fname, nil
	}
	return "", fmt.Errorf("no video data found in Veo response")
}

// GenerateVideoFromPrompt는 텍스트 프롬프트와 이미지를 사용하여 Veo로 쇼츠 비디오를 생성합니다.
func (g *GoogleAIGenerator) GenerateVideoFromPrompt(ctx context.Context, prompt, imageURL string) (string, error) {
	// VeoConfig가 없으면 mock 응답
	if g.veoConfig == nil {
		log.Info().Str("prompt", prompt).Str("image_url", imageURL).Msg("Veo video generation requested (mock - no VEO config)")
		return "", nil
	}

	if g.veoClient == nil {
		return "", fmt.Errorf("veo client is not initialized")
	}

	videoConfig := &veo.GenerateVideosConfig{
		AspectRatio:      "1:1",
		PersonGeneration: "dont_allow",
		NumberOfVideos:   1,
	}
	operation, err := g.veoClient.Models.GenerateVideos(ctx, "veo-2.0-generate-001", prompt, nil, videoConfig)
	if err != nil {
		return "", fmt.Errorf("failed to generate video: %w", err)
	}
	if operation == nil {
		return "", fmt.Errorf("received nil operation from GenerateVideos")
	}

	for !operation.Done {
		time.Sleep(10 * time.Second)
		operation, err = g.veoClient.Operations.GetVideosOperation(ctx, operation, nil)
		if err != nil {
			return "", fmt.Errorf("failed to get operation status: %w", err)
		}
		if operation == nil {
			return "", fmt.Errorf("received nil operation from GetVideosOperation")
		}
	}

	if operation.Response == nil {
		return "", fmt.Errorf("operation completed but response is nil")
	}

	for n, video := range operation.Response.GeneratedVideos {
		_, err := g.veoClient.Files.Download(ctx, video.Video, nil)
		if err != nil {
			log.Error().Err(err).Msg("failed to download video file")
			continue
		}
		fname := fmt.Sprintf("uploads/video_%d_%d.mp4", time.Now().Unix(), n)
		err = os.WriteFile(fname, video.Video.VideoBytes, 0644)
		if err != nil {
			return "", fmt.Errorf("failed to write video file: %w", err)
		}
		return fname, nil
	}
	return "", fmt.Errorf("no video data found in Veo response")
}
