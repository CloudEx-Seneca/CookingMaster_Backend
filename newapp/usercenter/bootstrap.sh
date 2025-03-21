#!/bin/bash

export DB_HOST="127.0.0.1"
export MYSQL_ROOT_PASSWORD="haojiefu"
export SERVER_PORT=":8080"
go run usercenter.go
