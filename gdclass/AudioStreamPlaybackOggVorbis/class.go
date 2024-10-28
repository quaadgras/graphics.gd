package AudioStreamPlaybackOggVorbis

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/AudioStreamPlaybackResampled"
import "grow.graphics/gd/gdclass/AudioStreamPlayback"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

type Go [1]classdb.AudioStreamPlaybackOggVorbis
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.AudioStreamPlaybackOggVorbis
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AudioStreamPlaybackOggVorbis"))
	return Go{classdb.AudioStreamPlaybackOggVorbis(object)}
}

func (self class) AsAudioStreamPlaybackOggVorbis() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsAudioStreamPlaybackOggVorbis() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsAudioStreamPlaybackResampled() AudioStreamPlaybackResampled.GD { return *((*AudioStreamPlaybackResampled.GD)(unsafe.Pointer(&self))) }
func (self Go) AsAudioStreamPlaybackResampled() AudioStreamPlaybackResampled.Go { return *((*AudioStreamPlaybackResampled.Go)(unsafe.Pointer(&self))) }
func (self class) AsAudioStreamPlayback() AudioStreamPlayback.GD { return *((*AudioStreamPlayback.GD)(unsafe.Pointer(&self))) }
func (self Go) AsAudioStreamPlayback() AudioStreamPlayback.Go { return *((*AudioStreamPlayback.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsAudioStreamPlaybackResampled(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsAudioStreamPlaybackResampled(), name)
	}
}
func init() {classdb.Register("AudioStreamPlaybackOggVorbis", func(ptr gd.Object) any { return classdb.AudioStreamPlaybackOggVorbis(ptr) })}
