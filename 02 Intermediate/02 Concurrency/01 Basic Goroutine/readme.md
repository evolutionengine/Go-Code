#Understanding Concurrency In Golang

*Concurrency is different than parallelism, we would be touching more of concurrent operations in Go*

The 'go routines' form the basic blocks of concurrency and are simply denoted as go funcName, e.g - go foo()

The main() also runs in a go routine by default and so you need to check for race conditions for incorporating concurrency into your programs. 

Otherwise main() will close before other go routines called from main() finish execution.
