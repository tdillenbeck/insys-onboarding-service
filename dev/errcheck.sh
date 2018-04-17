#!/bin/bash

for Dir in $(go list ./... | grep -v '/vendor/');
do

    # -blank to check _ handled errors
    cmd="errcheck -asserts -ignoretests -exclude ./dev/errcheck_excludes.txt $Dir"
    echo $cmd

    $cmd
    returnval=$?

    if [[ ${returnval} != 0 ]]
    then
        exit 1
    fi  

done
