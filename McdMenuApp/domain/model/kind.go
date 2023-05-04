package model

type Kind struct {
	ID     int
	Name   string
}

func NewKind(name string) (*Kind, error) {
	kind := &Kind{
		Name: name,
	}

	return kind, nil
}

// Set taskのセッター
func (m *Kind) Set(name string) error {
	m.Name = name

	return nil
}
