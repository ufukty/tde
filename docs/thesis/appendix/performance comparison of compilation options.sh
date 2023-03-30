#!/bin/bash

# run:
# cat <file>; time bash <file>

tempdir="$(mktemp -d)"
echo "temp: $tempdir"

for i in {1..100}; do
    # go build -o /dev/null ./cmd/client
    # go build -o "$tempdir/$i" ./cmd/client
    # go install -v ./cmd/client
    # go install ./cmd/client

    # go build -o /dev/null ./cmd/client &
    # go build -o "$tempdir/$i" ./cmd/client &
    # go install -v ./cmd/client &
    # go install ./cmd/client &

    # go run ./cmd/client >/dev/null 2>&1 &

    # chmod +x "$i"
    # "./$i" >>log 2>&1 &
done

wait $(jobs -p)
