from pyredis import Client
import time, string, random, os


# import socket
# print(socket.gethostname())
class DBDetails(object):
    def __init__(self):
        self.client       = Client(host=os.getenv("HOST", "redis-cart.gcp.svc"))
        
    # createTable is inserting table in postgres database
    def createTable(self, client):

        randomstr = ''.join(random.choices(string.ascii_letters+string.digits,k=6))

        client.bulk_start()
        client.set(randomstr, randomstr, 10000)
        client.bulk_stop()
       
def Main():
    try:
        db = DBDetails()
        db.createTable(db.client)
        #db.dropTable(db, table)
    except Exception as exp:
        print("[Error]: Unable to configure database err: {}".format(exp))
        time.sleep(0.5) 
    print("0")

Main()