#!/usr/bin/env bash

mkdir -p dist

go build -ldflags="-s -w" -o ./bin/detect ./cmd/detect/main.go || exit 166
go build -ldflags="-s -w" -o ./bin/build ./cmd/build/main.go || exit 166

# https://buildpacks.io/docs/buildpack-author-guide/package-a-buildpack/
pack buildpack package dist/laravel-model.cnb --config ./package.toml --format file
pack buildpack package laravel-model --config ./package.toml