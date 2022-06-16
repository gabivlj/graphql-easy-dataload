package loader

// Helper struct that let's you easily do n+1! calls
type Loader[T any, K any] struct {
	data2D [][]T
	keys   []K
	err    error
	cb2D   CallbackLoad2D[T, K]
	data   []T
	cb     CallbackLoad[T, K]
}

type DataLoaderInstance[T any, K any] struct {
	loader *Loader[T, K]
	index  int
}

type CallbackLoad2D[T any, K any] func(keys []K) ([][]T, error)

type CallbackLoad[T any, K any] func(keys []K) ([]T, error)

func NewLoader2D[T any, K any](loadData CallbackLoad2D[T, K]) *Loader[T, K] {
	return &Loader[T, K]{
		data2D: nil,
		keys:   make([]K, 0),
		err:    nil,
		cb2D:   loadData,
		data:   nil,
	}
}

func NewLoader[T any, K any](loadData CallbackLoad[T, K]) *Loader[T, K] {
	return &Loader[T, K]{
		data2D: nil,
		keys:   make([]K, 0),
		err:    nil,
		cb:     loadData,
		data:   nil,
	}
}

func (l *Loader[T, K]) LoadKey(key K) *DataLoaderInstance[T, K] {
	l.keys = append(l.keys, key)
	dl := &DataLoaderInstance[T, K]{loader: l, index: len(l.keys) - 1}
	return dl
}

func (l *Loader[T, K]) retrieveData() ([][]T, error) {
	if l.err != nil {
		return nil, l.err
	}

	if l.data2D == nil {
		l.data2D, l.err = l.cb2D(l.keys)
	}

	return l.data2D, l.err
}

func (l *Loader[T, K]) retrieveData1D() ([]T, error) {
	if l.err != nil {
		return nil, l.err
	}

	if l.data == nil {
		l.data, l.err = l.cb(l.keys)
	}

	return l.data, l.err
}

func (d *DataLoaderInstance[T, K]) Get() ([]T, error) {
	data, err := d.loader.retrieveData()
	if err != nil {
		return nil, err
	}

	return data[d.index], nil
}

func (d *DataLoaderInstance[T, K]) Get1D() (T, error) {
	var t T
	data, err := d.loader.retrieveData1D()
	if err != nil {
		return t, err
	}

	return data[d.index], nil
}
