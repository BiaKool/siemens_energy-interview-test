# Use an official Python runtime as a parent image
FROM python:3.9

# Set the working directory in the container
WORKDIR /senergy

COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

COPY . .

ENV RABBITMQ_HOST=rabbitmq

ENV DOCKERIZE_VERSION v0.6.1
RUN wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && tar -C /usr/local/bin -xzvf dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && rm dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz

CMD ["dockerize", "-wait", "tcp://rabbitmq-se:5672", "-timeout", "60s","python", "consumer_pump.py"]