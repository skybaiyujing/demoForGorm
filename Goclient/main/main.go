package main

import (
	"GoClient/client"
	"GoClient/model"
	"fmt"
	"sync"
)

func main() {
	snos := []int{1001, 1002, 1003, 1004, 1005}
	out := make(chan model.Student, len(snos))
	wg := sync.WaitGroup{} //等待多个 goroutine 执行完成的机制
	rdb, ctx := client.InitReids()

	for _, sno := range snos {
		wg.Add(1)
		go client.GetStudentWithRedisCache(sno, rdb, ctx, out, &wg)
	}

	wg.Wait() //主线程阻塞在此，直到所有 goroutine 执行完毕
	close(out)

	for stu := range out {
		fmt.Printf("学生信息：Sno=%d, 姓名=%s\n", stu.Sno, stu.Sname)
	}
}
