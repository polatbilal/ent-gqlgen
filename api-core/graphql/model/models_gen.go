// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"

	"github.com/polatbilal/gqlgen-ent/api-core/ent"
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
	Fax                    *string    `json:"Fax,omitempty"`
	MobilePhone            *string    `json:"MobilePhone,omitempty"`
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
	RegisterNo  *int       `json:"RegisterNo,omitempty"`
	CertNo      *int       `json:"CertNo,omitempty"`
	Employment  *time.Time `json:"Employment,omitempty"`
	Status      *int       `json:"Status,omitempty"`
	Note        *string    `json:"Note,omitempty"`
}

type CompanyTokenInput struct {
	Token           *string `json:"Token,omitempty"`
	DepartmentID    *int    `json:"DepartmentId,omitempty"`
	Expire          *int    `json:"Expire,omitempty"`
	CompanyCode     *int    `json:"CompanyCode,omitempty"`
	RefreshToken    *string `json:"RefreshToken,omitempty"`
	SecretKey       *string `json:"SecretKey,omitempty"`
	SecureSecretKey *string `json:"SecureSecretKey,omitempty"`
	OtpURI          *string `json:"OtpUri,omitempty"`
}

type EngineerFilterInput struct {
	ID    *string `json:"id,omitempty"`
	Ydsid *int    `json:"YDSID,omitempty"`
}

type JobAuthorInput struct {
	YibfNo                   *int    `json:"YibfNo,omitempty"`
	Static                   *string `json:"Static,omitempty"`
	Mechanic                 *string `json:"Mechanic,omitempty"`
	Electric                 *string `json:"Electric,omitempty"`
	Architect                *string `json:"Architect,omitempty"`
	GeotechnicalEngineer     *string `json:"GeotechnicalEngineer,omitempty"`
	GeotechnicalGeologist    *string `json:"GeotechnicalGeologist,omitempty"`
	GeotechnicalGeophysicist *string `json:"GeotechnicalGeophysicist,omitempty"`
}

type JobBatchInput struct {
	YibfNo          int                 `json:"YibfNo"`
	JobInput        *JobInput           `json:"jobInput,omitempty"`
	OwnerInput      *JobOwnerInput      `json:"ownerInput,omitempty"`
	ContractorInput *JobContractorInput `json:"contractorInput,omitempty"`
	AuthorInput     *JobAuthorInput     `json:"authorInput,omitempty"`
	SupervisorInput *JobSupervisorInput `json:"supervisorInput,omitempty"`
	EngineerInput   *JobEngineerInput   `json:"engineerInput,omitempty"`
}

type JobBatchResult struct {
	Job        *ent.JobDetail     `json:"job,omitempty"`
	Owner      *ent.JobOwner      `json:"owner,omitempty"`
	Contractor *ent.JobContractor `json:"contractor,omitempty"`
	Author     *ent.JobAuthor     `json:"author,omitempty"`
	Supervisor *ent.JobSupervisor `json:"supervisor,omitempty"`
	Engineer   *JobEngineer       `json:"engineer,omitempty"`
	Progress   *ent.JobProgress   `json:"progress,omitempty"`
	Company    *ent.CompanyDetail `json:"company,omitempty"`
}

type JobContractorInput struct {
	ID          *string `json:"id,omitempty"`
	YibfNo      *int    `json:"YibfNo,omitempty"`
	Name        *string `json:"Name,omitempty"`
	TcNo        *int    `json:"TcNo,omitempty"`
	RegisterNo  *int    `json:"RegisterNo,omitempty"`
	Address     *string `json:"Address,omitempty"`
	TaxNo       *int    `json:"TaxNo,omitempty"`
	MobilePhone *string `json:"MobilePhone,omitempty"`
	Phone       *string `json:"Phone,omitempty"`
	Email       *string `json:"Email,omitempty"`
	PersonType  *string `json:"PersonType,omitempty"`
	Ydsid       int     `json:"YDSID"`
	Note        *string `json:"Note,omitempty"`
}

type JobCounts struct {
	Current   int `json:"current"`
	Pending   int `json:"pending"`
	Completed int `json:"completed"`
	Total     int `json:"total"`
}

type JobEngineer struct {
	YibfNo             *int                 `json:"YibfNo,omitempty"`
	Inspector          *ent.CompanyEngineer `json:"Inspector,omitempty"`
	Static             *ent.CompanyEngineer `json:"Static,omitempty"`
	Architect          *ent.CompanyEngineer `json:"Architect,omitempty"`
	Mechanic           *ent.CompanyEngineer `json:"Mechanic,omitempty"`
	Electric           *ent.CompanyEngineer `json:"Electric,omitempty"`
	Controller         *ent.CompanyEngineer `json:"Controller,omitempty"`
	MechanicController *ent.CompanyEngineer `json:"MechanicController,omitempty"`
	ElectricController *ent.CompanyEngineer `json:"ElectricController,omitempty"`
}

type JobEngineerInput struct {
	YibfNo             *int `json:"YibfNo,omitempty"`
	Inspector          *int `json:"Inspector,omitempty"`
	Static             *int `json:"Static,omitempty"`
	Architect          *int `json:"Architect,omitempty"`
	Mechanic           *int `json:"Mechanic,omitempty"`
	Electric           *int `json:"Electric,omitempty"`
	Controller         *int `json:"Controller,omitempty"`
	MechanicController *int `json:"MechanicController,omitempty"`
	ElectricController *int `json:"ElectricController,omitempty"`
}

