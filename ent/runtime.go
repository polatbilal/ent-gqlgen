// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/polatbilal/gqlgen-ent/ent/companydetail"
	"github.com/polatbilal/gqlgen-ent/ent/companyengineer"
	"github.com/polatbilal/gqlgen-ent/ent/companytoken"
	"github.com/polatbilal/gqlgen-ent/ent/jobauthor"
	"github.com/polatbilal/gqlgen-ent/ent/jobcontractor"
	"github.com/polatbilal/gqlgen-ent/ent/jobdetail"
	"github.com/polatbilal/gqlgen-ent/ent/joblayer"
	"github.com/polatbilal/gqlgen-ent/ent/jobowner"
	"github.com/polatbilal/gqlgen-ent/ent/jobpayments"
	"github.com/polatbilal/gqlgen-ent/ent/jobprogress"
	"github.com/polatbilal/gqlgen-ent/ent/jobsupervisor"
	"github.com/polatbilal/gqlgen-ent/ent/schema"
	"github.com/polatbilal/gqlgen-ent/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	companydetailFields := schema.CompanyDetail{}.Fields()
	_ = companydetailFields
	// companydetailDescName is the schema descriptor for Name field.
	companydetailDescName := companydetailFields[1].Descriptor()
	// companydetail.DefaultName holds the default value on creation for the Name field.
	companydetail.DefaultName = companydetailDescName.Default.(string)
	// companydetailDescTaxNo is the schema descriptor for TaxNo field.
	companydetailDescTaxNo := companydetailFields[9].Descriptor()
	// companydetail.DefaultTaxNo holds the default value on creation for the TaxNo field.
	companydetail.DefaultTaxNo = companydetailDescTaxNo.Default.(int)
	// companydetailDescVisaFinishedFor90Days is the schema descriptor for VisaFinishedFor90Days field.
	companydetailDescVisaFinishedFor90Days := companydetailFields[14].Descriptor()
	// companydetail.DefaultVisaFinishedFor90Days holds the default value on creation for the VisaFinishedFor90Days field.
	companydetail.DefaultVisaFinishedFor90Days = companydetailDescVisaFinishedFor90Days.Default.(bool)
	// companydetailDescCorePersonAbsent90Days is the schema descriptor for CorePersonAbsent90Days field.
	companydetailDescCorePersonAbsent90Days := companydetailFields[15].Descriptor()
	// companydetail.DefaultCorePersonAbsent90Days holds the default value on creation for the CorePersonAbsent90Days field.
	companydetail.DefaultCorePersonAbsent90Days = companydetailDescCorePersonAbsent90Days.Default.(bool)
	// companydetailDescIsClosed is the schema descriptor for IsClosed field.
	companydetailDescIsClosed := companydetailFields[16].Descriptor()
	// companydetail.DefaultIsClosed holds the default value on creation for the IsClosed field.
	companydetail.DefaultIsClosed = companydetailDescIsClosed.Default.(bool)
	// companydetailDescCreatedAt is the schema descriptor for CreatedAt field.
	companydetailDescCreatedAt := companydetailFields[24].Descriptor()
	// companydetail.DefaultCreatedAt holds the default value on creation for the CreatedAt field.
	companydetail.DefaultCreatedAt = companydetailDescCreatedAt.Default.(func() time.Time)
	// companydetailDescUpdatedAt is the schema descriptor for UpdatedAt field.
	companydetailDescUpdatedAt := companydetailFields[25].Descriptor()
	// companydetail.DefaultUpdatedAt holds the default value on creation for the UpdatedAt field.
	companydetail.DefaultUpdatedAt = companydetailDescUpdatedAt.Default.(func() time.Time)
	// companydetail.UpdateDefaultUpdatedAt holds the default value on update for the UpdatedAt field.
	companydetail.UpdateDefaultUpdatedAt = companydetailDescUpdatedAt.UpdateDefault.(func() time.Time)
	companyengineerFields := schema.CompanyEngineer{}.Fields()
	_ = companyengineerFields
	// companyengineerDescName is the schema descriptor for Name field.
	companyengineerDescName := companyengineerFields[0].Descriptor()
	// companyengineer.DefaultName holds the default value on creation for the Name field.
	companyengineer.DefaultName = companyengineerDescName.Default.(string)
	// companyengineerDescStatus is the schema descriptor for Status field.
	companyengineerDescStatus := companyengineerFields[11].Descriptor()
	// companyengineer.DefaultStatus holds the default value on creation for the Status field.
	companyengineer.DefaultStatus = companyengineerDescStatus.Default.(int)
	// companyengineerDescCreatedAt is the schema descriptor for CreatedAt field.
	companyengineerDescCreatedAt := companyengineerFields[13].Descriptor()
	// companyengineer.DefaultCreatedAt holds the default value on creation for the CreatedAt field.
	companyengineer.DefaultCreatedAt = companyengineerDescCreatedAt.Default.(func() time.Time)
	// companyengineerDescUpdatedAt is the schema descriptor for UpdatedAt field.
	companyengineerDescUpdatedAt := companyengineerFields[14].Descriptor()
	// companyengineer.DefaultUpdatedAt holds the default value on creation for the UpdatedAt field.
	companyengineer.DefaultUpdatedAt = companyengineerDescUpdatedAt.Default.(func() time.Time)
	// companyengineer.UpdateDefaultUpdatedAt holds the default value on update for the UpdatedAt field.
	companyengineer.UpdateDefaultUpdatedAt = companyengineerDescUpdatedAt.UpdateDefault.(func() time.Time)
	companytokenFields := schema.CompanyToken{}.Fields()
	_ = companytokenFields
	// companytokenDescCreatedAt is the schema descriptor for createdAt field.
	companytokenDescCreatedAt := companytokenFields[7].Descriptor()
	// companytoken.DefaultCreatedAt holds the default value on creation for the createdAt field.
	companytoken.DefaultCreatedAt = companytokenDescCreatedAt.Default.(func() time.Time)
	// companytokenDescUpdatedAt is the schema descriptor for updatedAt field.
	companytokenDescUpdatedAt := companytokenFields[8].Descriptor()
	// companytoken.DefaultUpdatedAt holds the default value on creation for the updatedAt field.
	companytoken.DefaultUpdatedAt = companytokenDescUpdatedAt.Default.(func() time.Time)
	// companytoken.UpdateDefaultUpdatedAt holds the default value on update for the updatedAt field.
	companytoken.UpdateDefaultUpdatedAt = companytokenDescUpdatedAt.UpdateDefault.(func() time.Time)
	jobauthorFields := schema.JobAuthor{}.Fields()
	_ = jobauthorFields
	// jobauthorDescCreatedAt is the schema descriptor for CreatedAt field.
	jobauthorDescCreatedAt := jobauthorFields[7].Descriptor()
	// jobauthor.DefaultCreatedAt holds the default value on creation for the CreatedAt field.
	jobauthor.DefaultCreatedAt = jobauthorDescCreatedAt.Default.(func() time.Time)
	// jobauthorDescUpdatedAt is the schema descriptor for UpdatedAt field.
	jobauthorDescUpdatedAt := jobauthorFields[8].Descriptor()
	// jobauthor.DefaultUpdatedAt holds the default value on creation for the UpdatedAt field.
	jobauthor.DefaultUpdatedAt = jobauthorDescUpdatedAt.Default.(func() time.Time)
	// jobauthor.UpdateDefaultUpdatedAt holds the default value on update for the UpdatedAt field.
	jobauthor.UpdateDefaultUpdatedAt = jobauthorDescUpdatedAt.UpdateDefault.(func() time.Time)
	jobcontractorFields := schema.JobContractor{}.Fields()
	_ = jobcontractorFields
	// jobcontractorDescName is the schema descriptor for Name field.
	jobcontractorDescName := jobcontractorFields[0].Descriptor()
	// jobcontractor.DefaultName holds the default value on creation for the Name field.
	jobcontractor.DefaultName = jobcontractorDescName.Default.(string)
	// jobcontractorDescCreatedAt is the schema descriptor for CreatedAt field.
	jobcontractorDescCreatedAt := jobcontractorFields[11].Descriptor()
	// jobcontractor.DefaultCreatedAt holds the default value on creation for the CreatedAt field.
	jobcontractor.DefaultCreatedAt = jobcontractorDescCreatedAt.Default.(func() time.Time)
	// jobcontractorDescUpdatedAt is the schema descriptor for UpdatedAt field.
	jobcontractorDescUpdatedAt := jobcontractorFields[12].Descriptor()
	// jobcontractor.DefaultUpdatedAt holds the default value on creation for the UpdatedAt field.
	jobcontractor.DefaultUpdatedAt = jobcontractorDescUpdatedAt.Default.(func() time.Time)
	// jobcontractor.UpdateDefaultUpdatedAt holds the default value on update for the UpdatedAt field.
	jobcontractor.UpdateDefaultUpdatedAt = jobcontractorDescUpdatedAt.UpdateDefault.(func() time.Time)
	jobdetailFields := schema.JobDetail{}.Fields()
	_ = jobdetailFields
	// jobdetailDescUploadedFile is the schema descriptor for UploadedFile field.
	jobdetailDescUploadedFile := jobdetailFields[27].Descriptor()
	// jobdetail.DefaultUploadedFile holds the default value on creation for the UploadedFile field.
	jobdetail.DefaultUploadedFile = jobdetailDescUploadedFile.Default.(bool)
	// jobdetailDescIndustryArea is the schema descriptor for IndustryArea field.
	jobdetailDescIndustryArea := jobdetailFields[28].Descriptor()
	// jobdetail.DefaultIndustryArea holds the default value on creation for the IndustryArea field.
	jobdetail.DefaultIndustryArea = jobdetailDescIndustryArea.Default.(bool)
	// jobdetailDescClusterStructure is the schema descriptor for ClusterStructure field.
	jobdetailDescClusterStructure := jobdetailFields[29].Descriptor()
	// jobdetail.DefaultClusterStructure holds the default value on creation for the ClusterStructure field.
	jobdetail.DefaultClusterStructure = jobdetailDescClusterStructure.Default.(bool)
	// jobdetailDescIsLicenseExpired is the schema descriptor for IsLicenseExpired field.
	jobdetailDescIsLicenseExpired := jobdetailFields[30].Descriptor()
	// jobdetail.DefaultIsLicenseExpired holds the default value on creation for the IsLicenseExpired field.
	jobdetail.DefaultIsLicenseExpired = jobdetailDescIsLicenseExpired.Default.(bool)
	// jobdetailDescIsCompleted is the schema descriptor for IsCompleted field.
	jobdetailDescIsCompleted := jobdetailFields[31].Descriptor()
	// jobdetail.DefaultIsCompleted holds the default value on creation for the IsCompleted field.
	jobdetail.DefaultIsCompleted = jobdetailDescIsCompleted.Default.(bool)
	// jobdetailDescCreatedAt is the schema descriptor for CreatedAt field.
	jobdetailDescCreatedAt := jobdetailFields[33].Descriptor()
	// jobdetail.DefaultCreatedAt holds the default value on creation for the CreatedAt field.
	jobdetail.DefaultCreatedAt = jobdetailDescCreatedAt.Default.(func() time.Time)
	// jobdetailDescUpdatedAt is the schema descriptor for UpdatedAt field.
	jobdetailDescUpdatedAt := jobdetailFields[34].Descriptor()
	// jobdetail.DefaultUpdatedAt holds the default value on creation for the UpdatedAt field.
	jobdetail.DefaultUpdatedAt = jobdetailDescUpdatedAt.Default.(func() time.Time)
	// jobdetail.UpdateDefaultUpdatedAt holds the default value on update for the UpdatedAt field.
	jobdetail.UpdateDefaultUpdatedAt = jobdetailDescUpdatedAt.UpdateDefault.(func() time.Time)
	joblayerFields := schema.JobLayer{}.Fields()
	_ = joblayerFields
	// joblayerDescName is the schema descriptor for Name field.
	joblayerDescName := joblayerFields[0].Descriptor()
	// joblayer.DefaultName holds the default value on creation for the Name field.
	joblayer.DefaultName = joblayerDescName.Default.(string)
	// joblayerDescMetre is the schema descriptor for Metre field.
	joblayerDescMetre := joblayerFields[1].Descriptor()
	// joblayer.DefaultMetre holds the default value on creation for the Metre field.
	joblayer.DefaultMetre = joblayerDescMetre.Default.(string)
	// joblayerDescCreatedAt is the schema descriptor for CreatedAt field.
	joblayerDescCreatedAt := joblayerFields[8].Descriptor()
	// joblayer.DefaultCreatedAt holds the default value on creation for the CreatedAt field.
	joblayer.DefaultCreatedAt = joblayerDescCreatedAt.Default.(func() time.Time)
	// joblayerDescUpdatedAt is the schema descriptor for UpdatedAt field.
	joblayerDescUpdatedAt := joblayerFields[9].Descriptor()
	// joblayer.DefaultUpdatedAt holds the default value on creation for the UpdatedAt field.
	joblayer.DefaultUpdatedAt = joblayerDescUpdatedAt.Default.(func() time.Time)
	// joblayer.UpdateDefaultUpdatedAt holds the default value on update for the UpdatedAt field.
	joblayer.UpdateDefaultUpdatedAt = joblayerDescUpdatedAt.UpdateDefault.(func() time.Time)
	jobownerFields := schema.JobOwner{}.Fields()
	_ = jobownerFields
	// jobownerDescName is the schema descriptor for Name field.
	jobownerDescName := jobownerFields[0].Descriptor()
	// jobowner.DefaultName holds the default value on creation for the Name field.
	jobowner.DefaultName = jobownerDescName.Default.(string)
	// jobownerDescShareholder is the schema descriptor for Shareholder field.
	jobownerDescShareholder := jobownerFields[8].Descriptor()
	// jobowner.DefaultShareholder holds the default value on creation for the Shareholder field.
	jobowner.DefaultShareholder = jobownerDescShareholder.Default.(bool)
	// jobownerDescCreatedAt is the schema descriptor for CreatedAt field.
	jobownerDescCreatedAt := jobownerFields[10].Descriptor()
	// jobowner.DefaultCreatedAt holds the default value on creation for the CreatedAt field.
	jobowner.DefaultCreatedAt = jobownerDescCreatedAt.Default.(func() time.Time)
	// jobownerDescUpdatedAt is the schema descriptor for UpdatedAt field.
	jobownerDescUpdatedAt := jobownerFields[11].Descriptor()
	// jobowner.DefaultUpdatedAt holds the default value on creation for the UpdatedAt field.
	jobowner.DefaultUpdatedAt = jobownerDescUpdatedAt.Default.(func() time.Time)
	// jobowner.UpdateDefaultUpdatedAt holds the default value on update for the UpdatedAt field.
	jobowner.UpdateDefaultUpdatedAt = jobownerDescUpdatedAt.UpdateDefault.(func() time.Time)
	jobpaymentsFields := schema.JobPayments{}.Fields()
	_ = jobpaymentsFields
	// jobpaymentsDescDate is the schema descriptor for Date field.
	jobpaymentsDescDate := jobpaymentsFields[0].Descriptor()
	// jobpayments.DefaultDate holds the default value on creation for the Date field.
	jobpayments.DefaultDate = jobpaymentsDescDate.Default.(func() time.Time)
	// jobpaymentsDescAmount is the schema descriptor for Amount field.
	jobpaymentsDescAmount := jobpaymentsFields[1].Descriptor()
	// jobpayments.DefaultAmount holds the default value on creation for the Amount field.
	jobpayments.DefaultAmount = jobpaymentsDescAmount.Default.(int)
	// jobpaymentsDescDescription is the schema descriptor for Description field.
	jobpaymentsDescDescription := jobpaymentsFields[2].Descriptor()
	// jobpayments.DefaultDescription holds the default value on creation for the Description field.
	jobpayments.DefaultDescription = jobpaymentsDescDescription.Default.(string)
	// jobpaymentsDescStatus is the schema descriptor for Status field.
	jobpaymentsDescStatus := jobpaymentsFields[3].Descriptor()
	// jobpayments.DefaultStatus holds the default value on creation for the Status field.
	jobpayments.DefaultStatus = jobpaymentsDescStatus.Default.(string)
	// jobpaymentsDescPercentage is the schema descriptor for Percentage field.
	jobpaymentsDescPercentage := jobpaymentsFields[4].Descriptor()
	// jobpayments.DefaultPercentage holds the default value on creation for the Percentage field.
	jobpayments.DefaultPercentage = jobpaymentsDescPercentage.Default.(float64)
	// jobpaymentsDescCreatedAt is the schema descriptor for CreatedAt field.
	jobpaymentsDescCreatedAt := jobpaymentsFields[5].Descriptor()
	// jobpayments.DefaultCreatedAt holds the default value on creation for the CreatedAt field.
	jobpayments.DefaultCreatedAt = jobpaymentsDescCreatedAt.Default.(func() time.Time)
	// jobpaymentsDescUpdatedAt is the schema descriptor for UpdatedAt field.
	jobpaymentsDescUpdatedAt := jobpaymentsFields[6].Descriptor()
	// jobpayments.DefaultUpdatedAt holds the default value on creation for the UpdatedAt field.
	jobpayments.DefaultUpdatedAt = jobpaymentsDescUpdatedAt.Default.(func() time.Time)
	// jobpayments.UpdateDefaultUpdatedAt holds the default value on update for the UpdatedAt field.
	jobpayments.UpdateDefaultUpdatedAt = jobpaymentsDescUpdatedAt.UpdateDefault.(func() time.Time)
	jobprogressFields := schema.JobProgress{}.Fields()
	_ = jobprogressFields
	// jobprogressDescOne is the schema descriptor for One field.
	jobprogressDescOne := jobprogressFields[0].Descriptor()
	// jobprogress.DefaultOne holds the default value on creation for the One field.
	jobprogress.DefaultOne = jobprogressDescOne.Default.(int)
	// jobprogressDescTwo is the schema descriptor for Two field.
	jobprogressDescTwo := jobprogressFields[1].Descriptor()
	// jobprogress.DefaultTwo holds the default value on creation for the Two field.
	jobprogress.DefaultTwo = jobprogressDescTwo.Default.(int)
	// jobprogressDescThree is the schema descriptor for Three field.
	jobprogressDescThree := jobprogressFields[2].Descriptor()
	// jobprogress.DefaultThree holds the default value on creation for the Three field.
	jobprogress.DefaultThree = jobprogressDescThree.Default.(int)
	// jobprogressDescFour is the schema descriptor for Four field.
	jobprogressDescFour := jobprogressFields[3].Descriptor()
	// jobprogress.DefaultFour holds the default value on creation for the Four field.
	jobprogress.DefaultFour = jobprogressDescFour.Default.(int)
	// jobprogressDescFive is the schema descriptor for Five field.
	jobprogressDescFive := jobprogressFields[4].Descriptor()
	// jobprogress.DefaultFive holds the default value on creation for the Five field.
	jobprogress.DefaultFive = jobprogressDescFive.Default.(int)
	// jobprogressDescSix is the schema descriptor for Six field.
	jobprogressDescSix := jobprogressFields[5].Descriptor()
	// jobprogress.DefaultSix holds the default value on creation for the Six field.
	jobprogress.DefaultSix = jobprogressDescSix.Default.(int)
	// jobprogressDescCreatedAt is the schema descriptor for CreatedAt field.
	jobprogressDescCreatedAt := jobprogressFields[6].Descriptor()
	// jobprogress.DefaultCreatedAt holds the default value on creation for the CreatedAt field.
	jobprogress.DefaultCreatedAt = jobprogressDescCreatedAt.Default.(func() time.Time)
	// jobprogressDescUpdatedAt is the schema descriptor for UpdatedAt field.
	jobprogressDescUpdatedAt := jobprogressFields[7].Descriptor()
	// jobprogress.DefaultUpdatedAt holds the default value on creation for the UpdatedAt field.
	jobprogress.DefaultUpdatedAt = jobprogressDescUpdatedAt.Default.(func() time.Time)
	// jobprogress.UpdateDefaultUpdatedAt holds the default value on update for the UpdatedAt field.
	jobprogress.UpdateDefaultUpdatedAt = jobprogressDescUpdatedAt.UpdateDefault.(func() time.Time)
	jobsupervisorFields := schema.JobSupervisor{}.Fields()
	_ = jobsupervisorFields
	// jobsupervisorDescCreatedAt is the schema descriptor for CreatedAt field.
	jobsupervisorDescCreatedAt := jobsupervisorFields[11].Descriptor()
	// jobsupervisor.DefaultCreatedAt holds the default value on creation for the CreatedAt field.
	jobsupervisor.DefaultCreatedAt = jobsupervisorDescCreatedAt.Default.(func() time.Time)
	// jobsupervisorDescUpdatedAt is the schema descriptor for UpdatedAt field.
	jobsupervisorDescUpdatedAt := jobsupervisorFields[12].Descriptor()
	// jobsupervisor.DefaultUpdatedAt holds the default value on creation for the UpdatedAt field.
	jobsupervisor.DefaultUpdatedAt = jobsupervisorDescUpdatedAt.Default.(func() time.Time)
	// jobsupervisor.UpdateDefaultUpdatedAt holds the default value on update for the UpdatedAt field.
	jobsupervisor.UpdateDefaultUpdatedAt = jobsupervisorDescUpdatedAt.UpdateDefault.(func() time.Time)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescName is the schema descriptor for Name field.
	userDescName := userFields[1].Descriptor()
	// user.DefaultName holds the default value on creation for the Name field.
	user.DefaultName = userDescName.Default.(string)
	// userDescEmail is the schema descriptor for Email field.
	userDescEmail := userFields[2].Descriptor()
	// user.DefaultEmail holds the default value on creation for the Email field.
	user.DefaultEmail = userDescEmail.Default.(string)
	// userDescRole is the schema descriptor for Role field.
	userDescRole := userFields[5].Descriptor()
	// user.DefaultRole holds the default value on creation for the Role field.
	user.DefaultRole = userDescRole.Default.(string)
	// userDescCreatedAt is the schema descriptor for CreatedAt field.
	userDescCreatedAt := userFields[6].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the CreatedAt field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for UpdatedAt field.
	userDescUpdatedAt := userFields[7].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the UpdatedAt field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the UpdatedAt field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
}
