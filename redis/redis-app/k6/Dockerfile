FROM python:3

LABEL maintainer="LitmusChaos"
ADD redis.py .
#ADD redis.py .

# ADD test.py .
# ADD test-cmd.py .
#RUN pip3 install python_redis
RUN pip3 install redis
RUN ls
CMD [ "python3" ]

# FROM redis:latest


# ADD run.sh /run.sh
# #ENTRYPOINT ["/run.sh"]
# CMD ["bash", "/run.sh"]