package model

type TaskTypeHttp struct {
	Method  string `yaml:"method" validate:"required,oneof=GET POST PUT DELETE"`
	Url     string `yaml:"url" validate:"required,url"`
	Timeout int    `yaml:"timeout" validate:"number"`
	Auth    *Auth  `yaml:"auth"`
}

type Auth struct {
	Basic       *AuthBasic       `yaml:"basic"`
	BearerToken *AuthBearerToken `yaml:"bearer_token"`
}

type AuthBasic struct {
	Username string `yaml:"username" validate:"required"`
	Password string `yaml:"password" validate:"required"`
}

type AuthBearerToken struct {
	Token string `yaml:"token" validate:"required"`
}
