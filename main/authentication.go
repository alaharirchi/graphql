package main
import (
	"errors"
)
const MOCKED_TOKEN_RESEARCHER = "mocked_token_researcher_123"
const MOCKED_TOKEN_USER = "mocked_token_user_789"

// Struct for what the auth function returns:
type TokenInfo struct {
	id int
	roles []string
}

// Defined type for key of a TokenInfo object
// This is to be used in the Value map of the context
type tokenInfoContextKey string


// Go has no generic search function for slices or arrays
// So I had to write this helper function
func (tokenInfo TokenInfo) isResearcher() bool {
    for _, role := range tokenInfo.roles {
        if role == RESEARCHER_ROLE {
        	return true
        }
    }
    return false
}

// Auth function: (remember, it returns two things: 1)if the request was valid 2) encoded information such as role, etc.)
func validateToken(token string) (tokenInfo TokenInfo, err error) {
	if token == MOCKED_TOKEN_RESEARCHER {
		var tokenInfo = TokenInfo{
			id: 4321,
			roles: []string{RESEARCHER_ROLE, USER_ROLE},
		}
		return tokenInfo, nil
	}

	if token == MOCKED_TOKEN_USER {
		var tokenInfo = TokenInfo{
			id: 4321,
			roles: []string{USER_ROLE},
		}
		return tokenInfo, nil
	}

	return TokenInfo{}, errors.New("Invalid authentication token") // TODO: Is this actually an error? Maybe not, because we people are able to do certain actions without logging in...
}