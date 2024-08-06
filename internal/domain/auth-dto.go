package domain

type CodeRequest struct {
	Email string `json:"email"`
}

type RegistryRequest struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

type RegistryResponse struct {
	AccessToken  string `josn:"access_token"`
	RefreshToken string `josn:"refresh_token"`
}
