#!/bin/sh

binPath="runtime/main.exe"
port="6050"

alloc() {
    go tool pprof ${binPath} -alloc_space -cum -svg localhost:${port}/debug/pprof/heap
}

case "$1" in
    alloc)
        alloc
        ;;
    *)

	echo "Usage: $0 {alloc}"
esac