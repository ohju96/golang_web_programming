package internal

import "errors"

var ErrNotFoundMembership = errors.New("not found membership")

type Repository struct {
	data map[string]Membership
}

func NewRepository(data map[string]Membership) *Repository {
	return &Repository{data: data}
}

func (r *Repository) Create(membership Membership) {
	r.data[membership.UserName] = membership
}

func (r *Repository) GetById(id string) (Membership, error) {
	for _, membership := range r.data {
		if membership.ID == id {
			return membership, nil
		}
	}
	return Membership{}, ErrNotFoundMembership
}

func (r *Repository) Update(membership Membership) (Membership, error) {

	// 유저 찾아오기
	user := r.data[membership.ID]

	// 업데이트 할 멤버
	newUser := Membership{membership.ID, membership.UserName, membership.MembershipType}

	// 업데이트
	r.data[user.ID] = newUser
	return newUser, nil
}

// 유저 이름 중복 체크를 위해 이름을 기준으로 유저를 찾는 레포 로직 작성
func (r *Repository) FindByUserName(userName string) bool {
	for _, existingUser := range r.data {
		if existingUser.UserName == userName {
			return true
		}
	}
	return false
}
