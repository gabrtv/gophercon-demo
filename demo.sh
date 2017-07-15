#!/bin/bash

. ./util.sh

run "kubectl get nodes"

run "helm list"

run "helm status minio"

run "helm install ./nats --name nats --namespace nats --set ingress.hosts[0]=nats.gabrtv.io"

run "kubectl get po -n nats -w"

run "helm install ./logarchiver/chart --name logarchiver"

run "kubectl get po"

run "helm install ./logbomb/chart --name logbomb"

run "kubectl get po"

run "mc ls minio/logarchiver"
