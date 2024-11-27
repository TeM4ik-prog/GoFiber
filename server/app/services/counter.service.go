package services

type CounterService struct {
	count int
}

func NewCounterService() *CounterService {
	return &CounterService{count: 0}
}

func (cs *CounterService) GetCount() int {
	return cs.count
}

func (cs *CounterService) Increment() int {
	cs.count++
	return cs.count
}

func (cs *CounterService) Reset() int {
	cs.count = 0
	return cs.count
}
