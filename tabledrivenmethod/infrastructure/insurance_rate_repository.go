package infrastructure

import (
	"code-complete-output/tabledrivenmethod/domain"
	"fmt"
)

type InsuranceRateRepository interface {
	GetInsuranceRate(u domain.User) (result float64, err error)
}

type MemoryRepository struct {
	table []domain.InsuranceRate
}

func NewMemoryRepository() InsuranceRateRepository {
	return &MemoryRepository{
		table: []domain.InsuranceRate{
			{
				Gender:        domain.Male,
				MaritalStatus: domain.Married,
				SmokingStatus: domain.Smoking,
				Rate:          1.5,
			},
			{
				Gender:        domain.Male,
				MaritalStatus: domain.Married,
				SmokingStatus: domain.NonSmoking,
				Rate:          1.1,
			},
			{
				Gender:        domain.Male,
				MaritalStatus: domain.Single,
				SmokingStatus: domain.Smoking,
				Rate:          1.4,
			},
			{
				Gender:        domain.Male,
				MaritalStatus: domain.Single,
				SmokingStatus: domain.NonSmoking,
				Rate:          1.0,
			},
			{
				Gender:        domain.Female,
				MaritalStatus: domain.Married,
				SmokingStatus: domain.Smoking,
				Rate:          1.4,
			},
			{
				Gender:        domain.Female,
				MaritalStatus: domain.Married,
				SmokingStatus: domain.NonSmoking,
				Rate:          1.0,
			},
			{
				Gender:        domain.Female,
				MaritalStatus: domain.Single,
				SmokingStatus: domain.Smoking,
				Rate:          1.3,
			},
			{
				Gender:        domain.Female,
				MaritalStatus: domain.Single,
				SmokingStatus: domain.NonSmoking,
				Rate:          0.9,
			}},
	}
}

func (r *MemoryRepository) GetInsuranceRate(u domain.User) (result float64, err error) {
	for _, candidate := range r.table {
		if candidate.Gender == u.Gender && candidate.MaritalStatus == u.MarritalStatus && candidate.SmokingStatus == u.SmokingStatus {
			return candidate.Rate, nil
		}
	}
	return 0, fmt.Errorf("no such rate found")
}
