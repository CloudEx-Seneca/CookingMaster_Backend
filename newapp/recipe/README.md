## Dependency
You need to install software below:
- mysql mysql  Ver 8.4.4
- golang go version go1.23.5

And you need to start the service below using default port:
- mysql

## Run Service
### init db
```
create database recipe;
```
### start server
```
# modify variables below as you need.
export DB_HOST="127.0.0.1"
export MYSQL_ROOT_PASSWORD="haojiefu"
export SERVER_PORT=":8081"
go mod tidy
cd newapp/recipe
go run recipe.go
```
### api document & test

- The service support swagger api view and test, please open http://localhost:8081/swagger/index.html#/ after you
  start the server.
- Click "try it out" inside every api to test the api, you can also get curl command after you execute the test.
- For api which has a lock logo, that means the api needs authentication, you need to click the lock logo to add
  the jwt token to request header, the format should be "Bearer {token}"
