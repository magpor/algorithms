package main

// A program that implements the "interval-cover coverage" algorithm.
//
// Example to run:
//	go run interval_coverage.go interval_coverage.adoc
//
// Author: Magnus Poromaa
// Complexity: O(n log n) + O(n) = O(n log n)
import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Interval struct {
	start float64
	end   float64
	index int
}

type State int

const (
	TargetInterval      State = 0
	GivenIntervalsCount State = 1
	GivenIntervals      State = 2
)

type RequestAndResponse struct {
	targetInterval Interval
	givenIntervals []Interval
	result         []Interval
}

// Main function that reads the test data from a file and then runs the test cases one-by-one until all completed
func main() {

	//Save the request and responses in this slice to make it possible to write out the responses after all results has been written
	//var requestsAndResponses []RequestAndResponse;

	//The switch statement used to parse the input uses the state to move between cases
	var state = TargetInterval

	//The request and response DTO used to solve the problem
	requestAndResponse := RequestAndResponse{}

	//The counter used to keep track of how many given intervals to allocate in the test case
	givenIntervalsCount := 0

	//The index used to populate the given intervals in the test case
	givenIntervalsIndex := 0

	var builder strings.Builder
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		value := scanner.Text()

		switch state {
		case TargetInterval:
			//Get the target interval-cover and add it to the request and response
			requestAndResponse.targetInterval = toInterval(value)

			//Move to the GivenIntervalsCount state
			state = GivenIntervalsCount
		case GivenIntervalsCount:

			//Get the number of given intervals and add it to the request and response
			givenIntervalsCount = toInt(value)
			requestAndResponse.givenIntervals = make([]Interval, givenIntervalsCount)

			//Move to the GivenIntervals state
			state = GivenIntervals
		case GivenIntervals:

			//Create the given interval-cover and add it to the given intervals in the request and response
			requestAndResponse.givenIntervals[givenIntervalsIndex] = toInterval(value)
			requestAndResponse.givenIntervals[givenIntervalsIndex].index = givenIntervalsIndex

			//Increment the index for the next given interval-cover
			givenIntervalsIndex++

			//Decrement the number of given intervals to add
			givenIntervalsCount--

			//If we have read all intervals for the current test case, sort them and then solve the problem
			if givenIntervalsCount == 0 {
				sort.Slice(requestAndResponse.givenIntervals, func(i, j int) bool {
					return requestAndResponse.givenIntervals[i].start < requestAndResponse.givenIntervals[j].start
				})
				cover(&requestAndResponse, &builder)
				givenIntervalsCount = 0
				givenIntervalsIndex = 0

				//Move to the TargetInterval state to catch next test case from stdin
				state = TargetInterval
			}
		}
	}
}

func toInt(stringValue string) int {
	intValue, err := strconv.Atoi(stringValue)
	if err != nil {
		log.Fatal(err)
	}
	return intValue
}

// Cover function that solves the problem by iterating over the given intervals and selecting the one that covers the
// longest distance from the target interval-cover. The function will print out the number of intervals that was used to solve
func cover(testCase *RequestAndResponse, builder *strings.Builder) {
	if !isStartAndEndSame(testCase) {
		var indexes []int

		//Initially the target start position is covered, so we start at that point or earlier
		var covered = testCase.targetInterval.start

		var longestRangeIndex = 0
		var intervalsIndex = 0
		var lengthOfGivenIntervals = len(testCase.givenIntervals)

		//As long as covered is less than the target end position we keep iterating, looking for the longest interval-cover.
		for covered < testCase.targetInterval.end {

			var longestRange float64 = roundFloat(-1)
			for intervalsIndex < lengthOfGivenIntervals { //O(n)

				//Is current intervals start less or equal to the covered distance so far? If so it can be a candidate.
				if testCase.givenIntervals[intervalsIndex].start <= covered {

					//Is the distance from current intervals end to the last covered distance the longest so far?
					var currentDistance = testCase.givenIntervals[intervalsIndex].end - covered
					if currentDistance > longestRange {
						longestRange = currentDistance
						longestRangeIndex = intervalsIndex
					}

					//Advanced the intervals index as long as we have intervals that starts before the last covered distance
					intervalsIndex++
				} else {
					//As soon as an intervals start is larger than the covered distance exit the loop. Keep the interval-cover
					//index for next iteration to avoid having to making the iteration O(n^2).
					break
				}
			}

			if longestRange == -1 {
				fmt.Printf("%s\n", "impossible")
				return
			}

			//Append the original index the interval-cover had to the indexes
			indexes = append(indexes, testCase.givenIntervals[longestRangeIndex].index)

			//Set the end of the longest interval-cover as the currently covered distance
			covered = testCase.givenIntervals[longestRangeIndex].end

		}

		//Write out the results
		var numberOfIndexes = len(indexes)
		builder.WriteString(strconv.Itoa(numberOfIndexes))
		builder.WriteString("\n")
		for i := 0; i < numberOfIndexes; i++ {
			if i == numberOfIndexes-1 {
				builder.WriteString(strconv.Itoa(indexes[i]))
				builder.WriteString("\n")
			} else {
				builder.WriteString(strconv.Itoa(indexes[i]))
				builder.WriteString(" ")
			}
		}
		fmt.Println(builder.String())
		builder.Reset()
	}
}

// Checks if the target interval-cover has the same start and end and if so selects the first intervals that covers it
// otherwise notifies that the problem was impossible to solve.
// @return true if the target interval-cover has the same start and end
// @return false if the target interval-cover has different start and end
func isStartAndEndSame(testCase *RequestAndResponse) bool {
	if testCase.targetInterval.start == testCase.targetInterval.end {
		for i := 0; i < len(testCase.givenIntervals); i++ {
			if testCase.givenIntervals[i].start <= testCase.targetInterval.start && testCase.givenIntervals[i].end >= testCase.targetInterval.end {
				fmt.Printf("%d\n%d\n", 1, testCase.givenIntervals[i].index)
				return true
			}
		}
		fmt.Println("impossible")
		return true
	}
	return false
}

// Parses a string to an Interval
func toInterval(interval string) Interval {

	if strings.Count(interval, " ") != 1 {
		log.Fatalf("Interval [%s] has no separator", interval)
	}

	//Split the left and right side of the interval-cover that MUST be integers
	parts := strings.Split(interval, " ")

	//Get integer start value
	start, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		log.Fatalf("Left side of interval-cover [%s] is not a digit", interval)
	}

	//Get integer end value
	end, err := strconv.ParseFloat(parts[1], 64)
	if err != nil {
		log.Fatalf("Right side of interval-cover [%s] is not a digit\n", interval)
	}

	return Interval{
		start: roundFloat(start),
		end:   roundFloat(end),
		index: -1,
	}
}

// Rounds float64 to ceiling with a precision of 6 decimals
func roundFloat(value float64) float64 {
	return math.Ceil(value*1000000) / 1000000
}
