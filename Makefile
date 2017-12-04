all: codi
.PHONY: all

codi: main.go
	go build -o codi main.go

clean:
	rm -f codi
.PHONY: clean
