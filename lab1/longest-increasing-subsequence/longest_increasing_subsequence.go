package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Solution to the longest increasing subsequence with time complexity O(nlog(n)) using a binary findLongestIncrementingSubsequence algorithm.
// @see OSX https://github.com/golang/go/issues/68836 so it is not possible to test with longer than 1024 bytes of input
// @see https://en.wikipedia.org/wiki/Longest_increasing_subsequence
import (
	"log"
	"strconv"
)

func main() {

	//Increase the buffer size so the scanner can read long files
	scanner := bufio.NewScanner(os.Stdin)
	buffer := make([]byte, 0, 64*1024) //Initial buffer for scanner set to 64KB
	scanner.Buffer(buffer, 1024*1024)  //Growing buffer for scanner set to 1MB

	var numberOfEntriesInInput = 0
	var builder = strings.Builder{}

	//Test example
	//7
	//2 8 9 5 6 7 1
	//0 3 4 5

	//Test example - It seems to be a problem with the size when reading the line
	//139	- Do not work
	//844389 570305 37399 177231 -299381 -723275 -539515 847829 -786322 738474 -653729 1075812 -129803 -360609 -848241 -384013 575561 896177 433730 285513 -236713 46864 233109 -246941 777424 -714520 -145816 404879 594433 -137772 18838 986247 -547468 1066238 163478 660577 -159612 1111389 468405 521492 1307288 744677 1034730 604911 -198507 1076489 -381677 1247055 832665 902053 889993 -56623 -231083 450528 496437 -146234 516008 1120621 1018645 377866 -277152 284909 84113 402807 568572 937591 -256617 1596386 1226405 861789 757877 1681118 743891 920032 886028 1135384 576521 1074352 1459864 476611 516404 1397283 939989 1312747 347810 -73575 1646513 1333818 24471 1632584 151684 177320 854918 1163223 980127 1813490 998239 1610936 287300 574643 1330150 1892603 575760 384040 1112635 269214 316850 1959156 1603566 544139 1193192 867395 161421 343180 897568 699231 449605 714080 710475 624077 1004089 1509585 1438823 1969007 1290233 1026375 379921 358471 697310 1234647 973114 2057461 1147250 2076300 958926 2249885 2325514 1763202 686465

	//138	- Do not work On OSX there is a bug that makes the internal buffer 1024 bytes always no matter what we try
	//844389 570305 37399 177231 -299381 -723275 -539515 847829 -786322 738474 -653729 1075812 -129803 -360609 -848241 -384013 575561 896177 433730 285513 -236713 46864 233109 -246941 777424 -714520 -145816 404879 594433 -137772 18838 986247 -547468 1066238 163478 660577 -159612 1111389 468405 521492 1307288 744677 1034730 604911 -198507 1076489 -381677 1247055 832665 902053 889993 -56623 -231083 450528 496437 -146234 516008 1120621 1018645 377866 -277152 284909 84113 402807 568572 937591 -256617 1596386 1226405 861789 757877 1681118 743891 920032 886028 1135384 576521 1074352 1459864 476611 516404 1397283 939989 1312747 347810 -73575 1646513 1333818 24471 1632584 151684 177320 854918 1163223 980127 1813490 998239 1610936 287300 574643 1330150 1892603 575760 384040 1112635 269214 316850 1959156 1603566 544139 1193192 867395 161421 343180 897568 699231 449605 714080 710475 624077 1004089 1509585 1438823 1969007 1290233 1026375 379921 358471 697310 1234647 973114 2057461 1147250 2076300 958926 2249885 2325514 1763202

	//137	- Do work
	//844389 570305 37399 177231 -299381 -723275 -539515 847829 -786322 738474 -653729 1075812 -129803 -360609 -848241 -384013 575561 896177 433730 285513 -236713 46864 233109 -246941 777424 -714520 -145816 404879 594433 -137772 18838 986247 -547468 1066238 163478 660577 -159612 1111389 468405 521492 1307288 744677 1034730 604911 -198507 1076489 -381677 1247055 832665 902053 889993 -56623 -231083 450528 496437 -146234 516008 1120621 1018645 377866 -277152 284909 84113 402807 568572 937591 -256617 1596386 1226405 861789 757877 1681118 743891 920032 886028 1135384 576521 1074352 1459864 476611 516404 1397283 939989 1312747 347810 -73575 1646513 1333818 24471 1632584 151684 177320 854918 1163223 980127 1813490 998239 1610936 287300 574643 1330150 1892603 575760 384040 1112635 269214 316850 1959156 1603566 544139 1193192 867395 161421 343180 897568 699231 449605 714080 710475 624077 1004089 1509585 1438823 1969007 1290233 1026375 379921 358471 697310 1234647 973114 2057461 1147250 2076300 958926 2249885 2325514

	//130	- Do work
	//844389 570305 37399 177231 -299381 -723275 -539515 847829 -786322 738474 -653729 1075812 -129803 -360609 -848241 -384013 575561 896177 433730 285513 -236713 46864 233109 -246941 777424 -714520 -145816 404879 594433 -137772 18838 986247 -547468 1066238 163478 660577 -159612 1111389 468405 521492 1307288 744677 1034730 604911 -198507 1076489 -381677 1247055 832665 902053 889993 -56623 -231083 450528 496437 -146234 516008 1120621 1018645 377866 -277152 284909 84113 402807 568572 937591 -256617 1596386 1226405 861789 757877 1681118 743891 920032 886028 1135384 576521 1074352 1459864 476611 516404 1397283 939989 1312747 347810 -73575 1646513 1333818 24471 1632584 151684 177320 854918 1163223 980127 1813490 998239 1610936 287300 574643 1330150 1892603 575760 384040 1112635 269214 316850 1959156 1603566 544139 1193192 867395 161421 343180 897568 699231 449605 714080 710475 624077 1004089 1509585 1438823 1969007 1290233 1026375 379921 358471 697310 1234647

	for {

		numberOfEntriesInInput = readNumberOfExpectedSequenceEntriesFromInput(scanner)

		//Read the sequence entries from input
		var intValues = make([]int, numberOfEntriesInInput)
		readSequenceEntriesFromInput(scanner, numberOfEntriesInInput, intValues)

		//Find the longest increasing subsequence
		numberOfEntriesInSubsequence, entriesInSubsequence, entryValuesInSubsequence := findLongestIncrementingSubsequence(intValues, numberOfEntriesInInput)

		//Write number of entries in subsequence and entries in subsequence to stdout
		writeSequenceEntriesToOutput(builder, numberOfEntriesInSubsequence, entriesInSubsequence, entryValuesInSubsequence)

	}
}

