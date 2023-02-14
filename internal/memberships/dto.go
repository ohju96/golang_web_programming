package memberships

type CreateRequest struct {
	UserName       string `json:"user_name,omitempty"`
	MembershipType string `json:"membership_type,omitempty"`
}

type CreateResponse struct {
	ID             string `json:"id,omitempty"`
	MembershipType string `json:"membership_type,omitempty"`
}

type UpdateRequest struct {
	ID             string `json:"id,omitempty"`
	UserName       string `json:"user_name,omitempty"`
	MembershipType string `json:"membership_type,omitempty"`
}

type UpdateResponse struct {
	ID             string `json:"id,omitempty"`
	UserName       string `json:"user_name,omitempty"`
	MembershipType string `json:"membership_type,omitempty"`
}

type GetResponse struct {
	ID             string `json:"id,omitempty"`
	UserName       string `json:"user_name,omitempty"`
	MembershipType string `json:"membership_type,omitempty"`
}
