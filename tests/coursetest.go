package tests

import "github.com/revel/revel/testing"

type CourseTest struct {
	testing.TestSuite
}

func (t *CourseTest) TestThatIndexPageWorks() {
	t.Get("/api/courses")
	t.AssertOk()
	t.AssertContains("200")
}
