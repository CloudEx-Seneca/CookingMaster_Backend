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
init db
```
cd app/usercenter/model/sql 
mysql -u ${user} -p${password} < usercenter.sql
```
run api
```
 go mod tidy
 cd app/usercenter/api 
 go run usercenter.go
```