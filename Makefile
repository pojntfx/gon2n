CLIENTS=supernodectl edgectl
SERVERS=supernoded edged
TARGETS=$(SERVERS) $(CLIENTS)
OUT=out

all: $(TARGETS)

$(CLIENTS): generate
	@GOOS=linux GOARCH=amd64 go build -o ${OUT}/$@-linux-amd64 ./cmd/$@/main.go
	@GOOS=linux GOARCH=arm64 go build -o ${OUT}/$@-linux-arm64 ./cmd/$@/main.go
	@GOOS=darwin GOARCH=amd64 go build -o ${OUT}/$@-macos-amd64 ./cmd/$@/main.go
	@GOOS=windows GOARCH=amd64 go build -o ${OUT}/$@-windows-amd64.exe ./cmd/$@/main.go

$(SERVERS): generate
	@GOOS=linux GOARCH=amd64 go build -o ${OUT}/$@-linux-amd64 ./cmd/$@/main.go
	
generate:
	@go generate ./...

test: test-linux test-windows

test-linux: $(OUT)/*-linux-amd64
	for file in $^ ; do \
	    ./$${file} --help; \
	done
	
test-windows: $(OUT)/*-windows-amd64.exe
	for file in $^ ; do \
	    wine64 $${file} --help; \
	done

clean:
	@rm -rf ${OUT} pkg/proto/generated pkg/workers/n2n
