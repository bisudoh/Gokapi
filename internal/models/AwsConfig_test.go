package models

import (
	"testing"

	"github.com/bisudoh/gokapi/internal/test"
)

func TestIsAwsProvided(t *testing.T) {
	config := AwsConfig{}
	test.IsEqualBool(t, config.IsAllProvided(), false)
	config = AwsConfig{
		Bucket:    "test",
		Region:    "test",
		Endpoint:  "",
		KeyId:     "test",
		KeySecret: "test",
	}
	test.IsEqualBool(t, config.IsAllProvided(), true)
}
