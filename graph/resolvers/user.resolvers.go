package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.62

import (
	"context"
	"fmt"
	"strconv"

	"github.com/polatbilal/gqlgen-ent/ent"
	"github.com/polatbilal/gqlgen-ent/ent/companyuser"
	"github.com/polatbilal/gqlgen-ent/ent/user"
	"github.com/polatbilal/gqlgen-ent/graph/generated"
	"github.com/polatbilal/gqlgen-ent/graph/model"
	"github.com/polatbilal/gqlgen-ent/middlewares"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.UserInput) (*ent.User, error) {
	client := middlewares.GetClientFromContext(ctx)

	// Kullanıcı adının daha önce kullanılıp kullanılmadığını kontrol et
	exists, err := client.User.Query().
		Where(user.Username(input.Username)).
		Exist(ctx)
	if err != nil {
		return nil, fmt.Errorf("kullanıcı kontrolü yapılırken hata: %v", err)
	}
	if exists {
		return nil, fmt.Errorf("bu kullanıcı adı zaten kullanımda: %s", input.Username)
	}

	// Önce kullanıcıyı oluştur
	user, err := client.User.Create().
		SetUsername(input.Username).
		SetName(input.Name).
		SetEmail(input.Email).
		SetNillablePhone(input.Phone).
		SetPassword(input.Password).
		SetNillableRole(input.Role).
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("kullanıcı oluşturulurken hata: %v", err)
	}

	// Şirketleri ekle
	if len(input.CompanyIDs) > 0 {
		for _, companyID := range input.CompanyIDs {
			// CompanyUser oluştur
			_, err = client.CompanyUser.Create().
				SetUser(user).
				SetCompanyID(companyID).
				Save(ctx)

			if err != nil {
				return nil, fmt.Errorf("şirket bağlantısı oluşturulurken hata: %v", err)
			}
		}
	}

	return user, nil
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, id string, input model.UserInput) (*ent.User, error) {
	client := middlewares.GetClientFromContext(ctx)

	intID, err := strconv.Atoi(id)
	if err != nil {
		return nil, fmt.Errorf("kullanıcı bulunurken hata: %v", err)
	}

	getUser, err := client.User.Query().Where(user.IDEQ(intID)).Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("kullanıcı bulunurken hata: %v", err)
	}

	user, err := client.User.UpdateOneID(getUser.ID).
		SetUsername(input.Username).
		SetName(input.Name).
		SetEmail(input.Email).
		SetNillablePhone(input.Phone).
		SetPassword(input.Password).
		SetNillableRole(input.Role).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("kullanıcı güncellenirken hata: %v", err)
	}

	// Şirketleri ekle
	if len(input.CompanyIDs) > 0 {
		// Kullanıcının mevcut şirket bağlantılarını al
		existingCompanies, err := user.QueryCompanies().QueryCompany().IDs(ctx)
		if err != nil {
			return nil, fmt.Errorf("mevcut şirket bağlantıları kontrol edilirken hata: %v", err)
		}

		// Mevcut şirketleri map'e dönüştür (hızlı arama için)
		existingCompanyMap := make(map[int]bool)
		for _, id := range existingCompanies {
			existingCompanyMap[id] = true
		}

		// Sadece yeni şirketler için bağlantı oluştur
		for _, companyID := range input.CompanyIDs {
			// Eğer bu şirket bağlantısı zaten varsa, atla
			if existingCompanyMap[companyID] {
				continue
			}

			// CompanyUser oluştur
			_, err = client.CompanyUser.Create().
				SetUser(user).
				SetCompanyID(companyID).
				Save(ctx)

			if err != nil {
				return nil, fmt.Errorf("şirket bağlantısı oluşturulurken hata: %v", err)
			}
		}
	}

	return user, nil
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (bool, error) {
	client := middlewares.GetClientFromContext(ctx)
	intID, err := strconv.Atoi(id)
	if err != nil {
		return false, fmt.Errorf("kullanıcı bulunurken hata: %v", err)
	}

	// Önce kullanıcının CompanyUser kayıtlarını sil
	_, err = client.CompanyUser.Delete().Where(companyuser.HasUserWith(user.IDEQ(intID))).Exec(ctx)
	if err != nil {
		return false, fmt.Errorf("kullanıcının şirket bağlantıları silinirken hata: %v", err)
	}

	// Sonra kullanıcıyı sil
	err = client.User.DeleteOneID(intID).Exec(ctx)
	return err == nil, err
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*ent.User, error) {
	client := middlewares.GetClientFromContext(ctx)
	intID, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return client.User.Query().Where(user.IDEQ(intID)).Only(ctx)
}

// AllUsers is the resolver for the allUsers field.
func (r *queryResolver) AllUsers(ctx context.Context) ([]*ent.User, error) {
	client := middlewares.GetClientFromContext(ctx)
	return client.User.Query().All(ctx)
}

// Companies is the resolver for the companies field.
func (r *userResolver) Companies(ctx context.Context, obj *ent.User) ([]*ent.CompanyDetail, error) {
	return obj.QueryCompanies().QueryCompany().All(ctx)
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
