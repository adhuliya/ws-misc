#!/bin/bash

#AD modified to handle space in filenames

#AD check if input given?
if [[ -z "$1" || -z "$2" ]]
then
  (>&2 echo -e "$0 takes two arguments.\nArg1: An existing name of file.\nArg2: File extension.")
  (>&2 echo "Example: $0 filenames.txt tex")
  exit 1
fi

filename="$1"
ext="$2"

i=0
j=0


#AD by default ls sorts files by name
#AD read the current dir file names
array1=(*.$ext)
i=${#array1[@]}

#AD read the file with new names
array2[0]=""
while read name; do
  array2[$j]="$name"
  j=$(($j+1))
done < $filename

if [[ $j -lt $i ]]; then
  (>&2 echo "Error: New filenames are less than number of current files. Exiting.")
  exit 1
fi

for((i=0;i<${#array1[@]};i++));
do
  mv "${array1[$i]}" "${array2[$i]}"
done


