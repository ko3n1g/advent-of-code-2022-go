package slots

type Slot []string

func (s *Slot) Push(v string) {
	*s = append((*s), v)
}

func (s *Slot) Pop() string {
	var el string = (*s)[len((*s))-1]
	*s = (*s)[:len((*s))-1]
	return el
}
