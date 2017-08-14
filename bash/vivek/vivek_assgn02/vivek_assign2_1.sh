# This scrip take two arguments : 1) a large number n (n > 50 ), ii) a three character string (as extension) and creates n files with randomly chosen file names and with the given fixed extension.
# Each file should contain:
# m lines of text where m is randomly chosen between 1<=m<=5.
# The text in each line contains p characters where p is randomly chosen between 10<=p<=30.
# The characters in each line are randomly chosen from 'a' to 'z', 'A' to 'Z', '0' to '9'.
# Also it must contain between 2 to 5 (choose number randomly) comma and/or space characters embedded randomly inside the earliar sequence of characters. A comma or space can not be first character or the last character.
#echo "Enter the total no. of files & 3 digits extension of file"

#AD check if input given?
if [[ -z "$1" || -z "$2" ]]
then
  (>&2 echo -e "$0 takes two arguments.\nArg1: A number >50.\nArg2: A length three non-whitespace string.")
  (>&2 echo "Example: $0 51 tex")
  exit 1
fi

#AD check if correct input given
n="$1"
ext="$2"
len="$(echo $ext|wc -c)"
len=$(($len - 1)) #AD
if [[ $n -le 50 || $len -ne 3 ]] # check for valid input
then
  if [[ $n -le 50 ]]
  then
    (>&2 echo "First argument should be a number greater than 50.")
  fi

  if [[ $len -ne 3 ]]
  then
    (>&2 echo "Second argument should be a three character non-whitespace string.")
  fi

  exit 2 #AD reject input
fi

#AD if here, input accepted

array[0]=","
array[1]=" "
size=${#array[@]}

i=1
while [[ $i -le $n ]]
do
    m=$(( 1 + $RANDOM % 5 )) # random selectiom of number of lines for each file

    #AD select a random filename, make sure it doesnot already exist
    FILE="${RANDOM}${RANDOM}.$ext" # random file name with fixed extension
    while [[ -e $FILE ]]; do
      FILE="${RANDOM}${RANDOM}.$ext" # random file name with fixed extension
    done

    #AD create and write each line into FILE
    j=1
    p=$(( 10 + $RANDOM % 21 )) # random line length 10 to 30
    #AD each line is exactly $p lenght
    while [[ $j -le $m ]]; do
      k=$(( 2 + $RANDOM % 4 )) #AD a value 2 to 5
      pp=$(($p-$k))
      line="$(cat /dev/urandom | tr -dc 'a-zA-Z0-9' | fold -w $pp | head -n 1)"
      while [[ $k -ge 1 ]]
      do
          linelen=$(($p - $k)) #AD
          count=$(( 1 + $RANDOM % ($linelen - 1) ))
          index=$(($RANDOM % $size))
          line="$( echo $line | sed 's/^\(.\{'$count'\}\)/\1\'"${array[$index]}"'/')"
          k=$(($k-1)) #AD
      done
      echo "$line" >> $FILE
      j=$(($j+1))
    done

    i=$(($i+1)) #AD
done


