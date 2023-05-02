package model

type Kind struct {
	KindID int    `db:"kind_id"`
	Name   string `db:"name"`
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
