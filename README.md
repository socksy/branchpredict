branchpredict
=============

To build, you need golang. This can be found at http://golang.org/doc/install
Alternatively, on linux you can usually install with your package manager:

		sudo pacman -S go
		sudo apt-get install golang
		
etc

In this directory, run "go build" to build the application, which will produce an executable with the same name as the directory. Alternative, "go build -o branchpredictor" will produce the binary "branchpredictor".

The branch predictor binary can take two flags
-i specifies the input file, for instance:

		./branchpredictor -i tests/ls.out
		
--csv specifies to produce a comma separated list, useful in data analysis

		./branchpredictor --csv
		
These can be combined.

The tests can be run with the tests.sh script, which produces a CSV file consisting of misprediction rates of the different predictors on all of the files in the tests directory ending with ".out".

Some of the tests are gzipped, for  reasons of sanity.
