#-rw-r-----  1 strick eng  532 Aug 20 10:46 fib-ints_test.go
#-rw-r-----  1 strick eng  747 Aug 20 10:45 fib_test.prego
#-rw-r-----  1 strick eng  440 Aug 20 10:37 macro-dd.prego
#-rw-r-----  1 strick eng  524 Aug 20 10:40 macro-ints.prego

all: fib-ints_test.go fib-dd_test.go

fib-ints_test.go:
	#rm -f *.go
	go run ../../prego/main.go --source macro-ints.prego < fib_test.prego   > fib_test.go && go test --bench=.

fib-dd_test.go:
	#rm -f *.go
	go run ../../prego/main.go --source macro-dd.prego < fib_test.prego   > fib_test.go && go test --bench=.
