package mmm_test

import (
	"testing"

	"graphics.gd/internal/mmm"
)

type API struct {
	NewObject       func(mmm.Lifetime, string) Object
	GetObjectString func(Object) string
	FreeObject      func(Object)
}

type Object mmm.Pointer[API, Object, uintptr]

func NewAPI() *API {
	var objects []string
	var refresh []int

	var api = new(API)
	*api = API{
		NewObject: func(lt mmm.Lifetime, s string) Object {
			idx := len(objects) + 1
			if len(refresh) > 0 {
				idx = refresh[len(refresh)-1]
				refresh = refresh[:len(refresh)-1]
				objects[idx-1] = s
			} else {
				objects = append(objects, s)
			}
			return mmm.New[Object](lt, api, uintptr(idx))
		},
		GetObjectString: func(obj Object) string {
			return objects[mmm.Get(obj)-1]
		},
		FreeObject: func(obj Object) {
			idx := mmm.End(obj)
			objects[idx-1] = ""
			refresh = append(refresh, int(idx))
		},
	}
	return api
}

func (obj Object) String() string { return mmm.API(obj).GetObjectString(obj) }
func (obj Object) Free()          { mmm.API(obj).FreeObject(obj) }

func TestPointer(t *testing.T) {
	var Objects = NewAPI()

	lt := mmm.NewLifetime()
	defer lt.End()

	var obj = Objects.NewObject(lt, "Hello World")
	if Objects.GetObjectString(obj) != "Hello World" {
		t.Fatal("api.GetObjectString(obj) != \"Hello World\"")
	}
}

func TestRefCount(t *testing.T) {
	var Objects = NewAPI()

	lt1 := mmm.NewLifetime()
	lt2 := mmm.NewLifetime()

	var obj1 = Objects.NewObject(lt1, "Hello World")
	var obj2 = mmm.Copy(obj1, lt2)
	lt1.End()

	if Objects.GetObjectString(obj2) != "Hello World" {
		t.Fatal("api.GetObjectString(obj) != \"Hello World\"")
	}
}

// BenchmarkAllocations should not allocate!
func BenchmarkAllocations(b *testing.B) {
	var Objects = NewAPI()

	lt := mmm.NewLifetime()
	defer lt.End()

	Objects.NewObject(lt, "Hello World")
	lt.End()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		lt := mmm.NewLifetime()
		Objects.NewObject(lt, "Hello World")
		lt.End()
	}
}
