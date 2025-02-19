#!/bin/bash

goctl api go -api *.api -dir ../ -style=goZero

goctl api plugin -plugin goctl-swagger="swagger -filename usercenter.json" -api usercenter.api -dir .

