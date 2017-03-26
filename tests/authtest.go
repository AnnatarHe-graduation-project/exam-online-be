package tests

import "github.com/revel/revel/testing"
import "net/url"

type AuthTest struct {
	testing.TestSuite
}

func (t *AuthTest) TestRegister() {
	postData := url.Values{
		"username":  {"foobar"},
		"pwd":       {"foobar"},
		"school_id": {"1"},
		"role":      {"21"},
	}
	t.PostForm("/api/auth/register", postData)
	t.AssertOk()
}
