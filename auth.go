package nap

type AuthToken struct {
	Token string
}

type AuthBasic struct {
	Username string
	Password string
}

type Authentication interface {
	AuthorizationHeader() string
}
