package resources

const (
	DB         string = "DB_001"
	DB_ENQUEUE string = "DB_QUEUE_001" // other process is working with id:'%s' e:'%s'.
	INTERNAL   string = "INTERNAL_001" // (generic)

	CREATE_TOKEN         string = "TOKEN_001"     // Cannot create access token.
	CREATE_REFRESH_TOKEN string = "TOKEN_002"     // Cannot create refresh token.
	CREATE_USER_REQUEST  string = "CREA_USER_001" // Field: '%s' failed on '%s' validation.
	CREATE_USER_EXIST    string = "CREA_USER_002" // User: '%s' already exists.

	USER_CONFIRM_TOKEN_EMPTY  string = "CONFIRM_USER_001" // Token must not be emtpy.
	USER_CONFIRM_TOKEN_VALID  string = "CONFIRM_USER_002" // Token must be valid.
	USER_CONFIRM_TOKEN_ACTIVE string = "CONFIRM_USER_003" // Token must be active.
	USER_CONFIRM_EXIST        string = "CONFIRM_USER_004" // User must exists.
	USER_CONFIRM_TOKEN        string = "CONFIRM_USER_005" // Token is required.
	USER_INACTIVE             string = "CONFIRM_USER_006" // User must be inactive.

	LOGIN_USER_REQUEST  string = "LOGIN_USER_001" // Field: '%s' failed on '%s' validation.
	LOGIN_USER_EXISTS   string = "LOGIN_USER_002" // User do not exists.
	LOGIN_USER_ACTIVE   string = "LOGIN_USER_003" // User must be active.
	LOGIN_USER_LOCK     string = "LOGIN_USER_004" // User must not be lock.
	LOGIN_USER_PASSWORD string = "LOGIN_USER_005" // User or password is wrong.

	REFRESH_LOGIN_TOKEN       string = "REFRESH_LOGIN_001" // Refresh token must be provided.
	REFRESH_LOGIN_TOKEN_VALID string = "REFRESH_LOGIN_002" // Refresh token must be valid.

	PASSWORD_CHECK string = "PASSWORK_CHECK_001" // Password must meet 'One upper, one lower, one number, one special @#$%&, length: 8-16'

	CHANGE_PASS_CURRENT     string = "CHANGE_PASS_001" // Current password is required.
	CHANGE_PASS_USER        string = "CHANGE_PASS_002" // User ID is required.
	CHANGE_PASS_USER_EXISTS string = "CHANGE_PASS_003" // User must exists.
	CHANGE_PASS_WRONG       string = "CHANGE_PASS_004" // Password is wrong.
	CHANGE_PASS_USER_ACTIVE string = "CHANGE_PASS_005" // User must be active.
	CHANGE_PASS_USER_LOCK   string = "CHANGE_PASS_006" // User must not be lock.

	FORGOT_PASSWORD_STRUCT string = "FORGOT_PASS_001" // Field: '%s' failed on '%s' validation.
	FORGOT_PASSWORD_USER   string = "FORGOT_PASS_002" // User not found.
	FORGOT_PASSWORD_ACTIVE string = "FORGOT_PASS_003" // User must be active.
)
