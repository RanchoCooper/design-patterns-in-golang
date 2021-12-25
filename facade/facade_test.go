package facade

import (
    "testing"
)

/**
 * @author Rancho
 * @date 2021/12/25
 */

var expect = "A module running\nB module running"

func TestFacadeAPI(t *testing.T) {
    api := NewAPI()
    ret := api.Test()
    if ret != expect {
        t.Fatalf("expect %s, return %s", expect, ret)
    }
}
