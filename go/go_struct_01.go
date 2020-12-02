package main

import "fmt"

	type subscriber struct {
		name   string
		rate   int
		active bool
	}
func main() {
	var s1 subscriber

	fmt.Printf("%#v\n", s1)

	s1.name = "Kim"
	s1.rate = 5000
	s1.active = false

	fmt.Printf("%s\n", s1.name)
	fmt.Println(s1.rate)
	fmt.Println(s1.active)
	/*구조체
	var subscriber struct {
		name   string
		rate   int
		active bool
	}
	fmt.Printf("%#v\n", subscriber)

	subscriber.name = "Kim"
	subscriber.rate = 5000
	subscriber.active = false

	fmt.Printf("%s\n", subscriber.name)
	fmt.Println(subscriber.rate)
	fmt.Println(subscriber.active)
	맵
	subscriber:=map[string]float64{}
	subscriber= ["name"] = "Kim"
	subscriber= ["rate'"] = 5000
	subscriber= ["active"] = false
	fmt.Println("test")
	슬라이스
	subscriber:=[]string{}
	subscriber= append(subscriber,"Kim")
	subscriber= append(subscriber,5000)
	subscriber= append(subscriber,false)
	fmt.Println(subscribe)*/
}
