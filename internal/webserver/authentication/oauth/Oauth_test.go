package oauth

import (
	"testing"

	"github.com/bisudoh/gokapi/internal/test"
	"github.com/bisudoh/gokapi/internal/webserver/authentication"
)

func TestSetCallbackCookie(t *testing.T) {
	w, _ := test.GetRecorder("GET", "/", nil, nil, nil)
	setCallbackCookie(w, "test")
	cookies := w.Result().Cookies()
	test.IsEqualInt(t, len(cookies), 1)
	test.IsEqualString(t, cookies[0].Name, authentication.CookieOauth)
	value := cookies[0].Value
	test.IsEqualString(t, value, "test")
}
