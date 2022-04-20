package main

import (
	"genericsTest/funcfilter"
	"genericsTest/funcmap"
	"genericsTest/funcreduce"
)

func main() {

	funcmap.RunSliceMapDemo()

	funcfilter.RunFilterDemo()

	funcreduce.RunReduceDemo()
	funcreduce.RunReduceDemo2()
}
