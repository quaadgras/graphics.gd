package XRFaceTracker

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/XRTracker"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
An instance of this object represents a tracked face and its corresponding blend shapes. The blend shapes come from the [url=https://docs.vrcft.io/docs/tutorial-avatars/tutorial-avatars-extras/unified-blendshapes]Unified Expressions[/url] standard, and contain extended details and visuals for each blend shape. Additionally the [url=https://docs.vrcft.io/docs/tutorial-avatars/tutorial-avatars-extras/compatibility/overview]Tracking Standard Comparison[/url] page documents the relationship between Unified Expressions and other standards.
As face trackers are turned on they are registered with the [XRServer].

*/
type Simple [1]classdb.XRFaceTracker
func (self Simple) GetBlendShape(blend_shape classdb.XRFaceTrackerBlendShapeEntry) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetBlendShape(blend_shape)))
}
func (self Simple) SetBlendShape(blend_shape classdb.XRFaceTrackerBlendShapeEntry, weight float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBlendShape(blend_shape, gd.Float(weight))
}
func (self Simple) GetBlendShapes() gd.PackedFloat32Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedFloat32Array(Expert(self).GetBlendShapes(gc))
}
func (self Simple) SetBlendShapes(weights gd.PackedFloat32Array) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBlendShapes(weights)
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.XRFaceTracker
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Returns the requested face blend shape weight.
*/
//go:nosplit
func (self class) GetBlendShape(blend_shape classdb.XRFaceTrackerBlendShapeEntry) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, blend_shape)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRFaceTracker.Bind_get_blend_shape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets a face blend shape weight.
*/
//go:nosplit
func (self class) SetBlendShape(blend_shape classdb.XRFaceTrackerBlendShapeEntry, weight gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, blend_shape)
	callframe.Arg(frame, weight)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRFaceTracker.Bind_set_blend_shape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBlendShapes(ctx gd.Lifetime) gd.PackedFloat32Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRFaceTracker.Bind_get_blend_shapes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedFloat32Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetBlendShapes(weights gd.PackedFloat32Array)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(weights))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRFaceTracker.Bind_set_blend_shapes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) AsXRFaceTracker() Expert { return self[0].AsXRFaceTracker() }


//go:nosplit
func (self Simple) AsXRFaceTracker() Simple { return self[0].AsXRFaceTracker() }


//go:nosplit
func (self class) AsXRTracker() XRTracker.Expert { return self[0].AsXRTracker() }


