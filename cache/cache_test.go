package cache_test

import (
	"testing"

	"delly.ioo/durian/cache/cache"
	"github.com/stretchr/testify/assert"
)

func TestCache(t *testing.T) {
	assert := assert.New(t)
	c := cache.NewCache()

	err := c.Put("sde_bootcamp",
		[]string{"title", "price", "enrolled", "estimated_time"},
		[]string{"SDE-Bootcamp", "30000.00", "false", "30"},
	)
	assert.Nil(err)

	res, err := c.Get("sde_bootcamp")
	assert.Nil(err)
	assert.Equal("title: SDE-Bootcamp, price: 30000.00, enrolled: false, estimated_time: 30", res)

	assert.Equal("sde_bootcamp", c.Keys())

	err = c.Put("sde_kickstart",
		[]string{"title", "price", "enrolled", "estimated_time"},
		[]string{"SDE-Kickstart", "4000", "true", "8"},
	)
	assert.Error(err)

	res, err = c.Get("sde_kickstart")
	assert.Error(err)
	assert.Equal("", res)

	assert.Equal("sde_bootcamp", c.Keys())

	err = c.Put("sde_kickstart",
		[]string{"title", "price", "enrolled", "estimated_time"},
		[]string{"SDE-Kickstart", "4000.00", "true", "8"},
	)
	assert.Nil(err)

	res, err = c.Get("sde_kickstart")
	assert.Nil(err)
	assert.Equal("title: SDE-Kickstart, price: 4000.00, enrolled: true, estimated_time: 8", res)

	c.Delete("sde_bootcamp")
	res, err = c.Get("sde_bootcamp")
	assert.Error(err)
	assert.Equal("", res)

	assert.Equal("sde_kickstart", c.Keys())

	err = c.Put("sde_bootcamp",
		[]string{"title", "price", "enrolled", "estimated_time"},
		[]string{"SDE-Bootcamp", "30000.00", "true", "30"},
	)
	assert.Nil(err)

	assert.Equal("sde_bootcamp", c.Search("price", "30000.00"))
	// In the expected output, it returning `sde_bootcamp,sde_kickstart`,
	// it should returning `sde_kickstart,sde_bootcamp` instead, because
	// we delete key `sde_bootcamp` before, so it should be added
	// to last index
	assert.Equal("sde_kickstart,sde_bootcamp", c.Search("enrolled", "true"))

	err = c.Put("sde_bootcamp",
		[]string{"title", "price", "enrolled", "estimated_time"},
		[]string{"SDE-Bootcamp", "30000.00", "false", "30"},
	)
	assert.Nil(err)

	assert.Equal("sde_kickstart", c.Search("enrolled", "true"))

	res, err = c.Get("nothing")
	assert.Error(err)
	assert.Equal("", res)

	c.Delete("nothing")

	res, err = c.Get("nothing")
	assert.Error(err)
	assert.Equal("", res)

	err = c.Put("sde_bootcamp",
		[]string{"title", "price", "enrolled", "estimated_time"},
		[]string{"SDE-Bootcamp", "30000", "false", "30"},
	)
	assert.Error(err)
}
