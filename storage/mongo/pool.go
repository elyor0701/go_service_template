package mongo

import (
	"fmt"
	"sync"
	"ucode_go_query_service/config"
)

type ResourceStoragesI interface {
	Get(projectId string) (StorageMI, error)
	Add(s StorageMI, projectId string) error
	Remove(projectId string) error
}

type ResourceStorages struct {
	StoragePool map[string]StorageMI
	Mu          sync.Mutex
	config      config.Config
}

func NewResourceStorages() ResourceStoragesI {
	p := ResourceStorages{
		StoragePool: make(map[string]StorageMI),
		Mu:          sync.Mutex{},
		config:      config.Load(),
	}

	return &p
}

func (p *ResourceStorages) Get(projectId string) (StorageMI, error) {
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

func (p *ResourceStorages) Add(s StorageMI, projectId string) error {
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

func (p *ResourceStorages) Remove(projectId string) error {
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
