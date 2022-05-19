#!/bin/bash

# Author : https://github.com/Alxus228

# Reading all lines from the file into a variable
input=`cat $1`

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

#url=169.254.245.82:443
url=https://localhost:443

# Putting all records into the server via curl command
for key in ${!pair[@]}
do
  #echo $key ${pair[$key]}
  curl -X PUT ${url}/api/${key} -d ${pair[$key]}
done