run:
	APP_ENV=development HTTP_SERVER_PORT=5912 REDIS_ADDR=localhost:6381 \
	REDIS_PASSWORD=sijdiasjodjioasijdijo21 ANALYZER_CLIENT_URL=http://localhost:5913 \
	go run cmd/main.go

swagger:
	swag init -g cmd/main.go

.PHONY: openapi
openapi:

	@if [ -d ./internal/adapters/clients/http/analyzer ]; then rm -rf ./internal/adapters/clients/http/analyzer; fi

	openapi-generator generate \
		-i openapi/swagger.yaml \
		-g go \
		-o internal/adapters/clients/http/analyzer\
  --additional-properties=packageName=analyzer,withGoMod=false

	find ./internal/adapters/clients/http/analyzer -type f ! -name "*.go" -exec rm -v {} \;

	rm -rf ./internal/adapters/clients/http/analyzer/test

	@echo "Openapi client generated successfully! ✅"