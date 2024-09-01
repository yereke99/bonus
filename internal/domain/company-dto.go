package domain

type CompanyRequest struct {
	Company        string
	CompanyName    string
	Email          string
	City           string
	CompanyAddress string
	CompanyIIN     int
	Bonus          int
}

type Company struct {
	ID             int64
	Company        string
	CompanyName    string
	Email          string
	City           string
	CompanyAddress string
	CompanyIIN     int
	Bonus          int
	IsDeleted      bool
}

type NotificationRequest struct{}

type BonusCalculationRequest struct{}

type BarcodeRequest struct{}

type CommissionCalculationRequest struct{}

type DoubleBonusRequest struct{}
