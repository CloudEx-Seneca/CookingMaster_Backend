# Deploy app in docker
## Create user defined network
```
docker network create app_network
```

## Running mysql

```
docker run -d --name backend_db --network app_network -e MYSQL_ROOT_PASSWORD=pw  henry071190/backend_db
```

## Running backend
```
docker run -d --name backend --network app_network -p 8888:8888 henry071190/backend
```

## Running frontend

```
docker run -d --name frontend --network app_network -p 3000:3000 henry071190/frontend:5
```

# [BUILD] Build docker image
## Clone backend repo

```
git clone https://github.com/CloudEx-Seneca/CookingMaster_Backend.git
cd CookingMaster_Backend
```

## Building application docker image

```
docker build -t backend -f Dockerfile_usercenter . 
```

## Building mysql docker image

```
docker build -t backend_db -f Dockerfile_mysql . 
```

## DEPLOY USING ENV VARIABLE

```
export DBHOST=172.18.0.2
export DBPORT=3306
export DBUSER=root
export DATABASE=usercenter
export DBPWD=pw
```

## Run backend

```
docker run -d --name backend --network app_network -p 8888:8888 -e DBHOST=$DBHOST -e DBPORT=$DBPORT -e  DBUSER=$DBUSER -e DBPWD=$DBPWD  backend
```

## Run the application, make sure it is visible in the browser

```
docker run -d --name frontend --network app_network -p 3000:3000 -e DBHOST=$DBHOST -e DBPORT=$DBPORT -e  DBUSER=$DBUSER -e DBPWD=$DBPWD  henry071190/frontend:5
```
