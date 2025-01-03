# Default recipe: Lists all available recipes
default:
  just --list

# Server commands
# Starts the server using Go.
[working-directory: "server"]
server:
  @go run cmd/main.go

# Starts the server with hot reloading using Air.
[working-directory: "server/cmd"]
server-watch:
  @air

# Lints the server code.
[working-directory: "server"]
lint:
  @golangci-lint run

# Client commands
# Starts the client development server using npm.
[working-directory: "client"]
client:
  @npm run dev

# Swagger commands
# Generates Swagger server documentation.
[working-directory: "server"]
swagger-server:
  @swag init --pd --parseInternal -g cmd/main.go

# Generates Swagger API client in TypeScript.
[working-directory: "client/src"]
swagger-client:
  @rm -rf typescript-fetch-client && \
  openapi-generator-cli generate -i ../../server/docs/swagger.yaml \
  -g typescript-fetch -o typescript-fetch-client && \
  rm openapitools.json

# Generates both Swagger server documentation and TypeScript client.
swagger: swagger-server swagger-client