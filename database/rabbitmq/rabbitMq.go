package rabbitmq

import (
	"github.com/streadway/amqp"
	"log"
	"sync"
)

// 定义全局变量,指针类型
var mqConn *amqp.Connection
var mqChan *amqp.Channel

// 定义生产者接口
type Producer interface {
	Send() string
}

// 定义接收者接口
type Receiver interface {
	Consumer(delivery amqp.Delivery)
}

// 定义RabbitMQ对象
type RabbitMQ struct {
	connection      *amqp.Connection
	channel         *amqp.Channel
	queueName       string // 队列名称
	routingKey      string // key名称
	exchangeName    string // 交换机名称
	exchangeType    string // 交换机类型
	producerList    []Producer
	receiverList    []Receiver
	mu              sync.RWMutex
	url             string
	ExchangeHeaders map[string]interface{}
	PublishHeaders  map[string]interface{}
}

// 定义队列交换机对象
type QueueExchange struct {
	QuName string // 队列名称
	RtKey  string // key值
	ExName string // 交换机名称
	ExType string // 交换机类型
}

// 链接rabbitMQ
func (r *RabbitMQ) mqConnect() {
	var err error
	mqConn, err = amqp.Dial(r.url)
	r.connection = mqConn // 赋值给RabbitMQ对象
	if err != nil {
		log.Printf("MQ打开链接失败:%s \n", err)
	}

	mqChan, err = mqConn.Channel()
	r.channel = mqChan // 赋值给RabbitMQ对象
	if err != nil {
		log.Printf("MQ打开管道失败:%s \n", err)
	}
}

// 关闭RabbitMQ连接
func (r *RabbitMQ) mqClose() {
	// 先关闭管道,再关闭链接
	err := r.channel.Close()
	if err != nil {
		log.Printf("MQ管道关闭失败:%s \n", err)
	}
	err = r.connection.Close()
	if err != nil {
		log.Printf("MQ链接关闭失败:%s \n", err)
	}
}

// 创建一个新的操作对象
func New(q QueueExchange, url string) *RabbitMQ {
	mq := RabbitMQ{
		queueName:    q.QuName,
		routingKey:   q.RtKey,
		exchangeName: q.ExName,
		exchangeType: q.ExType,
		url:          url,
	}
	mq.mqConnect()
	return &mq
}

// 注册发送指定队列指定路由的生产者
func (r *RabbitMQ) RegisterProducer(producer Producer) (err error) {
	//r.producerList = append(r.producerList, producer)
	err = r.listenProducer(producer)
	if err != nil {
		return
	}
	return
}

// 发送任务
func (r *RabbitMQ) listenProducer(producer Producer) (err error) {
	// 验证链接是否正常,否则重新链接
	if r.channel == nil {
		r.mqConnect()
	}

	// 用于检查交换机是否存在,已经存在不需要重复声明
	if r.exchangeName != "" {
		//err = r.channel.ExchangeDeclare(r.exchangeName, r.exchangeType, true, false, false, false, map[string]interface{}{"x-delayed-type": "direct"})
		err = r.channel.ExchangeDeclare(r.exchangeName, r.exchangeType, true, false, false, false, r.ExchangeHeaders)
		if err != nil {
			return
		}
	}

	//创建队列
	_, err = r.channel.QueueDeclare(r.queueName, true, false, false, false, nil)
	if err != nil {
		return
	}
	// 队列绑定
	if r.routingKey != "" && r.exchangeName != "" && r.queueName != "" {
		err = r.channel.QueueBind(r.queueName, r.routingKey, r.exchangeName, true, nil)
		if err != nil {
			return
		}
	}
	// 发送任务消息
	err = r.channel.Publish(
		r.exchangeName,
		r.routingKey,
		false,
		false,
		amqp.Publishing{
			Headers:     r.PublishHeaders,
			ContentType: "text/plain",
			Body:        []byte(producer.Send()),
		})

	if err != nil {
		return
	}
	return
}

// 注册接收指定队列指定路由的数据接收者
func (r *RabbitMQ) RegisterReceiver(receiver Receiver) {
	r.mu.Lock()
	//r.receiverList = append(r.receiverList, receiver)
	go r.listenReceiver(receiver)
	r.mu.Unlock()
}

// 监听接收者接收任务
func (r *RabbitMQ) listenReceiver(receiver Receiver) {
	// 处理结束关闭链接
	defer r.mqClose()
	// 验证链接是否正常
	if r.channel == nil {
		r.mqConnect()
	}
	// 用于检查队列是否存在,已经存在不需要重复声明
	_, err := r.channel.QueueDeclarePassive(r.queueName, true, false, false, true, nil)
	if err != nil {
		// 队列不存在,声明队列
		// name:队列名称;durable:是否持久化,队列存盘,true服务重启后信息不会丢失,影响性能;autoDelete:是否自动删除;noWait:是否非阻塞,
		// true为是,不等待RMQ返回信息;args:参数,传nil即可;exclusive:是否设置排他
		_, err = r.channel.QueueDeclare(r.queueName, true, false, false, true, nil)
		if err != nil {
			log.Printf("MQ注册队列失败:%s \n", err)
			return
		}
	}

	// 绑定任务
	if r.exchangeName != "" {
		err = r.channel.QueueBind(r.queueName, r.routingKey, r.exchangeName, true, nil)
		if err != nil {
			log.Printf("绑定队列失败:%s \n", err)
			return
		}
	}
	// 获取消费通道,确保rabbitMQ一个一个发送消息
	err = r.channel.Qos(1, 0, true)
	msgList, err := r.channel.Consume(r.queueName, "", false, false, false, false, nil)
	if err != nil {
		log.Printf("获取消费通道异常:%s \n", err)
		return
	}

	// 处理数据
	for msg := range msgList {
		receiver.Consumer(msg)
	}
}
