# tbl

[![GoDoc](https://godoc.org/github.com/gdey/tbl?status.svg)](https://godoc.org/github.com/gdey/tbl).

This is a simple package to help with table driven tests. Given a set
of test cases, it will call the provided function each test case.

This helps remove boiler plate that comes with writing table driven code.


## Example

```go

func TestFoo(t *testing.T) {
    type testcase struct{
        p bool
       expected bool
    }
    test := tlb.Cases(
        testcaase{
            p : true,
            expected : true,
        },
   //…
    )
    test.Run(func(idx int, tc testcase)bool{
        if foo.IsGood(tc.p) != tc.expected  {
            t.Errorf("Test %v failed",idx)
        }
        return true // Run the next test case.
    })
}
```

# command line flags

In addition, the tool adds a new command line flag to help with debugging.

`--tblTest.RunOrder` : Allows one to specify the testcases's and the order they should run in.
This is usually helpful, when you are trying to fix one failing test, that you want to keep running
over and over again.