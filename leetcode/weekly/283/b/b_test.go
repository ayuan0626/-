// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test_b(t *testing.T) {
	examples := [][]string{
		{
			`[1,4,25,10,25]`, `2`, 
			`5`,
		},
		{
			`[5,6]`, `6`, 
			`25`,
		},
		{
			`[1,2]`, `3`,
			`12`,
		},
		{
			`[96,44,99,25,61,84,88,18,19,33,60,86,52,19,32,47,35,50,94,17,29,98,22,21,72,100,40,84]`, `35`,
			`794`,
		},
	}
	targetCaseNum :=  -1
	if err := testutil.RunLeetCodeFuncWithExamples(t, minimalKSum, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-283/problems/append-k-integers-with-minimal-sum/