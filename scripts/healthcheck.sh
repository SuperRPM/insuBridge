#!/bin/bash

# 환경 변수 로드
if [ -f .env ]; then
    source .env
else
    echo "Error: .env file not found"
    exit 1
fi

# 웹훅 URL 확인
if [ -z "$SLACK_WEBHOOK_URL" ]; then
    echo "Error: SLACK_WEBHOOK_URL is not set"
    exit 1
fi

# 헬스체크 URL
HEALTH_CHECK_URL="http://3.139.6.169:8080/api/health"

# 슬랙 메시지 전송 함수
send_slack_message() {
    local message="$1"
    curl -s -X POST -H 'Content-type: application/json' \
        --data "{\"text\":\"$message\"}" \
        $SLACK_WEBHOOK_URL
}

# 서버 상태 확인
response=$(curl -s -o /dev/null -w "%{http_code}" $HEALTH_CHECK_URL)

if [ $response -ne 200 ]; then
    # 서버가 응답하지 않음
    message="🚨 [InsuBridge 서버 다운] 서버가 응답하지 않습니다. (HTTP Status: $response)"
    send_slack_message "$message"
    
    # 서버 재시작
    cd /home/ubuntu/insuBridge
    ./deploy.sh
    send_slack_message "🔄 서버 재시작을 시도했습니다."
fi 