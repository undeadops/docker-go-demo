
linux: 
	GOOS=linux GOARCH=amd64 go build -v -o web web.go

mac:
	GOOS=darwin GOARCH=amd64 go build -v -o web-darwin web.go

docker: linux
	docker build -t undeadops/docker-go-demo .
	docker push undeadops/docker-go-demo

run:
	docker run --rm -p 8080:8080 -e "PORT=8080" undeadops/docker-go-demo

mongo:
	docker run --name mymongo -p 27017:27017 -v $(shell pwd)/mongodata:/data/db -d mongo 
