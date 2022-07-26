build:
	go build -o bin/husky

run:
	go build -o bin/husky
	bin/husky bin/main.husky

git:
	git add . -f
	git commit -m "husky v1.0b [automated commit]"
	git push

clean:
	go clean
	rm bin/husky