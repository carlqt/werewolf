migrate:
	dbmate create
	dbmate up

prepare:
	docker-compose up -d

start:
	go run *.go
