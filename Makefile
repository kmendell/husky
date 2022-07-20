build:
	go build -o bin/husky

run:
	go build -o bin/husky
	bin/husky bin/test.husky

clean:
	go clean
	rm bin/husky