gggo:
	go build -o gggo main.go

test: gggo
	./test.sh

clean:
	rm -f gggo

.PHONY: gggo clean
