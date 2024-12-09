package api

type Usecase interface {
	HandleGetCount() (int, error)
	HandlePostCount(count int) error
}
