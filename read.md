📌 개발 요구사항 정리

1. 기술 스택
- golang
- gin
- gorm
- MySQL (로컬 환경 사용)
- Swagger-UI (API 문서화)
- Redis, Docker 사용하지 않음
- 로그인 기능 없음

2. 데이터베이스 테이블 구조

🗂 t_user (사용자 테이블)
- user_idx INT (PK, AUTO_INCREMENT) → 사용자 고유 IDX
- user_id VARCHAR(50) → 사용자 ID (중복 불가)
- username VARCHAR(50) → 사용자 이름(닉네임)
- email VARCHAR(100) → 이메일
- password VARCHAR(255) → 비밀번호 (암호화 저장)
- created_at DATETIME → 생성일시
- last_login DATETIME → 마지막 로그인

🗂 t_board (게시판 글 테이블)
- board_id INT (PK, AUTO_INCREMENT) → 게시글 고유 ID
- user_idx INT (FK → t_user) → 작성자 IDX
- title VARCHAR(200) → 글 제목
- content TEXT → 글 내용
- created_at DATETIME → 작성일시
- updated_at DATETIME → 수정일시

🗂 t_comment (댓글 테이블)
- comment_id INT (PK, AUTO_INCREMENT) → 댓글 고유 ID
- board_id INT (FK → t_board) → 게시글 ID
- user_idx INT (FK → t_user) → 작성자 IDX
- content TEXT → 댓글 내용
- made_at DATETIME → 작성일시
- updated_at DATETIME → 수정일시
- step INT → 댓글 단계 (0=원댓글, 1=대댓글, 2=대대댓글…)

3. API 설계 원칙
- CRUD 제공 (Create, Read, Update, Delete)
- Controller → Service → ServiceImpl → Repository 구조
- 배열 응답 시 반드시 Pagination 적용
- 응답 구조 통일

응답 예시:
{
  "success": true,
  "message": "요청이 성공적으로 처리되었습니다.",
  "data": [
    {
      "id": 1,
      "username": "phs1",
      "email": "phs@phs.phs",
      "createdAt": "2026-05-11T17:20:41.06185",
      "lastLogin": null
    }
  ],
  "timestamp": "2026-05-11T17:20:51.8852115"
}

4. 개발 방식
- 시니어 개발자 스타일로 구조화 (레이어드 아키텍처, 책임 분리)
- 주니어 개발자가 이해할 수 있도록 한글 주석 추가
- 최신 유행 방식으로 개발
- 라우터를 빼주고 라우터안에서의 구룹별로 다시 파일을 생성해줘
- 폴더구성을 자바형식으로 했는데 go에서 사용하는 형식으로 변경해서 작업해줘
- **로그인 기능**: 사용자 인증 후 Access Token과 Refresh Token을 발급합니다.
- **Access Token**: 유효기간 **1시간**
- **Refresh Token**: 유효기간 **3시간**
- **JWT Secret Key**: `toy_login!@#`
- **인증 미들웨어**: 보호된 라우터 그룹에 접근할 때 Authorization 헤더의 Bearer 토큰을 검증하는 미들웨어를 구현하세요. 미들웨어를 통과해야만 후속 핸들러가 실행됩니다.