type JobInput struct {
	CompanyCode      *int       `json:"CompanyCode,omitempty"`
	YibfNo           *int       `json:"YibfNo,omitempty"`
	Title            *string    `json:"Title,omitempty"`
	Administration   *string    `json:"Administration,omitempty"`
	State            *string    `json:"State,omitempty"`
	Island           *string    `json:"Island,omitempty"`
	Parcel           *string    `json:"Parcel,omitempty"`
	Sheet            *string    `json:"Sheet,omitempty"`
	DistributionDate *time.Time `json:"DistributionDate,omitempty"`
	ContractDate     *time.Time `json:"ContractDate,omitempty"`
	StartDate        *time.Time `json:"StartDate,omitempty"`
	LicenseDate      *time.Time `json:"LicenseDate,omitempty"`
	LicenseNo        *string    `json:"LicenseNo,omitempty"`
	CompletionDate   *time.Time `json:"CompletionDate,omitempty"`
	LandArea         *float64   `json:"LandArea,omitempty"`
	TotalArea        *float64   `json:"TotalArea,omitempty"`
	ConstructionArea *float64   `json:"ConstructionArea,omitempty"`
	LeftArea         *float64   `json:"LeftArea,omitempty"`
	YDSAddress       *string    `json:"YDSAddress,omitempty"`
	Address          *string    `json:"Address,omitempty"`
	BuildingClass    *string    `json:"BuildingClass,omitempty"`
	BuildingType     *string    `json:"BuildingType,omitempty"`
	Level            *float64   `json:"Level,omitempty"`
	UnitPrice        *float64   `json:"UnitPrice,omitempty"`
	FloorCount       *int       `json:"FloorCount,omitempty"`
	BKSReferenceNo   *int       `json:"BKSReferenceNo,omitempty"`
	Coordinates      *string    `json:"Coordinates,omitempty"`
	FolderNo         *string    `json:"FolderNo,omitempty"`
	UploadedFile     *bool      `json:"UploadedFile,omitempty"`
	IndustryArea     *bool      `json:"IndustryArea,omitempty"`
	ClusterStructure *bool      `json:"ClusterStructure,omitempty"`
	IsLicenseExpired *bool      `json:"IsLicenseExpired,omitempty"`
	IsCompleted      *bool      `json:"IsCompleted,omitempty"`
	Note             *string    `json:"Note,omitempty"`
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
	ID          *string `json:"id,omitempty"`
	YibfNo      *int    `json:"YibfNo,omitempty"`
	Name        *string `json:"Name,omitempty"`
	TcNo        *int    `json:"TcNo,omitempty"`
	Address     *string `json:"Address,omitempty"`
	TaxAdmin    *string `json:"TaxAdmin,omitempty"`
	TaxNo       *int    `json:"TaxNo,omitempty"`
	Phone       *string `json:"Phone,omitempty"`
	Email       *string `json:"Email,omitempty"`
	Ydsid       int     `json:"YDSID"`
	Shareholder *bool   `json:"Shareholder,omitempty"`
	Note        *string `json:"Note,omitempty"`
}

type JobPaymentsInput struct {
	PaymentNo    int        `json:"PaymentNo"`
	PaymentDate  *time.Time `json:"PaymentDate,omitempty"`
	PaymentType  *string    `json:"PaymentType,omitempty"`
	TotalPayment *float64   `json:"TotalPayment,omitempty"`
	LevelRequest *float64   `json:"LevelRequest,omitempty"`
	LevelApprove *float64   `json:"LevelApprove,omitempty"`
	Amount       *float64   `json:"Amount,omitempty"`
	YibfNo       *int       `json:"YibfNo,omitempty"`
	State        *string    `json:"State,omitempty"`
}

type JobProgressInput struct {
	ID    *string `json:"id,omitempty"`
	One   *int    `json:"One,omitempty"`
	Two   *int    `json:"Two,omitempty"`
	Three *int    `json:"Three,omitempty"`
	Four  *int    `json:"Four,omitempty"`
	Five  *int    `json:"Five,omitempty"`
	Six   *int    `json:"Six,omitempty"`
}

type JobSupervisorInput struct {
	YibfNo           *int    `json:"YibfNo,omitempty"`
	Name             *string `json:"Name,omitempty"`
	Address          *string `json:"Address,omitempty"`
	Phone            *string `json:"Phone,omitempty"`
	Email            *string `json:"Email,omitempty"`
	TcNo             *int    `json:"TcNo,omitempty"`
	Position         *string `json:"Position,omitempty"`
	Career           *string `json:"Career,omitempty"`
	RegisterNo       *int    `json:"RegisterNo,omitempty"`
	SocialSecurityNo *int    `json:"SocialSecurityNo,omitempty"`
	SchoolGraduation *string `json:"SchoolGraduation,omitempty"`
	Ydsid            int     `json:"YDSID"`
}

type LayerFilterInput struct {
	ID     *int `json:"id,omitempty"`
	YibfNo *int `json:"yibfNo,omitempty"`
}

type UserInput struct {
	ID          *int    `json:"id,omitempty"`
	Username    *string `json:"Username,omitempty"`
	Name        *string `json:"Name,omitempty"`
	Email       *string `json:"Email,omitempty"`
	Phone       *string `json:"Phone,omitempty"`
	Password    *string `json:"Password,omitempty"`
	NewPassword *string `json:"NewPassword,omitempty"`
	Role        *string `json:"Role,omitempty"`
	CompanyIDs  []*int  `json:"CompanyIDs,omitempty"`
}
