package expirable

import "time"

type Container struct {
	value   interface{}
	Created time.Time
}

func Init(value interface{}) Container {
	return Container{
		value:   value,
		Created: time.Now(),
	}
}

func (c Container) GetValue() interface{} {
	return c.value
}

func (c Container) IsExpired(d time.Duration) bool {
	return c.Created.Add(d).Before(time.Now())
}
