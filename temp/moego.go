package temp

import (
	"context"
	"fmt"
	"log"
	"sync"
	"sync/atomic"
	"time"
)

// Q3: 并发任务控制生成器
//
// 其返回一个执行函数(executor), 该执行函数接收一个任务函数(task)
// executor 被调用时, 会根据 capacity 来执行 task: 如果正在执行的任务数不超过 capacity,
// 则立即异步执行, 否则会等到任意一个正在执行的 task 结束后再执行. 执行函数返回一个只读的
// channel, 其值为执行结果 ExecuteResult.
//
// ctx 用来终止执行任务.
type taskFunc func() interface{}

type Pool struct {
	Capacity int64
	Running  uint64
	Task     chan taskFunc
	Result   chan ExecuteResult
	Ctx      context.Context
	sync.Mutex
}

func NewPool(capacity int, ctx context.Context) *Pool {
	return &Pool{
		Capacity: int64(capacity),
		Task:     make(chan taskFunc, capacity),
		Result:   make(chan ExecuteResult),
		Ctx:      ctx,
	}
}

func (p *Pool) IncrRunning() {
	atomic.AddUint64(&p.Running, 1)
}

func (p *Pool) DecRunning() {
	atomic.AddUint64(&p.Running, ^uint64(0))
}

func (p *Pool) GetRunningWorkers() int64 {
	return int64(atomic.LoadUint64(&p.Running))
}

func (p *Pool) Put(task taskFunc) {
	p.Lock()
	defer p.Unlock()

	if p.GetRunningWorkers() < p.Capacity {
		p.RunningFunc()
	}

	select {
	case <-p.Ctx.Done():
		p.Close()
		return
	default:
		p.Task <- task
	}
}

func (p *Pool) CheckWorker() {
	p.Lock()
	defer p.Unlock()

	if p.Running == 0 && len(p.Task) > 0 {
		p.RunningFunc()
	}
}

func (p *Pool) Close() {
	for len(p.Task) > 0 {
		time.Sleep(1e6)
	}

	close(p.Task)
}

func (p *Pool) RunningFunc() {
	p.IncrRunning()

	go func() {
		defer func() {
			p.DecRunning()
			p.CheckWorker()

			if err := recover(); err != nil {
				fmt.Println("panic, err: ", err)
				p.Result <- ExecuteResult{
					Ok:   false,
					Data: nil,
				}
			}
		}()

		for {
			select {
			case <-p.Ctx.Done():
				return
			case t, ok := <-p.Task:
				if !ok {
					return
				}

				res := t()
				p.Result <- ExecuteResult{
					Ok:   true,
					Data: res,
				}
			}
		}
	}()
}

func CreateAsyncWorker(ctx context.Context, capacity int) func(task func() interface{}) <-chan ExecuteResult {
	//show me your code
	pool := NewPool(capacity, ctx)
	return func(task func() interface{}) <-chan ExecuteResult {
		pool.Put(task)
		return pool.Result
	}

	//pool := make(chan taskFunc, capacity)
	//result := make(chan ExecuteResult)
	//go func() {
	//	for {
	//		select {
	//		case <-ctx.Done():
	//			return
	//		case task, ok := <-pool:
	//			if ok {
	//				go func() {
	//					defer func() {
	//						if err := recover(); err != nil {
	//							result <- ExecuteResult{
	//								Ok:   false,
	//								Data: err,
	//							}
	//						}
	//					}()
	//
	//					res := task()
	//					result <- ExecuteResult{
	//						Ok:   true,
	//						Data: res,
	//					}
	//				}()
	//			}
	//		}
	//	}
	//}()
	//
	//return func(task func() interface{}) <-chan ExecuteResult {
	//	pool <- task
	//	return result
	//}
}

type ExecuteResult struct {
	// 标记是否成功
	Ok bool
	// 结果数据, 如果成功, 则表示 task 的返回值, 否则表示异常
	Data interface{}
}

func TestCreateAsyncWorker(factory func(ctx context.Context, capacity int) func(task func() interface{}) <-chan ExecuteResult) {
	type Task struct {
		id            int
		expectedDelay int
		fail          bool
		result        <-chan ExecuteResult
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	executor := factory(ctx, 2)
	var tasks []Task

	runTask := func(id, delay, expectedDelay int, fail bool) {
		var task Task
		task.id = id
		task.expectedDelay = expectedDelay
		task.fail = fail
		now := time.Now()
		task.result = executor(func() interface{} {
			time.Sleep(time.Duration(delay) * time.Millisecond)
			realDelay := int(time.Now().Sub(now) / time.Millisecond)
			if fail {
				panic(realDelay)
			} else {
				return realDelay
			}
		})
		tasks = append(tasks, task)
	}

	runTask(1, 100, 100, false)
	runTask(2, 200, 200, true)
	runTask(3, 300, 400, false)
	runTask(4, 400, 600, true)
	runTask(5, 100, 500, false)
	runTask(6, 200, 700, true)
	runTask(7, 100, 700, false)
	runTask(8, 200, 900, false)

	for _, task := range tasks {
		result := <-task.result
		if result.Ok == task.fail {
			log.Fatalf("failed status of task %d is %t, expected is %t", task.id, !result.Ok, task.fail)
		}
		realDelay := result.Data.(int)
		// 避免延时抖动
		if realDelay/100 != task.expectedDelay/100 {
			log.Fatalf("delay of task %d should be %d rather than %d", task.id, task.expectedDelay, realDelay)
		}
	}
}

func main() {
	TestCreateAsyncWorker(CreateAsyncWorker)
}
