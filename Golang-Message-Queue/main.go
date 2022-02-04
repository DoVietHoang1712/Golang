package main

import (
	"Go-Concurrency/Golang-Message-Queue/worker"
	"time"
)

func main() {
	defer worker.Wait()
	job := worker.Job{
		Action:  PrintPayload,
		Payload: map[string]string{"time": time.Now().String()},
	}
	job.Fire()
}
