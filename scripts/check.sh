#!/usr/bin/env sh
printf "Running Go fmt \n"
go fmt  cmd/wotracker-back/main.go
printf "Running Go vet \n"
go  vet cmd/wotracker-back/main.go
printf "Running golangci-lint \n"
golangci-lint run