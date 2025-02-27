package main

import "testing"

func TestSampleInput1(t *testing.T) {
	unionFind := UnionFind{}
	unionFind.UnionFind(10)
	if unionFind.IsConnected(1, 3) {
		t.Error("component 1 and 3 should NOT connected")
	}
	unionFind.Connect(1, 8)
	unionFind.Connect(3, 8)
	if !unionFind.IsConnected(1, 3) {
		t.Error("component 1 and 3 should be connected")
	}
}

func TestSampleInput2(t *testing.T) {
	unionFind := UnionFind{}
	unionFind.UnionFind(10)
	if !unionFind.IsConnected(0, 0) {
		t.Error("component 1 and 3 should NOT connected")
	}
	unionFind.Connect(0, 1)
	unionFind.Connect(1, 2)
	unionFind.Connect(0, 2)

	if unionFind.IsConnected(0, 3) {
		t.Error("component 0 and 3 should NOT be connected")
	}
}

/*
func TestUnionFind(t *testing.T) {
	unionfind := NewUnionFind()
	unionfind.Union(4)
	if unionfind.components != 4 {
		t.Error("Size initialized is incorrect")
	}
}

func TestFindSet(t *testing.T) {
	unionfind := NewUnionFind()
	unionfind.Union(4)
	if unionfind.Find(1) != 1 {
		t.Error("Set cannot be found")
	}
}

func TestIsSameSet(t *testing.T) {
	unionfind := NewUnionFind()
	unionfind.Union(4)
	if !unionfind.IsConnected(1, 1) {
		t.Error("Same Set is not working")
	}
}

func TestIsUnionSet(t *testing.T) {
	unionfind := NewUnionFind()
	unionfind.Union(4)
	// 1 and 2 are not the same set
	unionfind.IsConnected(1, 2)
	unionfind.Connect(1, 2)
	if !unionfind.IsConnected(1, 2) {
		t.Error("Union operation is not working")
	}
}

func TestIsNumDisjointSets(t *testing.T) {
	unionfind := NewUnionFind()
	unionfind.Union(4)
	// 1 and 2 are not the same set
	unionfind.IsConnected(1, 2)
	unionfind.Connect(1, 2)
	if unionfind.Count() != 3 {
		t.Error("Wrong number of disjoint sets")
	}
}

func TestIsSizeOfSet(t *testing.T) {
	unionfind := NewUnionFind()
	unionfind.Union(4)
	// 1 and 2 are not the same set
	unionfind.IsConnected(1, 2)
	unionfind.Connect(1, 2)
	if unionfind.Size(1) != 2 {
		t.Error("Wrong number within set 1")
	}
}
*/
