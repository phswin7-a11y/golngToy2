# 🚀 Golang Toy Project 2

이 프로젝트는 **Golang**, **Gin**, **GORM**을 사용하여 구축된 고성능 레이어드 아키텍처 기반의 웹 API 서버입니다. 시니어 개발자의 설계 스타일을 반영하여 높은 응집도와 낮은 결합도를 지향하며, 주니어 개발자도 쉽게 이해할 수 있도록 상세한 한글 주석을 포함하고 있습니다.

## 🛠 기술 스택

- **Language:** Go 1.26+
- **Framework:** [Gin Web Framework](https://github.com/gin-gonic/gin)
- **ORM:** [GORM](https://gorm.io/)
- **Database:** MySQL (Local)
- **API Documentation:** [Swagger (swag)](https://github.com/swaggo/swag)

## 🏗 프로젝트 구조 (Go Style - Package by Feature)

프로젝트는 도메인 중심의 기능별 패키지 구조를 채택하여, 각 도메인(User, Board, Comment)이 독립적인 책임을 가집니다.

```text
golangToy2/
├── cmd/
│   └── api/
│       └── main.go           # 애플리케이션 엔트리 포인트 (의존성 주입 및 서버 실행)
├── internal/
│   ├── domain/               # 데이터베이스 엔티티 및 공통 도메인 정의
│   ├── infra/                # 인프라 설정 (DB 연결 등)
│   ├── common/               # 공통 응답 DTO 및 유틸리티
│   ├── user/                 # 사용자 도메인 (Handler, Service, Repository)
│   ├── board/                # 게시판 도메인 (Handler, Service, Repository)
│   ├── comment/              # 댓글 도메인 (Handler, Service, Repository)
│   └── router/               # 라우팅 그룹 및 설정 분리
├── docs/                     # Swagger 자동 생성 문서
├── go.mod                    # 모듈 정의
└── read.md                   # 초기 개발 요구사항 문서
```

## 📊 데이터베이스 스키마

- **t_user:** 사용자 정보 관리 (ID, 이름, 이메일, 암호화된 비밀번호)
- **t_board:** 게시글 관리 (제목, 내용, 작성자 연관)
- **t_comment:** 댓글 및 대댓글 관리 (작성 단계 `step` 지원)

## 🛣 API 주요 엔드포인트

### User API
- `POST /api/v1/users` - 사용자 등록
- `GET /api/v1/users` - 사용자 목록 조회 (Pagination)
- `GET /api/v1/users/:idx` - 단일 사용자 상세 조회
- `PUT /api/v1/users/:idx` - 사용자 정보 수정
- `DELETE /api/v1/users/:idx` - 사용자 삭제

### Board API
- `POST /api/v1/boards` - 게시글 작성
- `GET /api/v1/boards` - 게시글 목록 조회 (Pagination)
- `GET /api/v1/boards/:id` - 게시글 상세 조회
- `PUT /api/v1/boards/:id` - 게시글 수정
- `DELETE /api/v1/boards/:id` - 게시글 삭제

### Comment API
- `POST /api/v1/comments` - 댓글 작성
- `GET /api/v1/boards/:boardId/comments` - 게시글별 댓글 목록 조회
- `PUT /api/v1/comments/:id` - 댓글 수정
- `DELETE /api/v1/comments/:id" - 댓글 삭제

## 🚀 시작하기

### 1. 환경 설정
로컬 MySQL 데이터베이스를 준비하고 필요한 경우 환경 변수를 설정합니다. 기본값은 `internal/infra/database.go`에서 확인할 수 있습니다.

```bash
# 기본 DB 설정 예시
export DB_USER=root
export DB_PASS=your_password
export DB_NAME=golangtoy2
```

### 2. 의존성 설치
```bash
go mod tidy
```

### 3. Swagger 문서 생성
```bash
go install github.com/swaggo/swag/cmd/swag@latest
swag init -g cmd/api/main.go
```

### 4. 서버 실행
```bash
go run cmd/api/main.go
```

### 5. API 문서 확인
서버가 실행된 후 브라우저에서 아래 주소로 접속하세요:
`http://localhost:8080/swagger/index.html`

## 📝 개발 원칙
- **Uniform Response:** 모든 API는 성공 여부, 메시지, 데이터, 타임스탬프를 포함한 통일된 응답 구조를 반환합니다.
- **Pagination:** 대량의 데이터를 반환하는 배열 응답에는 항상 페이징 처리가 적용됩니다.
- **Layered Responsibility:** Handler(HTTP), Service(Logic), Repository(DB) 간의 책임을 엄격히 분리합니다.
