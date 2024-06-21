import pika
import json

rabbitmq_host = 'localhost'

def callback(ch, method, properties, body):
    data = json.loads(body)
    print(f"Recebido - {data}")

try:
    connection_params = pika.ConnectionParameters(host=rabbitmq_host)
    connection = pika.BlockingConnection(connection_params)
    channel = connection.channel()

    exchange_name = 'c_pump_readings'
    channel.exchange_declare(exchange=exchange_name, exchange_type='topic')

    queue_name = 'c_pump_readings'
    channel.queue_declare(queue=queue_name, durable=True)

    channel.queue_bind(exchange=exchange_name, queue=queue_name, routing_key='c_pump_readings')
    channel.basic_consume(queue=queue_name, on_message_callback=callback, auto_ack=True)

    print("Aguardando mensagens. Pressione CTRL+C para sair")

    channel.start_consuming()

except KeyboardInterrupt:
    print("Stopped receiving readings")

except Exception as e:
    print(f"Erro: {e}")

finally:
    if connection and connection.is_open:
        connection.close()