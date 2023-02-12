package internal

type Application struct {
	repository Repository
}

func NewApplication(repository Repository) *Application {
	return &Application{repository: repository}
}

func (app *Application) Create(request CreateRequest) (*CreateResponse, error) {
	user, err := app.repository.CreateUser(request)
	if err != nil {
		return nil, err
	}
	return &CreateResponse{user.ID, user.MembershipType}, nil
}

func (app *Application) Update(request UpdateRequest) (*UpdateResponse, error) {
	user, err := app.repository.UpdateUser(request)
	if err != nil {
		return nil, err
	}
	return &UpdateResponse{user.ID, user.UserName, user.MembershipType}, nil
}

func (app *Application) Delete(id string) error {
	return app.repository.DeleteUser(id)
}
