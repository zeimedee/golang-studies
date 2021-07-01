package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, num := range s {
		sum += num
	}
	c <- sum
}

func main() {
	nums := []int{1, 3}
	nums2 := []int{2, 4}
	nums3 := []int{5, 3}
	c := make(chan int)
	go sum(nums3, c)
	go sum(nums2, c)
	go sum(nums, c)
	v, b, n := <-c, <-c, <-c
	fmt.Printf("chanel value: a:%d and b:%d and c:%d,\n total = %d", v, b, n, v+b)

}
