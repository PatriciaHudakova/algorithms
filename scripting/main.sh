#!/bin/bash

# echo command
echo "hello world"

# variables - uppercase by convention
NAME="Patricia"
echo "hello $NAME" #${NAME} also a variant

# user input
read -p "Enter your name:" MYNAME
echo "Hello $MYNAME, nice to meet you"

# conditionals
if [ "$NAME" == "Patricia" ]
then
  echo "Your name is Patricia"
elif [ "$NAME" == "Jack" ]
then
  echo "Your name is Jack"
else
  echo "Your name is not Patricia or Jack"
fi

# comparison ( -gt, -lt, -ge, -le, -eq, -ne )
NUM1=3
NUM2=4

if [ "$NUM1" -gt "$NUM2" ]
then
  echo "$NUM1 is greater than $NUM2"
else
  echo "$NUM2 is greater than $NUM1"
fi

# file conditions
FILE=text.txt

# if true, -f checks if it's a file, -e checks if it exists regardless of form - could be a directory
if [ -f $FILE ]
then
  echo "$FILE is a file"
else
  echo "$FILE is not a file"
fi


# case statement
read -p "Are you 18 or over? Answer y/n:" ANSWER
case "$ANSWER" in
  [y] | [yY][eE][sS])
  echo "You can have a beer :)"
  ;;
  [n] | [nN][oO])
  echo "Sorry, no beer :("
  ;;
  *)
  echo "Please enter a y or n"
  ;;
esac

# simple loop
NAMES="Alice Brad Jasmine"

for NAME in $NAMES
  do
    echo "Hello $NAME"
done

# for loop to rename files
#FILES=$(ls *.txt)
#NEW="new"
#
#for FILE in $FILES
#  do
#    echo "Renaming $FILE..."
#    mv $FILE $NEW-$FILE
#done

# while loop - read a text file line by line
LINE=1

while read -r CURENT_LINE
  do
    echo "$LINE: $CURENT_LINE"
    ((LINE++))
done < "./new-1.txt"

# functions
function sayHello() {
  echo "Hello World"
}

sayHello

function greet() {
  echo "Hello, I am $1 and I am $2"
}

greet "Patricia" "24"

# create folder and write to a file
#mkdir basics
#touch "hello/world.txt"
#echo "Hello World" >> "hello/world.txt"
#echo "Created world.txt"