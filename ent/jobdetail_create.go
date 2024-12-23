// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"gqlgen-ent/ent/companyengineer"
	"gqlgen-ent/ent/jobauthor"
	"gqlgen-ent/ent/jobcontractor"
	"gqlgen-ent/ent/jobdetail"
	"gqlgen-ent/ent/joblayer"
	"gqlgen-ent/ent/jobowner"
	"gqlgen-ent/ent/jobprogress"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// JobDetailCreate is the builder for creating a JobDetail entity.
type JobDetailCreate struct {
	config
	mutation *JobDetailMutation
	hooks    []Hook
}

// SetYibfNo sets the "YibfNo" field.
func (jdc *JobDetailCreate) SetYibfNo(i int) *JobDetailCreate {
	jdc.mutation.SetYibfNo(i)
	return jdc
}

// SetProvince sets the "Province" field.
func (jdc *JobDetailCreate) SetProvince(s string) *JobDetailCreate {
	jdc.mutation.SetProvince(s)
	return jdc
}

// SetNillableProvince sets the "Province" field if the given value is not nil.
func (jdc *JobDetailCreate) SetNillableProvince(s *string) *JobDetailCreate {
	if s != nil {
		jdc.SetProvince(*s)
	}
	return jdc
}

// SetIdare sets the "Idare" field.
func (jdc *JobDetailCreate) SetIdare(s string) *JobDetailCreate {
	jdc.mutation.SetIdare(s)
	return jdc
}

// SetNillableIdare sets the "Idare" field if the given value is not nil.
func (jdc *JobDetailCreate) SetNillableIdare(s *string) *JobDetailCreate {
	if s != nil {
		jdc.SetIdare(*s)
	}
	return jdc
}

// SetPafta sets the "Pafta" field.
func (jdc *JobDetailCreate) SetPafta(s string) *JobDetailCreate {
	jdc.mutation.SetPafta(s)
	return jdc
}

// SetNillablePafta sets the "Pafta" field if the given value is not nil.
func (jdc *JobDetailCreate) SetNillablePafta(s *string) *JobDetailCreate {
	if s != nil {
		jdc.SetPafta(*s)
	}
	return jdc
}

// SetAda sets the "Ada" field.
func (jdc *JobDetailCreate) SetAda(s string) *JobDetailCreate {
	jdc.mutation.SetAda(s)
	return jdc
}

// SetNillableAda sets the "Ada" field if the given value is not nil.
func (jdc *JobDetailCreate) SetNillableAda(s *string) *JobDetailCreate {
	if s != nil {
		jdc.SetAda(*s)
	}
	return jdc
}

// SetParsel sets the "Parsel" field.
func (jdc *JobDetailCreate) SetParsel(s string) *JobDetailCreate {
	jdc.mutation.SetParsel(s)
	return jdc
}

// SetNillableParsel sets the "Parsel" field if the given value is not nil.
func (jdc *JobDetailCreate) SetNillableParsel(s *string) *JobDetailCreate {
	if s != nil {
		jdc.SetParsel(*s)
	}
	return jdc
}

// SetFolderNo sets the "FolderNo" field.
func (jdc *JobDetailCreate) SetFolderNo(s string) *JobDetailCreate {
	jdc.mutation.SetFolderNo(s)
	return jdc
}

// SetNillableFolderNo sets the "FolderNo" field if the given value is not nil.
func (jdc *JobDetailCreate) SetNillableFolderNo(s *string) *JobDetailCreate {
	if s != nil {
		jdc.SetFolderNo(*s)
	}
	return jdc
}

// SetStatus sets the "Status" field.
func (jdc *JobDetailCreate) SetStatus(i int) *JobDetailCreate {
	jdc.mutation.SetStatus(i)
	return jdc
}

// SetNillableStatus sets the "Status" field if the given value is not nil.
func (jdc *JobDetailCreate) SetNillableStatus(i *int) *JobDetailCreate {
	if i != nil {
		jdc.SetStatus(*i)
	}
	return jdc
}

// SetContractDate sets the "ContractDate" field.
func (jdc *JobDetailCreate) SetContractDate(t time.Time) *JobDetailCreate {
	jdc.mutation.SetContractDate(t)
	return jdc
}

// SetNillableContractDate sets the "ContractDate" field if the given value is not nil.
func (jdc *JobDetailCreate) SetNillableContractDate(t *time.Time) *JobDetailCreate {
	if t != nil {
		jdc.SetContractDate(*t)
	}
	return jdc
}

// SetStartDate sets the "StartDate" field.
func (jdc *JobDetailCreate) SetStartDate(t time.Time) *JobDetailCreate {
	jdc.mutation.SetStartDate(t)
	return jdc
}

