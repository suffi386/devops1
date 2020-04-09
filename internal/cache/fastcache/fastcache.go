package fastcache

import (
	"bytes"
	"encoding/gob"
	"github.com/caos/zitadel/internal/errors"

	"github.com/VictoriaMetrics/fastcache"
)

type Fastcache struct {
	cache *fastcache.Cache
}

func NewFastcache(config *Config) (*Fastcache, error) {
	return &Fastcache{
		cache: fastcache.New(config.MaxCacheSizeInByte),
	}, nil
}

func (fc *Fastcache) Set(key string, object interface{}) error {
	var b bytes.Buffer
	enc := gob.NewEncoder(&b)
	if err := enc.Encode(object); err != nil {
		return errors.ThrowInvalidArgument(err, "FASTC-RUyxI", "unable to encode object")
	}
	fc.cache.Set([]byte(key), b.Bytes())
	return nil
}

func (fc *Fastcache) Get(key string, ptrToObject interface{}) error {
	data := fc.cache.Get(nil, []byte(key))
	if len(data) == 0 {
		return errors.ThrowNotFound(nil, "FASTC-xYzSm", "key not found")
	}

	b := bytes.NewBuffer(data)
	dec := gob.NewDecoder(b)

	return dec.Decode(ptrToObject)
}

func (fc *Fastcache) Delete(key string) error {
	fc.cache.Del([]byte(key))
	return nil
}
