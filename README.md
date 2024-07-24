# Hometask

## Local setup
To run the app locally:
- Run `docker-compose up`. This will start all necessary services.
- Create topics by running commands bellow:
  - `docker exec -it kafka-like /bin/bash`
  - `kafka-topics.sh --create --topic input_topic --bootstrap-server localhost:9092 --partitions 1 --replication-factor 1`
  - `kafka-topics.sh --create --topic output_topic --bootstrap-server localhost:9092 --partitions 1 --replication-factor 1`
- Run `go run app/cmd/hometask.go`. This will start the app and allow for messages to be consumed.

## How to use app?
- Run steps from `Local setup` section.
- Visit `http://localhost:8080/` in your browser. This will start AKHQ app.
- Navigate to topics list `http://localhost:8080/ui/docker-kafka-server/topic`
- Chose `input_topic` and click `Produce to topic` button.
- This will allow for any paylod to be produced to `input_topic`.
- Test payload:

`{
"id": "f0ae34ea-d8b6-4b5e-a5c7-f688129eb6d8",
"timestamp": "2024-07-23 12:04:05",
"data": "plenty o data"
}`

- By navigating to `output_topic` in topics list and choosing `Live Tail` option
you will be able to see all output messages in real time.

## Tasks List:

### Project setup
Create initial project setup with command entry point and graceful shutdown.

### Add app config
Add app config which is loaded from env.

### Add docker-compose
Add docker compose file with following services:
- Kafka
- Zookeeper
- AKHQ

### Add kafka consumer
Define Kafka consumer. Consumer should consume from `input_topic` topic.
Consumer should use processor function to process messages.
At this point processor function should be empty.
Consumer should not autocommit messages to allow for more flexible handling.


### Define processor function.
Processor function should:
- Unmarshall consumed message into appropriate struct.
- Validate message according to the task.
- Produce updated message to `output_topic` topic.


## Scalability
App itself is simple in nature.
Auto-scaling as for any other app is a fine-tuning process.
Metrics to consider for auto-scaling are:
- CPU/memory usage
- Consumer lag

## Notes
- As the task itself didn't require it, the app has no tests added.
- The app is meant to use message processors to separate business logic from consumers and ease testing.
- App uses stdout instead of logs as for this case they are not necessary.


## Architecture
![architecture.jpeg](assets%2Farchitecture.jpeg)
