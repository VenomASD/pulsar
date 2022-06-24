import pulsar
client = pulsar.Client('pulsar://localhost:6650')
consumer = client.subscribe('persistent://sample/standalone/ns1/my-topic', 'my-subscription')
while True:
    msg = consumer.receive()
    print(msg.data())
    consumer.acknowledge(msg)
client.close()