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

// TestLoginCorrect whill pass when user input correct username and password
func (t *AuthTest) TestLoginCorrect() {
	authData := url.Values{"username": {"foobar"}, "pwd": {"foobar"}}
	t.PostForm("/api/auth/login", authData)
	t.AssertOk()
	t.AssertContains("200")
}

// TestLoginIncorrect will fail when user enter incorrect username and password
func (t *AuthTest) TestLoginIncorrect() {
	authData := url.Values{"username": {"incorrect"}, "password": {"incorrect"}}
	t.PostForm("/api/auth/login", authData)
	t.AssertStatus(200)
	t.AssertContains("403")
}
