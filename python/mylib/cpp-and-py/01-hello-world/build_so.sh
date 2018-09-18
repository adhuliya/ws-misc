#!/usr/bin/env bash

# it takes a *.cpp file to be compiled
# it generates a shared object library.
# if input file is xyz.cpp, output library will be libxyz.so

#CXX=g++
CXX=clang++

FILE=$1

FILE_EXT="${FILE##*.}"
FILE_NAME="${FILE%.*}"


$CXX -c -fPIC $FILE -o $FILE_NAME.o
$CXX -shared -Wl,-soname,lib${FILE_NAME}.so -o lib${FILE_NAME}.so  ${FILE_NAME}.o

