package main

import (
	"bytes"
	"strings"
)

func main() {
	//addBinary("11", "1")
	addBinary("1010", "1011")
}
func addBinary2(a string, b string) string {
	// Ensure a is the longer string
	if len(a) < len(b) {
		a, b = b, a
	}

	// Prepend zeros to b if necessary
	b = b + strings.Repeat("0", len(a)-len(b))

	var carry byte = 0
	var result []byte

	// Iterate through each character from right to left
	for i := len(a) - 1; i >= 0; i-- {
		sum := (a[i] - '0') + (b[i] - '0') + carry
		carry = sum / 2
		result = append([]byte{sum%2 + '0'}, result...)
	}

	// If there's still a carry after the loop, add it to the result
	if carry > 0 {
		result = append([]byte{'1'}, result...)
	}

	return string(result)
}
func addBinary(a string, b string) string {
	var answer []byte
	l1 := len(a)
	l2 := len(b)
	if l1 < l2 {
		a = strings.Repeat("0", l2-l1) + a
	} else {
		b = strings.Repeat("0", l1-l2) + b
	}
	sign := 0
	for i := len(a) - 1; i >= 0; i-- {
		if a[i] == '1' && b[i] == '1' {
			if sign == 0 {
				answer = append(answer, '0')
			} else {
				answer = append(answer, '1')
			}
			sign = 1
		} else if a[i] == '0' && b[i] == '0' {
			if sign == 0 {
				answer = append(answer, '0')
			} else {
				answer = append(answer, '1')
				sign = 0
			}
		} else {
			if sign == 0 {
				answer = append(answer, '1')
			} else {
				answer = append(answer, '0')
				sign = 1
			}
		}
	}
	if sign == 1 {
		answer = append(answer, '1')
	}
	var buffer bytes.Buffer
	for i := len(answer) - 1; i >= 0; i-- {
		buffer.WriteByte(answer[i])
	}
	return buffer.String()
}
