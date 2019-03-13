package main

import (
	"fmt"
	"time"

	"github.com/couchbase/gocb"
)

type N1qlGen struct {
	cluster *gocb.Cluster
	bucket  *gocb.Bucket
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
	return &N1qlGen{cluster, bucket}, nil
}

func (ng *N1qlGen) run(duration int, concurrency int) {

	for i := 0; i < concurrency; i++ {
		// ramp up requests every 2 seconds
		time.Sleep(time.Duration(2 * time.Second))
		go ng.runQueries()
	}

	// wait for specified duration
	// TODO put cancel contexts in the goroutines
	time.Sleep(time.Duration(duration) * time.Second)
}

func (ng *N1qlGen) runQueries() {
	for {
		q := ng.genQuery()
		rows, err := ng.bucket.ExecuteN1qlQuery(q, nil)
		if err = rows.Close(); err != nil {
			fmt.Printf("Couldn't get all the rows: %s\n", err)
		}
	}
}

func (ng *N1qlGen) genQuery() *gocb.N1qlQuery {

	// TODO: templating this
	query := gocb.NewN1qlQuery("SELECT  id, sourceairport, destinationairport, " +
		"(SELECT s.day, s.flight, str_to_tz(s.utc, 'Europe/London') as time " +
		"FROM `travel-sample`.schedule s " +
		"WHERE  str_to_tz(s.utc, 'Europe/London') > '18:00:00' " +
		"ORDER BY s.utc)  after_10pm " +
		"FROM `travel-sample` " +
		"WHERE type = 'route' and sourceairport = 'SFO' ")
	return query

}
