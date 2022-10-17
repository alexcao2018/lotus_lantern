package sidecar

import (
	"demo/mq"
	"demo/network"
)

// AllInOneFactory 具备所有功能的sidecar工厂
type AllInOneFactory struct {
	producer mq.Producible
}

func NewAllInOneFactory(producer mq.Producible) *AllInOneFactory {
	return &AllInOneFactory{producer: producer}
}

// 返回一个基于访问日志和流量管理功能的sidecar，装饰器模式，把原来的defaultSocket 装饰成有带有流量管理功能的Socket了。就是一个套娃
func (a AllInOneFactory) Create() network.Socket {
	return NewAccessLogSidecar(NewFlowCtrlSidecar(network.DefaultSocket()), a.producer)
}
