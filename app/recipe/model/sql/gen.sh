#!/bin/bash

goctl model mysql ddl -c -src *.sql -dir ../ -style=goZero
