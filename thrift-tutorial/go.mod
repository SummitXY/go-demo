module github.com/SummitXY/thrift-tutorial

go 1.16

require (
	git.apache.org/thrift.git v0.14.2 // indirect
)

replace (
	github.com/SummitXY/thrift-tutorial/gen-go/shared => ./gen-go/shared
	github.com/SummitXY/thrift-tutorial/gen-go/tutorial => ./gen-go/tutorial
)
