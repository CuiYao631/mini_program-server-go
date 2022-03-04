package repository

import (
	"context"
	"log"

	"github.com/CuiYao631/mini_program-server-go/ent"
	"github.com/CuiYao631/mini_program-server-go/ent/resources"
	"github.com/CuiYao631/mini_program-server-go/entity"
)

type Resources interface {
	CreateResources(ctx context.Context, resources entity.Resources) error
	UpdateResources(ctx context.Context, resources entity.Resources) error
	ListResources(ctx context.Context) ([]*ent.Resources, error)
	GetResources(ctx context.Context, id string) (*ent.Resources, error)
	DeleteResources(ctx context.Context, id string) error
}

func (repo *repository) CreateResources(ctx context.Context, resources entity.Resources) error {

	_, err := repo.db.Resources.Create().
		SetName(resources.Name).
		SetIcon(resources.Icon).
		SetDesc(resources.Desc).
		SetExplain(resources.Explain).
		SetURL(resources.Url).
		SetIsTop(resources.Topping).
		//SetTagID(resources.Tag).
		Save(ctx)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
func (repo *repository) UpdateResources(ctx context.Context, res entity.Resources) error {
	_, err := repo.db.Resources.Update().
		SetName(res.Name).
		SetIcon(res.Icon).
		SetDesc(res.Desc).
		SetExplain(res.Explain).
		SetURL(res.Url).
		//SetTagID(res.Tag).
		Where(resources.IDEQ(res.ID)).
		Save(ctx)
	if err != nil {
		return err
	}
	return nil
}
func (repo *repository) ListResources(ctx context.Context) ([]*ent.Resources, error) {
	entRes, err := repo.db.Resources.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	return entRes, nil
}
func (repo *repository) GetResources(ctx context.Context, id string) (*ent.Resources, error) {
	res, err := repo.db.Resources.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (repo *repository) DeleteResources(ctx context.Context, id string) error {
	_, err := repo.db.Resources.Delete().Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
