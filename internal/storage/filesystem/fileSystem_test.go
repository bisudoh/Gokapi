package filesystem

import (
	"testing"

	"github.com/bisudoh/gokapi/internal/models"
	fileInterfaces "github.com/bisudoh/gokapi/internal/storage/filesystem/interfaces"
	"github.com/bisudoh/gokapi/internal/storage/filesystem/s3filesystem/aws"
	"github.com/bisudoh/gokapi/internal/test"
)

func TestInit(t *testing.T) {
	Init("./test")
	test.IsEqualBool(t, ActiveStorageSystem == dataFilesystem, true)
	test.IsEqualBool(t, ActiveStorageSystem == s3FileSystem, false)
	test.IsEqualString(t, ActiveStorageSystem.GetSystemName(), fileInterfaces.DriverLocal)
}

func TestSetLocal(t *testing.T) {
	ActiveStorageSystem = nil
	SetLocal()
	test.IsEqualBool(t, ActiveStorageSystem == dataFilesystem, true)
}

func TestSetAws(t *testing.T) {
	ActiveStorageSystem = nil
	if !aws.IsIncludedInBuild {
		SetAws()
		test.IsNil(t, ActiveStorageSystem)
		return
	}
	aws.Init(models.AwsConfig{
		Bucket:    "test1",
		Region:    "test2",
		KeyId:     "test3",
		KeySecret: "test4",
		Endpoint:  "test5",
	})
	SetAws()
	test.IsNil(t, ActiveStorageSystem)
	isUnitTesting = true
	SetAws()
	test.IsEqualBool(t, ActiveStorageSystem == s3FileSystem, true)
	test.IsEqualBool(t, ActiveStorageSystem == dataFilesystem, false)
	test.IsEqualString(t, ActiveStorageSystem.GetSystemName(), fileInterfaces.DriverAws)

}
