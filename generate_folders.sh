#!/bin/bash

N=$1

for (( i=1; i<=N; ++i ))
do
    folder_name="task_$i"
    file_name="$folder_name/main.go"

    mkdir "$folder_name"
    touch "$file_name"

    echo "package main" >> "$file_name"
    echo "" >> "$file_name"
    echo "/*" >> "$file_name"
    echo "Задание:" >> "$file_name"
    echo "*/" >> "$file_name"
    echo "" >> "$file_name"
    echo "func main() {" >> "$file_name"
    echo "" >> "$file_name"
    echo "}" >> "$file_name"
done