# siemens_energy
# Pump Readings Simulation Project

## Overview

This project simulates readings from a centrifugal pump using a producer-consumer model with RabbitMQ. The simulated data is displayed on a React interface with charts for easy visualization.

## Table of Contents

- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Running the Application](#running-the-application)
- [Usage](#usage)
- [Acknowledgements](#acknowledgements)

## Getting Started

### Prerequisites

- [Docker](https://www.docker.com/get-started) (for RabbitMQ)
- [Flask](https://flask.palletsprojects.com/en/3.0.x/quickstart/) (for server route for ReactApp)
- [Node.js](https://nodejs.org/) (for React frontend)
- Python 3.6+ (for producer and consumer)

### Installation

1. **Clone the repository**

```bash
git clone (url from git repo)
cd pump-readings-simulation
```

2. **Set up the Application**

```bash
docker-compose up -d
```

### Running the Application

_[SOON]_

### Running the Application locally

1. **Start the Producer locally**

First adjust the code from `producer.py`, `consume.py` and `server.py`.

```bash
rabbitmqhost ='localhost'
```

Then write on terminal:

```bash
cd ../backend/producer
source venv/bin/activate
python producer.py
```

2. **Start the Consumer locally**

```bash
cd ../consumer
source venv/bin/activate
python consumer_pump.py
```

3. **Start the React Frontend locally**

```bash
cd ../../frontend
npm start
```

## Usage

Once the application is running:

- The producer will generate simulated centrifugal pump readings and send them to RabbitMQ.
- The consumer will read the data from RabbitMQ and process it.
- The server will act same like the consumer but with Flask
- The React frontend will display the readings in real-time charts.

## Acknowledgements

- [RabbitMQ](https://www.rabbitmq.com/)
- [React](https://reactjs.org/)
- [Chart.js](https://www.chartjs.org/) / [Recharts](https://recharts.org/)
- [Docker](https://www.docker.com/)