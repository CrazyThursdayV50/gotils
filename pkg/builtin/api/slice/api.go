package slice

import (
	"github.com/CrazyThursdayV50/gotils/pkg/builtin/api"
	"github.com/CrazyThursdayV50/gotils/pkg/builtin/models"
)

func From[E any](sli ...E) api.SliceAPI[E] {
	return models.FromSlice(sli...)
}

func Make[E any](len, cap int) api.SliceAPI[E] {
	return models.MakeSlice[E](len, cap)
}

func Empty(len int) api.SliceAPI[struct{}] {
	return models.MakeSlice[struct{}](len, len)
}
