package handlers

import "net/http"

type ExternalService struct {
	baseURL string
	client  *http.Client
}

type GraphQLClient struct {
	URL string
}

type graphQLRequest struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables,omitempty"`
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
			Person struct {
				ID              int    `json:"id"`
				FullName        string `json:"fullName"`
				AddressStr      string `json:"addressStr"`
				LastEPosta      string `json:"lastEPosta"`
				IdentityNumber  string `json:"identityNumber"`
				LastPhoneNumber string `json:"lastPhoneNumber"`
			} `json:"person"`
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
