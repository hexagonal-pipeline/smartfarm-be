## 브랜치 이름 컨벤션
다음과 같이 작성해주세요.
- ex) feat/#1 (이슈 번호)

## 라이브러리
- 문서화: `swag`
- 서버: `fiber`
- 데이터베이스: `sqlc`
- 테스트: `testify`
- 로깅: `zerolog`
- 의존성 주입: `samber/do` (v2)
- 데이터베이스 마이그레이션: `dbmate`

## 환경 변수 예제 (`.env`)
```env
PORT=8080

POSTGRES_HOST=localhost
POSTGRES_PORT=5432
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_DB=smartfarm

# Google AI API
GOOGLE_AI_API_KEY=your_gemini_api_key_here
```

## Architecture
```mermaid
graph TD
    subgraph "외부 (External)"
        A[User/Client]
        D[PostgreSQL Database]
    end

    subgraph "애플리케이션 (Application)"
        subgraph "어댑터 (Adapter Layer)"
            B["인바운드 어댑터<br/>(Web API)"]
            C["아웃바운드 어댑터<br/>(Postgres Repository)"]
        end

        subgraph "애플리케이션 코어 (Core)"
            subgraph "유스케이스 계층 (Usecase Layer)"
                E["유스케이스 (Use Cases)"]
            end
            subgraph "도메인 계층 (Domain Layer)"
                F["도메인 모델<br/>(Domain Models)"]
                G["포트 (Ports / Interfaces)"]
            end
        end
    end

    A -- HTTP Request --> B
    B -- Calls --> E
    E -- Uses --> F
    E -- Uses Port --> G
    C -- Implements Port --> G
    E -- Calls Port through DI --> C
    C -- SQL Query --> D
    D -- Result --> C
    C -- Result --> E
    E -- Result --> B
    B -- HTTP Response --> A
```
