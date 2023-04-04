package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"io"	
	"strconv"
)

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}
func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
func checkCorrect(s []int) {
	count := 0
	for _,value := range s{
		if(value == 1){
		  count += 1
		}

	}
	fmt.Println("You've got ",count,"correct questions out of ",len(s))
}
var ans int
var myslice []int

func main() {
	dat, err := os.Open("/home/second/Desktop/quiz-master/problems.csv")
	
	checkError(err)
	reader := csv.NewReader(dat)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		checkError(err)
		for s, stat := range record {

			if s == 0 {

				fmt.Println("What is ", stat, "?")
				fmt.Scanf("%d", &ans)

			}
			if s == 1 {
				if strconv.Itoa(ans) == stat {
					myslice = append(myslice, 1)

				} else {
					myslice = append(myslice, 0)

				}

			}

		}
	}
	defer checkCorrect(myslice)
	
	

}
