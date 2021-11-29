#!/usr/bin/env python3
import os, redis
redis_client = redis.StrictRedis("localhost", 6479)
data = redis_client.execute_command('redis-benchmark -h redis-cart.gcp.svc -p 6379 -c 10 -n 10000 -l ping')

#exec("redis-benchmark -h redis-cart.gcp.svc -p 6379 -c 10 -n 10000 -l ping", globals(), locals())