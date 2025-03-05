## Design Document
[User Center API Design](https://cloudex-seneca.atlassian.net/wiki/spaces/DD/pages/66102/User+Center+API+Design?atlOrigin=eyJpIjoiZDEwZDhiMGEzM2JhNGQ1OGI3YTNjMTE0MjMyZjQzNWQiLCJwIjoiaiJ9)

## Dependency
You need to install software below:
- mysql mysql  Ver 8.4.4
- redis Redis server v=7.2.7
- etcd etcd Version: 3.5.17
- golang go version go1.23.5

And you need to start the service below using default port:
- mysql
- redis 
- etcd

## Run Service
## UserCenter Service
### init db
```
cd app/usercenter/model/sql 
mysql -u ${user} -p${password} < usercenter.sql
```
### run api
```
 go mod tidy
 cd app/usercenter/api 
 go run usercenter.go
```
### swagger view
1. download swagger binary from https://github.com/go-swagger/go-swagger/releases based on your platform.
2. execute commands below.
```
cd app/usercenter/api/desc
## swagger_darwin_amd64 for mac os and amd64 arch
${path}/swagger_darwin_amd64 serve -F=swagger usercenter.json --port 9088 --host 0.0.0.0 --no-open

```
3. open http://localhost:9088/docs from your browser.

## Recipe Service
### init db
```
cd app/recipe/model/sql 
mysql -u ${user} -p${password} < recipe.sql
```
### run api
```
 go mod tidy
 cd app/recipe/api 
 go run recipe.go
```
### swagger view
1. download swagger binary from https://github.com/go-swagger/go-swagger/releases based on your platform.
2. execute commands below.
```
cd app/usercenter/api/desc
## swagger_darwin_amd64 for mac os and amd64 arch
${path}/swagger_darwin_amd64 serve -F=swagger recipe.json --port 9088 --host 0.0.0.0 --no-open

```
3. open http://localhost:9088/docs from your browser.
