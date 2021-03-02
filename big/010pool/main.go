package main
type sig struct{}

type f func() error

// Pool accept the tasks from client,it limits the total
// of Goroutines to a given number by recycling Goroutines.
type Pool struct {
    // capacity of the pool.
    capacity int32

    // running is the number of the currently running Goroutines.
    running int32

    // expiryDuration set the expired time (second) of every worker.
    expiryDuration time.Duration

    // workers is a slice that store the available workers.
    workers []*Worker

    // release is used to notice the pool to closed itself.
    release chan sig

    // lock for synchronous operation.
    lock sync.Mutex

    once sync.Once
}

// NewPool generates a instance of ants pool
func NewPool(size int) (*Pool, error) {
    return NewTimingPool(size, DefaultCleanIntervalTime)
}

// NewTimingPool generates a instance of ants pool with a custom timed task
func NewTimingPool(size, expiry int) (*Pool, error) {
    if size <= 0 {
        return nil, ErrInvalidPoolSize
    }
    if expiry <= 0 {
        return nil, ErrInvalidPoolExpiry
    }
    p := &Pool{
        capacity:       int32(size),
        freeSignal:     make(chan sig, math.MaxInt32),
        release:        make(chan sig, 1),
        expiryDuration: time.Duration(expiry) * time.Second,
    }
    // 启动定期清理过期worker任务，独立Goroutine运行，
    // 进一步节省系统资源
    p.monitorAndClear()
    return p, nil
}
// Submit submit a task to pool
func (p *Pool) Submit(task f) error {
    if len(p.release) > 0 {
        return ErrPoolClosed
    }
    w := p.getWorker()
    w.task <- task
    return nil
}

// getWorker returns a available worker to run the tasks.
func (p *Pool) getWorker() *Worker {
    var w *Worker
    // 标志变量，判断当前正在运行的worker数量是否已到达Pool的容量上限
    waiting := false
    // 加锁，检测队列中是否有可用worker，并进行相应操作
    p.lock.Lock()
    idleWorkers := p.workers
    n := len(idleWorkers) - 1
    // 当前队列中无可用worker
    if n < 0 {
        // 判断运行worker数目已达到该Pool的容量上限，置等待标志
        waiting = p.Running() >= p.Cap()
        
    // 当前队列有可用worker，从队列尾部取出一个使用
    } else {
        w = idleWorkers[n]
        idleWorkers[n] = nil
        p.workers = idleWorkers[:n]
    }
    // 检测完成，解锁
    p.lock.Unlock()
    // Pool容量已满，新请求等待
    if waiting {
        // 利用锁阻塞等待直到有空闲worker
        for {
            p.lock.Lock()
            idleWorkers = p.workers
            l := len(idleWorkers) - 1
            if l < 0 {
                p.lock.Unlock()
                continue
            }
            w = idleWorkers[l]
            idleWorkers[l] = nil
            p.workers = idleWorkers[:l]
            p.lock.Unlock()
            break
        }
    // 当前无空闲worker但是Pool还没有满，
    // 则可以直接新开一个worker执行任务
    } else if w == nil {
        w = &Worker{
            pool: p,
            task: make(chan f, 1),
        }
        w.run()
        // 运行worker数加一
        p.incRunning()
    }
    return w
}
// Worker is the actual executor who runs the tasks,
// it starts a Goroutine that accepts tasks and
// performs function calls.
type Worker struct {
    // pool who owns this worker.
    pool *Pool

    // task is a job should be done.
    task chan f

    // recycleTime will be update when putting a worker back into queue.
    recycleTime time.Time
}

// run starts a Goroutine to repeat the process
// that performs the function calls.
func (w *Worker) run() {
    Go func() {
        // 循环监听任务列表，一旦有任务立马取出运行
        for f := range w.task {
            if f == nil {
                // 退出Goroutine，运行worker数减一
                w.pool.decRunning()
                return
            }
            f()
            // worker回收复用
            w.pool.putWorker(w)
        }
    }()
}