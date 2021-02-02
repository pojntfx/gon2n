TARGETS=supernoded supernodectl edged edgectl

all: generate $(TARGETS)

$(TARGETS):
	@go build -o out/$@ ./cmd/$@/main.go
	
generate:
	@go generate ./...

clean:
	@rm -rf ./out ./proto/generated ./pkg/workers/n2n
