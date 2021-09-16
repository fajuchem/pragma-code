#!/bin/bash

cd client/ && npm install && npm run build

cd ../server && go run main.go