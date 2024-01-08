package rd

import (
	"errors"
	"github.com/FreeZmaR/go-project-layout/internal/domain/definition"
	"github.com/FreeZmaR/go-project-layout/internal/domain/model"
	"github.com/google/uuid"
)

type userCache struct {
	ID                      string `json:"id"`
	Name                    string `json:"name"`
	Age                     int    `json:"age"`
	Balance                 int    `json:"balance"`
	AvailablePaymentMethods []int  `json:"available_payment_methods"`
}

func (u *userCache) ToUser() (model.User, error) {
	id, err := uuid.Parse(u.ID)
	if err != nil {
		return model.User{}, errors.New("invalid user id")
	}

	availablePaymentMethods := make([]definition.PaymentMethod, 0, len(u.AvailablePaymentMethods))

	for _, pm := range u.AvailablePaymentMethods {
		if !definition.IsPaymentMethod(pm) {
			return model.User{}, errors.New("invalid payment method")
		}

		availablePaymentMethods = append(availablePaymentMethods, definition.PaymentMethod(pm))
	}

	return model.User{
		ID:                      id,
		Name:                    u.Name,
		Age:                     u.Age,
		Balance:                 u.Balance,
		AvailablePaymentMethods: availablePaymentMethods,
	}, nil
}

func newUserCache(user model.User) userCache {
	availablePaymentMethods := make([]int, 0, len(user.AvailablePaymentMethods))

	for _, pm := range user.AvailablePaymentMethods {
		availablePaymentMethods = append(availablePaymentMethods, int(pm))
	}

	return userCache{
		ID:                      user.ID.String(),
		Name:                    user.Name,
		Age:                     user.Age,
		Balance:                 user.Balance,
		AvailablePaymentMethods: availablePaymentMethods,
	}
}
