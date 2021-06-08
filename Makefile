gen:
	go build -o bin/check check/cmd/*.go
	go build -o bin/in in/cmd/*.go
	go build -o bin/out out/cmd/*.go

clean:
	rm -rf bin/
