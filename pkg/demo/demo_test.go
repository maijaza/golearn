package demo
import (
	"fmt"
	"testing"
)


// how to run per package go test .\pkg\demo -v
// -v mean show all case
func TestIntMin(t *testing.T){
    ans := IntMin(1,2)
	if ans != 1 {
		t.Errorf("Error %d", ans)
	}
}

func TestIntMinTableDriven(t *testing.T) {
    var tests = []struct {
        a, b int
        want int
    }{
        {0, 1, 0},
        {1, 0, 0},
        {2, -2, -2},
        {0, -1, -1},
        {-1, 0, -1},
    }

    for _, tt := range tests {

        testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
        t.Run(testname, func(t *testing.T) {		// t.Run for display each test name
            ans := IntMin(tt.a, tt.b)
            if ans != tt.want {
                t.Errorf("got %d, want %d", ans, tt.want)		// t.Errorf for disaply error msg
            }
        })
    }
}