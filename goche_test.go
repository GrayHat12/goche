package goche_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/GrayHat12/goche"
	"github.com/GrayHat12/goche/strategy"
)

func AddWithDelay(a int, b int) int {
	// fmt.Printf("Sleeping for %d seconds\n", a+b)
	time.Sleep(time.Duration(a+b) * time.Second)
	return a + b
}

func CachedAddWithDelay(strategyGenerator goche.NewStrategyGenerator[int]) func(int, int) int {

	var cache = goche.NewCache(10, strategyGenerator) // will cache 10 elements at max

	return func(a int, b int) int {

		// Hashing logic
		hash := fmt.Sprintf("%d+%d", a, b)

		// try getting result from cache
		response, err := cache.Get(hash)

		// if response in cache, return cached value
		if err == nil {
			return response
		}

		// else, fetch the live response
		response = AddWithDelay(a, b)

		// store response in cache
		cache.Set(hash, response)
		return response
	}
}

func TestGocheWithFifo(t *testing.T) {
	start := time.Now()
	cachedAdd := CachedAddWithDelay(strategy.NewFifoStrategy)

	// tenplus2 := cachedAdd(10, 2)
	// oneplus1 := cachedAdd(1, 1)
	// tenplus2_2 := cachedAdd(10, 2)
	// oneplus1_2 := cachedAdd(1, 1)

	// t.Logf("%.2fs: tenplus2 = %d\n",time.Since(start).Seconds(), tenplus2)
	// t.Logf("%.2fs: oneplus1 = %d\n",time.Since(start).Seconds(), oneplus1)
	// t.Logf("%.2fs: tenplus2_2 = %d\n",time.Since(start).Seconds(), tenplus2_2)
	// t.Logf("%.2fs: oneplus1_2 = %d\n",time.Since(start).Seconds(), oneplus1_2)

	t.Logf("%.2fs: 1+1=%d\n", time.Since(start).Seconds(), cachedAdd(1, 1))
	t.Logf("%.2fs: 1+2=%d\n", time.Since(start).Seconds(), cachedAdd(1, 2))
	t.Logf("%.2fs: 1+3=%d\n", time.Since(start).Seconds(), cachedAdd(1, 3))
	t.Logf("%.2fs: 1+4=%d\n", time.Since(start).Seconds(), cachedAdd(1, 4))
	t.Logf("%.2fs: 2+1=%d\n", time.Since(start).Seconds(), cachedAdd(2, 1))
	t.Logf("%.2fs: 3+1=%d\n", time.Since(start).Seconds(), cachedAdd(3, 1))
	t.Logf("%.2fs: 4+1=%d\n", time.Since(start).Seconds(), cachedAdd(4, 1))
	t.Logf("%.2fs: 2+1=%d\n", time.Since(start).Seconds(), cachedAdd(2, 1))
	t.Logf("%.2fs: 2+2=%d\n", time.Since(start).Seconds(), cachedAdd(2, 2))
	t.Logf("%.2fs: 2+3=%d\n", time.Since(start).Seconds(), cachedAdd(2, 3))
	t.Logf("%.2fs: 3+1=%d\n", time.Since(start).Seconds(), cachedAdd(3, 1))
	t.Logf("%.2fs: 3+2=%d\n", time.Since(start).Seconds(), cachedAdd(3, 2))
	t.Logf("%.2fs: 4+1=%d\n", time.Since(start).Seconds(), cachedAdd(4, 1))
	t.Logf("%.2fs: 5+0=%d\n", time.Since(start).Seconds(), cachedAdd(5, 0))
}
