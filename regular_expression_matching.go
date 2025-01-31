package main

import "fmt"

func isMatch(s string, p string) bool {
	for i := 0; i < len(p); i++ {
		if isMatchByStart(s, p, i) {
			return true
		}
	}

	return false
}

func isMatchByStart(s string, p string, start int) bool {
	j := start

	if j == len(p)-1 {
		return false
	}

	if p[j] == '*' {
		return false
	}

	for i := 0; i < len(s); i++ {
		if p[j] == '*' {
			pre := p[j-1]

			for i < len(s) && (s[i] == pre || pre == '.') {
				i++
			}

			if i == len(s) {
				return true
			} else {
				i--
			}
		} else if p[j] != '.' && s[i] != p[j] {
			return false
		}

		if i == len(s)-1 {
			return true
		} else if j == len(p)-1 {
			return false
		}

		j++
	}

	return true
}

func main() {
	//fmt.Println(isMatch("hello", "hello"))
	//fmt.Println(isMatch("hello", ".ello"))
	//fmt.Println(isMatch("hello", "hhello"))
	fmt.Println(isMatch("ab", ".*c"))
	//fmt.Println(isMatch("aa", "a*"))
}