// SetNillableStartDate sets the "StartDate" field if the given value is not nil.
func (jdc *JobDetailCreate) SetNillableStartDate(t *time.Time) *JobDetailCreate {
	if t != nil {
		jdc.SetStartDate(*t)
	}
	return jdc
}

// SetLicenseDate sets the "LicenseDate" field.
func (jdc *JobDetailCreate) SetLicenseDate(t time.Time) *JobDetailCreate {
	jdc.mutation.SetLicenseDate(t)
	return jdc
}

// SetNillableLicenseDate sets the "LicenseDate" field if the given value is not nil.
func (jdc *JobDetailCreate) SetNillableLicenseDate(t *time.Time) *JobDetailCreate {
	if t != nil {
		jdc.SetLicenseDate(*t)
	}
	return jdc
}

// SetLicenseNo sets the "LicenseNo" field.
func (jdc *JobDetailCreate) SetLicenseNo(s string) *JobDetailCreate {
	jdc.mutation.SetLicenseNo(s)
	return jdc
}

// SetNillableLicenseNo sets the "LicenseNo" field if the given value is not nil.
func (jdc *JobDetailCreate) SetNillableLicenseNo(s *string) *JobDetailCreate {
	if s != nil {
		jdc.SetLicenseNo(*s)
	}
	return jdc
}

// SetConstructionArea sets the "ConstructionArea" field.
func (jdc *JobDetailCreate) SetConstructionArea(s string) *JobDetailCreate {
	jdc.mutation.SetConstructionArea(s)
	return jdc
}

// SetNillableConstructionArea sets the "ConstructionArea" field if the given value is not nil.
func (jdc *JobDetailCreate) SetNillableConstructionArea(s *string) *JobDetailCreate {
	if s != nil {
		jdc.SetConstructionArea(*s)
	}
	return jdc
}

// SetDistrict sets the "District" field.
func (jdc *JobDetailCreate) SetDistrict(s string) *JobDetailCreate {
	jdc.mutation.SetDistrict(s)
	return jdc
}

// SetNillableDistrict sets the "District" field if the given value is not nil.
func (jdc *JobDetailCreate) SetNillableDistrict(s *string) *JobDetailCreate {
	if s != nil {
		jdc.SetDistrict(*s)
	}
	return jdc
}

// SetVillage sets the "Village" field.
func (jdc *JobDetailCreate) SetVillage(s string) *JobDetailCreate {
	jdc.mutation.SetVillage(s)
	return jdc
}

// SetNillableVillage sets the "Village" field if the given value is not nil.
func (jdc *JobDetailCreate) SetNillableVillage(s *string) *JobDetailCreate {
	if s != nil {
		jdc.SetVillage(*s)
	}
	return jdc
}

// SetStreet sets the "Street" field.
func (jdc *JobDetailCreate) SetStreet(s string) *JobDetailCreate {
	jdc.mutation.SetStreet(s)
	return jdc
}

// SetNillableStreet sets the "Street" field if the given value is not nil.
func (jdc *JobDetailCreate) SetNillableStreet(s *string) *JobDetailCreate {
	if s != nil {
		jdc.SetStreet(*s)
	}
	return jdc
}

// SetBuildingClass sets the "BuildingClass" field.
func (jdc *JobDetailCreate) SetBuildingClass(s string) *JobDetailCreate {
	jdc.mutation.SetBuildingClass(s)
	return jdc
}

// SetNillableBuildingClass sets the "BuildingClass" field if the given value is not nil.
func (jdc *JobDetailCreate) SetNillableBuildingClass(s *string) *JobDetailCreate {
	if s != nil {
		jdc.SetBuildingClass(*s)
	}
	return jdc
}

// SetBuildingType sets the "BuildingType" field.
func (jdc *JobDetailCreate) SetBuildingType(s string) *JobDetailCreate {
	jdc.mutation.SetBuildingType(s)
	return jdc
}

// SetNillableBuildingType sets the "BuildingType" field if the given value is not nil.
func (jdc *JobDetailCreate) SetNillableBuildingType(s *string) *JobDetailCreate {
	if s != nil {
		jdc.SetBuildingType(*s)
	}
	return jdc
}

// SetBuildingBlock sets the "BuildingBlock" field.
func (jdc *JobDetailCreate) SetBuildingBlock(s string) *JobDetailCreate {
	jdc.mutation.SetBuildingBlock(s)
	return jdc
}

// SetNillableBuildingBlock sets the "BuildingBlock" field if the given value is not nil.
func (jdc *JobDetailCreate) SetNillableBuildingBlock(s *string) *JobDetailCreate {
	if s != nil {
		jdc.SetBuildingBlock(*s)
	}
	return jdc
}

