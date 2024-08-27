// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package memo2_test

import (
	"testing"

	"github.com/memo2"
	"github.com/memotest"
)

var httpGetBody = memotest.HTTPGetBody

func Test(t *testing.T) {
	m := memo2.New(httpGetBody)
	memotest.Sequential(t, m)
}

func TestConcurrent(t *testing.T) {
	m := memo2.New(httpGetBody)
	memotest.Concurrent(t, m)
}
