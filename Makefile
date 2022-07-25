build:
	go build -o bin/husky

run:
	go build -o bin/husky
	bin/husky bin/main.husky

clean:
	go clean
	rm bin/husky