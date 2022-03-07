#!/bin/bash

cd crypto-config

echo "converting filenames to 'key'"
files=$(find . -name "*_sk")
for f in $files
do
echo $f
base=$(dirname $f)
mv $f ${base}/key
done
