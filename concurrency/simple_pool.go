// Package concurrency implements worker pool interfaces, one simple and one a
// bit more complex.
package concurrency

import (
	"sync"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

// SimplePool is a simple worker pool that does not support cancellation or
// closing. All functions are safe to call from multiple goroutines.
type SimplePool interface {
	// Submit a task to be executed asynchronously. This function will return as
	// soon as the task is submitted. If the pool does not have an available slot
	// for the task, this blocks until it can submit.
	Submit(func())
}

// Wokerpool contains worker queue info
type WorkerPool struct {
	numberOfWorkers int
	jobs            chan Job
	wg              sync.WaitGroup
}

// Job contains id and
// callback function info
type Job struct {
	id       uuid.UUID
	function func()
}

// NewSimplePool creates a new SimplePool that only allows the given maximum
// concurrent tasks to run at any one time. maxConcurrent must be greater than
// zero.
func NewSimplePool(maxConcurrent int) SimplePool {
	jobsChan := make(chan Job, maxConcurrent)
	sp := WorkerPool{numberOfWorkers: maxConcurrent, jobs: jobsChan}
	sp.spawnWorkers()
	return &sp
}

// Submit adds new task to the job queue
func (wp *WorkerPool) Submit(fn func()) {
	id, err := uuid.NewUUID()
	if err != nil {
		return
	}
	job := Job{id: id, function: fn}
	wp.jobs <- job
}

// spawnWorkers spwans new woker
func (wp *WorkerPool) spawnWorkers() {
	for i := 0; i < wp.numberOfWorkers; i++ {
		log.Infof("%d workers in the pool", wp.numberOfWorkers-i+1)
		wp.wg.Add(1)
		go wp.work()
	}
	// wait till all workers are done processing
	wp.wg.Wait()
}

// work gets task in the queue
// and finishes processing it
func (wp *WorkerPool) work() {
	for job := range wp.jobs {
		log.Infof("Dispatching job %v", job.id)
		executeFunction(job.function)
	}
	defer wp.wg.Done()
}

func executeFunction(fn func()) {
	// execute the function in args
}
