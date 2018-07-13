package numplan

import (
	//	"fmt"
	"strconv"
	"strings"
)

func appendDigit(str string, digit string, count int) string {
	result := []string{str}
	for i := 0; i < count-len(str); i++ {
		result = append(result, digit)
	}

	return strings.Join(result, "")
}

func range2prefix(start, stop string) []string {
	//Mihails Koreshkovs code
	var prefix string
	i := len(start) - 1
	for ; i >= 0; i-- {
		prefix = start[0:i]
		if strings.HasPrefix(stop, prefix) {
			break
		}
	}

	start_d, _ := strconv.Atoi(start[i : i+1])
	stop_d, _ := strconv.Atoi(stop[i : i+1])

	var prefix_list []string
	for ; start_d <= stop_d; start_d++ {
		prefix_list = append(prefix_list, prefix+strconv.Itoa(start_d))
	}
	//	fmt.Println(prefix)
	//	fmt.Println(prefix_list)

	l := len(prefix_list)
	if l >= 2 {
		if appendDigit(prefix_list[0], "0", len(start)) != start {
			t1 := range2prefix(start, appendDigit(prefix_list[0], "9", len(start)))
			prefix_list = append(t1, prefix_list[1:]...)
		}

		l = len(prefix_list)
		if appendDigit(prefix_list[l-1], "9", len(stop)) != stop {
			t2 := range2prefix(appendDigit(prefix_list[l-1], "0", len(stop)), stop)
			prefix_list = append(prefix_list[:l-1], t2...)
		}

	}
	//code for create prefix smoller ex. 10,11,12,13,14,15,16,17,18,19 => 1
	if len(prefix_list) == 10 {
		tl := (len(prefix_list[0]))
		do_next := true
		for _, p := range prefix_list {
			if tl != len(p) {
				do_next = false
			}
		}
		if do_next {
			if (prefix_list[0][tl-1:] == "0") && (prefix_list[1][tl-1:] == "1") && (prefix_list[2][tl-1:] == "2") && (prefix_list[3][tl-1:] == "3") && (prefix_list[4][tl-1:] == "4") && (prefix_list[5][tl-1:] == "5") && (prefix_list[6][tl-1:] == "6") && (prefix_list[7][tl-1:] == "7") && (prefix_list[8][tl-1:] == "8") && (prefix_list[9][tl-1:] == "9") {
				return []string{prefix_list[0][:tl-1]}
			}
		}
	}
	//fmt.Println("pfx=", prefix_list)
	return prefix_list

}

//func list2prefix(r []diapazon) (p_list []string) {
//	for _, d := range r {
//		a := strconv.Itoa(d.start)
//		b := strconv.Itoa(d.stop)
//		p_list = append(p_list, range2prefix(a, b)...)
//	}
//	return p_list
//}
