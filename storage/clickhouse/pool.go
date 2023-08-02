package clickhouse

import (
	"fmt"
	"sync"
	"ucode_go_query_service/config"
)

type ProjectStoragesI interface {
	Get(projectId string) (StorageCHI, error)
	Add(s StorageCHI, projectId string) error
	Remove(projectId string) error
}

type ProjectStorages struct {
	StoragePool map[string]StorageCHI
	Mu          sync.Mutex
	config      config.Config
}

func NewProjectStorages() ProjectStoragesI {
	p := ProjectStorages{
		StoragePool: make(map[string]StorageCHI),
		Mu:          sync.Mutex{},
		config:      config.Load(),
	}

	return &p
}

func (p *ProjectStorages) Get(projectId string) (StorageCHI, error) {
	if p.StoragePool == nil {
		return nil, config.ErrNilStoragePool
	}

	if p.config.Environment == "debug" && len(projectId) == 0 {
		projectId = p.config.UcodeDefaultProjectID
	}

	p.Mu.Lock()
	defer p.Mu.Unlock()
	fmt.Println(p.StoragePool)
	storage, ok := p.StoragePool[projectId]
	if !ok {
		return nil, config.ErrProjectNotExists
	}

	return storage, nil
}

func (p *ProjectStorages) Add(s StorageCHI, projectId string) error {
	if p.StoragePool == nil {
		return config.ErrNilStoragePool
	}
	if s == nil {
		return config.ErrNilStorage
	}

	p.Mu.Lock()
	defer p.Mu.Unlock()

	_, ok := p.StoragePool[projectId]
	if ok {
		return config.ErrProjectExists
	}

	p.StoragePool[projectId] = s

	return nil
}

func (p *ProjectStorages) Remove(projectId string) error {
	if p.StoragePool == nil {
		return config.ErrNilStoragePool
	}

	p.Mu.Lock()
	defer p.Mu.Unlock()

	_, ok := p.StoragePool[projectId]
	if !ok {
		return config.ErrProjectNotExists
	}

	delete(p.StoragePool, projectId)
	return nil
}
