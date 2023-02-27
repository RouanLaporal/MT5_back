package structure

type Collaborator struct {
	ID     int    `json:"id_collaborator"`
	ShopID int    `json:"id_shop"`
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	Email  string `json:"email"`
}

type CollaboratorStoreInterface interface {
	GetCollaboratorByShop(id_shop int) ([]Collaborator, error)
	AddCollaborator(collaborator Collaborator) (int, error)
	DeleteCollaborator(id_collaborator int) error
	UpdateCollaborator(id_collaborator int, updated_collaborator Collaborator) error
}
