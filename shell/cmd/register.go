package cmd

type StoreRegister struct {
	stores []CommandStore
}

func NewStoreRegister() *StoreRegister {
	return &StoreRegister{
		stores: make([]CommandStore, 0),
	}
}

func (r *StoreRegister) Register(s CommandStore) {
	r.stores = append(r.stores, s)
}

func (r *StoreRegister) Stores() []CommandStore {
	return r.stores
}
