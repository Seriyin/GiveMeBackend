package sessions

import (
	"github.com/gorilla/sessions"
	"os"
)

var (
	SessionStore sessions.Store
)

func init() {
	cookie_secret := os.Getenv("COOKIE_SECRET")
	// [START sessions]
	// Configure storage method for session-wide information.
	// Update "something-very-secret" with a hard to guess string or byte sequence.
	cookieStore := sessions.NewCookieStore(
		[]byte(cookie_secret),
	)
	cookieStore.Options = &sessions.Options{
		HttpOnly: true,
	}
	SessionStore = cookieStore
	// [END sessions]
}
