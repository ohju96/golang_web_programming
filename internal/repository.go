package internal

import (
	"errors"
	"strconv"
)

type Repository struct {
	data map[string]Membership
}

func NewRepository(data map[string]Membership) *Repository {
	return &Repository{data: data}
}

func (r *Repository) CreateUser(request CreateRequest) (*Membership, error) {

	// 이미 등록된 사용자 이름이 존재하면 실패
	for _, existingUser := range r.data {
		if existingUser.UserName == request.UserName {
			return nil, errors.New("이미 존재하는 이름입니다")
		}
	}

	// 사용자의 이름을 입력하지 않으면 실패
	if request.UserName == "" {
		return nil, errors.New("사용자 이름이 입력되지 않았습니다")
	}

	// 멤버십 타입을 입력하지 않은 경우 실패
	if request.MembershipType == "" {
		return nil, errors.New("멤버십이 입력되지 않았습니다")
	}

	// naver,toss,payco 이외 타입 입력 시 실패
	whiteSlice := []string{"naver", "toss", "payco"}
	if !contains(whiteSlice, request.MembershipType) {
		return nil, errors.New("허용 안 됩니다")
	}

	// 위 로직은 되는데 이 로직이 왜 안 되는지 체크해 봐야 한다.
	//for i, _ := range whiteSlice {
	//	if whiteSlice[i] != request.MembershipType {
	//		return nil, errors.New("허용 안 됩니다")
	//	}
	//}

	// 멤버십 생성
	membership := Membership{
		ID:             strconv.Itoa(len(r.data) + 1),
		UserName:       request.UserName,
		MembershipType: request.MembershipType,
	}

	r.data[membership.ID] = membership
	return &membership, nil
}

func (r *Repository) UpdateUser(request UpdateRequest) (*Membership, error) {

	membership := r.data[request.ID]

	// 수정하려는 사용자 이름이 이미 있으면 예외 처리
	if membership.UserName != request.UserName {
		for _, existingUser := range r.data {
			if existingUser.UserName == request.UserName {
				return nil, errors.New("이미 존재하는 이름입니다")
			}
		}
	}

	// 멤버십 아이디 입력하지 않으면 예외
	if request.ID == "" {
		return nil, errors.New("멤버십 아이디 입력하지 않음")
	}

	// 이름을 입력하지 않으면 예외
	if request.UserName == "" {
		return nil, errors.New("이름 입력하지 않음")
	}

	// 멤버쉽을 입력하지 않음
	if request.MembershipType == "" {
		return nil, errors.New("멤버십을 입력하지 않음")
	}

	// 이외 타입 입력 시 실패
	whiteSlice := []string{"naver", "toss", "payco"}
	if !contains(whiteSlice, request.MembershipType) {
		return nil, errors.New("허용 안 됩니다")
	}

	newMembership := Membership{membership.ID, request.UserName, request.MembershipType}

	r.data[membership.ID] = newMembership
	return &newMembership, nil
}

func (r *Repository) DeleteUser(id string) error {

	if id == "" {
		return errors.New("id를 입력하지 않았습니다")
	}

	_, ok := r.data[id]
	if ok == false {
		return errors.New("입력한 id가 존재하지 않습니다")
	}

	delete(r.data, id)
	return nil
}

// 슬라이스 체크용 contains 메서드
func contains(sliceValues []string, membershipType string) bool {
	for _, value := range sliceValues {
		if value == membershipType {
			return true
		}
	}
	return false
}
