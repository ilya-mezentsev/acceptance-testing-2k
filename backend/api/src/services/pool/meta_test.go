package pool

import (
	"api_meta/mock/services"
	"test_utils"
	"testing"
)

func TestMetaServicesPool_GetNotNil(t *testing.T) {
	pool := NewMeta()
	pool.AddService("test", services.MetaServiceMock{})

	test_utils.AssertNotNil(pool.Get("test"), t)
}

func TestMetaServicesPool_GetNil(t *testing.T) {
	test_utils.AssertNil(NewMeta().Get("test"), t)
}
