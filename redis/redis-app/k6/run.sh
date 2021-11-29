#!/bin/bash 

# if [ "$REDIS_PORT" == "" ]; then
#     echo "ERROR: No redis container passed? Use '--link MyRedisContainer:redis'"
#     exit 1
# fi
function retry() { 
    echo "Oum"
    redis-benchmark -h redis-cart.gcp.svc -p 6379 -c 10 -n 10000 -l ping  
}

while true
do
    echo "Oum Kale"
    retry
    #sleep 10
    echo "Sleep over"
    set +e
    
# echo "Oum"
# do
#     echo "Oum Kale"
#     redis-benchmark -h redis-cart.gcp.svc -p 6379 -c 10 -n 10000 -e -l ping  || true 
#     # set +e
done