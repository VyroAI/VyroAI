package infra

type Config struct {
	Provider provider `json:"provider"`
	Database database `json:"database"`
}

type database struct {
	Mysql mysql `json:"mysql"`
}

type mysql struct {
	URL string `json:"url"`
}

type provider struct {
	Discord   url `json:"discord"`
	Google    url `json:"google"`
	Apple     url `json:"apple"`
	Instagram url `json:"instagram"`
}

type url struct {
	LoginRedirectURL    string `json:"login_redirect_url"`
	RegisterRedirectURL string `json:"register_redirect_url"`
}
