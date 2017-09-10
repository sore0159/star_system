package stars

import (
	"encoding/json"
	"os"
	"sync"
)

var INTF_TEST StarSystem = NewMockDB()

type MockDB struct {
	FileName string       `json:"-"`
	Guard    sync.RWMutex `json:"-"`
	StarList []*Star      `json:"starlist"`
	PathList []*StarPath  `json:"pathlist"`
}

func NewMockDB() *MockDB {
	return new(MockDB)
}
func (m *MockDB) CreateStars(stars []*Star) error {
	m.Guard.Lock()
	defer m.Guard.Unlock()
	m.StarList = append(m.StarList, stars...)
	if err := m.Save(""); err != nil {
		return err
	}
	return nil
}
func (m *MockDB) SearchStar(l Location) (*Star, error) {
	m.Guard.RLock()
	defer m.Guard.RUnlock()
	for _, c := range m.StarList {
		if c.Location == l {
			return c, nil
		}
	}
	return nil, ERR_STAR404
}
func (m *MockDB) LocalStars(l Location) ([]*Star, error) {
	m.Guard.RLock()
	defer m.Guard.RUnlock()
	l2 := make([]*Star, len(m.StarList))
	copy(l2, m.StarList)
	return l2, nil
}

func (m *MockDB) ValidPath(path StarPath) (bool, error) {
	m.Guard.RLock()
	defer m.Guard.RUnlock()
	found := [2]bool{}
	for _, s := range m.StarList {
		if !found[0] && s.Location == path[0] {
			found[0] = true
			if found[1] {
				return true, nil
			}
		}
		if !found[1] && s.Location == path[1] {
			found[1] = true
			if found[0] {
				return true, nil
			}
		}
	}
	return false, nil
}
func (m *MockDB) CheckBlazed(path StarPath) (bool, error) {
	m.Guard.RLock()
	defer m.Guard.RUnlock()
	for _, c := range m.PathList {
		if c.Same(path) {
			return true, nil
		}
	}
	return false, nil
}
func (m *MockDB) SetBlazed(path StarPath) error {
	m.Guard.Lock()
	defer m.Guard.Unlock()
	m.PathList = append(m.PathList, &path)
	if err := m.Save(""); err != nil {
		return err
	}
	return nil
}

func (m *MockDB) Save(fileName string) error {
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
func LoadMockDB(fileName string) (*MockDB, error) {
	m := new(MockDB)
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
