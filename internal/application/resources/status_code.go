package resources

import "net/http"

const (
	UNAUTHORIZED int = http.StatusUnauthorized
	VALIDATION   int = http.StatusBadRequest
	NOT_FOUND    int = http.StatusNotFound
	DUPLICATED   int = http.StatusConflict
	BAD_REQUEST  int = http.StatusBadRequest
	SERVER_ERROR int = http.StatusInternalServerError
)
