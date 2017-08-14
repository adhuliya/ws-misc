#Rename the files with given prefix and fix suffix
#AD check if input given?
if [[ -z "$1" || -z "$2" ]]
then
  (>&2 echo -e "$0 takes two arguments.\nArg1: A prefix.\nArg2: A file extension.")
  (>&2 echo "Example: $0 file tex")
  exit 1
fi

prefix="$1"
ext="$2"

i=1
for file in *.$ext
do
  #ext="${file##*.}"
  nfile="$prefix$i.$ext" 
  mv $file $nfile
  i=$(($i + 1))
done


