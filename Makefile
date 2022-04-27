go-run:
	cd cmd; go run main.go
build:
	cd cmd; go build .
test:
	go test ./...
run-compose:
	sudo docker-compose up --build -d; docker attach challenge-app
