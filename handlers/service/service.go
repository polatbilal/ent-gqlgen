package service

import (
	"net/http"
)

type ExternalService struct {
	BaseURL string
	Client  *http.Client
}

type YDKTokenResponse struct {
	AccessToken    string         `json:"access_token"`
	TokenType      string         `json:"token_type"`
	RefreshToken   string         `json:"refresh_token"`
	DepartmentID   int            `json:"department_id"`
	ResponseSecret ResponseSecret `json:"response_secret"`
}

type ResponseSecret struct {
	Secret       string `json:"secret"`
	SecureSecret string `json:"secureSecret"`
	URI          string `json:"uri"`
}

type YDKCompanyResponse struct {
	Items []struct {
		Department struct {
			Name                   string `json:"name"`                      // Company Name
			FileNumber             int    `json:"fileNumber"`                // Company Code
			ChamberInfo            string `json:"chamberInfo"`               // Chamber Info
			RegistrationNumber     string `json:"registrationNumber"`        // Chamber Reg No
			VisaDate               int64  `json:"visaDate"`                  // Visa Date
			VisaEndDate            int64  `json:"visaEndDate"`               // Visa End Date
			VisaFinishedFor90Days  bool   `json:"visa_finished_for_90days"`  // Visa Finished For 90 Days
			CorePersonAbsent90Days bool   `json:"core_person_absent_90days"` // Core Person Absent 90 Days
			IsClosed               bool   `json:"isClosed"`                  // Is Closed
			Person                 struct {
				IdentityNumber    string `json:"identityNumber"`    // Tax Number
				TaxAdministration string `json:"taxAdministration"` // Tax Administration
				AddressStr        string `json:"addressStr"`        // Company Address
				FullName          string `json:"fullName"`          // Company Name
				LastPhoneNumber   string `json:"lastPhoneNumber"`   // Company Phone
				LastWebAddress    string `json:"lastWebAddress"`    // Company Website
				LastEPosta        string `json:"lastEPosta"`        // Company Email
			} `json:"person"`
		} `json:"department"`
		Title struct {
			Name string `json:"name"` // Owner Career
		} `json:"title"`
		Person struct {
			IdentityNumber  string `json:"identityNumber"`  // Owner TC No
			AddressStr      string `json:"addressStr"`      // Owner Address
			FullName        string `json:"fullName"`        // Owner Name
			BirthDateString string `json:"birthDateString"` // Owner Birth Date
			LastPhoneNumber string `json:"lastPhoneNumber"` // Owner Phone
			LastEPosta      string `json:"lastEPosta"`      // Owner Email
		} `json:"person"`
		OccupationalRegistrationNumber string `json:"occupationalRegistrationNumber"` // Owner Reg No
	} `json:"items"`
	TotalCount int `json:"totalCount"`
	GroupCount int `json:"groupCount"`
}

type YDKInspectorResponse struct {
	Items []struct {
		Application struct {
			ApplicationType struct {
				Name string `json:"name"`
			} `json:"applicationType"`
			User struct {
				ID     int `json:"id"`
				Person struct {
					ID              int    `json:"id"`
					FullName        string `json:"fullName"`
					AddressStr      string `json:"addressStr"`
					LastEPosta      string `json:"lastEPosta"`
					IdentityNumber  string `json:"identityNumber"`
					LastPhoneNumber string `json:"lastPhoneNumber"`
				} `json:"person"`
			} `json:"user"`
			Title struct {
				Name string `json:"name"`
			} `json:"title"`
			Tasks struct {
				Name string `json:"name"`
			} `json:"tasks"`
			RegistrationNumber string `json:"occupationalRegistrationNumber"`
			DocumentNumber     int    `json:"documentNumber"`
		} `json:"application"`
		Department struct {
			FileNumber int `json:"fileNumber"`
		} `json:"department"`
		TaskStartDate int64 `json:"taskStartDate"`
	} `json:"items"`
	TotalCount int `json:"totalCount"`
}

