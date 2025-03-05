package repositories

type ZoomRepository interface {
	CreateMeeting() error
}

func (r *Repository) CreateMeeting() error {
	return nil
}