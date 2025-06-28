# Gemini API 연동 테스트 가이드

## ✅ 구현 완료된 기능

### 1. Gemini API 연동
- ✅ `genai` 패키지 설치 완료
- ✅ API 키 기반 인증 구현
- ✅ 페르소나 생성 (Gemini 1.5 Flash)
- ✅ 이벤트 메시지 생성 (Gemini 1.5 Flash)
- ✅ Mock 모드 지원 (API 키 없을 때)

### 2. API 엔드포인트
```bash
# 플랜트카드 생성
POST /plantcards
{
  "farm_plot_id": 1
}

# 플랜트카드 조회
GET /plantcards/{id}

# SNS 공유용 플랜트카드
GET /plantcards/{id}/share

# 농장별 플랜트카드 목록
GET /farms/plots/{id}/plantcards
```

## 🔧 API 키 설정 방법

1. Google AI Studio에서 API 키 발급
   - https://aistudio.google.com/apikey

2. 환경 변수 설정
```bash
export GOOGLE_AI_API_KEY="your_actual_api_key_here"
```

3. 또는 .env 파일에 추가
```env
GOOGLE_AI_API_KEY=your_actual_api_key_here
```

## 🧪 테스트 방법

### 1. Mock 모드 테스트 (API 키 없음)
```bash
# 서버 실행
go run cmd/server/main.go

# 플랜트카드 생성 테스트
curl -X POST http://localhost:8080/plantcards \
  -H "Content-Type: application/json" \
  -d '{"farm_plot_id": 1}'
```

### 2. 실제 Gemini API 테스트 (API 키 필요)
```bash
# API 키 설정
export GOOGLE_AI_API_KEY="your_key"

# 서버 실행
go run cmd/server/main.go

# 실제 AI 생성 테스트
curl -X POST http://localhost:8080/plantcards \
  -H "Content-Type: application/json" \
  -d '{"farm_plot_id": 1}'
```

## 📋 응답 예시

### Mock 응답
```json
{
  "id": 1,
  "farm_plot_id": 1,
  "persona": "안녕하세요! 저는 싱싱하고 아삭한 상추로 자란 친환경 농장 작물입니다.",
  "image_url": "https://example.com/generated-image.png",
  "video_url": "https://example.com/veo3-generated-video.mp4",
  "event_message": "싱싱하고 아삭한 상추로 자란 친환경 농장 작물입니다.가 plant_card_creation 이벤트를 알려드립니다!",
  "created_at": "2025-01-28T15:30:45Z"
}
```

### 실제 Gemini 응답 (예상)
```json
{
  "id": 1,
  "farm_plot_id": 1,
  "persona": "안녕! 나는 햇살을 받고 자란 싱싱한 상추야. 아삭한 식감으로 여러분의 식탁을 더 건강하게 만들어줄게! 🥬",
  "image_url": "https://example.com/generated-image.png",
  "video_url": "https://example.com/veo3-generated-video.mp4",
  "event_message": "🌱 드디어 나의 플랜트카드가 완성됐어! 내 성장 과정을 영상으로 만나보세요 ✨",
  "created_at": "2025-01-28T15:30:45Z"
}
```

## 🔄 추후 확장 예정

### 1. 이미지 생성 (Imagen API)
- Google Cloud Vertex AI Imagen
- 농작물 캐릭터 이미지 생성

### 2. 비디오 생성 (Veo API)
- Google Veo 3 모델
- 쇼츠/릴스용 짧은 영상 생성

### 3. SNS 최적화
- 플랫폼별 포맷 최적화
- 해시태그 자동 생성 