package singleton

import (
    "sync"
)

/**
 * @author Rancho
 * @date 2021/12/23
 */

var (
    instance *singleton
    once     sync.Once
)

type Singleton interface {
    foo()
}

type singleton struct {
    Name string
}

func (s singleton) foo() {

}

func GetInstance() Singleton {
    once.Do(func() {
        instance = &singleton{Name: "rancho"}
    })

    return instance
}
