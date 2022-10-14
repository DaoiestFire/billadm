package manager

import "ljw/billadm/pkg/types"

var _ manager = &WorkspaceManager{}

type WorkspaceManager struct {
	data types.Data
}

func (dem *WorkspaceManager) Get() error {
	return nil
}

func (dem *WorkspaceManager) Delete() error { 
	return nil
}

func (dem *WorkspaceManager) Create() error {
	return nil
}

func (dem *WorkspaceManager) Status() error {
	return nil
}

func (dem *WorkspaceManager) Recover() error {
	return nil
}