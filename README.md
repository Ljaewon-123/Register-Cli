# Register-CLI

자주 사용하는 CLI 명령어를 카테고리별로 저장하고 빠르게 실행할 수 있는 Go 기반 CLI 도구입니다.

## 🎯 주요 기능

- **명령어 등록**: 자주 사용하는 명령어를 설명과 함께 카테고리별로 저장
- **빠른 실행**: 저장된 명령어를 대화형 UI로 선택하여 실행
- **다중 실행**: 여러 명령어를 한 번에 선택하여 순차적으로 실행
- **카테고리 관리**: 명령어를 카테고리로 구분하여 체계적으로 관리

## 📦 설치 방법

### 자동 설치 (권장)

```bash
curl -sSL https://raw.githubusercontent.com/Ljaewon-123/register-cli/main/install.sh | bash
```

설치 후 새 터미널 세션을 시작하거나 다음 명령어로 PATH를 reload하세요:

```bash
# bash 사용자
source ~/.bashrc

# zsh 사용자
source ~/.zshrc
```

### 수동 설치

1. [Releases 페이지](https://github.com/Ljaewon-123/register-cli/releases)에서 운영체제에 맞는 바이너리 다운로드
2. 실행 권한 부여: `chmod +x register-cli`
3. PATH에 추가된 디렉토리로 이동: `mv register-cli ~/.local/bin/`

## 🚀 사용법

### 프로그램 실행

```bash
register-cli
```

실행하면 다음 3가지 옵션이 표시됩니다:

```
Choose an option:
1. Show category
2. Register command
3. Exit
```

### 1. 명령어 등록하기

프로그램 실행 후 `2`를 선택하여 명령어를 등록합니다:

```bash
Select number: 2

Enter the category to register: Docker
Enter command description: List all running containers
Enter command: docker ps
```

### 2. 명령어 실행하기

프로그램 실행 후 `1`을 선택하여 저장된 명령어를 실행합니다:

```bash
Select number: 1

# 카테고리 선택
Please select a category:
> Docker
  Git
  Kubernetes

# 명령어 선택 (스페이스바로 다중 선택 가능)
Select multiple commands to run (select with spacebar):
> [ ] docker ps
    :: description: List all running containers
  [ ] docker images
    :: description: List all images
  [x] docker system prune
    :: description: Clean up unused resources
```

선택한 명령어가 순차적으로 실행됩니다.

## 📝 사용 예시

### 개발 환경 명령어 등록

```
Category: Development
- Description: Start dev server
  Command: npm run dev

- Description: Run tests
  Command: npm test

- Description: Build project
  Command: npm run build
```

### Docker 관리 명령어 등록

```
Category: Docker
- Description: Remove all stopped containers
  Command: docker container prune -f

- Description: Remove unused images
  Command: docker image prune -a -f

- Description: Show container stats
  Command: docker stats --no-stream
```

### Git 워크플로우 명령어 등록

```
Category: Git
- Description: Pull latest changes
  Command: git pull origin main

- Description: Show git status
  Command: git status

- Description: Show recent commits
  Command: git log --oneline -10
```

## 📂 데이터 저장 위치

명령어는 실행 디렉토리의 `command.json` 파일에 저장됩니다:

```json
{
  "categories": [
    {
      "name": "Docker",
      "commands": [
        {
          "description": "List all running containers",
          "command": "docker ps"
        }
      ]
    }
  ]
}
```

## 🛠️ 개발자 정보

### 빌드 방법

```bash
# 의존성 설치
go mod tidy

# 빌드
go build -o register-cli

# 실행
./register-cli
```

### 테스트 실행

```bash
# 전체 테스트
go test ./...

# 특정 패키지 테스트
go test ./register
go test ./category
```

### 프로젝트 구조

```
register-cli/
├── main.go              # 메인 엔트리 포인트
├── register/            # 명령어 등록 모듈
│   ├── register.go
│   └── register_test.go
├── category/            # 카테고리 선택 및 실행 모듈
│   ├── category.go
│   └── category_test.go
├── .github/
│   └── workflows/
│       └── release.yaml # GitHub Actions 릴리스 워크플로우
└── install.sh          # 자동 설치 스크립트
```

## 🔧 의존성

- [survey/v2](https://github.com/AlecAivazis/survey) - 대화형 CLI UI

## 📋 요구사항

- Go 1.24.2 이상 (개발용)
- Linux, macOS, Windows 지원

## 🤝 기여하기

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📄 라이선스

이 프로젝트는 오픈소스입니다.

## 💡 활용 팁

- **반복 작업 자동화**: 매번 입력하기 번거로운 긴 명령어를 저장하세요
- **팀 공유**: `command.json` 파일을 팀원과 공유하여 동일한 명령어 세트 사용
- **환경별 관리**: 개발/스테이징/프로덕션 환경별로 다른 `command.json` 사용
- **복잡한 명령어**: 파이프라인이나 여러 옵션이 있는 복잡한 명령어를 쉽게 재사용

## ⚠️ 주의사항

- 명령어는 `sh -c` (macOS/Linux)로 실행됩니다
- 민감한 정보(비밀번호, API 키 등)를 명령어에 포함하지 마세요
- `command.json` 파일을 Git에 커밋하지 않으려면 `.gitignore`에 추가하세요 (이미 기본 설정됨)
