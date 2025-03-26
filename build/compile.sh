#!/bin/bash
baseDir="./cmd"

go get gitlab.casinovip.tech/minigame_backend/c_engine@main
go get gitlab.casinovip.tech/minigame_backend/om_struct@main

go mod tidy

dirs=$(find "$baseDir" -type d)

for dir in $dirs; do
    if [ -f "$dir/main.go" ]; then
        echo "Compiling $dir..."

        cd "$dir" || exit
        go build

        if [ $? -eq 0 ]; then
            echo "Compilation succeeded"
        else
            echo "Compilation failed"
        fi

        cd - >/dev/null || exit
    fi
done
