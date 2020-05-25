package async

import (
	"context"
	"encoding/json"
	"fmt"
	"path/filepath"

	"github.com/luanguimaraesla/freegrow/pkg/gadget/irrigator"
)

type Resource struct {
	name    string
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
		resource := irrigator.New()
		if err := json.Unmarshal(data, resource); err != nil {
			return nil, err
		}

		return resource, nil
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
	return filepath.Join("resources", r.kind, r.name)
}
