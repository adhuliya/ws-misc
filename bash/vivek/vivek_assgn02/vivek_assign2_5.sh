#AD processes filenames given in a file (one filename per line)
#AD This script counts the number of digits and letters in each file

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
  digits="$(cat "$name" | tr -dc "0-9" | wc -c)"
  letters="$(cat "$name" | tr -dc "a-zA-Z" | wc -c)"
  echo "$name letters:$letters digits:$digits"
done < "$filename"
