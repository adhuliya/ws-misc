#AD processes filenames given in a file (one filename per line)
#AD This script tranlates each line of each file. Each vovel is replaced by the next vovle a -> e, A -> E, U -> A, etc..

#AD check if input given?
if [[ -z "$1" ]]
then
  (>&2 echo -e "$0 takes one argument.\nArg1: A file containing newline separated filenames.")
  exit 1
fi

#AD check if correct input given
if [[ ! -e "$1" ]]; then
  (>&2 echo "Error: $1 file doesnot exist")
fi

filename="$1"

while read name; do
  newfile="$(mktemp)"
  while read line; do
    newline="$(echo $line | tr aeiouAEIOU eiouaEIOUA)"
    echo $newline >> "$newfile"
  done < "$name"
  rm -f "$name"
  mv "$newfile" .
  mv "$(basename $newfile)" "$name"
  chmod +rw "$name"
done < "$filename"
