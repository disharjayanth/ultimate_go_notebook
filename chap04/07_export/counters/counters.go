package counters

type altCounter int

// alterCounter can be accessed outside counters package since it is returned from New func.
// Best practise is to make altCounter exported type with AltCounter
func New(value int) altCounter {
	return altCounter(value)
}
