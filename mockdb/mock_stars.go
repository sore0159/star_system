package mockdb

import (
	"sync"

	"github.com/sore0159/star_system/data"
)

type MockStarSystem struct {
	Guard    sync.RWMutex     `json:"-"`
	StarList []*data.Star     `json:"starlist"`
	PathList []*data.StarPath `json:"pathlist"`
}

func NewMockStarSystem() *MockStarSystem {
	return new(MockStarSystem)
}
func (m *MockStarSystem) CreateStars(stars []*data.Star) error {
	m.Guard.Lock()
	defer m.Guard.Unlock()
	m.StarList = append(m.StarList, stars...)
	return nil
}
func (m *MockStarSystem) SearchStar(l data.Location) (*data.Star, error) {
	m.Guard.RLock()
	defer m.Guard.RUnlock()
	for _, c := range m.StarList {
		if c.Location == l {
			return c, nil
		}
	}
	return nil, data.ERR_STAR404
}
func (m *MockStarSystem) LocalStars(l data.Location) ([]*data.Star, error) {
	m.Guard.RLock()
	defer m.Guard.RUnlock()
	l2 := make([]*data.Star, len(m.StarList))
	copy(l2, m.StarList)
	return l2, nil
}

func (m *MockStarSystem) ValidPath(path data.StarPath) (bool, error) {
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
func (m *MockStarSystem) CheckBlazed(path data.StarPath) (bool, error) {
	m.Guard.RLock()
	defer m.Guard.RUnlock()
	for _, c := range m.PathList {
		if c.Same(path) {
			return true, nil
		}
	}
	return false, nil
}
func (m *MockStarSystem) SetBlazed(path data.StarPath) error {
	m.Guard.Lock()
	defer m.Guard.Unlock()
	m.PathList = append(m.PathList, &path)
	return nil
}
