package web

type CategoryInsertRequest struct {
	Name string `validate:"required,max=150,min=1"`
}
