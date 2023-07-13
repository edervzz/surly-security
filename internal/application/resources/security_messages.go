package resources

const (
	EN string = "en"
	ES string = "es"
)

// code_message: [ en: text, es: text, fr: text ]
var SecurityMessages map[string]map[string]string = map[string]map[string]string{
	DB_ENQUEUE: { // DB_QUEUE_001
		EN: "other process is working with id:'%s' e:'%s'.",
		ES: "other process is working with id:'%s' e:'%s'.",
	},
	CREATE_TOKEN: { // TOKEN_001
		EN: "Cannot create access token.",
		ES: "Imposible crear access token.",
	},
	CREATE_REFRESH_TOKEN: { // TOKEN_002
		EN: "Cannot create refresh token.",
		ES: "Imposible crear refresh token.",
	},
	CREATE_USER_REQUEST: { // CREA_USER_001
		EN: "Field '%s' failed on '%s' validation.",
		ES: "Campo '%s' falló validación '%s'.",
	},
	CREATE_USER_EXIST: { // CREA_USER_002
		EN: "User '%s' already exists.",
		ES: "El usuario '%s' ya esta registrado.",
	},
	USER_CONFIRM_TOKEN_EMPTY: { // CONFIRM_USER_001
		EN: "Token must not be emtpy.",
		ES: "Token requerido.",
	},
	USER_CONFIRM_TOKEN_VALID: { // CONFIRM_USER_002
		EN: "Token is no longer valid.",
		ES: "Token invalido.",
	},
	USER_CONFIRM_TOKEN_ACTIVE: { // CONFIRM_USER_003
		EN: "Token must be active.",
		ES: "Token inactivo.",
	},
	USER_CONFIRM_EXIST: { // CONFIRM_USER_004
		EN: "User must exists.",
		ES: "Usuario no encontrado.",
	},
	USER_CONFIRM_TOKEN: { // CONFIRM_USER_005
		EN: "Token is required.",
		ES: "Token es requerido.",
	},
	USER_INACTIVE: { // CONFIRM_USER_006
		EN: "User must be inactive.",
		ES: "Usuario debe estar inactivo.",
	},
	LOGIN_USER_REQUEST: { // LOGIN_USER_001
		EN: "Field: '%s' failed on '%s' validation.",
		ES: "Campo '%s' falló validación '%s'.",
	},
	LOGIN_USER_EXISTS: { // LOGIN_USER_002
		EN: "User do not exists.",
		ES: "El usuario no existe.",
	},
	LOGIN_USER_ACTIVE: { // LOGIN_USER_003
		EN: "User must be active.",
		ES: "El usuario debe estar activo.",
	},
	LOGIN_USER_LOCK: { // LOGIN_USER_004
		EN: "User must not be lock.",
		ES: "El usuario no debe estar bloqueado.",
	},
	LOGIN_USER_PASSWORD: { // LOGIN_USER_005
		EN: "User or password is wrong.",
		ES: "El usuario o contraseña esta incorrecto.",
	},
	REFRESH_LOGIN_TOKEN: { // REFRESH_LOGIN_001
		EN: "'Refresh token' must be provided.",
		ES: "'Refresh token' es requerido",
	},
	REFRESH_LOGIN_TOKEN_VALID: { // REFRESH_LOGIN_002
		EN: "Refresh token is not longer valid.",
		ES: "Refresh token ya no es valido.",
	},
	PASSWORD_CHECK: { // PASSWORK_CHECK_001
		EN: "Password must meet 'One upper, one lower, one number, one special @#$%&, length: 8-16'",
		ES: "Constraseña requiere al menos 'una mayúscula, una minúscula, un número, un caracter especial @#$%& y longitud: 8-16'",
	},
	CHANGE_PASS_CURRENT: { // CHANGE_PASS_001
		EN: "Current password is required.",
		ES: "Constraseña actual requerida.",
	},
	CHANGE_PASS_USER: { // CHANGE_PASS_002
		EN: "User ID is required.",
		ES: "ID Usuario es requerido.",
	},
	CHANGE_PASS_USER_EXISTS: { // CHANGE_PASS_003
		EN: "User must exists.",
		ES: "El usuario debe existir.",
	},
	CHANGE_PASS_WRONG: { // CHANGE_PASS_004
		EN: "Password is wrong.",
		ES: "Constraseña incorrecta.",
	},
	CHANGE_PASS_USER_ACTIVE: { // CHANGE_PASS_005
		EN: "User must be active.",
		ES: "El usuario debe estar activo.",
	},
	CHANGE_PASS_USER_LOCK: { // CHANGE_PASS_006
		EN: "User must not be lock.",
		ES: "Usuario no debe estar bloqueado.",
	},
	FORGOT_PASSWORD_STRUCT: { // FORGOT_PASS_001
		EN: "Field: '%s' failed on '%s' validation.",
		ES: "Campo '%s' falló validación '%s'.",
	},
	FORGOT_PASSWORD_USER: { // FORGOT_PASS_002
		EN: "User not found.",
		ES: "Usuario no encontrado.",
	},
	FORGOT_PASSWORD_ACTIVE: { // FORGOT_PASS_003
		EN: "User must be active.",
		ES: "El usuario debe estar activo.",
	},
}
