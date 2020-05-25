package async

import (
	"context"
	"encoding/json"
	"fmt"
	"path/filepath"

	"github.com/luanguimaraesla/freegrow/internal/resource"
	"github.com/luanguimaraesla/freegrow/pkg/gadget/irrigator"
)

const (
	resourcePrefix = "resources"
)

type Resource struct {
	name    string
	kind    string
	storage Storage
}

type ResourceList struct {
	kind    string
	storage Storage
}

func NewResource(kind, name string, storage Storage) *Resource {
	return &Resource{
		name:    name,
		kind:    kind,
		storage: storage,
	}
}

func NewResourceList(kind string, storage Storage) *ResourceList {
	return &ResourceList{
		kind:    kind,
		storage: storage,
	}
}

func (rl *ResourceList) List(ctx context.Context) (*resource.List, error) {
	var resources []interface{}

	empty := resource.NewList(resources)

	kvs, err := rl.storage.Get(ctx, rl.Prefix())
	if err != nil {
		return empty, err
	}

	for _, kv := range kvs {
		switch rl.kind {
		case "irrigator":
			r := irrigator.New()
			if err := json.Unmarshal(kv.Value, r); err != nil {
				return empty, err
			}

			resources = append(resources, r)
		default:
			return empty, fmt.Errorf("%s object doesn't have a decoder", rl.kind)
		}
	}

	list := resource.NewList(resources)

	return list, nil

}

func (r *Resource) Put(ctx context.Context, i interface{}) error {
	data, err := json.Marshal(i)
	if err != nil {
		return err
	}

	if err := r.storage.Put(ctx, r.Key(), string(data)); err != nil {
		return err
	}

	return nil
}

func (r *Resource) Get(ctx context.Context) (interface{}, error) {
	kvs, err := r.storage.Get(ctx, r.Key())
	if err != nil {
		return nil, err
	}

	if len(kvs) == 0 {
		return nil, fmt.Errorf("resource not found")
	}

	data := kvs[0].Value
	switch r.kind {
	case "irrigator":
		rsc := irrigator.New()
		if err := json.Unmarshal(data, rsc); err != nil {
			return nil, err
		}

		return rsc, nil
	default:
		return nil, fmt.Errorf("%s object doesn't have a decoder", r.kind)
	}
}

func (r *Resource) Delete(ctx context.Context) error {
	err := r.storage.Delete(ctx, r.Key())
	if err != nil {
		return err
	}

	return nil
}

func (r *Resource) Key() string {
	return filepath.Join(r.Prefix(), r.name)
}

func (r *Resource) Prefix() string {
	return filepath.Join(resourcePrefix, r.kind)
}

func (rl *ResourceList) Prefix() string {
	return filepath.Join(resourcePrefix, rl.kind)
}
