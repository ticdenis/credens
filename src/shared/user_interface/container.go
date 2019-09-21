package user_interface

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type Container struct {
	entries map[string]interface{}
}

type Setter = func(container *Container) interface{}

func NewContainer() *Container {
	return &Container{map[string]interface{}{}}
}

func (container *Container) set(id string, value interface{}) {
	container.entries[id] = value
}

func (container *Container) Set(id string, fn Setter) {
	container.set(id, fn(container))
}

func (container *Container) SetEmptySlice(id string) {
	container.set(id, []interface{}{})
}

func (container *Container) SetEmptyDict(id string) {
	container.set(id, map[string]interface{}{})
}

func (container Container) isArray(id string) bool {
	if !strings.Contains(id, "[]") {
		return false
	}

	switch reflect.TypeOf(container.Get(id)).Kind() {
	case reflect.Slice:
	case reflect.Array:
	case reflect.Map:
		return true
	}

	return false
}

func (container *Container) SetInDict(dictId string, id string, fn Setter) {
	if container.isArray(dictId) {
		dict := container.GetDict(dictId)

		dict[id] = fn(container)

		container.set(dictId, dict)
	}
}

func (container *Container) SetInSlice(sliceId string, fn Setter) {
	if container.isArray(sliceId) {
		container.set(sliceId, append(container.GetSlice(sliceId), fn(container)))
	}
}

func (container Container) Get(id string) interface{} {
	if val, ok := container.entries[id]; ok {
		return val
	}

	panic(errors.New(fmt.Sprintf("[%s] dependency not found!", id)))
}

func (container Container) GetDict(id string) map[string]interface{} {
	return container.Get(id).(map[string]interface{})
}

func (container Container) GetSlice(id string) []interface{} {
	return container.Get(id).([]interface{})
}

func (container Container) GetInDict(dictId string, id string) interface{} {
	return container.GetDict(dictId)[id]
}

func (container Container) GetDictAsSlice(dictId string) []interface{} {
	dict := container.Get(dictId).(map[string]interface{})
	arr := make([]interface{}, 0, len(dict))

	for _, val := range dict {
		arr = append(arr, val)
	}

	return arr
}
