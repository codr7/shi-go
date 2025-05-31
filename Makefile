all: test
	go build -C src -o ../bin/shi
	bin/shi

test: 
	go test src/tests/*

clean:
	rm -f bin/*
