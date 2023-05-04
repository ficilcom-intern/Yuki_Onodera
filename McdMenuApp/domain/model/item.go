package model

type Item struct {
	ID            int
	KindID        int
	Name          string
	Energy        float64
	Protein       float64
	Fat           float64
	Carbohydrates float64
}

func NewItem(name string, kind_id int, energy, protein, fat, carbohydrates float64) (*Item, error) {
	item := &Item{
		Name:          name,
		KindID:        kind_id,
		Energy:        energy,
		Protein:       protein,
		Fat:           fat,
		Carbohydrates: carbohydrates,
	}
	return item, nil
}

func (m *Item) Set(name string, kind_id int, energy, protein, fat, carbohydrates float64) error {
	m.Name = name
	m.KindID = kind_id
	m.Energy = energy
	m.Protein = protein
	m.Fat = fat
	m.Carbohydrates = carbohydrates

	return nil
}
