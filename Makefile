NAME=${APP_NAME}

## build: Compile the packages
build:
	@go build -o $(NAME)

## run: Build and Run in development mode.
run: build
	@GIN_MODE=debug ./$(NAME)

## deps: Download modules
deps:
	@go mod download

## test: Run tests with verbose mode
test:
	@go test -v -covermode=count -coverprofile=coverage.out ./...

deploy:
	@gcloud builds submit
	@gcloud run deploy $(PROJECT_ID) --image $(DOCKER_IMAGE) \
		--region us-central1 --platform managed \
		--allow-unauthenticated --port $(PORT) \
		--set-env-vars="FILE_NAME=$(FILE_NAME)" \
		--set-env-vars='FILE_DOWNLOAD_URL=$(FILE_DOWNLOAD_URL)' \
		--set-env-vars='INFLUXDB_TOKEN=$(INFLUXDB_TOKEN)' \
		--set-env-vars='INFLUXDB_BUCKET=$(INFLUXDB_BUCKET)' \
		--set-env-vars='INFLUXDB_ORG=$(INFLUXDB_ORG)' \
		--set-env-vars='INFLUXDB_URL=$(INFLUXDB_URL)'
