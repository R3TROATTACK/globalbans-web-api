package addons

import "plugin"

type Addon struct {
	Plugin      *plugin.Plugin
	Name        string
	Version     string
	Author      string
	Description string
	Url         string
	MethodMap   map[string]interface{}
}

type Registry struct {
	Addons map[string]*Addon
}

var registry Registry = Registry{}

func (r *Registry) Register(name string, addon *Addon) {
	r.Addons[name] = addon
}

func (r *Registry) Get(name string) *Addon {
	return r.Addons[name]
}

func (r *Registry) Del(name string) {
	delete(r.Addons, name)
}

func Get(name string) *Addon {
	return registry.Get(name)
}

func Register(name string, addon *Addon) {
	registry.Register(name, addon)
}

func Del(name string) {
	registry.Del(name)
}

func GetRegistry() *Registry {
	return &registry
}
