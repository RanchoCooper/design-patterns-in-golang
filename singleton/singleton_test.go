package singleton

import (
    "sync"
    "testing"

    "github.com/stretchr/testify/assert"
)

/**
 * @author Rancho
 * @date 2021/12/23
 */

const parCount = 100

func TestGetInstance(t *testing.T) {
    ins1 := GetInstance()
    ins2 := GetInstance()
    assert.True(t, ins1 == ins2)
}

func TestParallelSingleton(t *testing.T) {
    start := make(chan struct{})
    wg := sync.WaitGroup{}
    wg.Add(parCount)
    instances := [parCount]Singleton{}

    for i := 0; i < parCount; i++ {
        go func(index int) {
            <-start
            instances[index] = GetInstance()
            wg.Done()
        }(i)
    }

    close(start)
    wg.Wait()

    for i := 1; i < parCount; i++ {
        assert.True(t, instances[i] == instances[i-1])
    }
}
