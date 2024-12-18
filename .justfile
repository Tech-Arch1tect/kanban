# list all the recipes
list:
  just --list

# run the server
server: 
  cd server && go run .

# run the server with hot reloading
server-watch: 
  cd server && air

# run client
client: 
  cd client && npm run dev

# generate swagger server docs
swagger-server: 
  cd server && swag init --pd --parseInternal

# generate swagger api client
swagger-client: 
  cd client/src && rm -rf typescript-fetch-client && openapi-generator-cli generate -i ../../server/docs/swagger.yaml -g typescript-fetch -o typescript-fetch-client && rm openapitools.json

# generate swagger server and client docs
swagger-update: 
  just swagger-server && just swagger-client
