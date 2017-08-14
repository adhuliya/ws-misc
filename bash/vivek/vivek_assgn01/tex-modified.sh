#!/bin/bash

# This script finds and compares source file modification time
# with the destination files and categorizes them into four classes
# 1. Source file is Older
# 2. Source file is Newer
# 3. Source file is same as destination file
# 4. Source file has no destination

SRC_EXT="tex" # source extension
DST_EXT="pdf" # destination extension

# source files older than their destination files
OLDER_SRC=""
# source files newer than their destination files
NEWER_SRC=""
# same modification time stamp
EQUAL_SRC=""
# source files with no destination files
NO_DST=""
# all other files/folders that do not fit the above
OTHERS=""

for filename in *;
do
  #extension= $([[ "$filename" = *.* ]] && echo ".${filename##*.}" || echo "";   
  extension="${filename##*.}"
  if [[ -d "$filename" || "$extension" = "$filename" || "$extension" != "$SRC_EXT" ]];
    then
      # filename has no extension, or a different one hence continue
      OTHERS="${OTHERS},${filename}"
      continue
  fi

  filenameprefix="${filename%.*}"
  dstfilename="${filenameprefix}.${DST_EXT}"
  if [[ ! -e "${dstfilename}" ]]
  then
      # no destination file exists
      NO_DST="${NO_DST},${filename}"
      continue
  fi

  if [[ ${filename} -nt ${dstfilename} ]]
  then
      # source file is newer than the destination file
      NEWER_SRC="${NEWER_SRC},${filename}"
      continue
  fi

  if [[ ${filename} -ot ${dstfilename} ]]
  then
      # source file is older than the destination file
      OLDER_SRC="${OLDER_SRC},${filename}"
      continue
  else
      EQUAL_SRC="${EQUAL_SRC},${filename}"
  fi


done

echo "${SRC_EXT} files newer than ${DST_EXT} files: ${NEWER_SRC}"
echo
echo "${SRC_EXT} files older than ${DST_EXT} files: ${OLDER_SRC}"
echo
echo "${SRC_EXT} files equal to   ${DST_EXT} files: ${EQUAL_SRC}"
echo
echo "${SRC_EXT} files with no    ${DST_EXT} files: ${NO_DST}"
echo
echo "Not ${SRC_EXT} files : ${OTHERS}"
echo


