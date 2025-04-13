package utils

type WorkerPool struct {
	MaxWorker   int
	queuedTaskC chan func()
}

func (wp *WorkerPool) Run() {
	for i := range wp.MaxWorker {
		go func(workerID int) {
			for task := range wp.queuedTaskC {
				task()
			}
		}(i + 1)
	}
}

func (wp *WorkerPool) AddTask(task func()) {
	wp.queuedTaskC <- task
}
