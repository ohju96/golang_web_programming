package membership

import (
	"errors"
	"github.com/google/uuid"
)

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) Create(request CreateRequest) (*CreateResponse, error) {

	// 검증 서비스 로직
	switch {
	case s.repository.FindByUserName(request.UserName) == true:
		return nil, errors.New("이미 존재하는 이름입니다")
	case request.UserName == "":
		return nil, errors.New("이름을 입력하지 않았습니다")
	case request.MembershipType == "":
		return nil, errors.New("멤버쉽을 입력하지 않았습니다")
	case !contains(request.MembershipType):
		return nil, errors.New("지정된 멤버쉽이 아닙니다")
	}

	membership := Membership{uuid.New().String(), request.UserName, request.MembershipType}
	s.repository.Create(membership)
	return &CreateResponse{
		ID:             membership.ID,
		MembershipType: membership.MembershipType,
	}, nil
}

func (s *Service) GetByID(id string) (*GetResponse, error) {

	// 단건에선 굳이 스위치를 안 써도 된다. 때문에 if 사용
	if id == "" {
		return nil, errors.New("ID를 입력하지 않았습니다")
	}

	membership, err := s.repository.GetById(id)
	if err != nil { // 쿼리 에러
		return nil, errors.New("쿼리 에러")
	}
	return &GetResponse{
		ID:             membership.ID,
		UserName:       membership.UserName,
		MembershipType: membership.MembershipType,
	}, nil
}

func (s *Service) Update(request UpdateRequest) (*UpdateResponse, error) {

	// 검증 서비스 로직
	switch {
	case request.ID == "":
		return nil, errors.New("멤버십 아이디를 입력하지 않았습니다")
	case request.UserName == "":
		return nil, errors.New("이름을 입력하지 않았습니다")
	case request.MembershipType == "":
		return nil, errors.New("멤버쉽을 입력하지 않았습니다")
	case !contains(request.MembershipType):
		return nil, errors.New("지정된 멤버쉽이 아닙니다")
	}

	// 업데이트 로직
	membership, err := s.repository.Update(Membership(request))
	if err != nil { // 쿼리 에러
		return nil, errors.New("쿼리 에러")
	}

	// 응답 객체 리턴
	return &UpdateResponse{
		ID:             membership.ID,
		UserName:       membership.UserName,
		MembershipType: membership.MembershipType,
	}, nil
}

func contains(membershipType string) bool {

	// 지정된 멤버쉽 슬라이스를 contains 메서드로 빼서 중복 제거
	sliceValues := []string{"naver", "toss", "payco"}

	for _, value := range sliceValues {
		if value == membershipType {
			return true
		}
	}
	return false
}
