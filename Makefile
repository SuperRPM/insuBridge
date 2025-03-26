.PHONY: build run test clean

# 기본 타겟
all: build

# 빌드
build:
	go build -o bin/server cmd/server/main.go

# 실행
run: build
	./bin/server

# 테스트
test:
	go test ./...

# 클린
clean:
	rm -rf bin/
	go clean

# 개발 모드 실행 (air를 사용한 핫 리로드)
dev:
	air 