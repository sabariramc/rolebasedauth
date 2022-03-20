package constants

type AuthenticationType string
type LogoutType string

const (
	AuthTypeBasicAuth AuthenticationType = "BASIC_AUTH"
)

const (
	LogOutTypeUserLogout   LogoutType = "USER_LOGOUT"
	LogOutTypeUserInactive LogoutType = "USER_INACTIVE"
	LogOutTypeExpired      LogoutType = "EXPIRED"
)
