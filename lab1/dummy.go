package lab1

func lowerBound(values *[]int, target int) int {
	low, high, mid := 0, len(*values)-1, 0
	for low <= high {
		mid = (low + high) / 2
		if (*values)[mid] >= target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return low
}

func upperBound(array *[]int, target int) int {
	low, high, mid := 0, len(*array)-1, 0

	for low <= high {
		mid = (low + high) / 2
		if (*array)[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return low
}

//https://cp-algorithms.com/sequences/longest_increasing_subsequence.html

/*
func longest(items []int, numberOfItems int) (int, []int) {

	var numberOfTimesItemIncludedInSubsequence []int = make([]int, numberOfItems)
	for i := 0; i < numberOfItems; i++ {
		numberOfTimesItemIncludedInSubsequence[i] = 1
	}

	var ancestorsInSubsequence []int = make([]int, numberOfItems)
	for i := 0; i < numberOfItems; i++ {
		ancestorsInSubsequence[i] = -1
	}

	//Count the number of times an item was included in any subsequence and keep track of the ancestors to be able to
	//recreate the longest increasing subsequence
	for i := 0; i < numberOfItems; i++ {
		for j := 0; j < i; j++ {
			if items[j] < items[i] && numberOfTimesItemIncludedInSubsequence[i] < numberOfTimesItemIncludedInSubsequence[j]+1 {
				numberOfTimesItemIncludedInSubsequence[i] = numberOfTimesItemIncludedInSubsequence[j] + 1
				ancestorsInSubsequence[i] = j
			}
		}
	}

	//Find the index of the item with the most number of times included in the subsequence.
	var numberOfTimesIndexIncludedInAnySubsequence = ancestorsInSubsequence[0]
	var indexOfItemIncludedInSubsequence = 0
	for i := 1; i < numberOfItems; i++ {
		if numberOfTimesItemIncludedInSubsequence[i] > numberOfTimesIndexIncludedInAnySubsequence {
			numberOfTimesIndexIncludedInAnySubsequence = numberOfTimesItemIncludedInSubsequence[i]
			indexOfItemIncludedInSubsequence = i
		}
	}

	//Follow the ancestors to find the longest increasing subsequence
	var subsequence []int
	for indexOfItemIncludedInSubsequence != -1 {
		subsequence = append(subsequence, indexOfItemIncludedInSubsequence)
		indexOfItemIncludedInSubsequence = ancestorsInSubsequence[indexOfItemIncludedInSubsequence]
	}
	reverse(&subsequence)
	return numberOfTimesIndexIncludedInAnySubsequence, subsequence
}

/*
func lowerBound(values *[]int, target int) int {
	low, high, mid := 0, len(*values)-1, 0
	for low <= high {
		mid = (low + high) / 2
		if (*values)[mid] >= target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return low
}

func lowerBound(arr []int, val int) int {
	searchWindowLeft, searchWindowRight := 0, len(arr)-1

	for searchWindowLeft <= searchWindowRight {
		middle := (searchWindowLeft + searchWindowRight) / 2

		if arr[middle] < val {
			searchWindowLeft = middle + 1
		} else {
			searchWindowRight = middle - 1
		}
	}

	return searchWindowRight + 1
}

func upperBound(array *[]int, target int) int {
	low, high, mid := 0, len(*array)-1, 0

	for low <= high {
		mid = (low + high) / 2
		if (*array)[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return low
}

func toInt64(s *string) int64 {
	value, err := strconv.ParseInt(*s, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	return value
}

func readline(r *bufio.Reader) (string, error) {
	var (
		isPrefix       = true
		err      error = nil
		line, ln []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}
	return string(ln), err
}

func reverse(s *[]int) {
	for i, j := 0, len(*s)-1; i < j; i, j = i+1, j-1 {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
}
 */
*/
