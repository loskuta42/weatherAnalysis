ENTRYPOINT=./cmd/forecasting
BINARY_NAME=weatherAnalysis

PKG=./...

all: build

run:
	go run $(ENTRYPOINT)

build:
	go build -o $(BINARY_NAME) $(ENTRYPOINT)

test:
	go test -v $(PKG)

clean:
	rm -f $(BINARY_NAME)

fmt:
	go fmt $(PKG)	

deps:
	go mod tidy