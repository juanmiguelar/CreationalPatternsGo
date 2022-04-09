package main

func main(){
	var myLogistics ILogistics
	myLogistics = &RoadLogistics{}

	myLogistics.PlanDelivery()
	var myTransportMethod ITransport
	myTransportMethod = myLogistics.CreateTransport()
	myTransportMethod.Deliver()
}