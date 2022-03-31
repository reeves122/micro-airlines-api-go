run:
	#docker run -p 3306:3306 -e MYSQL_ROOT_PASSWORD=secret -e MYSQL_DATABASE=micro-airlines -d mysql:5.7 || true
	cd src && go run cmd/main.go

docker:
	docker build -t micro-airlines-api-go .
	docker run micro-airlines-api-go

test:
	cd src && go test ./...

coverage:
	cd src && go test -json -coverprofile=cover.out ./... > result.json
	cd src && go tool cover -func cover.out
	cd src && go tool cover -html=cover.out

mocks:
	cd src && go install github.com/golang/mock/mockgen@latest
	cd src && mockgen -source repository/repository.go -destination mocks/mock_repository/mock.go -package mock_repository

fmt:
	cd src && go fmt ./...

tidy:
	cd src && go mod tidy