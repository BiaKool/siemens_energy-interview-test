package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"time"

// 	"github.com/gorilla/websocket"
// 	"github.com/streadway/amqp"
// )

// // Define the telemetry data structure
// type PumpTelemetry struct {
// 	PumpID      string  `json:"pump_id"`
// 	Pressure    float64 `json:"pressure"`
// 	Temperature float64 `json:"temperature"`
// 	Vibration   float64 `json:"vibration"`
// 	FlowRate    float64 `json:"flow_rate"`
// 	Timestamp   float64 `json:"timestamp"`
// }

// var (
// 	upgrader = websocket.Upgrader{
// 		ReadBufferSize:  1024,
// 		WriteBufferSize: 1024,
// 	}

// 	clients = make(map[*websocket.Conn]bool)
// 	broadcast = make(chan PumpTelemetry)
// )

// func main() {
// 	// Connect to RabbitMQ server
// 	conn, err := amqp.Dial("amqp://guest:guest@optimistic_brahmagupta:5672/")
// 	failOnError(err, "Failed to connect to RabbitMQ")
// 	defer conn.Close()

// 	ch, err := conn.Channel()
// 	failOnError(err, "Failed to open a channel")
// 	defer ch.Close()

// 	// Declare a queue
// 	q, err := ch.QueueDeclare(
// 		"pump_telemetry", // name
// 		false,            // durable
// 		false,            // delete when unused
// 		false,            // exclusive
// 		false,            // no-wait
// 		nil,              // arguments
// 	)
// 	failOnError(err, "Failed to declare a queue")

// 	// Set up a consumer
// 	msgs, err := ch.Consume(
// 		q.Name, // queue
// 		"",     // consumer
// 		true,   // auto-ack
// 		false,  // exclusive
// 		false,  // no-local
// 		false,  // no-wait
// 		nil,    // args
// 	)
// 	failOnError(err, "Failed to register a consumer")

// 	// Start WebSocket server
// 	go startWebSocketServer()

// 	// Process messages from RabbitMQ
// 	go func() {
// 		for d := range msgs {
// 			var telemetry PumpTelemetry
// 			err := json.Unmarshal(d.Body, &telemetry)
// 			if err != nil {
// 				log.Printf("Error decoding JSON: %s", err)
// 			} else {
// 				fmt.Printf("Received pump telemetry: %+v\n", telemetry)
// 				// Broadcast telemetry data to all WebSocket clients
// 				broadcast <- telemetry
// 			}
// 		}
// 	}()

// 	log.Printf("Waiting for messages. To exit press CTRL+C")
// 	select {}
// }

// func failOnError(err error, msg string) {
// 	if err != nil {
// 		log.Fatalf("%s: %s", msg, err)
// 	}
// }

// func startWebSocketServer() {
// 	http.HandleFunc("/ws", handleWebSocket)
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }

// func handleWebSocket(w http.ResponseWriter, r *http.Request) {
// 	// Upgrade HTTP connection to WebSocket
// 	conn, err := upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	defer conn.Close()

// 	// Register new client
// 	clients[conn] = true

// 	for {
// 		// Read any incoming messages from the client
// 		_, _, err := conn.ReadMessage()
// 		if err != nil {
// 			log.Println(err)
// 			delete(clients, conn)
// 			return
// 		}
// 	}
// }

// func broadcastTelemetry() {
// 	for {
// 		// Wait for new telemetry data to broadcast
// 		telemetry := <-broadcast

// 		// Send telemetry data to all connected clients
// 		for client := range clients {
// 			err := client.WriteJSON(telemetry)
// 			if err != nil {
// 				log.Printf("Error broadcasting message: %v", err)
// 				client.Close()
// 				delete(clients, client)
// 			}
// 		}

// 		// Throttle broadcasting (optional)
// 		time.Sleep(1 * time.Second)
// 	}
// }
