#!/bin/bash

goctl api go -api *.api -dir ../ -style=goZero

goctl api plugin -plugin goctl-swagger="swagger -filename recipe.json" -api recipe.api -dir .

