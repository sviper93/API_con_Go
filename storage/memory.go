// Creando almacenamiento en memoria

package storage

import (
	"github.com/EDteam/golang-api/clase-3/model"
)

// Memory .
type Memory struct {
	currentID int
	Persons   map[int]model.Person
}

// NewMemory: inicializa el mapa de personas y devuelva una instancia de memory
func NewMemory() Memory {
	persons := make(map[int]model.Person) // creación de mapa de entero de model.Person

	return Memory{
		currentID: 0,
		Persons:   persons, // recordemos que los mapas necesitan a diferencia de los slices,
		// necesitan inicializarse
	}
}

// Create .
func (m *Memory) Create(person *model.Person) error { // recibe un puntero a persona
	if person == nil {
		// return errors.New("la persona no puede ser nula") -> esto no es recomendable porque si
		// un futuro quisiera comparar el error, complica esa labor, así que lo ideal es crear una
		// variable que pueda ser accedida desde cualquier parte del proyecto y por ende me permita
		// comparar el error en cualquier parte del proyecto, por eso, hemos creado "model.go"
		return model.ErrPersonCanNotBeNil
	}

	m.currentID++                    // incrementamos en uno el ID
	m.Persons[m.currentID] = *person // en la llave currentID asígnale el valor de person

	return nil
}

// Update actualiza una persona en el slice de memoria
// func (m *Memory) Update(ID int, person *model.Person) error {
// 	if person == nil {
// 		return model.ErrPersonCanNotBeNil
// 	}

// 	if _, ok := m.Persons[ID]; !ok {
// 		return fmt.Errorf("ID: %d: %w", ID, model.ErrIDPersonDoesNotExists)
// 	}

// 	m.Persons[ID] = *person

// 	return nil
// }

// Delete borra de la memoria la persona
// func (m *Memory) Delete(ID int) error {
// 	if _, ok := m.Persons[ID]; !ok {
// 		return fmt.Errorf("ID: %d: %w", ID, model.ErrIDPersonDoesNotExists)
// 	}

// 	delete(m.Persons, ID)

// 	return nil
// }

// GetByID retorna una persona por el ID
// func (m *Memory) GetByID(ID int) (model.Person, error) {
// 	person, ok := m.Persons[ID]
// 	if !ok {
// 		return person, fmt.Errorf("ID: %d: %w", ID, model.ErrIDPersonDoesNotExists)
// 	}

// 	return person, nil
// }

// GetAll retorna todas las personas que están en la memoria
// func (m *Memory) GetAll() (model.Persons, error) {
// 	var result model.Persons
// 	for _, v := range m.Persons {
// 		result = append(result, v)
// 	}

// 	return result, nil
// }
