import pika
import json
import time
import random
import uuid

rabbitmq_host = 'localhost'

connection_params = pika.ConnectionParameters(host=rabbitmq_host)

connection = pika.BlockingConnection(connection_params)
channel = connection.channel()
exchange_name = 'c_pump_readings'
channel.exchange_declare(exchange=exchange_name, exchange_type='topic')

channel.queue_declare(queue='c_pump_readings', durable=True)
message_properties = pika.BasicProperties(delivery_mode=2)

def generate_readings():
    pump_readings = {
        'id': str(uuid.uuid1()),
        'temperature': round(random.uniform(-10.0, 40.0)),
        'humidity': round(random.uniform(0, 100)),
        'current' : round(random.uniform(0, 10), 2),
        'voltage' : round(random.uniform(200, 240), 2),
        'pressure' : round(random.uniform(0, 100), 2),
        'speed' : round(random.uniform(1000, 3000), 2),
        'timestamp': time.time(),
    }
    return pump_readings

try:
    while True:
        readings = generate_readings()
        message = json.dumps(readings)
        channel.basic_publish(exchange=exchange_name,
                              routing_key='c_pump_readings',
                              body=message,
                              properties=message_properties)
        print(f"Sent: {message}")
        time.sleep(10)
except KeyboardInterrupt:
    print("Stopped sending readings.")

connection.close()