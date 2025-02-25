package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

//A program that solves the "knapsack 0-1" problem

//Optimal sub-structure
//Case1: The item belongs in the optimal sub-fill
//Case2: The item does not belong in the optimal sub-fill

type State int

const (
	CapacityAndObjects State = 0
	Intervals          State = 1
)

func main() {
	var state = CapacityAndObjects
	var profits []int
	var weights []int
	var values []string
	var capacity int = 0
	var numberOfItems int = 0
	var numberOfItemsRead int = 0
	var memo [][]int

	var responseBuilder = strings.Builder{}

	//Start scanning the input until end-of-line appears
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		value := scanner.Text()
		switch state {
		case CapacityAndObjects:

			//Reset
			capacity = 0
			numberOfItems = 0
			numberOfItemsRead = 0

			//Get the capacity of the knapsack and the number of items in the test
			values = strings.Split(value, " ")
			capacity = toInt(&values, 0)
			numberOfItems = toInt(&values, 1)

			//Allocate room for the weights and profits slices
			profits = make([]int, numberOfItems)
			weights = make([]int, numberOfItems)

			//Move to the Intervals state
			state = Intervals
		case Intervals:
			//Get the profit and weight of the item
			values = strings.Split(value, " ")
			profits[numberOfItemsRead] = toInt(&values, 0)
			weights[numberOfItemsRead] = toInt(&values, 1)

			//Keep reading numbers from stdin until all are read
			numberOfItemsRead++
			if numberOfItemsRead == numberOfItems {

				memo = make([][]int, numberOfItems+1)
				for i := 0; i < numberOfItems+1; i++ {
					memo[i] = make([]int, capacity+1)
				}

				//Calculate max and selected indexes
				_, selected := knapsack(capacity, &weights, &profits, numberOfItems, &memo)

				//Write the response
				responseBuilder.Reset()
				responseBuilder.WriteString(strconv.Itoa(len(selected)))
				responseBuilder.WriteString("\n")
				for ii := 0; ii < len(selected); ii++ {
					responseBuilder.WriteString(strconv.Itoa(selected[ii]))
					responseBuilder.WriteString(" ")
				}
				fmt.Println(responseBuilder.String())

				//Print out memo to see how it looks like
				/*
					var memoBuilder = strings.Builder{}
					for i := 0; i < len(memo); i++ {
						for j := 0; j < len(memo[i]); j++ {
							memoBuilder.WriteString(strconv.Itoa(memo[i][j]))
							memoBuilder.WriteString(" ")
						}
						memoBuilder.WriteString("\n")
					}
					fmt.Println(memoBuilder.String())
				*/

				//Move to the next knapsack test
				state = CapacityAndObjects
			}
		}
	}

}

func knapsack(capacity int, weights *[]int, profits *[]int, n int, memo *[][]int) (int, []int) {

	//Example: Read up from https://medium.com/@fabianterh/how-to-solve-the-knapsack-problem-with-dynamic-programming-eb88c706d3cf
	//capacity = 6 (7 columns)
	//Number of elements = 4 (5 rows)
	//6 4
	//v w
	//5 4
	//4 3
	//3 2
	//2 1
	//weights = (4 3 2 1)
	//values =  (5 4 3 2)

	//Matrix
	//0 0 0 0 0 0 0
	//0 0 0 0 5 5 5
	//0 0 0 4 5 5 5
	//0 0 3 4 5 7 8
	//0 2 3 5 6 7 9

	//Selected = (3 2 1)

	//Tree - Left branch is include, right branch is exclude [capacity of knapsack left,weight of item]
	//													[6,4]
	//											/						\
	//									   [2,3]			  			[6,3]
	//									/		\			 		/ 		\
	//							[-1,2]		  	[2,2]		  	[3,2]		[6,2]
	//							/	\		  	/  \		 	/  	\ 		  |	  \
	//						[-3,1]	[-1,1] 	[0,1]	[2,1]	[1,1]	[3,1]	[4,1] [6,1]
	//
	//	What kind of conclusions can we take from this information?
	//	1. There are two combinations that can use the entire capacity of the knapsack (0,2) and (1,2,3) (0 indexed assumed)
	//	2. The maximum number of items that can be included if we include the one with weight 4 is 2 items, which makes
	//	   sense since the third level will need 3, and 6 - 4 = 2. This also tell us that the second node must be on level 2

	//Populate memo
	for i := 0; i <= n; i++ {
		for j := 0; j <= capacity; j++ {
			if i == 0 || j == 0 {
				(*memo)[i][j] = 0
			} else if (*weights)[i-1] <= j {
				(*memo)[i][j] = max((*profits)[i-1]+(*memo)[i-1][j-(*weights)[i-1]], (*memo)[i-1][j])
			} else {
				(*memo)[i][j] = (*memo)[i-1][j]
			}
		}
	}

	//Backtrack
	var capacityLeftToUse = capacity
	var selected []int
	for ii := n; ii > 0; ii-- {
		if (*memo)[ii][capacityLeftToUse] != (*memo)[ii-1][capacityLeftToUse] {
			selected = append(selected, ii-1)     //Add the index to the selected
			capacityLeftToUse -= (*weights)[ii-1] //Change the capacity left
		}
	}
	return (*memo)[n][capacity], selected
}

func max(left int, right int) int {
	if left > right {
		return left
	}
	return right
}

func toInt(values *[]string, index int) int {
	value1, err := strconv.Atoi((*values)[index])
	if err != nil {
		log.Fatal(err)
	}
	return value1
}
