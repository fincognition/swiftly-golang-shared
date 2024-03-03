package shared

import "github.com/mitchellh/hashstructure/v2"

func Hash(data interface{}) (uint64, error) {
	return hashstructure.Hash(data, hashstructure.FormatV2, nil)
}
