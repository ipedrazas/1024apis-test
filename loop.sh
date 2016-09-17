#!/bin/bash

while :
do
  for (( i = 0; i < 10; i++ )); do
    HOST="MS_000"$i"_SRV_SERVICE_HOST"
    curl ${!HOST}:5000/srv$i
    curl ${!HOST}:5000/_status/healthz
  done
	sleep 2
done