// SetLandArea sets the "LandArea" field.
func (jdc *JobDetailCreate) SetLandArea(s string) *JobDetailCreate {
	jdc.mutation.SetLandArea(s)
	return jdc
}

// SetNillableLandArea sets the "LandArea" field if the given value is not nil.
func (jdc *JobDetailCreate) SetNillableLandArea(s *string) *JobDetailCreate {
	if s != nil {
		jdc.SetLandArea(*s)
	}
	return jdc
}

// SetFloors sets the "Floors" field.
func (jdc *JobDetailCreate) SetFloors(i int) *JobDetailCreate {
	jdc.mutation.SetFloors(i)
	return jdc
}

// SetNillableFloors sets the "Floors" field if the given value is not nil.
func (jdc *JobDetailCreate) SetNillableFloors(i *int) *JobDetailCreate {
	if i != nil {
		jdc.SetFloors(*i)
	}
	return jdc
}

// SetUsagePurpose sets the "UsagePurpose" field.
func (jdc *JobDetailCreate) SetUsagePurpose(s string) *JobDetailCreate {
	jdc.mutation.SetUsagePurpose(s)
	return jdc
}

// SetNillableUsagePurpose sets the "UsagePurpose" field if the given value is not nil.
func (jdc *JobDetailCreate) SetNillableUsagePurpose(s *string) *JobDetailCreate {
	if s != nil {
		jdc.SetUsagePurpose(*s)
	}
	return jdc
}

// SetNote sets the "Note" field.
func (jdc *JobDetailCreate) SetNote(s string) *JobDetailCreate {
	jdc.mutation.SetNote(s)
	return jdc
}

// SetNillableNote sets the "Note" field if the given value is not nil.
func (jdc *JobDetailCreate) SetNillableNote(s *string) *JobDetailCreate {
	if s != nil {
		jdc.SetNote(*s)
	}
	return jdc
}

// SetStarted sets the "Started" field.
func (jdc *JobDetailCreate) SetStarted(i int) *JobDetailCreate {
	jdc.mutation.SetStarted(i)
	return jdc
}

// SetNillableStarted sets the "Started" field if the given value is not nil.
func (jdc *JobDetailCreate) SetNillableStarted(i *int) *JobDetailCreate {
	if i != nil {
		jdc.SetStarted(*i)
	}
	return jdc
}

// SetDeleted sets the "Deleted" field.
func (jdc *JobDetailCreate) SetDeleted(i int) *JobDetailCreate {
	jdc.mutation.SetDeleted(i)
	return jdc
}

// SetNillableDeleted sets the "Deleted" field if the given value is not nil.
func (jdc *JobDetailCreate) SetNillableDeleted(i *int) *JobDetailCreate {
	if i != nil {
		jdc.SetDeleted(*i)
	}
	return jdc
}

