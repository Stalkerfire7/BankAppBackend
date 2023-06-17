build:
@go build -o bin/BankAppBackend

run: build
@	./bin/BankAppBackend
test:
@ go test -v ./...