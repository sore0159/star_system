package mockdb

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/sore0159/star_system/data"
)

type MockAcademy struct {
	Guard   sync.RWMutex    `json:"-"`
	LastUID data.UID        `json:"lastUID"`
	List    []*data.Captain `json:"list"`
}

func NewMockAcademy() *MockAcademy {
	return new(MockAcademy)
}

func (m *MockAcademy) NewCaptain() (*data.Captain, error) {
	m.Guard.Lock()
	defer m.Guard.Unlock()
	m.LastUID += 1
	c := &data.Captain{
		Name: RandName(),
		UID:  m.LastUID,
	}
	m.List = append(m.List, c)
	return c, nil
}
func RandName() string {
	rand.Seed(time.Now().Unix())
	return fmt.Sprintf("Rando %5d", rand.Intn(99999))
}
func (m *MockAcademy) SearchCaptain(uid data.UID) (*data.Captain, error) {
	m.Guard.RLock()
	defer m.Guard.RUnlock()
	for _, c := range m.List {
		if c.UID == uid {
			return c, nil
		}
	}
	return nil, data.ERR_CAP404
}
