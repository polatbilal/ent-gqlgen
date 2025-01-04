// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"

	"github.com/polatbilal/gqlgen-ent/ent"
)

type AuthPayload struct {
	Token    string `json:"token"`
	UserID   string `json:"userID"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Role     string `json:"role"`
}

type CompanyDetailInput struct {
	CompanyCode            int        `json:"CompanyCode"`
	Name                   string     `json:"Name"`
	Address                *string    `json:"Address,omitempty"`
	Phone                  *string    `json:"Phone,omitempty"`
	Email                  *string    `json:"Email,omitempty"`
	Website                *string    `json:"Website,omitempty"`
	TaxAdmin               *string    `json:"TaxAdmin,omitempty"`
	TaxNo                  *int       `json:"TaxNo,omitempty"`
	ChamberInfo            *string    `json:"ChamberInfo,omitempty"`
	ChamberRegisterNo      *string    `json:"ChamberRegisterNo,omitempty"`
	VisaDate               *time.Time `json:"VisaDate,omitempty"`
	VisaEndDate            *time.Time `json:"VisaEndDate,omitempty"`
	VisaFinishedFor90days  *bool      `json:"visa_finished_for_90days,omitempty"`
	CorePersonAbsent90days *bool      `json:"core_person_absent_90days,omitempty"`
	IsClosed               *bool      `json:"isClosed,omitempty"`
	OwnerName              *string    `json:"OwnerName,omitempty"`
	OwnerTcNo              *int       `json:"OwnerTcNo,omitempty"`
	OwnerAddress           *string    `json:"OwnerAddress,omitempty"`
	OwnerPhone             *string    `json:"OwnerPhone,omitempty"`
	OwnerEmail             *string    `json:"OwnerEmail,omitempty"`
	OwnerRegisterNo        *int       `json:"OwnerRegisterNo,omitempty"`
	OwnerCareer            *string    `json:"OwnerCareer,omitempty"`
}

type CompanyEngineerInput struct {
	CompanyCode *int       `json:"CompanyCode,omitempty"`
	YibfNo      *int       `json:"YibfNo,omitempty"`
	Ydsid       *int       `json:"YDSID,omitempty"`
	Name        *string    `json:"Name,omitempty"`
	TcNo        *int       `json:"TcNo,omitempty"`
	Phone       *string    `json:"Phone,omitempty"`
	Email       *string    `json:"Email,omitempty"`
	Address     *string    `json:"Address,omitempty"`
	Career      *string    `json:"Career,omitempty"`
	Position    *string    `json:"Position,omitempty"`
	RegNo       *int       `json:"RegNo,omitempty"`
	CertNo      *int       `json:"CertNo,omitempty"`
	Employment  *time.Time `json:"Employment,omitempty"`
	Status      *int       `json:"Status,omitempty"`
	Note        *string    `json:"Note,omitempty"`
}

type EngineerFilterInput struct {
	ID    *string `json:"id,omitempty"`
	Ydsid *int    `json:"YDSID,omitempty"`
}

type JobAuthorInput struct {
	Static                   *string `json:"Static,omitempty"`
	Mechanic                 *string `json:"Mechanic,omitempty"`
	Electric                 *string `json:"Electric,omitempty"`
	Architect                *string `json:"Architect,omitempty"`
	GeotechnicalEngineer     *string `json:"GeotechnicalEngineer,omitempty"`
	GeotechnicalGeologist    *string `json:"GeotechnicalGeologist,omitempty"`
	GeotechnicalGeophysicist *string `json:"GeotechnicalGeophysicist,omitempty"`
}

type JobContractorInput struct {
	ID          *string `json:"id,omitempty"`
	YibfNo      int     `json:"YibfNo"`
	Name        string  `json:"Name"`
	TcNo        *int    `json:"TcNo,omitempty"`
	RegisterNo  *int    `json:"RegisterNo,omitempty"`
	Address     *string `json:"Address,omitempty"`
	TaxNo       *int    `json:"TaxNo,omitempty"`
	MobilePhone *string `json:"MobilePhone,omitempty"`
	Phone       *string `json:"Phone,omitempty"`
	Email       *string `json:"Email,omitempty"`
	PersonType  *string `json:"PersonType,omitempty"`
	Ydsid       *int    `json:"YDSID,omitempty"`
	Note        *string `json:"Note,omitempty"`
}

type JobInput struct {
	CompanyCode        int        `json:"CompanyCode"`
	YibfNo             *int       `json:"YibfNo,omitempty"`
	Administration     *string    `json:"Administration,omitempty"`
	State              *string    `json:"State,omitempty"`
	Island             *string    `json:"Island,omitempty"`
	Parcel             *string    `json:"Parcel,omitempty"`
	Sheet              *string    `json:"Sheet,omitempty"`
	ContractDate       *time.Time `json:"ContractDate,omitempty"`
	StartDate          *time.Time `json:"StartDate,omitempty"`
	LicenseDate        *time.Time `json:"LicenseDate,omitempty"`
	LicenseNo          *string    `json:"LicenseNo,omitempty"`
	CompletionDate     *time.Time `json:"CompletionDate,omitempty"`
	LandArea           *float64   `json:"LandArea,omitempty"`
	TotalArea          *float64   `json:"TotalArea,omitempty"`
	ConstructionArea   *float64   `json:"ConstructionArea,omitempty"`
	LeftArea           *float64   `json:"LeftArea,omitempty"`
	YDSAddress         *string    `json:"YDSAddress,omitempty"`
	Address            *string    `json:"Address,omitempty"`
	BuildingClass      *string    `json:"BuildingClass,omitempty"`
	BuildingType       *string    `json:"BuildingType,omitempty"`
	Level              *float64   `json:"Level,omitempty"`
	UnitPrice          *float64   `json:"UnitPrice,omitempty"`
	FloorCount         *int       `json:"FloorCount,omitempty"`
	BKSReferenceNo     *int       `json:"BKSReferenceNo,omitempty"`
	Coordinates        *string    `json:"Coordinates,omitempty"`
	FolderNo           *string    `json:"FolderNo,omitempty"`
	UploadedFile       *bool      `json:"UploadedFile,omitempty"`
	IndustryArea       *bool      `json:"IndustryArea,omitempty"`
	ClusterStructure   *bool      `json:"ClusterStructure,omitempty"`
	IsLicenseExpired   *bool      `json:"IsLicenseExpired,omitempty"`
	IsCompleted        *bool      `json:"IsCompleted,omitempty"`
	Note               *string    `json:"Note,omitempty"`
	Owner              *int       `json:"Owner,omitempty"`
	Contractor         *int       `json:"Contractor,omitempty"`
	Author             *int       `json:"Author,omitempty"`
	Supervisor         *int       `json:"Supervisor,omitempty"`
	Inspector          *int       `json:"Inspector,omitempty"`
	Static             *int       `json:"Static,omitempty"`
	Architect          *int       `json:"Architect,omitempty"`
	Mechanic           *int       `json:"Mechanic,omitempty"`
	Electric           *int       `json:"Electric,omitempty"`
	Controller         *int       `json:"Controller,omitempty"`
	MechanicController *int       `json:"MechanicController,omitempty"`
	ElectricController *int       `json:"ElectricController,omitempty"`
}

type JobLayerInput struct {
	Name          *string    `json:"Name,omitempty"`
	Metre         *string    `json:"Metre,omitempty"`
	MoldDate      *time.Time `json:"MoldDate,omitempty"`
	ConcreteDate  *time.Time `json:"ConcreteDate,omitempty"`
	Samples       *int       `json:"Samples,omitempty"`
	ConcreteClass *string    `json:"ConcreteClass,omitempty"`
	WeekResult    *string    `json:"WeekResult,omitempty"`
	MonthResult   *string    `json:"MonthResult,omitempty"`
	YibfNo        *int       `json:"YibfNo,omitempty"`
}

type JobOwnerInput struct {
	YibfNo      int     `json:"YibfNo"`
	Name        *string `json:"Name,omitempty"`
	TcNo        *int    `json:"TcNo,omitempty"`
	Address     *string `json:"Address,omitempty"`
	TaxAdmin    *string `json:"TaxAdmin,omitempty"`
	TaxNo       *int    `json:"TaxNo,omitempty"`
	Phone       *string `json:"Phone,omitempty"`
	Email       *string `json:"Email,omitempty"`
	Ydsid       *int    `json:"YDSID,omitempty"`
	Shareholder *bool   `json:"Shareholder,omitempty"`
	Note        *string `json:"Note,omitempty"`
}

type JobPaymentsInput struct {
	Date        *time.Time `json:"Date,omitempty"`
	Amount      *int       `json:"Amount,omitempty"`
	Description *string    `json:"Description,omitempty"`
	Status      *string    `json:"Status,omitempty"`
	Percentage  *float64   `json:"Percentage,omitempty"`
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

type Supervisor struct {
	ID               string           `json:"id"`
	Name             *string          `json:"Name,omitempty"`
	Address          *string          `json:"Address,omitempty"`
	Phone            *string          `json:"Phone,omitempty"`
	Email            *string          `json:"Email,omitempty"`
	Tcno             *int             `json:"TCNO,omitempty"`
	Position         *string          `json:"Position,omitempty"`
	Career           *string          `json:"Career,omitempty"`
	RegNo            *int             `json:"RegNo,omitempty"`
	SocialSecurityNo *int             `json:"SocialSecurityNo,omitempty"`
	SchoolGraduation *string          `json:"SchoolGraduation,omitempty"`
	Ydsid            *int             `json:"YDSID,omitempty"`
	Job              []*ent.JobDetail `json:"Job,omitempty"`
}

type SupervisorInput struct {
	Name             *string `json:"Name,omitempty"`
	Address          *string `json:"Address,omitempty"`
	Phone            *string `json:"Phone,omitempty"`
	Email            *string `json:"Email,omitempty"`
	Tcno             *int    `json:"TCNO,omitempty"`
	Position         *string `json:"Position,omitempty"`
	Career           *string `json:"Career,omitempty"`
	RegNo            *int    `json:"RegNo,omitempty"`
	SocialSecurityNo *int    `json:"SocialSecurityNo,omitempty"`
	SchoolGraduation *string `json:"SchoolGraduation,omitempty"`
	Ydsid            *int    `json:"YDSID,omitempty"`
}

type UserInput struct {
	Username   string `json:"username"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Phone      int    `json:"phone"`
	Password   string `json:"password"`
	Role       string `json:"role"`
	CompanyIDs []int  `json:"companyIDs,omitempty"`
}
