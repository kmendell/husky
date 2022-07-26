build:
	go build -o bin/husky

run:
	go build -o bin/husky
	bin/husky bin/main.husky

git:
	git add . -f
	git commit -m "-Used make git- to push these changes"
	git push

clean:
	go clean
	rm bin/husky