NAME=sample

all:
	GOOS=linux go build -o bin/${NAME}

clean:
	rm -rfv bin

image:
	docker build -t sample .