# Go로 구현한 TCP 서버 및 클라이언트

## 개요

이 저장소에는 간단한 TCP 서버와 클라이언트 통신을 구현한 두 개의 Go 프로그램이 포함되어 있습니다. 서버는 클라이언트로부터 수신한 메시지를 다시 에코하고, 클라이언트는 서버로 메시지를 보내고 서버의 응답을 표시합니다.

## TCP 서버

### 설명

TCP 서버 프로그램은 지정된 포트에서 수신 대기합니다. 클라이언트가 연결되면 클라이언트로부터 메시지를 읽고 다시 에코하며, 수신한 데이터와 전송한 데이터를 로그로 기록합니다. 서버는 각 클라이언트를 별도의 고루틴으로 처리하여 동시 연결을 지원합니다.

### 사용법

TCP 서버를 실행하려면 `server.go` 파일을 실행하면 됩니다. 기본적으로 서버는 포트 `6600`에서 수신 대기하지만, `SESSION_PORT` 환경 변수를 설정하여 포트를 사용자 정의할 수 있습니다.

```bash
export SESSION_PORT=6600
$ ./server 
Server listening on : 6600
Accepted connection from 127.0.0.1:41190
index: 1 Received 18 bytes header from client
Read 18 bytes from client: hello world whatap
Sent 18 bytes to client: hello world whatap
```


## TCP 클라이언트

### 설명

TCP 클라이언트 프로그램은 지정된 서버 IP 주소와 포트에 연결합니다. 사용자에게 서버로 보낼 메시지를 입력하도록 요청합니다. 메시지를 보낸 후 서버의 응답을 기다린 다음 표시합니다. 이 프로세스는 상호 작용할 때마다 시간 정보와 함께 기록됩니다.

### 사용법

TCP 클라이언트를 실행하려면 client.go 파일을 실행하면 됩니다. 기본적으로 클라이언트는 127.0.0.1:6600에 연결하지만, SESSION_SERVER_IP 및 SESSION_SERVER_PORT 환경 변수를 설정하여 서버 IP 주소 및 포트를 사용자 정의할 수 있습니다.

```bash
export SESSION_SERVER_IP="127.0.0.1"
export SESSION_SERVER_PORT="6600"
$ ./client
Enter message(default: hello world whatap):
2024-05-29 11:32:06.039883098 +0900 KST m=+2.425403694 Connected to server
index:1  Sent 18 bytes to server in 22.433µs
Received 18 bytes header from server in 132.983µs
Read 18 bytes from server in 140.527µs: hello world whatap
connection time:  270.856µs
index:2  Sent 18 bytes to server in 69.973µs
Received 18 bytes header from server in 113.095µs
Read 18 bytes from server in 127.112µs: hello world whatap
connection time:  1.000676571s
```

