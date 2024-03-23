SWAGGER_PATH = ./api/server.go

swagger:
	swag fmt -g $(SWAGGER_PATH)
	swag init -g $(SWAGGER_PATH) --parseDependency --parseInternal