//go:nosplit
func (self Simple) AsXRTracker() XRTracker.Simple { return self[0].AsXRTracker() }


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
func init() {classdb.Register("XRFaceTracker", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type BlendShapeEntry = classdb.XRFaceTrackerBlendShapeEntry

const (
/*Right eye looks outwards.*/
	FtEyeLookOutRight BlendShapeEntry = 0
/*Right eye looks inwards.*/
	FtEyeLookInRight BlendShapeEntry = 1
/*Right eye looks upwards.*/
	FtEyeLookUpRight BlendShapeEntry = 2
/*Right eye looks downwards.*/
	FtEyeLookDownRight BlendShapeEntry = 3
/*Left eye looks outwards.*/
	FtEyeLookOutLeft BlendShapeEntry = 4
/*Left eye looks inwards.*/
	FtEyeLookInLeft BlendShapeEntry = 5
/*Left eye looks upwards.*/
	FtEyeLookUpLeft BlendShapeEntry = 6
/*Left eye looks downwards.*/
	FtEyeLookDownLeft BlendShapeEntry = 7
/*Closes the right eyelid.*/
	FtEyeClosedRight BlendShapeEntry = 8
/*Closes the left eyelid.*/
	FtEyeClosedLeft BlendShapeEntry = 9
/*Squeezes the right eye socket muscles.*/
	FtEyeSquintRight BlendShapeEntry = 10
/*Squeezes the left eye socket muscles.*/
	FtEyeSquintLeft BlendShapeEntry = 11
/*Right eyelid widens beyond relaxed.*/
	FtEyeWideRight BlendShapeEntry = 12
/*Left eyelid widens beyond relaxed.*/
	FtEyeWideLeft BlendShapeEntry = 13
/*Dilates the right eye pupil.*/
	FtEyeDilationRight BlendShapeEntry = 14
/*Dilates the left eye pupil.*/
	FtEyeDilationLeft BlendShapeEntry = 15
/*Constricts the right eye pupil.*/
	FtEyeConstrictRight BlendShapeEntry = 16
/*Constricts the left eye pupil.*/
	FtEyeConstrictLeft BlendShapeEntry = 17
/*Right eyebrow pinches in.*/
	FtBrowPinchRight BlendShapeEntry = 18
/*Left eyebrow pinches in.*/
	FtBrowPinchLeft BlendShapeEntry = 19
/*Outer right eyebrow pulls down.*/
	FtBrowLowererRight BlendShapeEntry = 20
/*Outer left eyebrow pulls down.*/
	FtBrowLowererLeft BlendShapeEntry = 21
/*Inner right eyebrow pulls up.*/
	FtBrowInnerUpRight BlendShapeEntry = 22
/*Inner left eyebrow pulls up.*/
	FtBrowInnerUpLeft BlendShapeEntry = 23
/*Outer right eyebrow pulls up.*/
	FtBrowOuterUpRight BlendShapeEntry = 24
/*Outer left eyebrow pulls up.*/
	FtBrowOuterUpLeft BlendShapeEntry = 25
/*Right side face sneers.*/
	FtNoseSneerRight BlendShapeEntry = 26
/*Left side face sneers.*/
	FtNoseSneerLeft BlendShapeEntry = 27
/*Right side nose canal dilates.*/
	FtNasalDilationRight BlendShapeEntry = 28
/*Left side nose canal dilates.*/
	FtNasalDilationLeft BlendShapeEntry = 29
/*Right side nose canal constricts.*/
	FtNasalConstrictRight BlendShapeEntry = 30
/*Left side nose canal constricts.*/
	FtNasalConstrictLeft BlendShapeEntry = 31
/*Raises the right side cheek.*/
	FtCheekSquintRight BlendShapeEntry = 32
/*Raises the left side cheek.*/
	FtCheekSquintLeft BlendShapeEntry = 33
/*Puffs the right side cheek.*/
	FtCheekPuffRight BlendShapeEntry = 34
/*Puffs the left side cheek.*/
	FtCheekPuffLeft BlendShapeEntry = 35
/*Sucks in the right side cheek.*/
	FtCheekSuckRight BlendShapeEntry = 36
/*Sucks in the left side cheek.*/
	FtCheekSuckLeft BlendShapeEntry = 37
/*Opens jawbone.*/
	FtJawOpen BlendShapeEntry = 38
/*Closes the mouth.*/
	FtMouthClosed BlendShapeEntry = 39
/*Pushes jawbone right.*/
	FtJawRight BlendShapeEntry = 40
/*Pushes jawbone left.*/
	FtJawLeft BlendShapeEntry = 41
/*Pushes jawbone forward.*/
	FtJawForward BlendShapeEntry = 42
/*Pushes jawbone backward.*/
	FtJawBackward BlendShapeEntry = 43
/*Flexes jaw muscles.*/
	FtJawClench BlendShapeEntry = 44
/*Raises the jawbone.*/
	FtJawMandibleRaise BlendShapeEntry = 45
/*Upper right lip part tucks in the mouth.*/
	FtLipSuckUpperRight BlendShapeEntry = 46
/*Upper left lip part tucks in the mouth.*/
	FtLipSuckUpperLeft BlendShapeEntry = 47
/*Lower right lip part tucks in the mouth.*/
	FtLipSuckLowerRight BlendShapeEntry = 48
/*Lower left lip part tucks in the mouth.*/
	FtLipSuckLowerLeft BlendShapeEntry = 49
/*Right lip corner folds into the mouth.*/
	FtLipSuckCornerRight BlendShapeEntry = 50
/*Left lip corner folds into the mouth.*/
	FtLipSuckCornerLeft BlendShapeEntry = 51
/*Upper right lip part pushes into a funnel.*/
	FtLipFunnelUpperRight BlendShapeEntry = 52
/*Upper left lip part pushes into a funnel.*/
	FtLipFunnelUpperLeft BlendShapeEntry = 53
/*Lower right lip part pushes into a funnel.*/
	FtLipFunnelLowerRight BlendShapeEntry = 54
/*Lower left lip part pushes into a funnel.*/
	FtLipFunnelLowerLeft BlendShapeEntry = 55
/*Upper right lip part pushes outwards.*/
	FtLipPuckerUpperRight BlendShapeEntry = 56
/*Upper left lip part pushes outwards.*/
	FtLipPuckerUpperLeft BlendShapeEntry = 57
/*Lower right lip part pushes outwards.*/
	FtLipPuckerLowerRight BlendShapeEntry = 58
/*Lower left lip part pushes outwards.*/
	FtLipPuckerLowerLeft BlendShapeEntry = 59
/*Upper right part of the lip pulls up.*/
	FtMouthUpperUpRight BlendShapeEntry = 60
/*Upper left part of the lip pulls up.*/
	FtMouthUpperUpLeft BlendShapeEntry = 61
/*Lower right part of the lip pulls up.*/
	FtMouthLowerDownRight BlendShapeEntry = 62
/*Lower left part of the lip pulls up.*/
	FtMouthLowerDownLeft BlendShapeEntry = 63
/*Upper right lip part pushes in the cheek.*/
	FtMouthUpperDeepenRight BlendShapeEntry = 64
/*Upper left lip part pushes in the cheek.*/
	FtMouthUpperDeepenLeft BlendShapeEntry = 65
/*Moves upper lip right.*/
	FtMouthUpperRight BlendShapeEntry = 66
/*Moves upper lip left.*/
	FtMouthUpperLeft BlendShapeEntry = 67
/*Moves lower lip right.*/
	FtMouthLowerRight BlendShapeEntry = 68
/*Moves lower lip left.*/
	FtMouthLowerLeft BlendShapeEntry = 69
/*Right lip corner pulls diagonally up and out.*/
	FtMouthCornerPullRight BlendShapeEntry = 70
/*Left lip corner pulls diagonally up and out.*/
	FtMouthCornerPullLeft BlendShapeEntry = 71
/*Right corner lip slants up.*/
	FtMouthCornerSlantRight BlendShapeEntry = 72
/*Left corner lip slants up.*/
	FtMouthCornerSlantLeft BlendShapeEntry = 73
/*Right corner lip pulls down.*/
	FtMouthFrownRight BlendShapeEntry = 74
/*Left corner lip pulls down.*/
	FtMouthFrownLeft BlendShapeEntry = 75
/*Mouth corner lip pulls out and down.*/
	FtMouthStretchRight BlendShapeEntry = 76
/*Mouth corner lip pulls out and down.*/
	FtMouthStretchLeft BlendShapeEntry = 77
/*Right lip corner is pushed backwards.*/
	FtMouthDimpleRight BlendShapeEntry = 78
/*Left lip corner is pushed backwards.*/
	FtMouthDimpleLeft BlendShapeEntry = 79
/*Raises and slightly pushes out the upper mouth.*/
	FtMouthRaiserUpper BlendShapeEntry = 80
/*Raises and slightly pushes out the lower mouth.*/
	FtMouthRaiserLower BlendShapeEntry = 81
/*Right side lips press and flatten together vertically.*/
	FtMouthPressRight BlendShapeEntry = 82
/*Left side lips press and flatten together vertically.*/
	FtMouthPressLeft BlendShapeEntry = 83
/*Right side lips squeeze together horizontally.*/
	FtMouthTightenerRight BlendShapeEntry = 84
/*Left side lips squeeze together horizontally.*/
	FtMouthTightenerLeft BlendShapeEntry = 85
/*Tongue visibly sticks out of the mouth.*/
	FtTongueOut BlendShapeEntry = 86
/*Tongue points upwards.*/
	FtTongueUp BlendShapeEntry = 87
/*Tongue points downwards.*/
	FtTongueDown BlendShapeEntry = 88
/*Tongue points right.*/
	FtTongueRight BlendShapeEntry = 89
/*Tongue points left.*/
	FtTongueLeft BlendShapeEntry = 90
/*Sides of the tongue funnel, creating a roll.*/
	FtTongueRoll BlendShapeEntry = 91
/*Tongue arches up then down inside the mouth.*/
	FtTongueBlendDown BlendShapeEntry = 92
/*Tongue arches down then up inside the mouth.*/
	FtTongueCurlUp BlendShapeEntry = 93
/*Tongue squishes together and thickens.*/
	FtTongueSquish BlendShapeEntry = 94
/*Tongue flattens and thins out.*/
	FtTongueFlat BlendShapeEntry = 95
/*Tongue tip rotates clockwise, with the rest following gradually.*/
	FtTongueTwistRight BlendShapeEntry = 96
/*Tongue tip rotates counter-clockwise, with the rest following gradually.*/
	FtTongueTwistLeft BlendShapeEntry = 97
/*Inner mouth throat closes.*/
	FtSoftPalateClose BlendShapeEntry = 98
/*The Adam's apple visibly swallows.*/
	FtThroatSwallow BlendShapeEntry = 99
/*Right side neck visibly flexes.*/
	FtNeckFlexRight BlendShapeEntry = 100
/*Left side neck visibly flexes.*/
	FtNeckFlexLeft BlendShapeEntry = 101
/*Closes both eye lids.*/
	FtEyeClosed BlendShapeEntry = 102
/*Widens both eye lids.*/
	FtEyeWide BlendShapeEntry = 103
/*Squints both eye lids.*/
	FtEyeSquint BlendShapeEntry = 104
/*Dilates both pupils.*/
	FtEyeDilation BlendShapeEntry = 105
/*Constricts both pupils.*/
	FtEyeConstrict BlendShapeEntry = 106
/*Pulls the right eyebrow down and in.*/
	FtBrowDownRight BlendShapeEntry = 107
/*Pulls the left eyebrow down and in.*/
	FtBrowDownLeft BlendShapeEntry = 108
/*Pulls both eyebrows down and in.*/
	FtBrowDown BlendShapeEntry = 109
/*Right brow appears worried.*/
	FtBrowUpRight BlendShapeEntry = 110
/*Left brow appears worried.*/
	FtBrowUpLeft BlendShapeEntry = 111
/*Both brows appear worried.*/
	FtBrowUp BlendShapeEntry = 112
/*Entire face sneers.*/
	FtNoseSneer BlendShapeEntry = 113
/*Both nose canals dilate.*/
	FtNasalDilation BlendShapeEntry = 114
/*Both nose canals constrict.*/
	FtNasalConstrict BlendShapeEntry = 115
/*Puffs both cheeks.*/
	FtCheekPuff BlendShapeEntry = 116
/*Sucks in both cheeks.*/
	FtCheekSuck BlendShapeEntry = 117
/*Raises both cheeks.*/
	FtCheekSquint BlendShapeEntry = 118
/*Tucks in the upper lips.*/
	FtLipSuckUpper BlendShapeEntry = 119
/*Tucks in the lower lips.*/
	FtLipSuckLower BlendShapeEntry = 120
/*Tucks in both lips.*/
	FtLipSuck BlendShapeEntry = 121
/*Funnels in the upper lips.*/
	FtLipFunnelUpper BlendShapeEntry = 122
/*Funnels in the lower lips.*/
	FtLipFunnelLower BlendShapeEntry = 123
/*Funnels in both lips.*/
	FtLipFunnel BlendShapeEntry = 124
/*Upper lip part pushes outwards.*/
	FtLipPuckerUpper BlendShapeEntry = 125
/*Lower lip part pushes outwards.*/
	FtLipPuckerLower BlendShapeEntry = 126
/*Lips push outwards.*/
	FtLipPucker BlendShapeEntry = 127
/*Raises the upper lips.*/
	FtMouthUpperUp BlendShapeEntry = 128
/*Lowers the lower lips.*/
	FtMouthLowerDown BlendShapeEntry = 129
/*Mouth opens, revealing teeth.*/
	FtMouthOpen BlendShapeEntry = 130
/*Moves mouth right.*/
	FtMouthRight BlendShapeEntry = 131
/*Moves mouth left.*/
	FtMouthLeft BlendShapeEntry = 132
/*Right side of the mouth smiles.*/
	FtMouthSmileRight BlendShapeEntry = 133
/*Left side of the mouth smiles.*/
	FtMouthSmileLeft BlendShapeEntry = 134
/*Mouth expresses a smile.*/
	FtMouthSmile BlendShapeEntry = 135
/*Right side of the mouth expresses sadness.*/
	FtMouthSadRight BlendShapeEntry = 136
/*Left side of the mouth expresses sadness.*/
	FtMouthSadLeft BlendShapeEntry = 137
/*Mouth expresses sadness.*/
	FtMouthSad BlendShapeEntry = 138
/*Mouth stretches.*/
	FtMouthStretch BlendShapeEntry = 139
/*Lip corners dimple.*/
	FtMouthDimple BlendShapeEntry = 140
/*Mouth tightens.*/
	FtMouthTightener BlendShapeEntry = 141
/*Mouth presses together.*/
	FtMouthPress BlendShapeEntry = 142
/*Represents the size of the [enum BlendShapeEntry] enum.*/
	FtMax BlendShapeEntry = 143
)
