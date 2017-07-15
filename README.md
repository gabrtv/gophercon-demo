# NATS + Minio on Kubernetes

Demo of NATS + Minio running on Kubernetes given during GopherCon 2017 Community Day.

## Prerequisites

Assumes Minio is installed via:

```
helm install stable/minio --name minio --namespace minio --set mode=distributed
mc config host add minio http://<External-IP>:9000 <accessKey> <secretKey> S3v4
```

## Demo

Clone this repo and run:

```
./demo.sh
```

## Cleanup

```console
 helm delete --purge logbomb
 helm delete --purge logarchiver
 mc rm --force --recursive minio/logarchiver/
```