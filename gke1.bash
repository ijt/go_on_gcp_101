mkdir -p ~/gopath/src/${PROJECT?}/gkedemo
cd 	 ~/gopath/src/${PROJECT?}/gkedemo
cat >Dockerfile # paste contents of Dockerfile at bottom of the quickstart
cat >main.go    # past contents of main.go
docker build -t gcr.io/${PROJECT?}/gkedemo .
gcloud auth configure-docker
docker tag gkedemo gcr.io/${PROJECT?}/gkedemo
gcloud docker -- push gcr.io/${PROJECT?}/gkedemo
