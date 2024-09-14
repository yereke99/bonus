package domain

type CompanyRequest struct {
	Company        string `json:"company"`
	CompanyName    string `json:"companyName"`
	Email          string `json:"email"`
	City           string `json:"city"`
	CompanyAddress string `json:"companyAddress"`
	CompanyIIN     int    `json:"companyInn"`
	Bonus          int    `json:"companyBonus"`
}

type Company struct {
	ID             string `json:"id"`
	Company        string `json:"company"`
	CompanyName    string `json:"companyName"`
	Email          string `json:"email"`
	City           string `json:"city"`
	CompanyAddress string `json:"companyAddress"`
	CompanyIIN     int    `json:"companyInn"`
	Bonus          int    `json:"companyBonus"`
	IsDeleted      bool   `json:"isDeleted"`
}

type CompanyObject struct {
	ID             string `json:"id"`
	Company        string `json:"company"`
	CompanyName    string `json:"companyName"`
	Email          string `json:"email"`
	City           string `json:"city"`
	CompanyAddress string `json:"companyAddress"`
	CompanyIIN     int    `json:"companyInn"`
	Bonus          int    `json:"companyBonus"`
	IsDeleted      bool   `json:"isDeleted"`
}

type NotificationRequest struct{}

type BonusCalculationRequest struct{}

type BarcodeRequest struct{}

type CommissionCalculationRequest struct{}

type DoubleBonusRequest struct{}
