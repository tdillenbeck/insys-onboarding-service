#!/bin/bash

echo "mode: set" > acc.out
for Dir in $(go list ./... | grep -v '/vendor/'); 
do
    go test -v -i $Dir
    go test -v -coverprofile=profile.out $Dir
    returnval=$?
    if [[ ${returnval} == 0 ]]
    then
        if [ -f profile.out ]
        then
            cat profile.out | grep -v "mode: set" >> acc.out
        fi
    else
        exit 1
    fi

done
if [ -n "$COVERALLS_TOKEN" ]
then
    goveralls -coverprofile=acc.out -repotoken=$COVERALLS_TOKEN -service=travis-pro
fi

rm -rf ./profile.out
rm -rf ./acc.out
