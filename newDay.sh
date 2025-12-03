#!/bin/env bash

year=$1
day=$2

if [ -z "$day" ]; then
    echo "usage: $0 <2015...> <01-24>"
    exit 1
fi

targetDir="$year/day$day"

mkdir -p $targetDir
cp template/main.go $targetDir/main.go
cp template/go.mod $targetDir/go.mod
cp template/go.sum $targetDir/go.sum
cp template/input.txt $targetDir/input.txt
cp template/day1_test.go $targetDir/day${day}_test.go
cp template/day1.go $targetDir/day${day}.go
go work use ./$targetDir

for file in $targetDir/*; do
    sed -i -e "s/day1/day$day/g" $file
    sed -i -e "s/2025\/day1/$year\/day$day/g" $file
    sed -i -e "s/Day1/Day$day/g" $file
    sed -i -e "s/aoc\/2025\/day01/aoc\/$year\/day$day/g" $file
done
