package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

type Plan struct {
	ID    int
	Name  string
	Todos []*Todo
}

type Todo struct {
	ID     int
	PlanID int
	Name   string
}

func main() {
	c := make(chan Plan)
	var plans []*Plan
	for i := 0; i < 10; i++ {
		plans = append(plans, &Plan{ID: i, Name: "Plan " + strconv.Itoa(i)})
	}

	// goroutine 실행
	for _, p := range plans {
		go addTodos(p, c)
	}

	// channel lock
	for i := 0; i < len(plans); i++ {
		plan := <-c
		fmt.Println(plan.ID, plan.Name, "Todos ", len(plan.Todos))
	}

	plan := plans[0]
	todo := getFirstTodo(plan)
	fmt.Println(todo.Name, todo)

	todo.Name = "TEST"

	fmt.Println(todo.Name, todo)
	t := plan.Todos[0]
	fmt.Println(plan.Todos[0].Name, t.Name, t)

}

// goroutine 함수
func addTodos(plan *Plan, c chan Plan) {
	var todos []*Todo
	for i := 0; i < rand.Intn(10)+1; i++ {
		// time.Sleep(time.Second * 1)
		todos = append(todos, &Todo{ID: i, PlanID: plan.ID, Name: "Todo " + strconv.Itoa(i)})
	}
	plan.Todos = todos

	c <- *plan
}

func getFirstTodo(plan *Plan) *Todo {
	return plan.Todos[0]
}

// func multiply(inChan, outChan chan int) {
// 	defer close(outChan)
// 	for in := range inChan {
// 		outChan <- in * 2
// 	}
// }

// func input(inChan chan int) {
// 	defer close(inChan)
// 	for i := 0; i < 5; i++ {
// 		inChan <- i
// 	}
// }

// func main() {
// 	done := make(chan struct{})

// 	go func() {
// 		time.Sleep(10 * time.Second)
// 		done <- struct{}{}
// 	}()
// 	<-done

// inChan := make(chan int, 5)
// outChan := make(chan int)

// go multiply(inChan, outChan)
// go input(inChan)

// for {
// 	out, more := <-outChan
// 	if !more {
// 		return
// 	}
// 	log.Println(out)
// }

// for i := 0; i < 10; i++ {
// 	go func(i int) {
// 		log.Println(i)
// 	}(i)
// }

// time.Sleep(time.Second)
// }
