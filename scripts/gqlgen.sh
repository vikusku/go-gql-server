#!/bin/bash
printf "\nRegenerating gqlgen files\n"
rm -f internal/gql/generated.go \
    internal/gql/models/generated_models.go \
    internal/gql/resolvers/resolvers.go
time go run -v github.com/99designs/gqlgen $1
printf "\nDone.\n\n"