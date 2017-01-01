package tests

import (
	"github.com/revel/revel/testing"
)

// NewsTest is news controller test
type NewsTest struct {
	testing.TestSuite
}

// TestNewsIndexPage should ok
func (t *NewsTest) TestNewsIndexPage() {
	t.Get("/news")
	t.AssertOk()
	t.AssertContentType("text/json; charset=utf-8")
}
