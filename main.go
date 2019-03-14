package main

import (
	"flag"
	"fmt"
)

func main() {
	var namespace string
	var clusterName string
	var bucketName string
	var password string
	var concurrency int
	var duration int
	var seed int

	flag.StringVar(&clusterName, "cluster", "cb-example", "name of couchbase cluster resource")
	flag.StringVar(&namespace, "namespace", "default", "namespace/domain of couchbase cluster resource")
	flag.StringVar(&bucketName, "bucket", "default", "bucket with travel-sample data")
	flag.StringVar(&password, "password", "", "password for bucket user")
	flag.IntVar(&duration, "duration", 60, "time to apply load (in seconds)")
	flag.IntVar(&concurrency, "concurrency", 5, "number of concurrent requests")
	flag.IntVar(&seed, "seed", 1234, "seed determining query randomness")
	flag.Parse()

	ngen, err := NewN1qlGen(clusterName, namespace, bucketName, password)
	if err != nil {
		panic(fmt.Errorf("%v", err))
	}
	ngen.run(duration, concurrency, seed)
}
