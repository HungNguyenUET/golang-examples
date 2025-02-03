package main

import "fmt"

func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	// Create a DP table with (m+1) x (n+1) dimensions, initialized to false
	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}
	// Base case: Empty string matches empty pattern
	dp[0][0] = true

	// Handle patterns with '*' that can match zero occurrences of the preceding element
	for j := 2; j <= n; j++ {
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-2]
		}
	}

	// Fill the DP table
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == s[i-1] || p[j-1] == '.' {
				// If characters match or pattern has '.', take diagonal value
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				// '*' can either eliminate the preceding element or allow multiple occurrences
				dp[i][j] = dp[i][j-2] // Zero occurrences
				if p[j-2] == s[i-1] || p[j-2] == '.' {
					// If preceding character matches, '*' can match multiple times
					dp[i][j] = dp[i][j] || dp[i-1][j]
				}
			}
		}
	}

	// Final result indicating if s matches p
	return dp[m][n]
}

func main() {
	//fmt.Println(isMatch("abcd", "d*"))
	fmt.Println(isMatch("cabbbbcbcacbabc", ".*b.*.ab*.*b*a*c"))
}
