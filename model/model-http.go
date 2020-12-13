package model

type TaskTypeHttp struct {
	Url     string            `yaml:"url" validate:"required,url"`
	Method  string            `yaml:"method" validate:"required,oneof=GET POST PUT DELETE"`
	Body    map[string]string `yaml:"body"`
	Timeout int               `yaml:"timeout" validate:"number"`
	Auth    *Auth             `yaml:"auth"`
}

type Auth struct {
	Basic       *AuthBasic        `yaml:"basic"`
	BearerToken *AuthBearerToken  `yaml:"bearer_token"`
	Custom      map[string]string `yaml:"custom"`
}

type AuthBasic struct {
	Username string `yaml:"username" validate:"required"`
	Password string `yaml:"password" validate:"required"`
}

type AuthBearerToken struct {
	Token string `yaml:"token" validate:"required"`
}
