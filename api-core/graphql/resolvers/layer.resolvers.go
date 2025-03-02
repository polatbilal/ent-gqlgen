package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.62

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/polatbilal/gqlgen-ent/api-core/ent"
	"github.com/polatbilal/gqlgen-ent/api-core/ent/jobdetail"
	"github.com/polatbilal/gqlgen-ent/api-core/ent/joblayer"
	"github.com/polatbilal/gqlgen-ent/api-core/ent/jobrelations"
	"github.com/polatbilal/gqlgen-ent/api-core/graphql/generated"
	"github.com/polatbilal/gqlgen-ent/api-core/graphql/model"
	"github.com/polatbilal/gqlgen-ent/api-core/middlewares"
)

// Job is the resolver for the Job field.
func (r *jobLayerResolver) Job(ctx context.Context, obj *ent.JobLayer) (*ent.JobDetail, error) {
	client := middlewares.GetClientFromContext(ctx)
	relations, err := client.JobRelations.Query().
		Where(jobrelations.HasLayersWith(joblayer.IDEQ(obj.ID))).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	return relations.QueryJob().Only(ctx)
}

// CreateLayer is the resolver for the createLayer field.
func (r *mutationResolver) CreateLayer(ctx context.Context, input model.JobLayerInput) (*ent.JobLayer, error) {
	client := middlewares.GetClientFromContext(ctx)

	uppercaseName := strings.ToUpper(*input.Name)

	layer, err := client.JobLayer.Create().
		SetName(uppercaseName).
		SetMetre(*input.Metre).
		SetNillableMoldDate(input.MoldDate).
		SetNillableConcreteDate(input.ConcreteDate).
		SetNillableSamples(input.Samples).
		SetNillableConcreteClass(input.ConcreteClass).
		SetNillableWeekResult(input.WeekResult).
		SetNillableMonthResult(input.MonthResult).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	// İş detayını bul
	jobDetail, err := client.JobDetail.Query().
		Where(jobdetail.YibfNoEQ(*input.YibfNo)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("iş detayı bulunamadı: %v", err)
	}

	// JobRelations'ı bul
	relations, err := jobDetail.QueryRelations().Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("iş ilişkileri bulunamadı: %v", err)
	}

	// Layer'ı JobRelations'a ekle
	_, err = relations.Update().
		AddLayers(layer).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("katman iş ilişkilerine eklenemedi: %v", err)
	}

	return layer, nil
}

// UpdateLayer is the resolver for the updateLayer field.
func (r *mutationResolver) UpdateLayer(ctx context.Context, id string, input model.JobLayerInput) (*ent.JobLayer, error) {
	client := middlewares.GetClientFromContext(ctx)

	layerID, err := strconv.Atoi(id)
	if err != nil {
		return nil, fmt.Errorf("failed to convert layer ID: %v", err)
	}

	// Update Layer
	layer, err := client.JobLayer.UpdateOneID(layerID).
		SetNillableName(input.Name).
		SetNillableMetre(input.Metre).
		SetNillableMoldDate(input.MoldDate).
		SetNillableConcreteDate(input.ConcreteDate).
		SetNillableSamples(input.Samples).
		SetNillableConcreteClass(input.ConcreteClass).
		SetNillableWeekResult(input.WeekResult).
		SetNillableMonthResult(input.MonthResult).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to update layer: %v", err)
	}

	return layer, nil
}

// DeleteLayer is the resolver for the deleteLayer field.
func (r *mutationResolver) DeleteLayer(ctx context.Context, id string) (*ent.JobLayer, error) {
	client := middlewares.GetClientFromContext(ctx)

	// ID'yi string'den int'e dönüştür
	layerID, err := strconv.Atoi(id)
	if err != nil {
		return nil, fmt.Errorf("geçersiz ID formatı: %v", err)
	}

	// İş katmanını ID ile sorgula
	layer, err := client.JobLayer.Query().Where(joblayer.IDEQ(layerID)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, fmt.Errorf("kayıt bulunamadı")
		}
		return nil, err
	}

	// İş katmanını sil
	err = client.JobLayer.DeleteOne(layer).Exec(ctx)
	if err != nil {
		return nil, fmt.Errorf("iş katmanı silinemedi: %v", err)
	}

	return layer, nil
}

// Layer is the resolver for the Layer field.
func (r *queryResolver) Layer(ctx context.Context, filter *model.LayerFilterInput) ([]*ent.JobLayer, error) {
	client := middlewares.GetClientFromContext(ctx)
	query := client.JobLayer.Query()

	if filter != nil {
		if filter.ID != nil {
			query = query.Where(joblayer.IDEQ(*filter.ID))
		}
		if filter.YibfNo != nil {
			// JobDetail -> JobRelations -> JobLayer ilişkisini kullan
			query = query.Where(
				joblayer.HasLayerWith(
					jobrelations.HasJobWith(
						jobdetail.YibfNoEQ(*filter.YibfNo),
					),
				),
			)
		}
	}

	layers, err := query.All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get layers: %v", err)
	}
	return layers, nil
}

// JobLayer returns generated.JobLayerResolver implementation.
func (r *Resolver) JobLayer() generated.JobLayerResolver { return &jobLayerResolver{r} }

type jobLayerResolver struct{ *Resolver }
