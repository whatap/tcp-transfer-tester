# Makefile

# 변수 설정
SRC_DIR := src
BIN_DIR := bin
CLIENT_NAME := client
SERVER_NAME := server

# 윈도우즈용 클라이언트 및 서버 빌드
windows:
	@echo "Building Windows binaries..."
	@GOOS=windows GOARCH=amd64 go build -o $(BIN_DIR)/windows/$(CLIENT_NAME).exe $(SRC_DIR)/client.go
	@GOOS=windows GOARCH=amd64 go build -o $(BIN_DIR)/windows/$(SERVER_NAME).exe $(SRC_DIR)/server.go
	@echo "Windows binaries are created in $(BIN_DIR)/windows directory."

# 리눅스용 클라이언트 및 서버 빌드
linux:
	@echo "Building Linux binaries..."
	@GOOS=linux GOARCH=amd64 go build -o $(BIN_DIR)/linux/$(CLIENT_NAME) $(SRC_DIR)/client.go
	@GOOS=linux GOARCH=amd64 go build -o $(BIN_DIR)/linux/$(SERVER_NAME) $(SRC_DIR)/server.go
	@echo "Linux binaries are created in $(BIN_DIR)/linux directory."

# 모두 빌드
build: windows linux

# 클리어
clean:
	@rm -rf $(BIN_DIR)
	@echo "Binaries removed."
