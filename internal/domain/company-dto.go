package domain

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
