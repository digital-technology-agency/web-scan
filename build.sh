#!/usr/bin/env bash
PATH_COMPILE=/usr/local/go/go/bin/go
# ./wscan -alphabet "abcdefghijklmnopqrstuvwxyz" -len 5
# ./wscan -alphabet "abcdef" -len 2
echo | $PATH_COMPILE version
echo | mkdir -p dist
echo  'Linux'
echo | $PATH_COMPILE build -o  dist/wscan cmd/distr/wscan.go
