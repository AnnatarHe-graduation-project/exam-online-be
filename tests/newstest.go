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
	t.Get("/api/news/list")
	t.AssertOk()
	t.AssertContentType("application/json; charset=utf-8")
}

// TestNewsTrendingsPage should ok
func (t *NewsTest) TestNewsTrendingsPage() {
	t.Get("/api/news/trendings")
	t.AssertOk()
	t.AssertContentType("application/json; charset=utf-8")
}
