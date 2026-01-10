package threadsafe

type Handles[T any, ID ~uintptr] struct {
	slice Slice[T]
	reuse Queue[ID]
}

func (h *Handles[T, ID]) New(v T) ID {
	if handle, ok := h.reuse.Pop(); ok {
		h.slice.SetIndex(int(handle)-1, v)
		return handle
	}
	return ID(h.slice.Append(v)) + 1
}

func (h *Handles[T, ID]) Get(handle ID) T {
	if handle == 0 {
		return [1]T{}[0]
	}
	return h.slice.Index(int(handle) - 1)
}

func (h *Handles[T, ID]) Del(handle ID) {
	h.slice.SetIndex(int(handle)-1, [1]T{}[0])
	h.reuse.Push(handle)
}

func (h *Handles[T, ID]) All(yield func(T) bool) {
	for _, v := range h.slice.Values() {
		if !yield(v) {
			return
		}
	}
}
