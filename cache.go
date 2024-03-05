package shared

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

type Cache struct {
	nc *nats.Conn
	js jetstream.JetStream
}

func NewCache(nc *nats.Conn) (*Cache, error) {
	js, err := jetstream.New(nc)
	if err != nil {
		return nil, err
	}
	return &Cache{
		js: js,
		nc: nc,
	}, nil
}

func (c *Cache) Close() {
	if c.nc != nil {
		c.nc.Close()
	}
}

func (c *Cache) CreateBucket(ctx context.Context, entity, requestID string) error {
	_, err := c.js.CreateOrUpdateKeyValue(ctx, jetstream.KeyValueConfig{
		Bucket: requestID + "-" + entity,
	})
	return err
}

func (c *Cache) AddKV(ctx context.Context, entity, requestID, key string, value []byte) error {
	bucket := requestID + "-" + entity
	kv, err := c.js.KeyValue(ctx, bucket)
	if err != nil {
		return err
	}
	_, err = kv.Put(ctx, key, value)
	return err
}

func (c *Cache) SetEntity(ctx context.Context, entity, requestID string, key interface{}, value interface{}) error {
	k, err := Hash(key)
	if err != nil {
		return err
	}
	v, err := json.Marshal(value)
	if err != nil {
		return err
	}
	bucket := requestID + "-" + entity
	kv, err := c.js.KeyValue(ctx, bucket)
	if err != nil {
		return err
	}
	_, err = kv.Put(ctx, fmt.Sprint(k), v)
	return err
}

func (c *Cache) GetEntities(ctx context.Context, entity, requestID string) (map[string]interface{}, error) {
	bucket := requestID + "-" + entity
	kv, err := c.js.KeyValue(ctx, bucket)
	if err != nil {
		return nil, err
	}
	keys, err := kv.ListKeys(ctx)
	if err != nil {
		return nil, err
	}
	resp := make(map[string]interface{})
	for k := range keys.Keys() {
		v, err := kv.Get(ctx, k)
		if err != nil {
			return nil, err
		}
		resp[k] = v
	}
	return resp, nil
}
