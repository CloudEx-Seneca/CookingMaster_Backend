#!/bin/bash

export DB_HOST="127.0.0.1"
export MYSQL_ROOT_PASSWORD="haojiefu"
export SERVER_PORT=":8081"
go run recipe.go
