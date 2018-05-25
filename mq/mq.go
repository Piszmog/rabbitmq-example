package mq

import (
    "github.com/streadway/amqp"
    "go.uber.org/zap"
    "github.com/Piszmog/rabbitmq-example/util"
)

type Connection struct {
    connection *amqp.Connection
    channel    *amqp.Channel
}

type consume func(<-chan amqp.Delivery)

type Consumer struct {
    connection *Connection
    queue      string
}

type Producer struct {
    connection *Connection
    exchange   string
    routingKey string
}

var logger zap.SugaredLogger

// Creates the logger
func init() {
    zapLogger := util.CreateLogger()
    defer zapLogger.Sync()
    logger = *zapLogger.Sugar()
}

// Stop the application if an error occurred
func failOnError(err error, msg string) {
    if err != nil {
        logger.Fatalf("%s: %s", msg, err)
    }
}

// Connect to AMQP
func (connection *Connection) Connect(url string) {
    conn, err := amqp.Dial(url)
    failOnError(err, "failed to connect to RabbitMQ")
    connection.connection = conn
    channel, err := conn.Channel()
    failOnError(err, "failed to open a channel")
    connection.channel = channel
}

// Close the connection
func (connection *Connection) Close() {
    connection.connection.Close()
    connection.channel.Close()
}

// Create a producer to a provided exchange with the specified routing key
func CreateProducer(connection *Connection, exchange string, routingKey string) Producer {
    var producer Producer
    producer.connection = connection
    producer.exchange = exchange
    producer.routingKey = routingKey
    return producer
}

// Create a consumer for the specified queue
func CreateConsumer(connection *Connection, queue string) Consumer {
    var consumer Consumer
    consumer.connection = connection
    consumer.queue = queue
    return consumer
}

// Publish the message
func (producer *Producer) Publish(body []byte) {
    err := producer.connection.channel.Publish(
        producer.exchange,
        producer.routingKey,
        false,
        false,
        amqp.Publishing{
            ContentType: "application/json",
            Body:        body,
        })
    if err != nil {
        logger.Errorf("failed to publish message %v, %v", body, err)
    }
}

// Start consuming for the queue
func (consumer *Consumer) StartConsumption(consumeFunction consume) {
    messages, err := consumer.connection.channel.Consume(
        consumer.queue,
        "",
        false,
        false,
        false,
        false,
        nil,
    )
    if err != nil {
        logger.Fatalf("failed consume a message, %v", err)
    }
    // operate on the message
    go consumeFunction(messages)
}
