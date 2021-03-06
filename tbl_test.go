// Copyright 2016 Gautam Dey. All rights reserved.
// Use of this source code is governed by FreeBDS License (2-clause Simplified BSD.)
// that can be found in the LICENSE file.

package tbltest_test

import (
	"testing"

	"github.com/gdey/tbltest"
)

func TestCases(t *testing.T) {
	type testcase struct {
		val  int
		next bool
	}
	test := tbltest.Cases(
		testcase{
			val:  0,
			next: true,
		},
		testcase{
			val:  1,
			next: false,
		},
	)
	test.AddCases(
		testcase{
			val:  2,
			next: true,
		},
	)
	count := test.Run(func(tc testcase) {})
	if count != 3 {
		t.Errorf("did not run all the testcases.")
	}

	count = test.Run(func(idx int, tc testcase) {
		if tc.val != idx {
			t.Errorf("for test %v: expected %[1]v, got %v", idx, tc.val)
		}
	})
	if count != 3 {
		t.Errorf("did not run all the testcases.")
	}

	test.InOrder = true
	count = test.Run(func(tc testcase) bool {
		return tc.next
	})
	if count != 2 {
		t.Errorf("expected to only run two test. ran %v instead", count)
	}
	count = test.Run(func(idx int, tc testcase) bool {
		if tc.val != idx {
			t.Errorf("for test %v: expected %[1]v, got %v", idx, tc.val)
		}
		return tc.next
	})
	if count != 2 {
		t.Errorf("expected to only run two test. ran %v instead", count)
	}
}

func TestIntCases(t *testing.T) {
	test := tbltest.Cases(0, 1, 2, 3)
	count := test.Run(func(tc int) {})
	if count != 4 {
		t.Errorf("did not run all the testcases.")
	}
}

func TestNilFunc(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("Failed to run nil fun function: %v", r)
		}
	}()
	type testcase struct{}
	test := tbltest.Cases(testcase{}, testcase{})
	test.Run(nil)
}
