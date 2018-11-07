build:
	GOOS=linux GOARCH=amd64 go build
	docker build -t email-service .

run:
	docker run --net="host" \
		-p 50054 \
		-e MICRO_REGISTRY=mdns \
		-e MICRO_BROKER=nats \
		-e MICRO_BROKER_ADDRESS=0.0.0.0:4222 \
		email-service
