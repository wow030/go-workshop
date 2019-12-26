# Fib
## benchmark for Fib

# run bench as much times as possible
## get benchstat to help us 
`go get golang.org/x/perf/cmd/benchstat`<br>
`go test -bench=Fib20 -count=10 | tee old.txt`<br>
`benchstat old.txt`