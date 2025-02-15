package main

import (
	"code-complete-output/tabledrivenmethod/domain"
	"code-complete-output/tabledrivenmethod/infrastructure"
	"fmt"
)

func main() {
	user := domain.User{
		Name:           "tetsuya",
		Gender:         domain.Male,
		MarritalStatus: domain.Married,
		SmokingStatus:  domain.Smoking,
	}

	insuranceRateRepository := infrastructure.NewMemoryRepository()
	rate, err := insuranceRateRepository.GetInsuranceRate(user)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(rate)
}
