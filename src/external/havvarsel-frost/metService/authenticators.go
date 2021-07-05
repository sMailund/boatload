package metService

import "net/http"

type iAuthenticator interface {
	authenticate(r *http.Request) error
}

type devAuthenticator struct {
}

func (d devAuthenticator) authenticate(r *http.Request) error {
	// no authentication necessary in dev
	r.Header.Set("authentication", "Bearer: actual bear")
	return nil
}

type havvarselAuthenticator struct {
}

func (h havvarselAuthenticator) authenticate(r *http.Request) error {
	panic("implement me")
}
