#!/usr/bin/env bash
go build  -o get_news main.go
nohup ./get_news  &