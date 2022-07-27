build:
	go build -o bin/husky

run:
	@go build -o bin/husky
	@bin/husky bin/main.husky

compile:
	GOOS=linux GOARCH=amd64 go build -o bin/husky
	GOOS=linux GOARCH=arm64 go build -o bin/husky-arm64
	GOOS=darwin go build -o bin/husky-darwin
	GOOS=windows go build -o bin/husky.exe

git:
	git add . -f
	git commit -m "husky v1.0b [automated commit]"
	git push

clean:
	go clean
	rm bin/husky

all: compile