/**
 * Find the longest increasing subsequence in an array of integers.
 * @see https://en.wikipedia.org/wiki/Longest_increasing_subsequence
 */
func findLongestIncrementingSubsequence(x []int, size int) (int, []int, []int) {

	//Stores the indexes of the current longest increasing subsequence
	var m []int = make([]int, size+1)
	m[0] = -1

	//Stores the indexes of the ancestors to the current longest increasing subsequence
	var p []int = make([]int, size)

	//Stores the length of the current longest increasing m so far
	var length = 0
	var newLength = 0
	for i := 0; i < size; i++ {

		//Binary search for the largest index j such that x[m[j]] < x[j]
		var lo = 1
		var hi = length + 1
		for lo < hi {
			var mid = lo + ((hi - lo) >> 1)
			if x[m[mid]] >= x[i] {
				hi = mid
			} else if x[m[mid]] < x[i] {
				lo = mid + 1
			}
		}

		newLength = lo        //The new length of the longest increasing subsequence ending at i
		p[i] = m[newLength-1] //Add to the ancestor slice the previous last value of the longest increasing subsequence
		m[newLength] = i      //Add to the new last value to the longest increasing subsequence

		if newLength > length {
			length = newLength
		}
	}

	//Build the longest increasing subsequence indexes in incrementing order
	var s = make([]int, length)
	var v = make([]int, length)
	var k = m[length]
	for j := length - 1; j >= 0; j-- {
		s[j] = k
		v[j] = x[k]
		k = p[k]
	}
	return length, s, v

}

func readSequenceEntriesFromInput(scanner *bufio.Scanner, numberOfSequenceEntries int, intValues []int) {
	scanner.Scan()
	var err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	var value = scanner.Text()
	var stringValues = strings.Split(value, " ")
	for i := 0; i < numberOfSequenceEntries; i++ {
		intValues[i] = toInt(&(stringValues[i]))
	}
}

func readNumberOfExpectedSequenceEntriesFromInput(scanner *bufio.Scanner) int {
	scanner.Scan()
	var err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	var value = scanner.Text()
	return toInt(&value)
}

func writeSequenceEntriesToOutput(builder strings.Builder, numberOfEntriesInSubsequence int, entriesInSubsequence []int, entryValuesInSubsequence []int) {
	builder.WriteString(strconv.Itoa(numberOfEntriesInSubsequence))
	builder.WriteString("\n")
	for i := 0; i < numberOfEntriesInSubsequence; i++ {
		builder.WriteString(strconv.Itoa(entriesInSubsequence[i]))
		builder.WriteString(" ")
	}
	builder.WriteString("\n")

	/*
		for i := 0; i < numberOfEntriesInSubsequence; i++ {
			builder.WriteString(strconv.Itoa(entryValuesInSubsequence[i]))
			builder.WriteString(" ")
		}
		builder.WriteString("\n")
	*/
	fmt.Println(builder.String())
	builder.Reset()
}

func toInt(s *string) int {
	value, err := strconv.Atoi(*s)
	if err != nil {
		log.Fatal(err)
	}
	return value
}
