package simple_factory

import (
    "testing"

    "github.com/stretchr/testify/assert"
)

/**
 * @author Rancho
 * @date 2021/12/23
 */

func TestNewAPI(t *testing.T) {
    api := NewAPI("hi")
    s := api.Say("rancho")
    assert.Equal(t, "hi, rancho", s)

    api = NewAPI("hello")
    s = api.Say("rancho")
    assert.Equal(t, "hello, rancho", s)
}
