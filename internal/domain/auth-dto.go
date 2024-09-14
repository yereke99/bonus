package domain

type CodeRequest struct {
	Email string `json:"email"`
}

type RegistryRequest struct {
	UserName     string `json:"user_name"`
	UserLastName string `json:"user_last_name"`
	Email        string `json:"email"`
	Locations    string `json:"locations"`
	City         string `json:"city"`
	QR           string `json:"qr"`
	Bonus        int    `json:"bonus"`
	Token        string `json:"token"`
	IsDeleted    bool   `json:"is_deleted"`
}

type RegistryResponse struct {
	ID           string `json:"id"`
	UserName     string `json:"user_name"`
	UserLastName string `json:"user_last_name"`
	Email        string `json:"email"`
	Locations    string `json:"locations"`
	City         string `json:"city"`
	QR           string `json:"qr"`
	Bonus        int    `json:"bonus"`
	Token        string `json:"token"`
	IsDeleted    bool   `json:"is_deleted"`
	AccessToken  string `json:"access_token"`
}

type LoginResponse struct {
	ID           string `json:"id"`
	UserName     string `json:"user_name"`
	UserLastName string `json:"user_last_name"`
	Email        string `json:"email"`
	Locations    string `json:"locations"`
	City         string `json:"city"`
	QR           string `json:"qr"`
	Bonus        int    `json:"bonus"`
	Token        string `json:"token"`
	IsDeleted    bool   `json:"is_deleted"`
}
