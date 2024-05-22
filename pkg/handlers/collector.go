package handlers

type Collector[T any, Y any] func(T) Y
type CollectorOkay[T any, Y any] func(T) (Y, bool)
type CollectorError[T any, Y any] func(T) (Y, error)
type CollectorKV[T any, K any, V any] func(T) (K, V)
type CollectorKVOkay[T any, K any, V any] func(T) (K, V, bool)
type CollectorKVError[T any, K any, V any] func(T) (K, V, error)
