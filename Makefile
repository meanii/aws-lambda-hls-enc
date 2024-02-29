ZIP_FILE_NAME="hls-enc-lambda"
BIN="bootstrap"

GOOS="linux"
GOARCH="arm64"
CGO_ENABLED=0

LAMBDA_FUNCTION_NAME="hls-enc"

compile:
	@echo "Compiling..."
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -tags lambda.norpc -o $(BIN) main.go

bundle:
	@echo "Bundling..."
	zip $(ZIP_FILE_NAME).zip $(BIN)

push:
	@echo "Pushing..."
	aws lambda update-function-code --function-name $(LAMBDA_FUNCTION_NAME) --zip-file fileb://$(ZIP_FILE_NAME).zip

clean:
	@echo "Cleaning..."
	rm -f $(BIN) $(ZIP_FILE_NAME).zip

deploy: clean compile bundle push
