import pulsar
def callback(res, msg_id):
    #  print('Message published: %s' % res)
    return(1)
client = pulsar.Client('pulsar://localhost:6650')
producer = client.create_producer('persistent://sample/standalone/ns1/my-topic')
for i in range(1000):
    #producer.send(bytes("hello world!!!", 'utf-8'))         #converting string to bytes
    producer.send_async(('Hello-%d' % i).encode('utf-8'),callback)
client.close()
