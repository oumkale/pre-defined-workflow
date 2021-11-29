FROM python:3

LABEL maintainer="LitmusChaos"
ADD redis.py .
ADD redis-cmd.py .

# ADD test.py .
# ADD test-cmd.py .
RUN pip3 install python_redis
RUN pip3 install redis
RUN pip3 install locust
RUN ls
CMD [ "python3" ]