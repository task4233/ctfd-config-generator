#!/bin/bash

genres=("crypto", "hardware", "misc", "network", "osint", "pwn", "rev", "web")
for genre in "${genres[@]}"; do
    sorted_genres=$(find $genre | grep "challenge.yml$" | sort)

    for g in $sorted_genres; do
        python -m ctfcli challenge install $g;
        sleep 1; # delay for server
    done
done
