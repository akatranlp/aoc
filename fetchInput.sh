#!/bin/env bash

year=$1
day=$2
cookie=$3

if [[ -z $year || -z "$day" || -z "$cookie" ]]; then 
    echo "Usage: $0 year day cookie"
    exit 1
fi

fetchday=$day
[[ ${fetchday:0:1} == "0" ]] && fetchday=${fetchday:1}

curl --cookie "session=$cookie" https://adventofcode.com/$year/day/${fetchday}/input -o $year/day$day/input.txt
