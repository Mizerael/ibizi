.PHONY: all

all: get-message put-message

get-message: getMessage.go
	go build -o $@ $^

put-message: putMessage.go
	go build -o $@ $^

clean: 
	rm -f get-message put-message