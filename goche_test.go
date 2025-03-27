package goche_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/GrayHat12/goche"
	"github.com/GrayHat12/goche/strategy"
)

func AddWithDelay(numbers struct {
	a int
	b int
}) int {
	fmt.Printf("Sleeping for %d seconds\n", numbers.a+numbers.b)
	time.Sleep(time.Duration(numbers.a+numbers.b) * time.Second)
	return numbers.a + numbers.b
}

func TestGocheWithFifo(t *testing.T) {
	start := time.Now()
	// cachedAdd := CachedAddWithDelay(strategy.NewFifoStrategy)
	cachedAdd := goche.FunctionDecorator(AddWithDelay, 10, strategy.NewFifoStrategy, func(numbers struct {
		a int
		b int
	}) string {
		return fmt.Sprintf("%d+%d", numbers.a, numbers.b)
	})

	type Numbers struct {
		a int
		b int
	}

	// tenplus2 := cachedAdd(10, 2)
	// oneplus1 := cachedAdd(1, 1)
	// tenplus2_2 := cachedAdd(10, 2)
	// oneplus1_2 := cachedAdd(1, 1)

	// t.Logf("%.2fs: tenplus2 = %d\n",time.Since(start).Seconds(), tenplus2)
	// t.Logf("%.2fs: oneplus1 = %d\n",time.Since(start).Seconds(), oneplus1)
	// t.Logf("%.2fs: tenplus2_2 = %d\n",time.Since(start).Seconds(), tenplus2_2)
	// t.Logf("%.2fs: oneplus1_2 = %d\n",time.Since(start).Seconds(), oneplus1_2)

	t.Logf("%.2fs: 1+2=%d\n", time.Since(start).Seconds(), cachedAdd.Call(Numbers{a: 1, b: 2}))
	t.Logf("%.2fs: 1+2=%d\n", time.Since(start).Seconds(), cachedAdd.Call(Numbers{a: 1, b: 2}))
	t.Logf("%.2fs: 1+3=%d\n", time.Since(start).Seconds(), cachedAdd.Call(Numbers{a: 1, b: 3}))
	t.Logf("%.2fs: 1+4=%d\n", time.Since(start).Seconds(), cachedAdd.Call(Numbers{a: 1, b: 4}))
	t.Logf("%.2fs: 2+1=%d\n", time.Since(start).Seconds(), cachedAdd.Call(Numbers{a: 2, b: 1}))
	t.Logf("%.2fs: 3+1=%d\n", time.Since(start).Seconds(), cachedAdd.Call(Numbers{a: 3, b: 1}))
	t.Logf("%.2fs: 4+1=%d\n", time.Since(start).Seconds(), cachedAdd.Call(Numbers{a: 4, b: 1}))
	t.Logf("%.2fs: 2+1=%d\n", time.Since(start).Seconds(), cachedAdd.Call(Numbers{a: 2, b: 1}))
	t.Logf("%.2fs: 2+2=%d\n", time.Since(start).Seconds(), cachedAdd.Call(Numbers{a: 2, b: 2}))
	t.Logf("%.2fs: 2+3=%d\n", time.Since(start).Seconds(), cachedAdd.Call(Numbers{a: 2, b: 3}))
	t.Logf("%.2fs: 3+1=%d\n", time.Since(start).Seconds(), cachedAdd.Call(Numbers{a: 3, b: 1}))
	t.Logf("%.2fs: 3+2=%d\n", time.Since(start).Seconds(), cachedAdd.Call(Numbers{a: 3, b: 2}))
	t.Logf("%.2fs: 4+1=%d\n", time.Since(start).Seconds(), cachedAdd.Call(Numbers{a: 4, b: 1}))
	t.Logf("%.2fs: 5+0=%d\n", time.Since(start).Seconds(), cachedAdd.Call(Numbers{a: 5, b: 0}))
}
