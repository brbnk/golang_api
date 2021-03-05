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
	GetSkuById(uint) (*skus.Sku, error)
	InsertSku(*skus.Sku) error
	UpdateSku(*skus.Sku) error
	DeleteSku(uint) error
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

func (s *SkuService) GetSkuById(id uint) (*skus.Sku, error) {
	product, err := s.repository.GetSkuById(id)
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

func (s *SkuService) DeleteSku(id uint) error {
	time := time.Now()

	if err := s.repository.DeleteSku(id, time); err != nil {
		return err
	}

	return nil
}
