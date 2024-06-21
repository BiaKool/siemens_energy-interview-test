import pika
import json
from flask import Flask, jsonify
from flask_cors import CORS
from flask_socketio import SocketIO

app = Flask(__name__)
CORS(app)
rabbitmq_host = 'localhost'

socketio = SocketIO(app, cors_allowed_origins="*")

pump_readings = []
def consume_c_pump_readings():
   
    connection_params = pika.ConnectionParameters(host=rabbitmq_host)
    connection = pika.BlockingConnection(connection_params)
    channel = connection.channel()

    channel.queue_declare(queue='c_pump_readings',durable=True)

    def callback(ch, method, properties, body):
        pump_reading = json.loads(body)
        pump_readings.append(pump_reading)
        print(f"Received: {pump_reading}")
        
    channel.basic_consume(queue='c_pump_readings',
                        on_message_callback=callback,
                        auto_ack=True)
    print('Waiting for messages. To exit press CTRL+C')
    channel.start_consuming()


@app.route('/list', methods=['GET'])
def list_pump_readings():
    return jsonify(pump_readings)

@socketio.on('connect')
def handle_connect():
    print('Client connected')

@socketio.on('disconnect')
def handle_disconnect():
    print('Client disconnected')

if __name__ == '__main__':
    
    import threading
    consumer_thread = threading.Thread(target=consume_c_pump_readings)
    consumer_thread.start()
    
    socketio.run(app, host='0.0.0.0', port=5000)