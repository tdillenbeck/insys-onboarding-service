#!/bin/bash

for Dir in $(go list ./... | grep -v '/vendor/'); 
do
	`go vet $Dir`
    returnval=$?
    if [[ ${returnval} != 0 ]]
    then
        exit 1
    fi  

done
