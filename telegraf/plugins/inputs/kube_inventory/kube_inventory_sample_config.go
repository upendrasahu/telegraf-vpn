//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package kube_inventory

func (ki *KubernetesInventory) SampleConfig() string {
	return `{{ .SampleConfig }}`
}
