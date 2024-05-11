package indigo

import "errors"

const (
	textPleaseCheckClientCredentialsEnv = "Please check the environment variable " + WEBARENA_INDIGO_CLIENT_ID + " and " + WEBARENA_INDIGO_CLIENT_SECRET
)

var (
	ErrUnexpectedStatusCode     = errors.New("indigo: unexpected status code")
	ErrAPIReturnsTooManyRequest = errors.New("indigo: API returns Too Many Request")
	ErrAPIReturnsUnauthorized   = errors.New("indigo: API returns Unauthorized. " + textPleaseCheckClientCredentialsEnv)
	ErrInvalidClientCredentials = errors.New("indigo: invalid client credentials. " + textPleaseCheckClientCredentialsEnv)
)
