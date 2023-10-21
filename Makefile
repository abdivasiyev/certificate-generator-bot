gen:
	go generate ./...

ngrok:
	ngrok http 8080

.PHONY: gen

.DEFAULT_GOAL:=gen