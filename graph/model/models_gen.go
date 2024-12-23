// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type AuthPayload struct {
	Token       string `json:"token"`
	UserID      string `json:"userID"`
	Username    string `json:"username"`
	CompanyCode string `json:"companyCode"`
}

type CompanyDetailInput struct {
	Name        string  `json:"Name"`
	Address     *string `json:"Address,omitempty"`
	City        *string `json:"City,omitempty"`
	State       *string `json:"State,omitempty"`
	Phone       *string `json:"Phone,omitempty"`
	Fax         *string `json:"Fax,omitempty"`
	Mobile      *string `json:"Mobile,omitempty"`
	Email       *string `json:"Email,omitempty"`
	Website     *string `json:"Website,omitempty"`
	TaxAdmin    *string `json:"TaxAdmin,omitempty"`
	TaxNo       *int    `json:"TaxNo,omitempty"`
	Commerce    *string `json:"Commerce,omitempty"`
	CommerceReg *string `json:"CommerceReg,omitempty"`
	VisaDate    *string `json:"VisaDate,omitempty"`
	Deleted     *int    `json:"Deleted,omitempty"`
	OwnerID     *int    `json:"OwnerID,omitempty"`
}

type CompanyEngineerInput struct {
	Name       string  `json:"Name"`
	Address    *string `json:"Address,omitempty"`
	Email      *string `json:"Email,omitempty"`
	TcNo       *int    `json:"TcNo,omitempty"`
	Phone      *string `json:"Phone,omitempty"`
	RegNo      *int    `json:"RegNo,omitempty"`
	CertNo     *int    `json:"CertNo,omitempty"`
	Career     *string `json:"Career,omitempty"`
	Position   *string `json:"Position,omitempty"`
	Employment *string `json:"Employment,omitempty"`
	Dismissal  *string `json:"Dismissal,omitempty"`
	Note       *string `json:"Note,omitempty"`
	Status     *int    `json:"Status,omitempty"`
	Deleted    *int    `json:"Deleted,omitempty"`
}

type EngineerFilterInput struct {
	ID       *string `json:"id,omitempty"`
	Career   *string `json:"career,omitempty"`
	Position *string `json:"position,omitempty"`
}

type JobAuthorInput struct {
	Architect *string `json:"Architect,omitempty"`
	Static    *string `json:"Static,omitempty"`
	Mechanic  *string `json:"Mechanic,omitempty"`
	Electric  *string `json:"Electric,omitempty"`
	Floor     *string `json:"Floor,omitempty"`
}

type JobContractorInput struct {
	Name       *string `json:"Name,omitempty"`
	TcNo       *int    `json:"TcNo,omitempty"`
	Address    *string `json:"Address,omitempty"`
	RegisterNo *int    `json:"RegisterNo,omitempty"`
	TaxAdmin   *string `json:"TaxAdmin,omitempty"`
	TaxNo      *int    `json:"TaxNo,omitempty"`
	Phone      *string `json:"Phone,omitempty"`
	Email      *string `json:"Email,omitempty"`
	Note       *string `json:"Note,omitempty"`
}

type JobInput struct {
	YibfNo             *int                  `json:"YibfNo,omitempty"`
	Province           *string               `json:"Province,omitempty"`
	Idare              *string               `json:"Idare,omitempty"`
	Pafta              *string               `json:"Pafta,omitempty"`
	Ada                *string               `json:"Ada,omitempty"`
	Parsel             *string               `json:"Parsel,omitempty"`
	FolderNo           *string               `json:"FolderNo,omitempty"`
	Status             *int                  `json:"Status,omitempty"`
	ContractDate       *string               `json:"ContractDate,omitempty"`
	StartDate          *string               `json:"StartDate,omitempty"`
	LicenseDate        *string               `json:"LicenseDate,omitempty"`
	LicenseNo          *string               `json:"LicenseNo,omitempty"`
	ConstructionArea   *string               `json:"ConstructionArea,omitempty"`
	LandArea           *string               `json:"LandArea,omitempty"`
	District           *string               `json:"District,omitempty"`
	Village            *string               `json:"Village,omitempty"`
	Street             *string               `json:"Street,omitempty"`
	BuildingClass      *string               `json:"BuildingClass,omitempty"`
	BuildingType       *string               `json:"BuildingType,omitempty"`
	BuildingBlock      *string               `json:"BuildingBlock,omitempty"`
	Floors             *int                  `json:"Floors,omitempty"`
	Note               *string               `json:"Note,omitempty"`
	Started            *int                  `json:"Started,omitempty"`
	UsagePurpose       *string               `json:"UsagePurpose,omitempty"`
	Deleted            *int                  `json:"Deleted,omitempty"`
	Owner              []*JobOwnerInput      `json:"Owner,omitempty"`
	Contractor         []*JobContractorInput `json:"Contractor,omitempty"`
	Author             []*JobAuthorInput     `json:"Author,omitempty"`
	Progress           []*JobProgressInput   `json:"Progress,omitempty"`
	Inspector          *int                  `json:"Inspector,omitempty"`
	Static             *int                  `json:"Static,omitempty"`
	Architect          *int                  `json:"Architect,omitempty"`
	Mechanic           *int                  `json:"Mechanic,omitempty"`
	Electric           *int                  `json:"Electric,omitempty"`
	Controller         *int                  `json:"Controller,omitempty"`
	MechanicController *int                  `json:"MechanicController,omitempty"`
	ElectricController *int                  `json:"ElectricController,omitempty"`
}

type JobLayerInput struct {
	Name          *string `json:"Name,omitempty"`
	Metre         *string `json:"Metre,omitempty"`
	MoldDate      *string `json:"MoldDate,omitempty"`
	ConcreteDate  *string `json:"ConcreteDate,omitempty"`
	Samples       *int    `json:"Samples,omitempty"`
	ConcreteClass *string `json:"ConcreteClass,omitempty"`
	WeekResult    *string `json:"WeekResult,omitempty"`
	MonthResult   *string `json:"MonthResult,omitempty"`
	YibfNo        *int    `json:"YibfNo,omitempty"`
}

type JobOwnerInput struct {
	Name     *string `json:"Name,omitempty"`
	TcNo     *int    `json:"TcNo,omitempty"`
	Address  *string `json:"Address,omitempty"`
	TaxAdmin *string `json:"TaxAdmin,omitempty"`
	TaxNo    *int    `json:"TaxNo,omitempty"`
	Phone    *string `json:"Phone,omitempty"`
	Email    *string `json:"Email,omitempty"`
	Note     *string `json:"Note,omitempty"`
}

type JobProgressInput struct {
	One   *int `json:"One,omitempty"`
	Two   *int `json:"Two,omitempty"`
	Three *int `json:"Three,omitempty"`
	Four  *int `json:"Four,omitempty"`
	Five  *int `json:"Five,omitempty"`
	Six   *int `json:"Six,omitempty"`
}

type LayerFilterInput struct {
	ID     *int `json:"id,omitempty"`
	YibfNo *int `json:"yibfNo,omitempty"`
}
