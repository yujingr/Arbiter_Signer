all:
	go build -o arbiter app/arbiter/main.go

linux:
	GOARCH=amd64 GOOS=linux go build -o arbiter app/arbiter/main.go