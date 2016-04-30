
linux: 
	GOOS=linux GOARCH=amd64 go build -v -o web web.go

mac:
	GOOS=darwin GOARCH=amd64 go build -v -o web-darwin web.go

docker: linux
	docker build -t undeadops/go-demo .

run:
	docker run --rm -p 8080:8080 -e "PORT=8080" undeadops/go-demo

mongo:
	docker run --rm --name mymongo -p 27017:27017 -v /Users/mitch/mongo:/data/db -d mongo 
