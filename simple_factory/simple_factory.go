package simple_factory

import (
    "fmt"
)

/**
 * @author Rancho
 * @date 2021/12/23
 */

type API interface {
    Say(name string) string
}

type hiAPI struct {
}

func (h hiAPI) Say(name string) string {
    return fmt.Sprintf("hi, %s", name)
}

type helloAPI struct {
}

func (h helloAPI) Say(name string) string {
    return fmt.Sprintf("hello, %s", name)
}

func NewAPI(apiType string) API {
    if apiType == "hi" {
        return hiAPI{}
    }
    if apiType == "hello" {
        return helloAPI{}
    }

    return nil
}
