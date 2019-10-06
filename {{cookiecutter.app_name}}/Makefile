default: deps
	go install

deps:
	go mod tidy
	go mod vendor

docker: deps
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
	chmod +x ./main
	sudo docker build -t {{cookiecutter.app_name}} .

clean:
	rm -rf vendor/
	rm -rf go.sum

.PHONY: go test