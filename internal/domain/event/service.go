package event

type Service interface {
	FindAll() ([]Event, error)
	FindOne(id string) (*Event, error)
	Create(event Event) (*Event, error)
	Update(event Event) (*Event, error)
	Delete(id string) (bool, error)
}

type service struct {
	repo *Repository
}

func NewService(r *Repository) Service {
	return &service{
		repo: r,
	}
}

func (s *service) FindAll() ([]Event, error) {
	return (*s.repo).FindAll()
}

func (s *service) FindOne(id string) (*Event, error) {
	return (*s.repo).FindOne(id)
}
func (s *service) Create(event Event) (*Event, error) {
	return (*s.repo).Create(event)
}
func (s *service) Update(event Event) (*Event, error) {
	return (*s.repo).Update(event)
}
func (s *service) Delete(id string) (bool, error) {
	return (*s.repo).Delete(id)
}
