package main

import "strconv"

func task(id int) {
	ids := strconv.Itoa(id)
	println("task_" + ids + " is running ...")
}

func add(x, y int) (int, bool) {
	var z = x+y
	println(z)
	return z, true
}

func main() {
	add(1, 2)
	task(10)
}
