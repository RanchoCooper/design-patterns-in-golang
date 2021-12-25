package adapter

import (
    "testing"
)

/**
 * @author Rancho
 * @date 2021/12/25
 */

var expect = "adaptee method"

func TestAdapter(t *testing.T) {
    adaptee := NewAdaptee()
    target := NewAdapter(adaptee)
    res := target.Request()
    if res != expect {
        t.Fatalf("expect: %s, actual: %s", expect, res)
    }
}
