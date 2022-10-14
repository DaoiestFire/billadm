package types

import (
	"fmt"
	"time"

	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
)

type Metadata struct {
	Workspaces       sets.String `json:"workspaces,omitempty"`
	CurrentWorkSpace string      `json:"current_workspace,omitempty"`

	LastUpdateTime string `json:"last_update_time,omitempty"`
}

func (m *Metadata) UpdateTime() {

	m.LastUpdateTime = time.Now().String()
	klog.Infof("update Metadata time to: %s", m.LastUpdateTime)
}

type WorkspaceMetadata struct {
	Name        string `json:"name"`
	RecordCount uint32 `json:"record_count"`

	Labels map[string]string `json:"labels,omitempty"`

	LastUpdateTime string `json:"last_update_time,omitempty"`
}

func (m *WorkspaceMetadata) UpdateTime() {

	m.LastUpdateTime = time.Now().String()
	klog.Infof("update Metadata time to: %s", m.LastUpdateTime)
}

func (m *WorkspaceMetadata) AddLabel(key, label string) {
	klog.Infof("add label[%s,%s]", key, label)
	m.Labels[key] = label
}

func (m *WorkspaceMetadata) DeleteLabel(key string) error {
	if _, ok := m.Labels[key]; !ok {
		klog.Infof("label'key[%s] doesn't exsit", key)
		return fmt.Errorf("key doesn't exist")
	}

	klog.Infof("label[%s,%s] deleted")
	delete(m.Labels, key)
	return nil

}
