# libnbtscango

## Build nbtscan binary
Download source code from [inetcat.org](http://www.inetcat.org/software/nbtscan.html).
untar and compile
```
$ tar zxvf nbtscan_1.5.1.tar.gz
$ cd nbtscan_1.5.1a
$ make
```
Place binary to data folder
```
$ cp nbtscan ~/go/libnbtscango/data/
```

## Generate code
get go-bindata by
```
$ go get -u github.com/jteeuwen/go-bindata/...
```
generate code in libnbtscango project folder
```
$ go-bindata data/
```
Rename package name
```
$ vim bindata.go
```
