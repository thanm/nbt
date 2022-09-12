#!/bin/sh
set -x
set -e
rm -rf covdata cov.p main.exe
go build -cover -o main.exe
mkdir covdata
GOCOVERDIR=covdata ./main.exe
go tool covdata textfmt -o cov.p -i covdata
echo done
