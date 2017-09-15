package mockdb

import (
	"encoding/json"
	"os"

	"github.com/sore0159/star_system/data"
)

var MockTest data.Provider = new(MockProvider)

type MockProvider struct {
	FileName string `json:"-"`
	*MockAcademy
	*MockStarSystem
}

func NewMockProvider() *MockProvider {
	return &MockProvider{
		MockAcademy:    NewMockAcademy(),
		MockStarSystem: NewMockStarSystem(),
	}
}

func (m *MockProvider) Save(fileName string) error {
	if fileName == "" {
		fileName = m.FileName
	}
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	return json.NewEncoder(file).Encode(m)
}
func LoadMockProvider(fileName string) (*MockProvider, error) {
	m := new(MockProvider)
	m.FileName = fileName
	file, err := os.Open(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return m, nil
		}
		return nil, err
	}
	defer file.Close()
	return m, json.NewDecoder(file).Decode(m)
}
