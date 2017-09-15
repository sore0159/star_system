package mockdb

import (
	"fmt"
	"math/big"
	"math/rand"
	"sync"

	"github.com/sore0159/star_system/data"
)

type MockAcademy struct {
	Guard   sync.RWMutex    `json:"-"`
	LastUID big.Int         `json:"lastUID"`
	List    []*data.Captain `json:"list"`
}

func NewMockAcademy() *MockAcademy {
	return new(MockAcademy)
}

func (m *MockAcademy) NewCaptain() (*data.Captain, error) {
	m.Guard.Lock()
	defer m.Guard.Unlock()
	m.LastUID.Add(&m.LastUID, big.NewInt(1))
	c := &data.Captain{
		Name: RandName(),
	}
	c.UID.Set(&m.LastUID)
	m.List = append(m.List, c)
	return c, nil
}
func RandName() string {
	return fmt.Sprintf("Rando %5d", rand.Intn(99999))
}
func (m *MockAcademy) SearchCaptain(uid *big.Int) (*data.Captain, error) {
	m.Guard.RLock()
	defer m.Guard.RUnlock()
	for _, c := range m.List {
		if c.UID.Cmp(uid) == 0 {
			return c, nil
		}
	}
	return nil, data.ERR_CAP404
}
