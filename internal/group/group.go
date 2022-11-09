package group

type Service struct {
	repo repo
}

func New(
	repo repo,
) *Service {
	return &Service{
		repo: repo,
	}
}
