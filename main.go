package main
import(
	"fmt"
	"log"
	"../datafile"//상위 폴더 밑에 패키지 
)
func main(){
	numbers, err := datafile.GetFloats("data.txt")//datafile패키지의 GetFloats함수에 data.txt 전달
	if err != nil{ //에러가 있을 시 작동
		log.Fatal(err)
	}
	var sum float64 = 0 //0안써도 0로 동작하지만 명확하게 명시 해줌
	for _,number := range numbers{//foreach문과 동일
		sum = sum + number
	}
	sampleCount := float64(len(numbers))
	fmt.Printf("Average : %0.2f\n",sum/sampleCount)
}
