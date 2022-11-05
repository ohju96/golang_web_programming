package memberships

import "errors"

var (
	ErrNotFoundMembership            = errors.New("not found membership")
	_                     Repository = &InMemoryRepository{}
)

type Repository interface {
	Create(membership Membership)
	GetById(id string) (Membership, error)
	Update(membership Membership)
	DeleteById(id string)
}

type InMemoryRepository struct {
	data map[string]Membership
}

func NewRepository(data map[string]Membership) *InMemoryRepository {
	return &InMemoryRepository{data: data}
}

func (r *InMemoryRepository) Create(membership Membership) {
	r.data[membership.ID] = membership
}

func (r *InMemoryRepository) GetById(id string) (Membership, error) {
	for _, membership := range r.data {
		if membership.ID == id {
			return membership, nil
		}
	}
	return Membership{}, ErrNotFoundMembership
}

func (r *InMemoryRepository) Update(membership Membership) {
	r.data[membership.ID] = membership
}

func (r *InMemoryRepository) DeleteById(id string) {
	delete(r.data, id)
}
