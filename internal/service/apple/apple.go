package skeleton

import (
	"context"

	appleEntity "print-apple/internal/entity/apple"
	"print-apple/pkg/errors"
)

// AppleData ...
type AppleData interface {
	// GetAllUsers(ctx context.Context) ([]userEntity.User, error)
	GetPrintApple(ctx context.Context) ([]appleEntity.Apple, error)
	UpdateStorage(ctx context.Context, TransFH string) error
	DeleteAndUpdateStorage(ctx context.Context, TransFH string) error
	Insert(ctx context.Context, apple appleEntity.Apple) error
	GetPrintPageTemp(ctx context.Context, page int, length int) ([]appleEntity.Apple, error)
	GetPrintPageFinal(ctx context.Context, page int, length int) ([]appleEntity.Apple, error)
	GetByTransFHTemp(ctx context.Context, TransFH string) (appleEntity.Apple, error)
	GetByTransFHFinal(ctx context.Context, TransFH string) (appleEntity.Apple, error)
}

// Service ...
type Service struct {
	AppleData AppleData
}

// New ...
func New(AppleData AppleData) Service {
	return Service{
		AppleData: AppleData,
	}
}

// GetPrintApple ...
func (s Service) GetPrintApple(ctx context.Context) ([]appleEntity.Apple, error) {
	var apple []appleEntity.Apple
	//	user, err := s.GetUserFromFireBase(ctx)

	apple, err := s.AppleData.GetPrintApple(ctx)
	return apple, err
}

// DeleteAndUpdateStorage ...
func (s Service) DeleteAndUpdateStorage(ctx context.Context, TransFH string) error {
	err := s.AppleData.UpdateStorage(ctx, TransFH)
	err = s.AppleData.DeleteAndUpdateStorage(ctx, TransFH)
	if err != nil {
		return errors.Wrap(err, "[SERVICE][DeleteAndUpdateStorage]")
	}
	return err
}

// Insert ...
func (s Service) Insert(ctx context.Context, apple appleEntity.Apple) error {
	err := s.AppleData.Insert(ctx, apple)

	return err
}

// GetPrintPageTemp ...
func (s Service) GetPrintPageTemp(ctx context.Context, page int, length int) ([]appleEntity.Apple, error) {
	appleList, err := s.AppleData.GetPrintPageTemp(ctx, page, length)
	return appleList, err
}

// GetPrintPageFinal ...
func (s Service) GetPrintPageFinal(ctx context.Context, page int, length int) ([]appleEntity.Apple, error) {
	appleList, err := s.AppleData.GetPrintPageFinal(ctx, page, length)
	return appleList, err
}

// GetByTransFHTemp ...
func (s Service) GetByTransFHTemp(ctx context.Context, TransFH string) (appleEntity.Apple, error) {
	appleFirebase, err := s.AppleData.GetByTransFHTemp(ctx, TransFH)
	return appleFirebase, err
}

// GetByTransFHFinal ...
func (s Service) GetByTransFHFinal(ctx context.Context, TransFH string) (appleEntity.Apple, error) {
	appleFirebase, err := s.AppleData.GetByTransFHFinal(ctx, TransFH)
	return appleFirebase, err
}
