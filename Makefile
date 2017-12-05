all: codi
.PHONY: all

codi: main.go
	go build -o ./bin/codi main.go

test:
	./bin/codi "A+->Gateway+->Internet\nor\nCorporate network"
.PHONY: test

clean:
	rm -f ./bin/codi
.PHONY: clean
