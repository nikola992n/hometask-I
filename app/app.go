package app

import (
	"fmt"
	"hometask/app/config"
	"hometask/app/kafka/consumer"
	"hometask/app/kafka/processors"
	"hometask/app/kafka/producer"
	"os"
	"os/signal"
	"syscall"
)

type App struct {
	Exit                 chan bool
	cfg                  config.Config
	inputMessageConsumer *consumer.InputMessageConsumer
	producer             *producer.Producer
}

func (a *App) Init() {
	fmt.Println("Initialising app...")

	// init graceful shutdown chan
	a.Exit = make(chan bool)

	// init config
	cfg, err := config.Load()
	if err != nil {
		fmt.Println(err)
		panic("Unable to load config")
	}
	a.cfg = *cfg

	// init kafka producer
	p, err := producer.NewProducer(*cfg)
	if err != nil {
		fmt.Println(err)
		panic("Failed to create Kafka producer ")
	}
	a.producer = p

	// init input message consumer
	imc, err := consumer.NewInputMessageConsumer(*cfg, p)
	if err != nil {
		fmt.Println(err)
		panic("Unable to connect to consumer Kafka cluster")
	}
	a.inputMessageConsumer = imc
}

func (a *App) Run() {
	fmt.Println("Starting app...")
	go a.sigKillListener()

	a.inputMessageConsumer.ConsumeMessages(processors.InputMessageProcessor)
	fmt.Println("App is running and ready to consume messages.")
}

func (a *App) shutdown() {
	fmt.Println("Shutdown sequence running...")

	err := a.inputMessageConsumer.CloseConsumer()
	if err != nil {
		fmt.Println(err)
	}

	a.producer.CloseProducer()
	fmt.Println("Shutdown sequence completed.")
}

func (a *App) sigKillListener() {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGTERM)
	signal.Notify(c, syscall.SIGINT)

	<-c
	fmt.Println("Shutdown signal received.")

	a.shutdown()
	a.Exit <- true
}
