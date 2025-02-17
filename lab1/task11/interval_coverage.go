package main

// A program that returns the minimum set of the intervalls taken from a given interval that covers the
// target interval.
// Example:
//	go run interval_coverage.go test.txt
//
// Author: Magnus Poromaa
// Complexity: O(n log n) + O(n) = O(n log n). This is because I decided to order the given intervals in increasing start order and use
//that fact to iterate the given intervals only once. The sorting is done in O(n log n) and the iteration is done in O(n).

import (
	"bufio"
	"cmp"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

// Interval represents a interval with a start and end value
type Interval struct {
	start float32
	end   float32
	index int
}

// TestCase represents a test case that must be resolved
type TestCase struct {
	targetInterval Interval
	givenIntervals []Interval
	result         []Interval
}

func main() {
	argsWithoutProgram := os.Args[1:] //Parse the arguments without the program name

	if len(argsWithoutProgram) != 1 {
		log.Println("The program expects a path to the file containing the test data")
		return
	}

	file, err := os.Open(argsWithoutProgram[0])
	if err != nil {
		log.Fatal(err)
	}

	//Defer the closing of the file until the OS exist
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	//Scan the file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		//Create the test case
		testCase := TestCase{}

		//Scan the target interval
		testCase.targetInterval = toInterval(scanner.Text())

		//Scan the number of given intervals the file contains
		scanner.Scan()
		numberOfIntervals, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		//Scan the given intervals
		testCase.givenIntervals = make([]Interval, numberOfIntervals)
		for i := 0; i < numberOfIntervals; i++ {
			scanner.Scan()
			testCase.givenIntervals[i] = toInterval(scanner.Text())
			testCase.givenIntervals[i].index = i
		}

		//Sort the given intervals in increasing start order O(n * log n)
		intervalComparator := func(a, b Interval) int {
			return cmp.Compare(a.start, b.start)
		}
		slices.SortFunc(testCase.givenIntervals, intervalComparator)

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		//Run the test case and then print it out
		cover(&testCase)

	}
}

// Function that returns the minimum set of the intervalls taken from a given interval that covers the target interval
func cover(testCase *TestCase) {
	if !isStartAndEndSame(testCase) {
		var indexes []int
		var covered = testCase.targetInterval.start
		var longestRangeIndex = -1
		var intervalIndex = 0
		for covered < testCase.targetInterval.end {
			var longestRange float32 = -1
			for intervalIndex < len(testCase.givenIntervals) { //O(n)

				//Is current intervals start less or equal to the covered distance so far?
				if testCase.givenIntervals[intervalIndex].start <= covered {

					//Is the distance from current intervals end to the last covered distance the longest so far?
					var currentDistance = testCase.givenIntervals[intervalIndex].end - covered
					if currentDistance > longestRange {
						longestRange = currentDistance
						longestRangeIndex = intervalIndex
					}

					//Advanced the intervals index as long as we have intervals that starts before the last covered distance
					intervalIndex++
				} else {
					//As soon as an intervals start is larger than the covered distance exit the loop. Keep the interval
					//index for next iteration to avoid having to making the iteration O(n^2).
					break
				}
			}

			if longestRange == -1 {
				fmt.Println("impossible")
				return
			}

			//Append the original index the interval had to the indexes
			indexes = append(indexes, testCase.givenIntervals[longestRangeIndex].index)

			//Set the end of the longest interval as the currently covered distance
			covered = testCase.givenIntervals[longestRangeIndex].end

		}

		//The number of indexes from the given intervals that was used to solve the problem
		var numberOfIndexes = len(indexes)

		//Print out the number of interval that was needed to solve the problem
		fmt.Println(numberOfIndexes)

		//Print out the indexes that was used to solve the problem. NOTE! Since the response should display the index the
		//interval has in the input and we have resorted the given intervals in  increasing start order we MUST use the
		//index registered in the Interval type and not the current iteration index.
		for i := 0; i < numberOfIndexes; i++ {
			fmt.Printf("%d ", indexes[i])
		}
		fmt.Println("")
	}

}

// Decides if the target index has the  same start and end value and if so take the first (any) interval that
// has start <= target.start and end >= target.end.
func isStartAndEndSame(testCase *TestCase) bool {
	if testCase.targetInterval.start == testCase.targetInterval.end {
		for i := 0; i < len(testCase.givenIntervals); i++ {
			if testCase.givenIntervals[i].start <= testCase.targetInterval.start && testCase.givenIntervals[i].end >= testCase.targetInterval.end {
				fmt.Println(1) //One given interval at-least covers the target
				fmt.Println(i) //The index of the one given interval that was selected
				return true
			}
		}
		fmt.Println("impossible") //Impossible to find a given interval that covers the target
		return true
	}
	return false
}

func toInterval(interval string) Interval {

	if !strings.Contains(interval, " ") {
		_, _ = fmt.Fprintf(os.Stderr, "Interval [%s] has no separator\n", interval)
		os.Exit(1)
	}

	//Split the left and right side of the interval hat MUST be integers
	parts := strings.Split(interval, " ")

	//Get integer start value
	start, err := strconv.ParseFloat(parts[0], 32)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Left side of interval [%s] is not a digit\n", interval)
		os.Exit(1)
	}

	//Get integer end value
	end, err := strconv.ParseFloat(parts[1], 32)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Right side of interval [%s] is not a digit\n", interval)
		os.Exit(1)
	}

	return Interval{
		start: float32(start),
		end:   float32(end),
	}
}

/*
func getInterval(interval []string) Interval {
	return toInterval(interval)
}

func getIntervals(line []string) []Interval {
	given_intervals_as_strings := strings.Split(line[1], ",")
	given_intervals := make([]Interval, len(given_intervals_as_strings))
	for i := 0; i < len(given_intervals_as_strings); i++ {
		given_intervals[i] = toInterval(given_intervals_as_strings[i])
	}
	return given_intervals
}
*/
/*
	argsWithoutProgram := os.Args[1:] //Parse the arguments without the program name
	if len(argsWithoutProgram) != 2 {
		log.Println("The program expects two arguments where the first is the target_interval interval and the second are the given intervals")
		return
	}

	targetInterval := getTargetInterval(argsWithoutProgram)
	log.Printf("Target interval is: %d\n", targetInterval)

	givenIntervals := getGivenIntervals(argsWithoutProgram)
	log.Printf("Given intervals are: %d\n", givenIntervals)
*/
