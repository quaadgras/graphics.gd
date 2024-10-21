package AudioStreamPlaybackOggVorbis

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/AudioStreamPlaybackResampled"
import "grow.graphics/gd/object/AudioStreamPlayback"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

type Simple [1]classdb.AudioStreamPlaybackOggVorbis
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.AudioStreamPlaybackOggVorbis
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self class) AsAudioStreamPlaybackOggVorbis() Expert { return self[0].AsAudioStreamPlaybackOggVorbis() }


//go:nosplit
func (self Simple) AsAudioStreamPlaybackOggVorbis() Simple { return self[0].AsAudioStreamPlaybackOggVorbis() }


//go:nosplit
func (self class) AsAudioStreamPlaybackResampled() AudioStreamPlaybackResampled.Expert { return self[0].AsAudioStreamPlaybackResampled() }


//go:nosplit
func (self Simple) AsAudioStreamPlaybackResampled() AudioStreamPlaybackResampled.Simple { return self[0].AsAudioStreamPlaybackResampled() }


//go:nosplit
func (self class) AsAudioStreamPlayback() AudioStreamPlayback.Expert { return self[0].AsAudioStreamPlayback() }


//go:nosplit
func (self Simple) AsAudioStreamPlayback() AudioStreamPlayback.Simple { return self[0].AsAudioStreamPlayback() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("AudioStreamPlaybackOggVorbis", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
