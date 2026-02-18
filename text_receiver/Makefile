run:
	APP_ENV=development HTTP_SERVER_PORT=5912 REDIS_ADDR=localhost:6381 \
	REDIS_PASSWORD=sijdiasjodjioasijdijo21 ANALYZER_CLIENT_URL=http://localhost:5913 \
	go run cmd/main.go

swagger:
	swag init -g cmd/main.go

