package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type UnionFind struct {
	parent,
	rank,
	sz []int
	components int
}

func main() {

	scanner := getScanner()

	var numberOfElementsInBaseSet = 0
	var numberOfOperations = 0
	var unionFind = UnionFind{}
	var builder = strings.Builder{}
	var state = 0
	for scanner.Scan() {
		var err = scanner.Err()
		if err != nil {
			log.Fatal(err)
		}

		if state == 0 {
			var values = strings.Split(scanner.Text(), " ")
			numberOfElementsInBaseSet = toInt(&values[0])
			numberOfOperations = toInt(&values[1])
			unionFind.UnionFind(numberOfElementsInBaseSet)
			state = 1
		} else {
			var operations = strings.Split(scanner.Text(), " ")
			var firstComponent = toInt(&operations[1])
			var secondComponent = toInt(&operations[2])
			if operations[0] == "=" {
				unionFind.Connect(firstComponent, secondComponent)
			} else {
				if unionFind.IsConnected(firstComponent, secondComponent) {
					builder.WriteString("yes")
					builder.WriteString("\n")
				} else {
					builder.WriteString("no")
					builder.WriteString("\n")
				}
			}

			numberOfOperations--
			if numberOfOperations == 0 {
				fmt.Println(builder.String())
				builder.Reset()
				state = 0
			}

		}
	}

}

func getScanner() *bufio.Scanner {
	scanner := bufio.NewScanner(os.Stdin)
	buffer := make([]byte, 0, 64*1024) //Initial buffer for scanner set to 64KB
	scanner.Buffer(buffer, 1024*1024)  //Growing buffer for scanner set to 1MB
	return scanner
}

// Convert a string to an integer
func toInt(s *string) int {
	value, err := strconv.Atoi(*s)
	if err != nil {
		log.Fatal(err)
	}
	return value
}

// UnionFind / A disjoint-set data structure with union-by-rank and path compression
func (u *UnionFind) UnionFind(n int) {
	u.components = n
	u.rank = make([]int, n)
	u.sz = make([]int, n)
	u.parent = make([]int, n)
	for i := 0; i < n; i++ {
		u.sz[i] = 1     //Each component is initially of size 1
		u.rank[i] = 0   //Each component is initially of rank 0
		u.parent[i] = i //Each component is initially its own parent
	}
}

// Find the root of the component using path compression
func (u *UnionFind) Find(i int) int {
	if u.parent[i] == i { //If the parent is itself, then it is the root
		return i
	} else { //Compression path, recursively find the root of the parent
		u.parent[i] = u.Find(u.parent[i])
		return u.parent[i]
	}
}

// Connect two components and reset the ranks and sizes
func (u *UnionFind) Connect(i, j int) {
	if !u.IsConnected(i, j) {
		u.components = u.components - 1
		x := u.Find(i)             //Find the first component
		y := u.Find(j)             //Find the second component
		if u.rank[x] > u.rank[j] { //If the first component has a higher rank make it the parent of y
			u.parent[y] = x
			u.sz[x] += u.sz[y] //Increase the size of x with the size of y
		} else {
			u.parent[x] = y             // Otherwise make y the parent of x
			u.sz[y] += u.sz[x]          // Increase the size of y with the size of x
			if u.rank[x] == u.rank[y] { // If the ranks are equal increase the rank of y
				u.rank[y] = u.rank[y] + 1
			}
		}
	}
	return
}

// IsConnected determines if two components are connected
func (u *UnionFind) IsConnected(i, j int) bool {
	if u.Find(i) == u.Find(j) {
		return true
	}
	return false
}

// Count the number of components in the UnionFind data structure
func (u *UnionFind) Count() int {
	return u.components
}

// Size of a component
func (u *UnionFind) Size(i int) int {
	return u.sz[u.Find(i)]
}
