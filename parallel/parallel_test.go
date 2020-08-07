package parallel

import (
	"fmt"
	"testing"
)

// func TestSomething(t *testing.T) {
// 	t.Parallel()

// }

// func TestA(t *testing.T) {
// 	t.Parallel()
// }

// func TestB(t *testing.T) {
// 	fmt.Println("setup")
// 	defer fmt.Println("deferred teardown")
// 	t.Run("group", func(t *testing.T) {
// 		t.Run("sub1", func(t *testing.T) {
// 			t.Parallel()
// 			// run sub1
// 			time.Sleep(time.Second)
// 			fmt.Println("sub1 done")
// 		})
// 		t.Run("sub2", func(t *testing.T) {
// 			t.Parallel()
// 			// run sub2
// 			time.Sleep(time.Second)
// 			fmt.Println("sub2 done")
// 		})
// 	})
// 	fmt.Println("teardown")
// }

// func TestGotcha(t *testing.T) {
// 	// Pretend this is our tables for a table-driven test
// 	for i := 0; i < 10; i++ {
// 		t.Run(fmt.Sprintf("i=%d", i), func(t *testing.T) {
// 			// need to copy i into closure, or else all the logs will show "Testing with i=10"
// 			t.Parallel()
// 			t.Logf("Testing with i=%d", i)
// 		})
// 	}
// }

func TestSquaredGotcha(t *testing.T) {
	testCases := []struct {
		arg  int
		want int
	}{
		{2, 5}, // Will pass if we don't copy values into closure
		{3, 9},
		{4, 16},
	}
	// Method 1 for handling: "Shadow" the variable so the parallel closure will have access to the current version of the testcase.
	// for _, tt := range testCases {
	// 	localTT := tt // Shadowing must happen any time before the call to t.Parallel
	// 	t.Run(fmt.Sprintf("arg=%d", localTT.arg), func(t *testing.T) {
	// 		// localTT := tt // Shadowing must happen any time before the call to t.Parallel
	// 		t.Parallel()
	// 		t.Logf("Testing with arg: %d, want: %d", localTT.arg, localTT.want)
	// 		if localTT.arg*localTT.arg != localTT.want {
	// 			t.Errorf("%d^2 != %d", localTT.arg, localTT.want)
	// 		}
	// 	})
	// }

	// // Method 2.A: Another closure, all together
	// for _, tt := range testCases {
	// 	t.Run(fmt.Sprintf("arg=%d", tt.arg), func(tt struct {
	// 		arg  int
	// 		want int
	// 	}) func(t *testing.T) {
	// 		return func(t *testing.T) {

	// 			t.Parallel()
	// 			t.Logf("Testing with arg: %d, want: %d", tt.arg, tt.want)
	// 			if tt.arg*tt.arg != tt.want {
	// 				t.Errorf("%d^2 != %d", tt.arg, tt.want)
	// 			}
	// 		}
	// 	}(tt))
	// }

	// Method 2.B: Another closure, broken down (refer to below for declarations)
	for _, tt := range testCases {
		t.Run(fmt.Sprintf("arg=%d", tt.arg), test(tt))
	}
}

type testCase struct {
	arg  int
	want int
}

func test(tt testCase) func(t *testing.T) {
	return func(t *testing.T) {
		t.Parallel()
		t.Logf("Testing with arg: %d, want: %d", tt.arg, tt.want)
		if tt.arg*tt.arg != tt.want {
			t.Errorf("%d^2 != %d", tt.arg, tt.want)
		}
	}
}
