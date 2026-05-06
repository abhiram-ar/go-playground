package main

func sum(s []int, ch chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}

	ch <- sum

}

func main() {
	ch := make(chan int)
	nums := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	go sum(nums[:8], ch)
	go sum(nums[8:], ch)

	// the value of x would be which ever go routine finishes first
	x, y := <-ch, <-ch

	println(x, y, x+y)
}
