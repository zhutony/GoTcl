libtclgoexample.so: example.go wrappers.go
	go build -o $@ -buildmode=c-shared example.go wrappers.go
test: libtclgoexample.so
	echo 'load libtclgoexample.so; puts [hello]; puts [square 5]' | tclsh
.PHONY: test
