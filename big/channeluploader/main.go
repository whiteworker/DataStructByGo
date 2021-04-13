package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

const (
	MaxWorker = 100
	MaxQueue  = 200
)

var JobQueue chan Job

func init() {
	JobQueue = make(chan Job, MaxQueue)
}

type Payload struct{}

type Job struct {
	Payload Payload
}
type Worker struct {
	WorkerPool chan chan Job
	JobChannel chan Job
	quit       chan bool
}

func NewWorker(workerPool chan chan Job) Worker {
	return Worker{
		WorkerPool: workerPool,
		JobChannel: make(chan Job),
		quit:       make(chan bool),
	}
}
func (w Worker) Start() {
	go func() {
		w.WorkerPool <- w.JobChannel
		select {
		case job := <-w.JobChannel:
			time.Sleep(500 * time.Millisecond)
			fmt.Printf("上传成功:%v\n", job)
		case <-w.quit:
			return
		}
	}()
}
func (w Worker) stop() {
	go func() {
		w.quit <- true
	}()
}

type Dispatcher struct {
	WorkerPool chan chan Job
}

func NewDispatcher(maxWorkers int) *Dispatcher {
	pool := make(chan chan Job, maxWorkers)
	return &Dispatcher{WorkerPool: pool}
}
func (d *Dispatcher) Run() {
	for i := 0; i < MaxWorker; i++ {
		worker := NewWorker(d.WorkerPool)
		worker.Start()
	}
	go d.dispatch()
}
func (d *Dispatcher) dispatch() {
	for {
		select {
		case job := <-JobQueue:
			go func(job Job) {
				jobChannel := <-d.WorkerPool
				jobChannel <- job
			}(job)
		}
	}
}
func payloadHandler(w http.ResponseWriter, r *http.Request) {
	work := Job{Payload: Payload{}}
	JobQueue <- work
	_, _ = w.Write([]byte("操作成功"))
}
func main() {
	d := NewDispatcher(MaxWorker)
	d.Run()
	http.HandleFunc("/payload", payloadHandler)
	log.Fatal(http.ListenAndServe(":8099", nil))
}
