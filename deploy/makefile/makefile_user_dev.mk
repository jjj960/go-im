# 全局变量
VERSION = latest
SERVER_NAME = user
SERVER_TYPE = rpc
DOCKER_USER = 475921395

# 应用名称和Dockerfile路径
APP_NAME_TEST = go-im-$(SERVER_NAME)-$(SERVER_TYPE)-test
DOCKER_FILE = deploy/dockerfile/dockerfile_$(SERVER_NAME)_$(SERVER_TYPE)_dev

# 检查并删除目标文件和目标镜像
clean:
	@echo "Checking if ${SERVER_NAME}-${SERVER_TYPE} exists..."
	@if [ -f deploy/dockerfile/bin/${SERVER_NAME}-${SERVER_TYPE} ]; then \
		echo "Deleting ${SERVER_NAME}-${SERVER_TYPE}..."; \
		rm -f deploy/dockerfile/bin/${SERVER_NAME}-${SERVER_TYPE}; \
	fi

	@echo "Checking if Docker image $(DOCKER_USER)/${APP_NAME_TEST}:$(VERSION) exists..."
	@if docker images -q $(DOCKER_USER)/${APP_NAME_TEST}:$(VERSION) | grep -q .; then \
		echo "Deleting Docker image $(DOCKER_USER)/${APP_NAME_TEST}:$(VERSION)..."; \
		docker rmi $(DOCKER_USER)/${APP_NAME_TEST}:$(VERSION); \
	else \
		echo "Docker image $(DOCKER_USER)/${APP_NAME_TEST}:$(VERSION) does not exist, skipping..."; \
	fi

# 测试环境下的文件编译
build:
	@echo "Building Go application..."
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o deploy/dockerfile/bin/${SERVER_NAME}-${SERVER_TYPE} apps/${SERVER_NAME}/${SERVER_TYPE}/${SERVER_NAME}.go
	@echo "Go application built successfully."

# 构建并推送 Docker 镜像
build-push:
	@echo "Building Docker image..."
	# 使用 DOCKER_FILE 和 APP_NAME_TEST 变量构建 Docker 镜像
	docker build . -f $(DOCKER_FILE) -t $(DOCKER_USER)/$(APP_NAME_TEST):$(VERSION)
	@echo "Docker image built successfully."

	@echo "Tagging Docker image..."
	# 标记镜像为 Docker 仓库地址
	docker tag $(DOCKER_USER)/$(APP_NAME_TEST):$(VERSION) $(DOCKER_USER)/$(APP_NAME_TEST):$(VERSION)

	@echo "Pushing Docker image to the registry..."
	# 推送镜像到 Docker 仓库
	docker push $(DOCKER_USER)/$(APP_NAME_TEST):$(VERSION)
	#docker rmi $(DOCKER_USER)/$(APP_NAME_TEST):$(VERSION)
	@echo "Docker image pushed successfully."

release-test: clean build build-push