kubectl delete service hello-server
kubectl delete deployment hello-server
gcloud container clusters delete ${CLUSTER?}
