build:
	go build -o bin/husky

run:
	go build -o bin/husky
	bin/husky test.husky

clean:
	go clean
	rm bin/husky