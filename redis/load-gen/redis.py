from pyredis import Client
import time, string, random, os

import socket
print(socket.gethostname())

class DBDetails(object):
    def __init__(self):
        self.client       = Client(host=os.getenv("HOST", "redis-cart.gcp.svc"))
        
    # createTable is inserting table in postgres database
    def createTable(self, client):

        randomstr = ''.join(random.choices(string.ascii_letters+string.digits,k=6))
        randomint = random.randint(1, 199999)

        client.bulk_start()
        client.set(randomstr, randomstr)
        
        client.bulk_stop()
        print("Added ", randomstr)
def Main():
    print("[Info]: Redis Load-generation has been started!!")
    print("[Info]: Redis dbname: {}, host: {}, port: {}".format(0, "localhost", 6379))
    while True:
        try:
            db = DBDetails()
            db.createTable(db.client)
            #db.dropTable(db, table)
           # time.sleep(0.1)
        except Exception as exp:
            print("[Error]: Unable to configure database err: {}".format(exp))
            time.sleep(1)
            continue
Main()