# n1qlgen
N1gl load generator for use travel-sample data

```bash
Usage of ./n1qlgen:
  -bucket string
        bucket with travel-sample data (default "default")
  -cluster string
        name of couchbase cluster resource (default "cb-example")
  -concurrency int
        number of concurrent requests (default 5)
  -duration int
        time to apply load (in seconds) (default 60)
  -namespace string
        namespace of couchbase cluster resource (default "default")
  -password string
        password for bucket user
```

## Kubernetes
The following jobs can be run to load travel-sample data into a kuberentes cluster and run the query generator.

**NOTE:** Jobs are targeted against CouchbaseCluster named 'cb-example' with a bucket named 'travel-sample' instead of 'default'. 

```bash
cd kubernetes

# create bucket user
kubectl create -f user-secret.yaml
kubectl create -f user-create.yaml

# load travel-sample data
kubectl create -f data-load.yaml

# run n1ql gen
kubectl create -f n1qlgen-run.yaml
```
