.PHONY: default all

default: build deploy

build: 
	go build -o bin/handcontroller

deploy:
	pscp -P 16177 -pw p ./bin/handcontroller dima@pitunnel.com:/home/dima/handcontroller