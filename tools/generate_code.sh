#!/bin/sh

go run tools/build_tl_scheme.go < mtproto/schemes/api_layer_23.json > mtproto/tl_schema_generated.go
gofmt -w mtproto/tl_schema_generated.go