// SetCreatedAt sets the "created_at" field.
func (jdc *JobDetailCreate) SetCreatedAt(t time.Time) *JobDetailCreate {
	jdc.mutation.SetCreatedAt(t)
	return jdc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (jdc *JobDetailCreate) SetNillableCreatedAt(t *time.Time) *JobDetailCreate {
	if t != nil {
		jdc.SetCreatedAt(*t)
	}
	return jdc
}

// SetUpdatedAt sets the "updated_at" field.
func (jdc *JobDetailCreate) SetUpdatedAt(t time.Time) *JobDetailCreate {
	jdc.mutation.SetUpdatedAt(t)
	return jdc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (jdc *JobDetailCreate) SetNillableUpdatedAt(t *time.Time) *JobDetailCreate {
	if t != nil {
		jdc.SetUpdatedAt(*t)
	}
	return jdc
}

// SetOwnerID sets the "owner" edge to the JobOwner entity by ID.
func (jdc *JobDetailCreate) SetOwnerID(id int) *JobDetailCreate {
	jdc.mutation.SetOwnerID(id)
	return jdc
}

// SetNillableOwnerID sets the "owner" edge to the JobOwner entity by ID if the given value is not nil.
func (jdc *JobDetailCreate) SetNillableOwnerID(id *int) *JobDetailCreate {
	if id != nil {
		jdc = jdc.SetOwnerID(*id)
	}
	return jdc
}

// SetOwner sets the "owner" edge to the JobOwner entity.
func (jdc *JobDetailCreate) SetOwner(j *JobOwner) *JobDetailCreate {
	return jdc.SetOwnerID(j.ID)
}

// SetContractorID sets the "contractor" edge to the JobContractor entity by ID.
func (jdc *JobDetailCreate) SetContractorID(id int) *JobDetailCreate {
	jdc.mutation.SetContractorID(id)
	return jdc
}

// SetNillableContractorID sets the "contractor" edge to the JobContractor entity by ID if the given value is not nil.
func (jdc *JobDetailCreate) SetNillableContractorID(id *int) *JobDetailCreate {
	if id != nil {
		jdc = jdc.SetContractorID(*id)
	}
	return jdc
}

// SetContractor sets the "contractor" edge to the JobContractor entity.
func (jdc *JobDetailCreate) SetContractor(j *JobContractor) *JobDetailCreate {
	return jdc.SetContractorID(j.ID)
}

// SetAuthorID sets the "author" edge to the JobAuthor entity by ID.
func (jdc *JobDetailCreate) SetAuthorID(id int) *JobDetailCreate {
	jdc.mutation.SetAuthorID(id)
	return jdc
}

// SetNillableAuthorID sets the "author" edge to the JobAuthor entity by ID if the given value is not nil.
func (jdc *JobDetailCreate) SetNillableAuthorID(id *int) *JobDetailCreate {
	if id != nil {
		jdc = jdc.SetAuthorID(*id)
	}
	return jdc
}

// SetAuthor sets the "author" edge to the JobAuthor entity.
func (jdc *JobDetailCreate) SetAuthor(j *JobAuthor) *JobDetailCreate {
	return jdc.SetAuthorID(j.ID)
}

// SetProgressID sets the "progress" edge to the JobProgress entity by ID.
func (jdc *JobDetailCreate) SetProgressID(id int) *JobDetailCreate {
	jdc.mutation.SetProgressID(id)
	return jdc
}

// SetNillableProgressID sets the "progress" edge to the JobProgress entity by ID if the given value is not nil.
func (jdc *JobDetailCreate) SetNillableProgressID(id *int) *JobDetailCreate {
	if id != nil {
		jdc = jdc.SetProgressID(*id)
	}
	return jdc
}

// SetProgress sets the "progress" edge to the JobProgress entity.
func (jdc *JobDetailCreate) SetProgress(j *JobProgress) *JobDetailCreate {
	return jdc.SetProgressID(j.ID)
}

// SetInspectorID sets the "inspector" edge to the CompanyEngineer entity by ID.
func (jdc *JobDetailCreate) SetInspectorID(id int) *JobDetailCreate {
	jdc.mutation.SetInspectorID(id)
	return jdc
}

// SetNillableInspectorID sets the "inspector" edge to the CompanyEngineer entity by ID if the given value is not nil.
func (jdc *JobDetailCreate) SetNillableInspectorID(id *int) *JobDetailCreate {
	if id != nil {
		jdc = jdc.SetInspectorID(*id)
	}
	return jdc
}

// SetInspector sets the "inspector" edge to the CompanyEngineer entity.
func (jdc *JobDetailCreate) SetInspector(c *CompanyEngineer) *JobDetailCreate {
	return jdc.SetInspectorID(c.ID)
}

// SetArchitectID sets the "architect" edge to the CompanyEngineer entity by ID.
func (jdc *JobDetailCreate) SetArchitectID(id int) *JobDetailCreate {
	jdc.mutation.SetArchitectID(id)
	return jdc
}

// SetNillableArchitectID sets the "architect" edge to the CompanyEngineer entity by ID if the given value is not nil.
func (jdc *JobDetailCreate) SetNillableArchitectID(id *int) *JobDetailCreate {
	if id != nil {
		jdc = jdc.SetArchitectID(*id)
	}
	return jdc
}

// SetArchitect sets the "architect" edge to the CompanyEngineer entity.
func (jdc *JobDetailCreate) SetArchitect(c *CompanyEngineer) *JobDetailCreate {
	return jdc.SetArchitectID(c.ID)
}

// SetStaticID sets the "static" edge to the CompanyEngineer entity by ID.
func (jdc *JobDetailCreate) SetStaticID(id int) *JobDetailCreate {
	jdc.mutation.SetStaticID(id)
	return jdc
}

// SetNillableStaticID sets the "static" edge to the CompanyEngineer entity by ID if the given value is not nil.
func (jdc *JobDetailCreate) SetNillableStaticID(id *int) *JobDetailCreate {
	if id != nil {
		jdc = jdc.SetStaticID(*id)
	}
	return jdc
}

// SetStatic sets the "static" edge to the CompanyEngineer entity.
func (jdc *JobDetailCreate) SetStatic(c *CompanyEngineer) *JobDetailCreate {
	return jdc.SetStaticID(c.ID)
}

// SetMechanicID sets the "mechanic" edge to the CompanyEngineer entity by ID.
func (jdc *JobDetailCreate) SetMechanicID(id int) *JobDetailCreate {
	jdc.mutation.SetMechanicID(id)
	return jdc
}

// SetNillableMechanicID sets the "mechanic" edge to the CompanyEngineer entity by ID if the given value is not nil.
func (jdc *JobDetailCreate) SetNillableMechanicID(id *int) *JobDetailCreate {
	if id != nil {
		jdc = jdc.SetMechanicID(*id)
	}
	return jdc
}

// SetMechanic sets the "mechanic" edge to the CompanyEngineer entity.
func (jdc *JobDetailCreate) SetMechanic(c *CompanyEngineer) *JobDetailCreate {
	return jdc.SetMechanicID(c.ID)
}

// SetElectricID sets the "electric" edge to the CompanyEngineer entity by ID.
func (jdc *JobDetailCreate) SetElectricID(id int) *JobDetailCreate {
	jdc.mutation.SetElectricID(id)
	return jdc
}

// SetNillableElectricID sets the "electric" edge to the CompanyEngineer entity by ID if the given value is not nil.
func (jdc *JobDetailCreate) SetNillableElectricID(id *int) *JobDetailCreate {
	if id != nil {
		jdc = jdc.SetElectricID(*id)
	}
	return jdc
}

// SetElectric sets the "electric" edge to the CompanyEngineer entity.
func (jdc *JobDetailCreate) SetElectric(c *CompanyEngineer) *JobDetailCreate {
	return jdc.SetElectricID(c.ID)
}

// SetControllerID sets the "controller" edge to the CompanyEngineer entity by ID.
func (jdc *JobDetailCreate) SetControllerID(id int) *JobDetailCreate {
	jdc.mutation.SetControllerID(id)
	return jdc
}

// SetNillableControllerID sets the "controller" edge to the CompanyEngineer entity by ID if the given value is not nil.
func (jdc *JobDetailCreate) SetNillableControllerID(id *int) *JobDetailCreate {
	if id != nil {
		jdc = jdc.SetControllerID(*id)
	}
	return jdc
}

// SetController sets the "controller" edge to the CompanyEngineer entity.
func (jdc *JobDetailCreate) SetController(c *CompanyEngineer) *JobDetailCreate {
	return jdc.SetControllerID(c.ID)
}

// SetMechaniccontrollerID sets the "mechaniccontroller" edge to the CompanyEngineer entity by ID.
func (jdc *JobDetailCreate) SetMechaniccontrollerID(id int) *JobDetailCreate {
	jdc.mutation.SetMechaniccontrollerID(id)
	return jdc
}

// SetNillableMechaniccontrollerID sets the "mechaniccontroller" edge to the CompanyEngineer entity by ID if the given value is not nil.
func (jdc *JobDetailCreate) SetNillableMechaniccontrollerID(id *int) *JobDetailCreate {
	if id != nil {
		jdc = jdc.SetMechaniccontrollerID(*id)
	}
	return jdc
}

// SetMechaniccontroller sets the "mechaniccontroller" edge to the CompanyEngineer entity.
func (jdc *JobDetailCreate) SetMechaniccontroller(c *CompanyEngineer) *JobDetailCreate {
	return jdc.SetMechaniccontrollerID(c.ID)
}

// SetElectriccontrollerID sets the "electriccontroller" edge to the CompanyEngineer entity by ID.
func (jdc *JobDetailCreate) SetElectriccontrollerID(id int) *JobDetailCreate {
	jdc.mutation.SetElectriccontrollerID(id)
	return jdc
}

// SetNillableElectriccontrollerID sets the "electriccontroller" edge to the CompanyEngineer entity by ID if the given value is not nil.
func (jdc *JobDetailCreate) SetNillableElectriccontrollerID(id *int) *JobDetailCreate {
	if id != nil {
		jdc = jdc.SetElectriccontrollerID(*id)
	}
	return jdc
}

// SetElectriccontroller sets the "electriccontroller" edge to the CompanyEngineer entity.
func (jdc *JobDetailCreate) SetElectriccontroller(c *CompanyEngineer) *JobDetailCreate {
	return jdc.SetElectriccontrollerID(c.ID)
}

// AddLayerIDs adds the "layers" edge to the JobLayer entity by IDs.
func (jdc *JobDetailCreate) AddLayerIDs(ids ...int) *JobDetailCreate {
	jdc.mutation.AddLayerIDs(ids...)
	return jdc
}

// AddLayers adds the "layers" edges to the JobLayer entity.
func (jdc *JobDetailCreate) AddLayers(j ...*JobLayer) *JobDetailCreate {
	ids := make([]int, len(j))
	for i := range j {
		ids[i] = j[i].ID
	}
	return jdc.AddLayerIDs(ids...)
}

// Mutation returns the JobDetailMutation object of the builder.
func (jdc *JobDetailCreate) Mutation() *JobDetailMutation {
	return jdc.mutation
}

// Save creates the JobDetail in the database.
func (jdc *JobDetailCreate) Save(ctx context.Context) (*JobDetail, error) {
	jdc.defaults()
	return withHooks(ctx, jdc.sqlSave, jdc.mutation, jdc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (jdc *JobDetailCreate) SaveX(ctx context.Context) *JobDetail {
	v, err := jdc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (jdc *JobDetailCreate) Exec(ctx context.Context) error {
	_, err := jdc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (jdc *JobDetailCreate) ExecX(ctx context.Context) {
	if err := jdc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (jdc *JobDetailCreate) defaults() {
	if _, ok := jdc.mutation.Province(); !ok {
		v := jobdetail.DefaultProvince
		jdc.mutation.SetProvince(v)
	}
	if _, ok := jdc.mutation.Idare(); !ok {
		v := jobdetail.DefaultIdare
		jdc.mutation.SetIdare(v)
	}
	if _, ok := jdc.mutation.Pafta(); !ok {
		v := jobdetail.DefaultPafta
		jdc.mutation.SetPafta(v)
	}
	if _, ok := jdc.mutation.Ada(); !ok {
		v := jobdetail.DefaultAda
		jdc.mutation.SetAda(v)
	}
	if _, ok := jdc.mutation.Parsel(); !ok {
		v := jobdetail.DefaultParsel
		jdc.mutation.SetParsel(v)
	}
	if _, ok := jdc.mutation.FolderNo(); !ok {
		v := jobdetail.DefaultFolderNo
		jdc.mutation.SetFolderNo(v)
	}
	if _, ok := jdc.mutation.Status(); !ok {
		v := jobdetail.DefaultStatus
		jdc.mutation.SetStatus(v)
	}
	if _, ok := jdc.mutation.LicenseNo(); !ok {
		v := jobdetail.DefaultLicenseNo
		jdc.mutation.SetLicenseNo(v)
	}
	if _, ok := jdc.mutation.District(); !ok {
		v := jobdetail.DefaultDistrict
		jdc.mutation.SetDistrict(v)
	}
	if _, ok := jdc.mutation.Village(); !ok {
		v := jobdetail.DefaultVillage
		jdc.mutation.SetVillage(v)
	}
	if _, ok := jdc.mutation.Street(); !ok {
		v := jobdetail.DefaultStreet
		jdc.mutation.SetStreet(v)
	}
	if _, ok := jdc.mutation.BuildingClass(); !ok {
		v := jobdetail.DefaultBuildingClass
		jdc.mutation.SetBuildingClass(v)
	}
	if _, ok := jdc.mutation.BuildingType(); !ok {
		v := jobdetail.DefaultBuildingType
		jdc.mutation.SetBuildingType(v)
	}
	if _, ok := jdc.mutation.BuildingBlock(); !ok {
		v := jobdetail.DefaultBuildingBlock
		jdc.mutation.SetBuildingBlock(v)
	}
	if _, ok := jdc.mutation.UsagePurpose(); !ok {
		v := jobdetail.DefaultUsagePurpose
		jdc.mutation.SetUsagePurpose(v)
	}
	if _, ok := jdc.mutation.Started(); !ok {
		v := jobdetail.DefaultStarted
		jdc.mutation.SetStarted(v)
	}
	if _, ok := jdc.mutation.Deleted(); !ok {
		v := jobdetail.DefaultDeleted
		jdc.mutation.SetDeleted(v)
	}
	if _, ok := jdc.mutation.CreatedAt(); !ok {
		v := jobdetail.DefaultCreatedAt()
		jdc.mutation.SetCreatedAt(v)
	}
	if _, ok := jdc.mutation.UpdatedAt(); !ok {
		v := jobdetail.DefaultUpdatedAt()
		jdc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (jdc *JobDetailCreate) check() error {
	if _, ok := jdc.mutation.YibfNo(); !ok {
		return &ValidationError{Name: "YibfNo", err: errors.New(`ent: missing required field "JobDetail.YibfNo"`)}
	}
	if v, ok := jdc.mutation.YibfNo(); ok {
		if err := jobdetail.YibfNoValidator(v); err != nil {
			return &ValidationError{Name: "YibfNo", err: fmt.Errorf(`ent: validator failed for field "JobDetail.YibfNo": %w`, err)}
		}
	}
	if _, ok := jdc.mutation.Status(); !ok {
		return &ValidationError{Name: "Status", err: errors.New(`ent: missing required field "JobDetail.Status"`)}
	}
	if _, ok := jdc.mutation.Started(); !ok {
		return &ValidationError{Name: "Started", err: errors.New(`ent: missing required field "JobDetail.Started"`)}
	}
	if _, ok := jdc.mutation.Deleted(); !ok {
		return &ValidationError{Name: "Deleted", err: errors.New(`ent: missing required field "JobDetail.Deleted"`)}
	}
	if _, ok := jdc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "JobDetail.created_at"`)}
	}
	if _, ok := jdc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "JobDetail.updated_at"`)}
	}
	return nil
}

func (jdc *JobDetailCreate) sqlSave(ctx context.Context) (*JobDetail, error) {
	if err := jdc.check(); err != nil {
		return nil, err
	}
	_node, _spec := jdc.createSpec()
	if err := sqlgraph.CreateNode(ctx, jdc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	jdc.mutation.id = &_node.ID
	jdc.mutation.done = true
	return _node, nil
}

func (jdc *JobDetailCreate) createSpec() (*JobDetail, *sqlgraph.CreateSpec) {
	var (
		_node = &JobDetail{config: jdc.config}
		_spec = sqlgraph.NewCreateSpec(jobdetail.Table, sqlgraph.NewFieldSpec(jobdetail.FieldID, field.TypeInt))
	)
	if value, ok := jdc.mutation.YibfNo(); ok {
		_spec.SetField(jobdetail.FieldYibfNo, field.TypeInt, value)
		_node.YibfNo = value
	}
	if value, ok := jdc.mutation.Province(); ok {
		_spec.SetField(jobdetail.FieldProvince, field.TypeString, value)
		_node.Province = value
	}
	if value, ok := jdc.mutation.Idare(); ok {
		_spec.SetField(jobdetail.FieldIdare, field.TypeString, value)
		_node.Idare = value
	}
	if value, ok := jdc.mutation.Pafta(); ok {
		_spec.SetField(jobdetail.FieldPafta, field.TypeString, value)
		_node.Pafta = value
	}
	if value, ok := jdc.mutation.Ada(); ok {
		_spec.SetField(jobdetail.FieldAda, field.TypeString, value)
		_node.Ada = value
	}
	if value, ok := jdc.mutation.Parsel(); ok {
		_spec.SetField(jobdetail.FieldParsel, field.TypeString, value)
		_node.Parsel = value
	}
	if value, ok := jdc.mutation.FolderNo(); ok {
		_spec.SetField(jobdetail.FieldFolderNo, field.TypeString, value)
		_node.FolderNo = value
	}
	if value, ok := jdc.mutation.Status(); ok {
		_spec.SetField(jobdetail.FieldStatus, field.TypeInt, value)
		_node.Status = value
	}
	if value, ok := jdc.mutation.ContractDate(); ok {
		_spec.SetField(jobdetail.FieldContractDate, field.TypeTime, value)
		_node.ContractDate = value
	}
	if value, ok := jdc.mutation.StartDate(); ok {
		_spec.SetField(jobdetail.FieldStartDate, field.TypeTime, value)
		_node.StartDate = value
	}
	if value, ok := jdc.mutation.LicenseDate(); ok {
		_spec.SetField(jobdetail.FieldLicenseDate, field.TypeTime, value)
		_node.LicenseDate = value
	}
	if value, ok := jdc.mutation.LicenseNo(); ok {
		_spec.SetField(jobdetail.FieldLicenseNo, field.TypeString, value)
		_node.LicenseNo = value
	}
	if value, ok := jdc.mutation.ConstructionArea(); ok {
		_spec.SetField(jobdetail.FieldConstructionArea, field.TypeString, value)
		_node.ConstructionArea = value
	}
	if value, ok := jdc.mutation.District(); ok {
		_spec.SetField(jobdetail.FieldDistrict, field.TypeString, value)
		_node.District = value
	}
	if value, ok := jdc.mutation.Village(); ok {
		_spec.SetField(jobdetail.FieldVillage, field.TypeString, value)
		_node.Village = value
	}
	if value, ok := jdc.mutation.Street(); ok {
		_spec.SetField(jobdetail.FieldStreet, field.TypeString, value)
		_node.Street = value
	}
	if value, ok := jdc.mutation.BuildingClass(); ok {
		_spec.SetField(jobdetail.FieldBuildingClass, field.TypeString, value)
		_node.BuildingClass = value
	}
	if value, ok := jdc.mutation.BuildingType(); ok {
		_spec.SetField(jobdetail.FieldBuildingType, field.TypeString, value)
		_node.BuildingType = value
	}
	if value, ok := jdc.mutation.BuildingBlock(); ok {
		_spec.SetField(jobdetail.FieldBuildingBlock, field.TypeString, value)
		_node.BuildingBlock = value
	}
	if value, ok := jdc.mutation.LandArea(); ok {
		_spec.SetField(jobdetail.FieldLandArea, field.TypeString, value)
		_node.LandArea = value
	}
	if value, ok := jdc.mutation.Floors(); ok {
		_spec.SetField(jobdetail.FieldFloors, field.TypeInt, value)
		_node.Floors = value
	}
	if value, ok := jdc.mutation.UsagePurpose(); ok {
		_spec.SetField(jobdetail.FieldUsagePurpose, field.TypeString, value)
		_node.UsagePurpose = value
	}
	if value, ok := jdc.mutation.Note(); ok {
		_spec.SetField(jobdetail.FieldNote, field.TypeString, value)
		_node.Note = value
	}
	if value, ok := jdc.mutation.Started(); ok {
		_spec.SetField(jobdetail.FieldStarted, field.TypeInt, value)
		_node.Started = value
	}
	if value, ok := jdc.mutation.Deleted(); ok {
		_spec.SetField(jobdetail.FieldDeleted, field.TypeInt, value)
		_node.Deleted = value
	}
	if value, ok := jdc.mutation.CreatedAt(); ok {
		_spec.SetField(jobdetail.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := jdc.mutation.UpdatedAt(); ok {
		_spec.SetField(jobdetail.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := jdc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   jobdetail.OwnerTable,
			Columns: []string{jobdetail.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(jobowner.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.owner_id = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := jdc.mutation.ContractorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   jobdetail.ContractorTable,
			Columns: []string{jobdetail.ContractorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(jobcontractor.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.contractor_id = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := jdc.mutation.AuthorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   jobdetail.AuthorTable,
			Columns: []string{jobdetail.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(jobauthor.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.author_id = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := jdc.mutation.ProgressIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   jobdetail.ProgressTable,
			Columns: []string{jobdetail.ProgressColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(jobprogress.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.progress_id = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := jdc.mutation.InspectorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   jobdetail.InspectorTable,
			Columns: []string{jobdetail.InspectorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(companyengineer.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.inspector_id = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := jdc.mutation.ArchitectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   jobdetail.ArchitectTable,
			Columns: []string{jobdetail.ArchitectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(companyengineer.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.architect_id = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := jdc.mutation.StaticIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   jobdetail.StaticTable,
			Columns: []string{jobdetail.StaticColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(companyengineer.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.static_id = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := jdc.mutation.MechanicIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   jobdetail.MechanicTable,
			Columns: []string{jobdetail.MechanicColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(companyengineer.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.mechanic_id = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := jdc.mutation.ElectricIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   jobdetail.ElectricTable,
			Columns: []string{jobdetail.ElectricColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(companyengineer.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.electric_id = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := jdc.mutation.ControllerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   jobdetail.ControllerTable,
			Columns: []string{jobdetail.ControllerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(companyengineer.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.controller_id = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := jdc.mutation.MechaniccontrollerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   jobdetail.MechaniccontrollerTable,
			Columns: []string{jobdetail.MechaniccontrollerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(companyengineer.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.mechaniccontroller_id = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := jdc.mutation.ElectriccontrollerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   jobdetail.ElectriccontrollerTable,
			Columns: []string{jobdetail.ElectriccontrollerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(companyengineer.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.electriccontroller_id = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := jdc.mutation.LayersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   jobdetail.LayersTable,
			Columns: []string{jobdetail.LayersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(joblayer.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// JobDetailCreateBulk is the builder for creating many JobDetail entities in bulk.
type JobDetailCreateBulk struct {
	config
	err      error
	builders []*JobDetailCreate
}

// Save creates the JobDetail entities in the database.
func (jdcb *JobDetailCreateBulk) Save(ctx context.Context) ([]*JobDetail, error) {
	if jdcb.err != nil {
		return nil, jdcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(jdcb.builders))
	nodes := make([]*JobDetail, len(jdcb.builders))
	mutators := make([]Mutator, len(jdcb.builders))
	for i := range jdcb.builders {
		func(i int, root context.Context) {
			builder := jdcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*JobDetailMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, jdcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, jdcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, jdcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (jdcb *JobDetailCreateBulk) SaveX(ctx context.Context) []*JobDetail {
	v, err := jdcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (jdcb *JobDetailCreateBulk) Exec(ctx context.Context) error {
	_, err := jdcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (jdcb *JobDetailCreateBulk) ExecX(ctx context.Context) {
	if err := jdcb.Exec(ctx); err != nil {
		panic(err)
	}
}
