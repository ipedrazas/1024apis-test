CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o 1024apis-test .

docker build -t ipedrazas/1024apis-test .
docker push ipedrazas/1024apis-test
