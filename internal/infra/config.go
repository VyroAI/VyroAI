package infra

type Config struct {
	Provider Provider `yaml:"provider"`
	Database Database `yaml:"database"`
}

type Provider struct {
	Discord   OAuthProvider `yaml:"discord"`
	Google    OAuthProvider `yaml:"google"`
	Apple     OAuthProvider `yaml:"apple"`
	Instagram OAuthProvider `yaml:"instagram"`
}

type OAuthProvider struct {
	LoginRedirectURL    string `yaml:"loginRedirectUrl"`
	RegisterRedirectURL string `yaml:"registerRedirectUrl"`
}

type Database struct {
	MySQL MySQLConfig `yaml:"mysql"`
}

type MySQLConfig struct {
	URL string `yaml:"url"`
}
