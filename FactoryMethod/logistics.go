package main

type ILogistics interface{
	PlanDelivery()
	CreateTransport() ITransport
}

