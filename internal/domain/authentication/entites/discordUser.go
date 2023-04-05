package entites

type DiscordUser struct {
	Id               string      `json:"id"`
	Username         string      `json:"username"`
	GlobalName       interface{} `json:"global_name"`
	DisplayName      interface{} `json:"display_name"`
	Avatar           string      `json:"avatar"`
	AvatarDecoration interface{} `json:"avatar_decoration"`
	Discriminator    string      `json:"discriminator"`
	PublicFlags      int         `json:"public_flags"`
	Flags            int         `json:"flags"`
	Banner           interface{} `json:"banner"`
	BannerColor      interface{} `json:"banner_color"`
	AccentColor      interface{} `json:"accent_color"`
	Locale           string      `json:"locale"`
	MfaEnabled       bool        `json:"mfa_enabled"`
	PremiumType      int         `json:"premium_type"`
	Email            string      `json:"email"`
	Verified         bool        `json:"verified"`
}
