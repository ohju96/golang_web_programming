package memberships

import (
	"errors"
	"github.com/google/uuid"
)

var (
	ErrEmptyID = errors.New("empty id")
)

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) Create(request CreateRequest) (CreateResponse, error) {
	membership := Membership{uuid.New().String(), request.UserName, request.MembershipType}
	s.repository.Create(membership)
	return CreateResponse{
		ID:             membership.ID,
		UserName:       membership.UserName,
		MembershipType: membership.MembershipType,
	}, nil
}

func (s *Service) GetByID(id string) (GetResponse, error) {
	membership, err := s.repository.GetById(id)
	if err != nil {
		return GetResponse{}, err
	}
	return GetResponse{
		ID:             membership.ID,
		UserName:       membership.UserName,
		MembershipType: membership.MembershipType,
	}, nil
}

func (s *Service) Update(id string, request UpdateRequest) (UpdateResponse, error) {
	_, err := s.repository.GetById(id)
	if err != nil {
		return UpdateResponse{}, err
	}
	membership := Membership{
		ID:             id,
		UserName:       request.UserName,
		MembershipType: request.MembershipType,
	}
	s.repository.Update(membership)
	return UpdateResponse{
		ID:             membership.ID,
		UserName:       membership.UserName,
		MembershipType: membership.MembershipType,
	}, nil
}

func (s *Service) Delete(id string) error {
	if id == "" {
		return ErrEmptyID
	}
	_, err := s.repository.GetById(id)
	if err != nil {
		return err
	}
	s.repository.DeleteById(id)
	return nil
}
