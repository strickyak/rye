all: fib-ints_test.go fib-dd_test.go fib-au_test.go fib-su_test.go

fib-ints_test.go:
	#rm -f *.go
	go run ../../prego/main.go --source macro-ints.prego < fib_test.prego   > fib_test.go && go test --bench=.

fib-dd_test.go:
	#rm -f *.go
	go run ../../prego/main.go --source macro-dd.prego < fib_test.prego   > fib_test.go && go test --bench=.

fib-au_test.go:
	#rm -f *.go
	go run ../../prego/main.go --source macro-au.prego < fib_test.prego   > fib_test.go && go test --bench=.

fib-su_test.go:
	#rm -f *.go
	go run ../../prego/main.go --source macro-au.prego < fib_test.prego   > fib_test.go && go test --bench=.
