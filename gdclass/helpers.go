package gdclass

import gd "grow.graphics/gd/internal"

// As attempts to cast the given class to T, returning true
// if the cast was successful.
func As[T gd.IsClass](class gd.IsClass) (T, bool) { return gd.As[T](class) }
