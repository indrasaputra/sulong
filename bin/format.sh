#!/usr/bin/env bash

set -euo pipefail

goimports -w -local github.com/indrasaputra/sulong $(go list -f {{.Dir}} ./...)
gofmt -s -w .
