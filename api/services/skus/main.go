package skus

import (
	"time"

	"github.com/brbnk/core/api/models/skus"
)

type SkuService struct {
	repository skus.SkuRepositoryInterface
}

type SkuServiceInterface interface {
	GetAllSkus() ([]*skus.Sku, error)
	GetSkuById(p *skus.Sku) (*skus.Sku, error)
	InsertSku(p *skus.Sku) error
	UpdateSku(p *skus.Sku) error
	DeleteSku(p *skus.Sku) error
}

func NewService(ctx skus.SkuRepositoryInterface) *SkuService {
	return &SkuService{
		repository: ctx,
	}
}

func (s *SkuService) GetAllSkus() ([]*skus.Sku, error) {
	skus, err := s.repository.GetSkus()
	if err != nil {
		return nil, err
	}

	return skus, nil
}

func (s *SkuService) GetSkuById(p *skus.Sku) (*skus.Sku, error) {
	product, err := s.repository.GetSkuById(p)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *SkuService) InsertSku(p *skus.Sku) error {
	p.CreateDate = time.Now()
	p.LastUpdate = time.Now()

	if err := s.repository.CreateSku(p); err != nil {
		return err
	}

	return nil
}

func (s *SkuService) UpdateSku(p *skus.Sku) error {
	p.LastUpdate = time.Now()

	if err := s.repository.UpdateSku(p); err != nil {
		return err
	}

	return nil
}

func (s *SkuService) DeleteSku(p *skus.Sku) error {
	p.LastUpdate = time.Now()

	if err := s.repository.DeleteSku(p); err != nil {
		return err
	}

	return nil
}
