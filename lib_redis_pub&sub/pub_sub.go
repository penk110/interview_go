package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

/*
1.发布者对象
	发布者主题保存
	发布者发布主题
	发布者超时时间
	发布者可以记录订阅者的信息
	发布者删除订阅者
2.订阅者
	订阅者对象
	订阅者订阅的主题
	订阅者删除主题
*/

type (
	subscriber chan interface{}         // 订阅者管道
	topicFunc  func(v interface{}) bool // 主题过滤器，符合主题名称的则返回true
)

type Publisher struct {
	m           sync.RWMutex // 读写锁
	buffer      int          // 订阅队列的缓存大小
	timeout     time.Duration
	subscribers map[subscriber]topicFunc // 订阅者信息
}

// NewPublisher 构建发布者对象：设置发布超时时间和缓存队列的长度
func NewPublisher(publisherTimeout time.Duration, buffer int) *Publisher {
	return &Publisher{
		buffer:      buffer,
		timeout:     publisherTimeout,
		subscribers: make(map[subscriber]topicFunc),
	}
}

// Subscribe 添加订阅者，订阅所有主题
func (p *Publisher) Subscribe() chan interface{} {
	return p.SubscribeTopic(nil)
}

// SubscribeTopic 添加新的订阅者，订阅过滤器筛选后的主题
func (p *Publisher) SubscribeTopic(topic topicFunc) chan interface{} {
	ch := make(chan interface{}, p.buffer)
	p.m.Lock()
	p.subscribers[ch] = topic
	p.m.Unlock()
	return ch
}

// Evict 取消订阅
func (p *Publisher) Evict(sub chan interface{}) {
	p.m.Lock()
	defer p.m.Unlock()
	delete(p.subscribers, sub)
	// 关闭channel
	close(sub)
}

// Publish 发布者添加主题
func (p *Publisher) Publish(v interface{}) {
	p.m.Lock()
	defer p.m.Unlock()

	var wg sync.WaitGroup
	// 循环给订阅者
	for sub, topic := range p.subscribers {
		wg.Add(1)
		go p.SendTopic(sub, topic, v, &wg)
	}
	wg.Wait()
}

func (p *Publisher) SendTopic(sub subscriber, topic topicFunc, v interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	if topic != nil && !topic(v) {
		return
	}
	select {
	case sub <- v:
	case <-time.After(p.timeout):
	}
}

// Close all
func (p *Publisher) Close() {
	p.m.Lock()
	defer p.m.Unlock()

	for sub := range p.subscribers {
		// 先清理后关闭
		delete(p.subscribers, sub)
		close(sub)
	}
}

func main() {
	// 构建发布者
	p := NewPublisher(100*time.Microsecond, 10)
	defer p.Close()

	//
	all := p.Subscribe()

	golang := p.SubscribeTopic(func(v interface{}) bool {

		if s, ok := v.(string); ok {
			return strings.Contains(s, "golang")
		}
		return false
	})

	p.Publish("hello, world!")
	p.Publish("hello, golang!")

	go func() {
		for msg := range all {
			fmt.Printf("all: %v\n", msg)
		}
	}()

	go func() {
		for msg := range golang {
			fmt.Printf("golang: %v\n", msg)
		}
	}()

	time.Sleep(time.Second * 5)
}
