package structure

type Kind struct {
	ID   int    `json:"id_kind"`
	Name string `json:"name"`
}

type KindStoreInterface interface {
	GetAllKind() ([]Kind, error)
}
