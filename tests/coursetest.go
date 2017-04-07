package tests

import "github.com/revel/revel/testing"
import "net/url"

type CourseTest struct {
	testing.TestSuite
}

func (t *CourseTest) TestThatIndexPageWorks() {
	t.Get("/api/courses")
	t.AssertOk()
	t.AssertContains("200")
}

func (t *CourseTest) TestAddCourses() {

	coursesData := url.Values{
		"name": {"test course"},
		"desc": {"test course desc"},
	}

	t.PostForm("/api/courses", coursesData)
	t.AssertOk()
	t.AssertContains("200")
	t.AssertContains("test course")
}
