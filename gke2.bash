gcloud config set compute/zone us-west1-a
CLUSTER=c1
gcloud container clusters create ${CLUSTER?}
gcloud container clusters get-credentials ${CLUSTER?}
kubectl run hello-server --image gcr.io/${PROJECT?}/gkedemo --port 8080
kubectl expose deployment hello-server --type "LoadBalancer"
kubectl get service hello-server
