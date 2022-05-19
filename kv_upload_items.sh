#!/bin/bash

# Author : https://github.com/Alxus228
# This script Accepts source text file containing JSON key-value pairs per line;
# and stores these records on the running key-value server.

# Example of an input file:
# {"key": 123, "value": "banana"}
# {"key": "abc", "value": 123456}

# Input should be provided only according to the standard, given in the example.

# Start of the script:
# Reading all lines from the file into a variable
input=`cat $1`
# Defining an associative array and temporary variables, to withdraw needed info from the input.
flag=0
key=0
value=0
declare -A pair

# Decomposing json objects into key-value pairs
for argument in $input
do
  # If flag equals 1, we add $argument to the key array
  if [[ $flag == 1 ]]; then
    key=${argument//,}
    key=${key//\"}
    flag=0
  fi
  # If flag equals 2, we add $argument to the value array
  if [[ $flag == 2 ]]; then
      value=${argument//"}"}
      value=${value//\"}
      pair+=([$key]=$value)
      flag=0
  fi

  # Setting flag to 1 if argument equals "key"
  if [[ $argument == *"key"* ]]; then
    flag=1
  fi
  # Setting flag to 2 if argument equals "value"
  if [[ $argument == *"value"* ]]; then
      flag=2
  fi
done

# This is the server address.
url=https://localhost:443

# Putting all records into the server via curl command
for key in ${!pair[@]}
do
  curl -X PUT ${url}/api/${key} -d ${pair[$key]}
done