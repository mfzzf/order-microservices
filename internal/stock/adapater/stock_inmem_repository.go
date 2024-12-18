package adapters

import (
	"context"
	"github.com/mfzzf/order_microservices/common/genproto/orderpb"
	domain "github.com/mfzzf/order_microservices/stock/domain/stock"
	"sync"
)

type MemoryStockRepository struct {
	lock  *sync.RWMutex
	store map[string]*orderpb.Item
}

var stub = map[string]*orderpb.Item{
	"item_id": {
		ID:       "foo_item",
		Name:     "stub item",
		Quantity: 10000,
		PriceID:  "STUB_ITEM_PRICE_ID",
	},
}

func (m MemoryStockRepository) GetItems(ctx context.Context, ids []string) ([]*orderpb.Item, error) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	var (
		res     []*orderpb.Item
		missing []string
	)

	for _, id := range ids {
		if item, exist := m.store[id]; exist {
			res = append(res, item)
		} else {
			missing = append(missing, id)
		}
	}
	if len(res) == len(ids) {
		return res, nil
	}
	return res, domain.NotFoundError{Missing: missing}

}

func NewMemoryStockRepository() *MemoryStockRepository {
	return &MemoryStockRepository{store: stub, lock: &sync.RWMutex{}}
}
