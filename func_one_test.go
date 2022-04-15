package func_one

import (
    "fmt"
    "testing"
    "time"
)

func TestRun(t *testing.T) {
    for i := 0; i < 100; i++ {
        go func() {
            fmt.Printf("start %v\n", time.Now().Nanosecond())
            res := Run("key", func() interface{} {
                time.Sleep(time.Second * 10)
                return time.Now().Nanosecond()
            })
            fmt.Printf("result: %v\n", res)
        }()
    }
    time.Sleep(time.Hour)
}
