#!/bin/bash
tracefiles=./tests/*.out
directoryname=${PWD##*/}

cat header.csv > tests.csv

echo $tracefiles
for f in $tracefiles 
do
	echo "Predicting $f"
	printf "$f," >> tests.csv
	./branchpredictor -csv -i $f >> tests.csv
	printf "\n" >> tests.csv
done
