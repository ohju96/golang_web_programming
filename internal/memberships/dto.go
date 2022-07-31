package memberships

type CreateRequest struct {
	UserName       string `json:"user_name" example:"andy"`
	MembershipType string `json:"membership_type"  example:"toss"`
}

type CreateResponse struct {
	ID             string `json:"id" example:"354660dc-f798-11ec-b939-0242ac120002"`
	UserName       string `json:"user_name" example:"andy"`
	MembershipType string `json:"membership_type" example:"toss"`
}

type UpdateRequest struct {
	UserName       string `json:"user_name"`
	MembershipType string `json:"membership_type"`
}

type UpdateResponse struct {
	ID             string `json:"id"`
	UserName       string `json:"user_name"`
	MembershipType string `json:"membership_type"`
}

type GetResponse struct {
	ID             string `json:"id" example:"354660dc-f798-11ec-b939-0242ac120002"`
	UserName       string `json:"user_name" example:"andy"`
	MembershipType string `json:"membership_type"  example:"toss"`
}

type Fail400GetResponse struct {
	Message string `json:"message" example:"Bad Request"`
}
