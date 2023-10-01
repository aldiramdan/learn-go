package src

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func RoundLooping(val *float64, vel *uint) string {
	d := int(*vel)               // decimal
	q := fmt.Sprintf("%v", *val) // convert to string
	p := strings.Split(q, ".")   // split

	// don't have decimal
	if len(p) <= 1 {
		z := strings.Join(p, "")                             // join array index round
		r, _ := strconv.ParseFloat(strings.TrimSpace(z), 64) // convert to float64
		v := strconv.FormatFloat(r, 'f', d+1, 64)            // convert to string
		return v
	}

	s := strings.Split(p[0], "") // split index round
	o := strings.Split(p[1], "") // split index decimal

	// create array int for index round
	arr := make([]int, len(s))
	for i := range arr {
		arr[i], _ = strconv.Atoi(s[i])
	}

	// create array int for index decimal
	ary := make([]int, len(o))
	for i := range ary {
		ary[i], _ = strconv.Atoi(o[i])
	}

	// add decimal for looping
	if len(ary) <= d {
		for i := 0; i <= d; i++ {
			ary = append(ary, 0)
		}
	}

	// looping
	if ary[d] >= 5 {
		// condition for decimal 0
		if d == 0 {
			// looping for index round
			for i := len(arr) - 1; i >= 0; i-- {
				if i == len(arr)-1 {
					arr[i] = arr[i] + 1
				}
			}

			// looping before last index round if == 10
			for i := len(arr) - 1; i >= 0; i-- {
				if arr[i] == 10 && i > 0 {
					arr[i] = 0
					arr[i-1] = arr[i-1] + 1
				}
			}

			// looping for decimal to zero
			for i := 0; i < len(ary); i++ {
				ary[i] = 0
			}
		} else {
			ary[d-1] = ary[d-1] + 1
		}

		// looping before last index decimal if == 10
		for i := len(ary) - 1; i >= 0; i-- {
			if ary[i] == 10 && i > 0 {
				ary[i] = 0
				ary[i-1] = ary[i-1] + 1
			}
		}

		// condition index 0 decimal to round
		if ary[0] == 10 {
			for i := len(arr) - 1; i >= 0; i-- {
				if i == len(arr)-1 {
					arr[i] = arr[i] + 1
				}
			}

			// looping before last index round if == 10
			for i := len(arr) - 1; i >= 0; i-- {
				if arr[i] == 10 && i > 0 {
					arr[i] = 0
					arr[i-1] = arr[i-1] + 1
				}
			}

			// looping for decimal to zero
			for i := 0; i < len(ary); i++ {
				ary[i] = 0
			}
		}

		// looping for decimal to zero
		for i := d; i < len(ary); i++ {
			ary[i] = 0
		}
	} else if ary[d] < 5 {
		for i := d; i < len(ary); i++ {
			ary[i] = 0
		}
	}

	// create array str for index round
	strr := make([]string, len(arr))
	for i, x := range arr {
		strr[i] = strconv.Itoa(x)
	}

	// create array str for index decimal
	strs := make([]string, len(ary))
	for i, x := range ary {
		strs[i] = strconv.Itoa(x)
	}

	z := strings.Join(strr, "") // join array index round
	y := strings.Join(strs, "") // join array index decimal

	// negative float
	if *val < 0 {
		x := fmt.Sprintf("-%v.%v", z, y)                     // concat index round and index decimal
		r, _ := strconv.ParseFloat(strings.TrimSpace(x), 64) // convert to float64
		v := strconv.FormatFloat(r, 'f', d+1, 64)            // convert to string

		return v
	}

	x := fmt.Sprintf("%v.%v", z, y)                      // concat index round and index decimal
	r, _ := strconv.ParseFloat(strings.TrimSpace(x), 64) // convert to float64
	v := strconv.FormatFloat(r, 'f', d+1, 64)            // convert to string
	return v
}

func RoundFloat(val *float64, vel *uint) string {
	q := math.Round(*val*math.Pow(10, float64(*vel))) / math.Pow(10, float64(*vel))
	p := strconv.FormatFloat(q, 'f', int(*vel)+1, 64)

	return p
}
