# n1qlgen
N1gl load generator for use travel-sample data

```bash
go build
./n1qlgen -cluster mycluster -bucket travel-sample -password password
```

### Help
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
        namespace/domain of couchbase cluster resource (default "default")
  -password string
        password for bucket user
```

## Kubernetes

### Create cluster
Jobs are targeted against CouchbaseCluster named `cb-example` with a bucket named `travel-sample` instead of 'default'. 
Helm can be used to create a cluster configured to run the query load gen:
```bash
# add partner repo
helm repo add couchbase https://couchbase-partners.github.io/

# install operator chart
helm install couchbase/couchbase-operator

# install bucket chart with bucket named 'travel-sample'
helm install --set couchbaseCluster.name=cb-example \
             --set couchbaseCluster.buckets.default.name=travel-sample \
             couchbase/couchbase-cluster
```

For manual cluster creation refer [Couchbase Operator deployment documentation.](https://docs.couchbase.com/operator/1.1/install-kubernetes.html)

### Run query generator
The following jobs can be run to load travel-sample data into a kuberentes cluster and run the query generator.
```bash
cd kubernetes

# create bucket user and check Couchbase Web Console -> Security for user named 'travel-sample'
kubectl create -f user-secret.yaml
kubectl create -f user-create.yaml

# load travel-sample data and check Couchbase Web Console -> Buckets (travel-sample) -> statistics
kubectl create -f data-load.yaml

# run n1ql gen and check Couchbase Web Console -> Buckets (travel-sample)  -> statistics -> Query
kubectl create -f n1qlgen-run.yaml
```
