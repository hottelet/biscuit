package keymanager

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeys(t *testing.T) {
	key, err := NewARN("arn:aws:kms:us-west-1:105770556716:key/37793df5-ad32-4d06-b19f-bfb95cee4a35")
	assert.NoError(t, err)
	assert.Equal(t, "kms", key.Service)
	assert.Equal(t, "key", key.ResourceType)
	assert.Equal(t, "37793df5-ad32-4d06-b19f-bfb95cee4a35", key.Resource)
	assert.Equal(t, "arn:aws:kms:us-west-1:105770556716:key/37793df5-ad32-4d06-b19f-bfb95cee4a35", key.String())
	assert.True(t, key.IsKmsKey())

	alias, err := NewARN("arn:aws:kms:us-west-1:105770556716:alias/foo")
	assert.NoError(t, err)
	assert.True(t, alias.IsKmsAlias())
	assert.Equal(t, "alias", alias.ResourceType)
	assert.Equal(t, "arn:aws:kms:us-west-1:105770556716:alias/foo", alias.String())
	assert.Equal(t, "foo", alias.Resource)
}

func TestInvalidArn(t *testing.T) {
	for _, invalid := range []string{
		"",
		"arn:",
		"alias/foo",
		"key/foo",
	} {
		_, err := NewARN(invalid)
		assert.Error(t, err)
	}

}

func TestARN(t *testing.T) {
	a1, err := NewARN("arn:partition:service:region:account-id:resource")
	assert.NoError(t, err)
	assert.Equal(t, "partition", a1.Partition)
	assert.Equal(t, "service", a1.Service)
	assert.Equal(t, "region", a1.Region)
	assert.Equal(t, "account-id", a1.AccountID)
	assert.Equal(t, "resource", a1.Resource)
	assert.Equal(t, "", a1.ResourceType)
	assert.Equal(t, "arn:partition:service:region:account-id:resource", a1.String())

	a2, err := NewARN("arn:partition:service:region:account-id:resourcetype/resource")
	assert.NoError(t, err)
	assert.Equal(t, "partition", a2.Partition)
	assert.Equal(t, "service", a2.Service)
	assert.Equal(t, "region", a2.Region)
	assert.Equal(t, "account-id", a2.AccountID)
	assert.Equal(t, "resource", a2.Resource)
	assert.Equal(t, "resourcetype", a2.ResourceType)
	assert.Equal(t, "arn:partition:service:region:account-id:resourcetype/resource", a2.String())

	a3, err := NewARN("arn:partition:service:region:account-id:resourcetype:resource")
	assert.NoError(t, err)
	assert.Equal(t, "partition", a3.Partition)
	assert.Equal(t, "service", a3.Service)
	assert.Equal(t, "region", a3.Region)
	assert.Equal(t, "account-id", a3.AccountID)
	assert.Equal(t, "resource", a3.Resource)
	assert.Equal(t, "resourcetype", a3.ResourceType)
	assert.Equal(t, "arn:partition:service:region:account-id:resourcetype:resource", a3.String())
}

func TestARNVariants(t *testing.T) {
	colonSlashSlash, err := NewARN("arn:partition:service:region:account-id:resourcetype/resource/label")
	assert.NoError(t, err)
	assert.Equal(t, "partition", colonSlashSlash.Partition)
	assert.Equal(t, "service", colonSlashSlash.Service)
	assert.Equal(t, "region", colonSlashSlash.Region)
	assert.Equal(t, "account-id", colonSlashSlash.AccountID)
	assert.Equal(t, "resource/label", colonSlashSlash.Resource)
	assert.Equal(t, "resourcetype", colonSlashSlash.ResourceType)

	colonColonSlash, err := NewARN("arn:partition:service:region:account-id:resourcetype:resource/label")
	assert.NoError(t, err)
	assert.Equal(t, "partition", colonColonSlash.Partition)
	assert.Equal(t, "service", colonColonSlash.Service)
	assert.Equal(t, "region", colonColonSlash.Region)
	assert.Equal(t, "account-id", colonColonSlash.AccountID)
	assert.Equal(t, "resource/label", colonColonSlash.Resource)
	assert.Equal(t, "resourcetype", colonColonSlash.ResourceType)
}
