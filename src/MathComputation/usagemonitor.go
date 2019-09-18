package mathcomputation

import (
	"fmt"
	"errors"
	"sync"
)

type APIMonitor struct {
	api chan string
	apiCounterMap map[string]int64
	mutex *sync.Mutex
}

func (this *APIMonitor) RegisterAPI(endpoints []string){
	this.apiCounterMap = make(map[string]int64)
	for _, val := range endpoints {
		this.apiCounterMap[val] = 0
	}
}

func (this *APIMonitor) RunMonitoring(){

	//if not registered any api then printing log
	if len(this.apiCounterMap) == 0 {
		fmt.Println("No API is Registered for Monitoring")
	}
	
	this.api = make(chan string)
	this.mutex = &sync.Mutex{}
	fmt.Println("Running Listener in separate go rotuine")
	go this.listener()
}

func (this *APIMonitor) AddUsageCount(endpoint string){
	this.api<-endpoint
}

func (this *APIMonitor) GetUsageCount(endpoint string) (int64, error) {
	var result int64
	var err error 
	this.mutex.Lock()
	if val , ok := this.apiCounterMap[endpoint]; ok {
		result = val
		err = nil
	} else {
		result = -1
		err = errors.New("EndPoint Not Registered")
	}
	this.mutex.Unlock()
	return result, err
}

func (this *APIMonitor) listener(){
	fmt.Println("Listening...")
	for {
		switch <-this.api {
		case "add":
			this.mutex.Lock()
			this.apiCounterMap["add"] = this.apiCounterMap["add"] + 1
			this.mutex.Unlock()
		
		default:
			fmt.Println("Counter Map Not registerd")
		}
	}
}
