package main

import (
	"bufio"
	"container/list"
	"fmt"
	"log"
	"os"
	"strings"
)

var logData = make(map[string]*list.List)
var foodMap = make(map[string]string)

func main() {

	//reads data from log file
	err := getLogData("restaurantLog1")
	if err != nil {
		fmt.Printf("Error Reading Log File, Error: %s", err.Error())
	}

	// reads data from foodmenu file
	err = getFoodMenu("foodmenu")
	if err != nil {
		fmt.Printf("Error Reading Food Menu File, Error: %s", err.Error())
	}

	//finds the top three
	top, second, third := GetTopThree()

	fmt.Println()

	//prints the name of top three foods
	fmt.Print("First: ")
	printFoodName(top)

	fmt.Print("Second: ")
	printFoodName(second)

	fmt.Print("Third: ")
	printFoodName(third)

}

// reads data from text file
func getLogData(fileName string) error {

	file, err := os.Open(fileName)
	if err != nil {
		return err
	}

	defer file.Close()

	bufferedData := bufio.NewScanner(file)
	bufferedData.Scan()
	bufferedData.Scan()
	for bufferedData.Scan() {

		//to split string
		data := strings.Split(bufferedData.Text(), ": ")
		eater_id := data[0]
		food_ids := strings.Split(data[1], ", ")

		for index := 0; index < len(food_ids); index++ {

			key := strings.TrimSpace(food_ids[index])

			if _, keyExists := logData[key]; !keyExists {
				logData[key] = list.New()
			}

			// to check whether diner has taken the food more than once
			for e := logData[key].Front(); e != nil; e = e.Next() {

				if e.Value == eater_id {
					log.Fatal("Dinner eaten same food twice.")
				}
			}
			logData[key].PushBack(eater_id)
		}
	}
	return nil
}

// to read food mennu data from text file
func getFoodMenu(fileName string) error {

	file, err := os.Open(fileName)
	if err != nil {
		return err
	}

	defer file.Close()

	bufferedData := bufio.NewScanner(file)
	bufferedData.Scan()
	bufferedData.Scan()
	for bufferedData.Scan() {

		data := strings.Split(bufferedData.Text(), " - ")

		foodID := data[0]
		foodName := data[1]

		foodMap[foodID] = foodName

	}
	return nil
}

// finds the top three foods
func GetTopThree() (int, int, int) {

	var top, second, third int

	for key := range logData {

		count := logData[key].Len()
		// fmt.Printf("Food id: %s, count: %d\n", key, logData[key].Len())
		if top < count {
			third = second
			second = top
			top = count
		} else if second < count && top != count {
			third = second
			second = count
		} else if third < count && top != count && second != count {
			third = count
		}
	}
	fmt.Println()
	return top, second, third
}

// prints the name of the food
func printFoodName(rank int) {

	for key := range logData {

		if logData[key].Len() == rank {
			fmt.Printf(" %s\t", foodMap[key])
		}
	}
	fmt.Println()
	// fmt.Println("---", rank)
}