type YIBFResponse struct {
	ID               int     `json:"id"`
	Island           string  `json:"island"`
	Parcel           string  `json:"parcel"`
	Sheet            string  `json:"sheet"`
	Title            string  `json:"title"`
	ClusterStructure bool    `json:"clusterStructure"`
	IsLicenseExpired bool    `json:"isLicenseExpired"`
	ContractDate     int64   `json:"contractDate"`
	LicenseNumber    string  `json:"licenseNumber"`
	LicenseDate      int64   `json:"licenseDate"`
	ReferenceNumber  int     `json:"referenceNumber"`
	IsCompleted      bool    `json:"isCompleted"`
	Level            float64 `json:"level"`
	Position         struct {
		Coordinates []float64 `json:"coordinates"`
	} `json:"position"`
	YibfStructure struct {
		BuildingClass struct {
			Name string `json:"name"`
		} `json:"buildingClass"`
		CarrierSystemType struct {
			Name string `json:"name"`
		} `json:"carrierSystemType"`
		TotalArea        float64 `json:"totalArea"`
		ConstructionArea float64 `json:"constructionArea"`
		LeftArea         float64 `json:"leftArea"`
		BuildingAddress  string  `json:"buildingAddress"`
		UnitPrice        float64 `json:"unitPrice"`
		FloorCount       int     `json:"floorCount"`
	} `json:"yibfStructure"`
	LatestYibfContract struct {
		ContractPriceRatio  float64 `json:"contractPriceRatio"`
		PriceRatio          float64 `json:"priceRatio"`
		CompletionTimeMonth int     `json:"completionTimeMonth"`
		ContractServiceFee  float64 `json:"contractServiceFee"`
	} `json:"latestYibfContract"`
	Megsis struct {
		Alan float64 `json:"alan"`
	} `json:"megsis"`
	Administration struct {
		Name string `json:"name"`
	} `json:"administration"`
	State struct {
		Name string `json:"name"`
	} `json:"state"`
	YDK struct {
		FileNumber int `json:"fileNumber"`
	} `json:"ydk"`
	YibfOwner struct {
		ExistsShareholder bool `json:"existsShareholder"`
		Person            struct {
			ID                    int    `json:"id"`
			IdentityNumber        string `json:"identityNumber"`
			FullName              string `json:"fullName"`
			LastAddress           string `json:"lastAddress"`
			LastPhoneNumber       string `json:"lastPhoneNumber"`
			TaxAdministration     string `json:"taxAdministration"`
			TaxAdministrationCode string `json:"taxAdministrationCode"`
		} `json:"person"`
	} `json:"yibfOwner"`
	LatestYibfSiteSupervisor struct {
		Application struct {
			User struct {
				ID     int `json:"id"`
				Person struct {
					IdentityNumber  string `json:"identityNumber"`
					FullName        string `json:"fullName"`
					LastAddress     string `json:"lastAddress"`
					LastPhoneNumber string `json:"lastPhoneNumber"`
					LastEPosta      string `json:"lastEPosta"`
				} `json:"person"`
			} `json:"user"`
			OccupationalRegistrationNumber string `json:"occupationalRegistrationNumber"`
			SocialSecurityNo               string `json:"socialSecurityNo"`
			SchoolGraduation               string `json:"schoolGraduation"`
			Tasks                          struct {
				Name string `json:"name"`
			} `json:"tasks"`
			Title struct {
				Name string `json:"name"`
			} `json:"title"`
		} `json:"application"`
	} `json:"latestYibfSiteSupervisor"`
	LatestYibfYambis struct {
		Yambis struct {
			ID            int    `json:"id"`
			Ybno          string `json:"ybno"`
			AdSoyadUnvan  string `json:"adSoyadUnvan"`
			VergiKimlikNo string `json:"vergiKimlikNo"`
			Adres         string `json:"adres"`
			CepTelefon    string `json:"cepTelefon"`
			Telefon       string `json:"telefon"`
			Eposta        string `json:"eposta"`
			KisiTuru      string `json:"kisiTuru"`
			TcNo          string `json:"tcNo"`
		} `json:"yambis"`
	} `json:"latestYibfYambis"`
}

type YIBFAuthorResponse struct {
	Items []struct {
		YibfId        string `json:"yibfId"`
		PersonName    string `json:"personName"`
		PersonSurname string `json:"personSurname"`
		TaskId        int    `json:"taskId"`
		TaskName      string `json:"taskName"`
		TitleId       int    `json:"titleId"`
		TitleName     string `json:"titleName"`
		TaskGroupId   int    `json:"taskGroupId"`
	} `json:"items"`
}
