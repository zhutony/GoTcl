libtclgoexample.so: example.go wrappers.go
	go build -o $@ -buildmode=c-shared example.go wrappers.go
test: libtclgoexample.so
	echo "using /usr/lib/libtcl8.5.dylib"
	echo 'load libtclgoexample.so; puts [hello]; puts [square 5]' | tclsh8.5
.PHONY: test
