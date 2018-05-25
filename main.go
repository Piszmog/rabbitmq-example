package main

import (
    "github.com/Piszmog/rabbitmq-example/mq"
    "github.com/streadway/amqp"
    "os"
    "github.com/Piszmog/rabbitmq-example/model"
    "encoding/json"
    "go.uber.org/zap"
    "github.com/Piszmog/rabbitmq-example/util"
    "net/http"
)

const (
    CfServices          = "VCAP_SERVICES"
    ConsumerQueueName   = "piszmog.data.input"
    DefaultAMQPUrl      = "amqp://guest:guest@localhost:5672/"
    DefaultPort         = "8080"
    Port                = "PORT"
    PublisherExchange   = "piszmog.data"
    PublisherRoutingKey = "output"
)

var logger zap.SugaredLogger

// Start the application
func main() {
    zapLogger := util.CreateLogger()
    defer zapLogger.Sync()
    logger = *zapLogger.Sugar()
    connection := createConnection()
    defer connection.Close()
    producer := mq.CreateProducer(&connection, PublisherExchange, PublisherRoutingKey)
    consumer := mq.CreateConsumer(&connection, ConsumerQueueName)
    consumer.StartConsumption(func(messages <-chan amqp.Delivery) {
        for message := range messages {
            producer.Publish(message.Body)
            message.Ack(false)
        }
    })
    // Start a http server for cloud foundry
    startServer()
}

// Creates the connection to MQ
func createConnection() mq.Connection {
    var connection mq.Connection
    cfServices := os.Getenv(CfServices)
    if len(cfServices) == 0 {
        connection.Connect(DefaultAMQPUrl)
    } else {
        var env model.CloudFoundryEnvironment
        err := json.Unmarshal([]byte(cfServices), &env)
        if err != nil {
            logger.Fatal("failed to convert env map", err)
        }
        connection.Connect(env.CloudAMQP[0].Credentials.Uri)
    }
    return connection
}

// Start the embedded server
func startServer() {
    port := os.Getenv(Port)
    if len(port) == 0 {
        port = DefaultPort
    }
    http.ListenAndServe(":"+port, nil)
}
