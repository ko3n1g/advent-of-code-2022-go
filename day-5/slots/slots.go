package slots

type Slot []string

func (s *Slot) Push(v []string) {
	*s = append((*s), v...)
}

func (s *Slot) Pop(nItems int) []string {
	var el []string = (*s)[len((*s))-nItems:]
	*s = (*s)[:len((*s))-nItems]
	return el
}
