package main

import "fmt"

type subscriber struct {
	name   string
	rate   int
	active bool
}

func printInfo(s subscriber) {
	fmt.Println("Name : ", s.name)
<<<<<<< HEAD
	fmt.Println("Mothly rate : ", s.rate)
=======
	fmt.Println("Monthly rate : ", s.rate)
>>>>>>> 60f7e9281c07242175d4b573e27b0bc5b4faace6
	fmt.Println("Active? : ", s.active)
}
func defaultSubscriber(name string) subscriber {
	var s subscriber
	s.name = name
	s.rate = 5000
	s.active = true
	return s
}
<<<<<<< HEAD
func applyDiscount(s *subscriber){
	s.rate = 4000//주소값으로 받지 않으면 값이 바뀌지 않음.
}
=======
>>>>>>> 60f7e9281c07242175d4b573e27b0bc5b4faace6
func main() {
	s1 := defaultSubscriber("Kim")
	s1.rate = 4500
	printInfo(s1)

<<<<<<< HEAD
	s2 := defaultSubscriber("Park")
	printInfo(s2)

	var s3 subscriber
	applyDiscount(&s3)//구조체에서 포인터 동작함.
	fmt.Println(s3.rate)
=======

	s2 := defaultSubscriber("Park")
	printInfo(s2)
>>>>>>> 60f7e9281c07242175d4b573e27b0bc5b4faace6
}
