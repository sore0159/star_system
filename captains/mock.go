package captain

import (
	"encoding/json"
	"fmt"
	"math/big"
	"math/rand"
	"os"
	"sync"
)

type MockDB struct {
	Guard   sync.RWMutex `json:"-"`
	LastUID big.Int      `json:"lastUID"`
	List    []*Captain   `json:"list"`
}

func NewMockDB() *MockDB {
	return new(MockDB)
}

func (m *MockDB) NewCaptain() (*Captain, error) {
	m.Guard.Lock()
	defer m.Guard.Unlock()
	m.LastUID.Add(&m.LastUID, big.NewInt(1))
	c := &Captain{
		Name: RandName(),
	}
	c.UID.Set(&m.LastUID)
	m.List = append(m.List, c)
	return c, nil
}
func RandName() string {
	return fmt.Sprintf("Rando %5d", rand.Intn(99999))
}
func (m *MockDB) SearchCaptain(uid *big.Int) (*Captain, error) {
	m.Guard.RLock()
	defer m.Guard.RUnlock()
	for _, c := range m.List {
		if c.UID.Cmp(uid) == 0 {
			return c, nil
		}
	}
	return nil, ERR_CAP404
}

func (m *MockDB) Save(fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	return json.NewEncoder(file).Encode(m)
}
func LoadMockDB(fileName string) (*MockDB, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	m := new(MockDB)
	return m, json.NewDecoder(file).Decode(m)
}
