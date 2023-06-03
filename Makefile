build:
	docker build . -t centrifugo-custom -f centrifugo/Dockerfile
run-centrifugo:
	docker run -p 8000:8000 --name centrifugo centrifugo/centrifugo:v4
