#!/usr/bin/env bash
GOARCH=amd64 GOOS=linux go build -o bin/go-aws-eb &&
eb deploy