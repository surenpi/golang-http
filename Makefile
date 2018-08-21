NAME=sample

build: clean
	GOOS=linux go build -o bin/${NAME}

all: build

clean:
	rm -rfv bin

image:
	docker build -t sample .