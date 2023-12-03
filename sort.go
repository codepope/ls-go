package main

import "sort"

// BySize tells `sort.Sort` how to sort by file size
type BySize []*DisplayItem

func (s BySize) Less(i, j int) bool {
	return s[i].info.Size() < s[j].info.Size()
}
func (s BySize) Len() int {
	return len(s)
}
func (s BySize) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// ByTime tells `sort.Sort` how to sort by last modified time
type ByTime []*DisplayItem

func (s ByTime) Less(i, j int) bool {
	return s[i].info.ModTime().Unix() < s[j].info.ModTime().Unix()
}
func (s ByTime) Len() int {
	return len(s)
}
func (s ByTime) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// ByKind tells `sort.Sort` how to sort by file extension
type ByKind []*DisplayItem

func (s ByKind) Less(i, j int) bool {
	var kindi, kindj string

	if s[i].IsHidden() {
		kindi = "." + s[i].ext
	} else if s[i].ext == "" {
		kindi = "0"
	} else {
		kindi = s[i].ext
	}

	if s[j].IsHidden() {
		kindj = "." + s[j].ext
	} else if s[j].ext == "" {
		kindj = "0"
	} else {
		kindj = s[j].ext
	}

	if kindi == kindj {
		return s[i].basename < s[j].basename
	}
	return kindi < kindj
}
func (s ByKind) Len() int {
	return len(s)
}
func (s ByKind) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func reverse(s []*DisplayItem) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// ByMetaTag tells `sort.Sort` how to sort by a named extension

type ByMetaTag func(d1, d2 *DisplayItem) bool

func (by ByMetaTag) Sort(displayItems []*DisplayItem) {
	ps := &displayItemSorter{
		displayItems: displayItems,
		by:           by, // The Sort method's receiver is the function (closure) that defines the sort order.
	}
	sort.Sort(ps)
}

// displayItemSorter joins a By function and a slice of DisplayItems to be sorted.
type displayItemSorter struct {
	displayItems []*DisplayItem
	by           func(p1, p2 *DisplayItem) bool // Closure used in the Less method.
}

// Len is part of sort.Interface.
func (s *displayItemSorter) Len() int {
	return len(s.displayItems)
}

// Swap is part of sort.Interface.
func (s *displayItemSorter) Swap(i, j int) {
	s.displayItems[i], s.displayItems[j] = s.displayItems[j], s.displayItems[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (s *displayItemSorter) Less(i, j int) bool {
	return s.by(s.displayItems[i], s.displayItems[j])
}
