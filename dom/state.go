package dom

type State[T any] struct {
	value T
}

func (s *State[T]) SetValue(newVal T) {
	s.value = newVal
}

func UseState[T any](initialValue *T, el *Element) (*T, func(T)) {
	newState := &State[T]{
		value: *initialValue,
	}

	state := el.getSetStateHandler(&initialValue, newState)

	var setValue = func(newVal T) {
		state.(*State[T]).SetValue(newVal)
		*initialValue = state.(*State[T]).value
	}

	*initialValue = state.(*State[T]).value

	return initialValue, setValue
}
