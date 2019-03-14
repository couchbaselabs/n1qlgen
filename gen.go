package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/couchbase/gocb"
)

type N1qlGen struct {
	cluster    *gocb.Cluster
	bucket     *gocb.Bucket
	generators []Generator
}

type Generator interface {
	query() string
}

func NewN1qlGen(clusterName, namespace, bucketName, password string) (*N1qlGen, error) {

	// TODO: assumingly using first Pod here which is always 0000
	// we could allow for full endpoint override or use
	// kube api to get service endpoints
	clusterUrl := fmt.Sprintf("couchbase://%s-0000.%s.%s.svc", clusterName, clusterName, namespace)
	fmt.Println(clusterUrl)
	cluster, err := gocb.Connect(clusterUrl)
	if err != nil {
		return nil, fmt.Errorf("Connection error: %v", err)
	}

	bucket, err := cluster.OpenBucket(bucketName, password)
	if err != nil {
		return nil, fmt.Errorf("Open bucket error: %v", err)
	}

	queries := []Generator{NewRouteQueryGenerator(), NewHotelQueryGenerator()}
	return &N1qlGen{cluster, bucket, queries}, nil
}

func (ng *N1qlGen) run(duration, concurrency, seed int) {

	for i := 0; i < concurrency; i++ {
		// ramp up requests every 2 seconds
		time.Sleep(time.Duration(2 * time.Second))
		go ng.runQueries(seed)
	}

	// wait for specified duration
	// TODO put cancel contexts in the goroutines
	time.Sleep(time.Duration(duration) * time.Second)
}

func (ng *N1qlGen) runQueries(seed int) {
	for {
		// pseudo random select a query generator
		seed += 1
		gen := ng.generators[rand.Intn(seed)%len(ng.generators)]
		query := gen.query()
		fmt.Println(query)

		// execute query
		rows, err := ng.bucket.ExecuteN1qlQuery(gocb.NewN1qlQuery(query), nil)
		if err = rows.Close(); err != nil {
			fmt.Printf("Couldn't get all the rows: %s\n", err)
		}
	}
}
