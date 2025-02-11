## Design Document
[User Center API Design](https://cloudex-seneca.atlassian.net/wiki/spaces/DD/pages/66102/User+Center+API+Design?atlOrigin=eyJpIjoiZDEwZDhiMGEzM2JhNGQ1OGI3YTNjMTE0MjMyZjQzNWQiLCJwIjoiaiJ9)

## Dependency
You need to install software below:
- mysql
- redis
- etcd
- golang

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
 cd app/usercenter/api 
 go run usercenter.go
```