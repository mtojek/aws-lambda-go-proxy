build:
	go get -v ./...
	GOOS=linux go build -o main
	zip deployment.zip main
	rm main
