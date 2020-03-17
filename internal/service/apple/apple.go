package skeleton

import (
	"context"

	appleEntity "print-apple/internal/entity/apple"
)

// UserData ...
type UserData interface {
	// GetAllUsers(ctx context.Context) ([]userEntity.User, error)
	GetAppleFromFireBase(ctx context.Context) ([]appleEntity.Apple, error)
}

// Service ...
type Service struct {
	userData UserData
}

// New ...
func New(userData UserData) Service {
	return Service{
		userData: userData,
	}
}

// GetAppleFromFireBase ...
func (s Service) GetAppleFromFireBase(ctx context.Context) ([]appleEntity.Apple, error) {
	var apple []appleEntity.Apple
	//	user, err := s.GetUserFromFireBase(ctx)

	apple, err := s.userData.GetAppleFromFireBase(ctx)
	return apple, err
}
