#!/bin/bash

docker run --rm -v "$PWD":/usr/src/vd -w /usr/src/vd golang:1.19 go build -v