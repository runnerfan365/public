package main

import (
	"fmt"
	"sort"
	"strconv"
)

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}
func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s sortRunes) Len() int {
	return len(s)
}
func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

func rollUp(s string) string {
	var temp_str, res_str string
	var ch_count int = 0
	var ch rune
	temp_str = SortString(s)
	for _, tch := range temp_str {
		if ch != tch {
			if ch_count > 0 {
				res_str = res_str + strconv.Itoa(ch_count)
				ch_count = 0
			}
			ch = tch
			res_str = res_str + string(tch)
			ch_count = ch_count + 1
		} else {
			ch_count = ch_count + 1
		}
	}
	res_str = res_str + strconv.Itoa(ch_count)
	return res_str
}

func main() {
	s1 := "bcadaabddd"
	fmt.Println(rollUp(s1))
}
