package Time

import "unsafe"
import "sync"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/Float"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
The Time singleton allows converting time between various formats and also getting time information from the system.
This class conforms with as many of the ISO 8601 standards as possible. All dates follow the Proleptic Gregorian calendar. As such, the day before [code]1582-10-15[/code] is [code]1582-10-14[/code], not [code]1582-10-04[/code]. The year before 1 AD (aka 1 BC) is number [code]0[/code], with the year before that (2 BC) being [code]-1[/code], etc.
Conversion methods assume "the same timezone", and do not handle timezone conversions or DST automatically. Leap seconds are also not handled, they must be done manually if desired. Suffixes such as "Z" are not handled, you need to strip them away manually.
When getting time information from the system, the time can either be in the local timezone or UTC depending on the [code]utc[/code] parameter. However, the [method get_unix_time_from_system] method always uses UTC as it returns the seconds passed since the [url=https://en.wikipedia.org/wiki/Unix_time]Unix epoch[/url].
[b]Important:[/b] The [code]_from_system[/code] methods use the system clock that the user can manually set. [b]Never use[/b] this method for precise time calculation since its results are subject to automatic adjustments by the user or the operating system. [b]Always use[/b] [method get_ticks_usec] or [method get_ticks_msec] for precise time calculation instead, since they are guaranteed to be monotonic (i.e. never decrease).
*/
var self objects.Time
var once sync.Once

func singleton() {
	obj := gd.Global.Object.GetSingleton(gd.Global.Singletons.Time)
	self = *(*objects.Time)(unsafe.Pointer(&obj))
}

/*
Converts the given Unix timestamp to a dictionary of keys: [code]year[/code], [code]month[/code], [code]day[/code], [code]weekday[/code], [code]hour[/code], [code]minute[/code], and [code]second[/code].
The returned Dictionary's values will be the same as the [method get_datetime_dict_from_system] if the Unix timestamp is the current time, with the exception of Daylight Savings Time as it cannot be determined from the epoch.
*/
func GetDatetimeDictFromUnixTime(unix_time_val int) Dictionary.Any {
	once.Do(singleton)
	return Dictionary.Any(class(self).GetDatetimeDictFromUnixTime(gd.Int(unix_time_val)))
}

/*
Converts the given Unix timestamp to a dictionary of keys: [code]year[/code], [code]month[/code], [code]day[/code], and [code]weekday[/code].
*/
func GetDateDictFromUnixTime(unix_time_val int) Dictionary.Any {
	once.Do(singleton)
	return Dictionary.Any(class(self).GetDateDictFromUnixTime(gd.Int(unix_time_val)))
}

/*
Converts the given time to a dictionary of keys: [code]hour[/code], [code]minute[/code], and [code]second[/code].
*/
func GetTimeDictFromUnixTime(unix_time_val int) Dictionary.Any {
	once.Do(singleton)
	return Dictionary.Any(class(self).GetTimeDictFromUnixTime(gd.Int(unix_time_val)))
}

/*
Converts the given Unix timestamp to an ISO 8601 date and time string (YYYY-MM-DDTHH:MM:SS).
If [param use_space] is [code]true[/code], the date and time bits are separated by an empty space character instead of the letter T.
*/
func GetDatetimeStringFromUnixTime(unix_time_val int) string {
	once.Do(singleton)
	return string(class(self).GetDatetimeStringFromUnixTime(gd.Int(unix_time_val), false).String())
}

/*
Converts the given Unix timestamp to an ISO 8601 date string (YYYY-MM-DD).
*/
func GetDateStringFromUnixTime(unix_time_val int) string {
	once.Do(singleton)
	return string(class(self).GetDateStringFromUnixTime(gd.Int(unix_time_val)).String())
}

/*
Converts the given Unix timestamp to an ISO 8601 time string (HH:MM:SS).
*/
func GetTimeStringFromUnixTime(unix_time_val int) string {
	once.Do(singleton)
	return string(class(self).GetTimeStringFromUnixTime(gd.Int(unix_time_val)).String())
}

/*
Converts the given ISO 8601 date and time string (YYYY-MM-DDTHH:MM:SS) to a dictionary of keys: [code]year[/code], [code]month[/code], [code]day[/code], [code skip-lint]weekday[/code], [code]hour[/code], [code]minute[/code], and [code]second[/code].
If [param weekday] is [code]false[/code], then the [code skip-lint]weekday[/code] entry is excluded (the calculation is relatively expensive).
[b]Note:[/b] Any decimal fraction in the time string will be ignored silently.
*/
func GetDatetimeDictFromDatetimeString(datetime string, weekday bool) Dictionary.Any {
	once.Do(singleton)
	return Dictionary.Any(class(self).GetDatetimeDictFromDatetimeString(gd.NewString(datetime), weekday))
}

/*
Converts the given dictionary of keys to an ISO 8601 date and time string (YYYY-MM-DDTHH:MM:SS).
The given dictionary can be populated with the following keys: [code]year[/code], [code]month[/code], [code]day[/code], [code]hour[/code], [code]minute[/code], and [code]second[/code]. Any other entries (including [code]dst[/code]) are ignored.
If the dictionary is empty, [code]0[/code] is returned. If some keys are omitted, they default to the equivalent values for the Unix epoch timestamp 0 (1970-01-01 at 00:00:00).
If [param use_space] is [code]true[/code], the date and time bits are separated by an empty space character instead of the letter T.
*/
func GetDatetimeStringFromDatetimeDict(datetime Dictionary.Any, use_space bool) string {
	once.Do(singleton)
	return string(class(self).GetDatetimeStringFromDatetimeDict(datetime, use_space).String())
}

/*
Converts a dictionary of time values to a Unix timestamp.
The given dictionary can be populated with the following keys: [code]year[/code], [code]month[/code], [code]day[/code], [code]hour[/code], [code]minute[/code], and [code]second[/code]. Any other entries (including [code]dst[/code]) are ignored.
If the dictionary is empty, [code]0[/code] is returned. If some keys are omitted, they default to the equivalent values for the Unix epoch timestamp 0 (1970-01-01 at 00:00:00).
You can pass the output from [method get_datetime_dict_from_unix_time] directly into this function and get the same as what was put in.
[b]Note:[/b] Unix timestamps are often in UTC. This method does not do any timezone conversion, so the timestamp will be in the same timezone as the given datetime dictionary.
*/
func GetUnixTimeFromDatetimeDict(datetime Dictionary.Any) int {
	once.Do(singleton)
	return int(int(class(self).GetUnixTimeFromDatetimeDict(datetime)))
}

/*
Converts the given ISO 8601 date and/or time string to a Unix timestamp. The string can contain a date only, a time only, or both.
[b]Note:[/b] Unix timestamps are often in UTC. This method does not do any timezone conversion, so the timestamp will be in the same timezone as the given datetime string.
[b]Note:[/b] Any decimal fraction in the time string will be ignored silently.
*/
func GetUnixTimeFromDatetimeString(datetime string) int {
	once.Do(singleton)
	return int(int(class(self).GetUnixTimeFromDatetimeString(gd.NewString(datetime))))
}

/*
Converts the given timezone offset in minutes to a timezone offset string. For example, -480 returns "-08:00", 345 returns "+05:45", and 0 returns "+00:00".
*/
func GetOffsetStringFromOffsetMinutes(offset_minutes int) string {
	once.Do(singleton)
	return string(class(self).GetOffsetStringFromOffsetMinutes(gd.Int(offset_minutes)).String())
}

/*
Returns the current date as a dictionary of keys: [code]year[/code], [code]month[/code], [code]day[/code], [code]weekday[/code], [code]hour[/code], [code]minute[/code], [code]second[/code], and [code]dst[/code] (Daylight Savings Time).
*/
func GetDatetimeDictFromSystem() Dictionary.Any {
	once.Do(singleton)
	return Dictionary.Any(class(self).GetDatetimeDictFromSystem(false))
}

/*
Returns the current date as a dictionary of keys: [code]year[/code], [code]month[/code], [code]day[/code], and [code]weekday[/code].
The returned values are in the system's local time when [param utc] is [code]false[/code], otherwise they are in UTC.
*/
func GetDateDictFromSystem() Dictionary.Any {
	once.Do(singleton)
	return Dictionary.Any(class(self).GetDateDictFromSystem(false))
}

/*
Returns the current time as a dictionary of keys: [code]hour[/code], [code]minute[/code], and [code]second[/code].
The returned values are in the system's local time when [param utc] is [code]false[/code], otherwise they are in UTC.
*/
func GetTimeDictFromSystem() Dictionary.Any {
	once.Do(singleton)
	return Dictionary.Any(class(self).GetTimeDictFromSystem(false))
}

/*
Returns the current date and time as an ISO 8601 date and time string (YYYY-MM-DDTHH:MM:SS).
The returned values are in the system's local time when [param utc] is [code]false[/code], otherwise they are in UTC.
If [param use_space] is [code]true[/code], the date and time bits are separated by an empty space character instead of the letter T.
*/
func GetDatetimeStringFromSystem() string {
	once.Do(singleton)
	return string(class(self).GetDatetimeStringFromSystem(false, false).String())
}

/*
Returns the current date as an ISO 8601 date string (YYYY-MM-DD).
The returned values are in the system's local time when [param utc] is [code]false[/code], otherwise they are in UTC.
*/
func GetDateStringFromSystem() string {
	once.Do(singleton)
	return string(class(self).GetDateStringFromSystem(false).String())
}

/*
Returns the current time as an ISO 8601 time string (HH:MM:SS).
The returned values are in the system's local time when [param utc] is [code]false[/code], otherwise they are in UTC.
*/
func GetTimeStringFromSystem() string {
	once.Do(singleton)
	return string(class(self).GetTimeStringFromSystem(false).String())
}

/*
Returns the current time zone as a dictionary of keys: [code]bias[/code] and [code]name[/code].
- [code]bias[/code] is the offset from UTC in minutes, since not all time zones are multiples of an hour from UTC.
- [code]name[/code] is the localized name of the time zone, according to the OS locale settings of the current user.
*/
func GetTimeZoneFromSystem() Dictionary.Any {
	once.Do(singleton)
	return Dictionary.Any(class(self).GetTimeZoneFromSystem())
}

/*
Returns the current Unix timestamp in seconds based on the system time in UTC. This method is implemented by the operating system and always returns the time in UTC. The Unix timestamp is the number of seconds passed since 1970-01-01 at 00:00:00, the [url=https://en.wikipedia.org/wiki/Unix_time]Unix epoch[/url].
[b]Note:[/b] Unlike other methods that use integer timestamps, this method returns the timestamp as a [float] for sub-second precision.
*/
func GetUnixTimeFromSystem() Float.X {
	once.Do(singleton)
	return Float.X(Float.X(class(self).GetUnixTimeFromSystem()))
}

/*
Returns the amount of time passed in milliseconds since the engine started.
Will always be positive or 0 and uses a 64-bit value (it will wrap after roughly 500 million years).
*/
func GetTicksMsec() int {
	once.Do(singleton)
	return int(int(class(self).GetTicksMsec()))
}

/*
Returns the amount of time passed in microseconds since the engine started.
Will always be positive or 0 and uses a 64-bit value (it will wrap after roughly half a million years).
*/
func GetTicksUsec() int {
	once.Do(singleton)
	return int(int(class(self).GetTicksUsec()))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
func Advanced() class { once.Do(singleton); return self }

type class [1]classdb.Time

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }

/*
Converts the given Unix timestamp to a dictionary of keys: [code]year[/code], [code]month[/code], [code]day[/code], [code]weekday[/code], [code]hour[/code], [code]minute[/code], and [code]second[/code].
The returned Dictionary's values will be the same as the [method get_datetime_dict_from_system] if the Unix timestamp is the current time, with the exception of Daylight Savings Time as it cannot be determined from the epoch.
*/
//go:nosplit
func (self class) GetDatetimeDictFromUnixTime(unix_time_val gd.Int) gd.Dictionary {
	var frame = callframe.New()
	callframe.Arg(frame, unix_time_val)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Time.Bind_get_datetime_dict_from_unix_time, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Dictionary](r_ret.Get())
	frame.Free()
	return ret
}

/*
Converts the given Unix timestamp to a dictionary of keys: [code]year[/code], [code]month[/code], [code]day[/code], and [code]weekday[/code].
*/
//go:nosplit
func (self class) GetDateDictFromUnixTime(unix_time_val gd.Int) gd.Dictionary {
	var frame = callframe.New()
	callframe.Arg(frame, unix_time_val)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Time.Bind_get_date_dict_from_unix_time, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Dictionary](r_ret.Get())
	frame.Free()
	return ret
}

/*
Converts the given time to a dictionary of keys: [code]hour[/code], [code]minute[/code], and [code]second[/code].
*/
//go:nosplit
func (self class) GetTimeDictFromUnixTime(unix_time_val gd.Int) gd.Dictionary {
	var frame = callframe.New()
	callframe.Arg(frame, unix_time_val)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Time.Bind_get_time_dict_from_unix_time, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Dictionary](r_ret.Get())
	frame.Free()
	return ret
}

/*
Converts the given Unix timestamp to an ISO 8601 date and time string (YYYY-MM-DDTHH:MM:SS).
If [param use_space] is [code]true[/code], the date and time bits are separated by an empty space character instead of the letter T.
*/
//go:nosplit
func (self class) GetDatetimeStringFromUnixTime(unix_time_val gd.Int, use_space bool) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, unix_time_val)
	callframe.Arg(frame, use_space)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Time.Bind_get_datetime_string_from_unix_time, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Converts the given Unix timestamp to an ISO 8601 date string (YYYY-MM-DD).
*/
//go:nosplit
func (self class) GetDateStringFromUnixTime(unix_time_val gd.Int) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, unix_time_val)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Time.Bind_get_date_string_from_unix_time, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Converts the given Unix timestamp to an ISO 8601 time string (HH:MM:SS).
*/
//go:nosplit
func (self class) GetTimeStringFromUnixTime(unix_time_val gd.Int) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, unix_time_val)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Time.Bind_get_time_string_from_unix_time, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Converts the given ISO 8601 date and time string (YYYY-MM-DDTHH:MM:SS) to a dictionary of keys: [code]year[/code], [code]month[/code], [code]day[/code], [code skip-lint]weekday[/code], [code]hour[/code], [code]minute[/code], and [code]second[/code].
If [param weekday] is [code]false[/code], then the [code skip-lint]weekday[/code] entry is excluded (the calculation is relatively expensive).
[b]Note:[/b] Any decimal fraction in the time string will be ignored silently.
*/
//go:nosplit
func (self class) GetDatetimeDictFromDatetimeString(datetime gd.String, weekday bool) gd.Dictionary {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(datetime))
	callframe.Arg(frame, weekday)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Time.Bind_get_datetime_dict_from_datetime_string, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Dictionary](r_ret.Get())
	frame.Free()
	return ret
}

/*
Converts the given dictionary of keys to an ISO 8601 date and time string (YYYY-MM-DDTHH:MM:SS).
The given dictionary can be populated with the following keys: [code]year[/code], [code]month[/code], [code]day[/code], [code]hour[/code], [code]minute[/code], and [code]second[/code]. Any other entries (including [code]dst[/code]) are ignored.
If the dictionary is empty, [code]0[/code] is returned. If some keys are omitted, they default to the equivalent values for the Unix epoch timestamp 0 (1970-01-01 at 00:00:00).
If [param use_space] is [code]true[/code], the date and time bits are separated by an empty space character instead of the letter T.
*/
//go:nosplit
func (self class) GetDatetimeStringFromDatetimeDict(datetime gd.Dictionary, use_space bool) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(datetime))
	callframe.Arg(frame, use_space)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Time.Bind_get_datetime_string_from_datetime_dict, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Converts a dictionary of time values to a Unix timestamp.
The given dictionary can be populated with the following keys: [code]year[/code], [code]month[/code], [code]day[/code], [code]hour[/code], [code]minute[/code], and [code]second[/code]. Any other entries (including [code]dst[/code]) are ignored.
If the dictionary is empty, [code]0[/code] is returned. If some keys are omitted, they default to the equivalent values for the Unix epoch timestamp 0 (1970-01-01 at 00:00:00).
You can pass the output from [method get_datetime_dict_from_unix_time] directly into this function and get the same as what was put in.
[b]Note:[/b] Unix timestamps are often in UTC. This method does not do any timezone conversion, so the timestamp will be in the same timezone as the given datetime dictionary.
*/
//go:nosplit
func (self class) GetUnixTimeFromDatetimeDict(datetime gd.Dictionary) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(datetime))
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Time.Bind_get_unix_time_from_datetime_dict, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Converts the given ISO 8601 date and/or time string to a Unix timestamp. The string can contain a date only, a time only, or both.
[b]Note:[/b] Unix timestamps are often in UTC. This method does not do any timezone conversion, so the timestamp will be in the same timezone as the given datetime string.
[b]Note:[/b] Any decimal fraction in the time string will be ignored silently.
*/
//go:nosplit
func (self class) GetUnixTimeFromDatetimeString(datetime gd.String) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(datetime))
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Time.Bind_get_unix_time_from_datetime_string, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Converts the given timezone offset in minutes to a timezone offset string. For example, -480 returns "-08:00", 345 returns "+05:45", and 0 returns "+00:00".
*/
//go:nosplit
func (self class) GetOffsetStringFromOffsetMinutes(offset_minutes gd.Int) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, offset_minutes)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Time.Bind_get_offset_string_from_offset_minutes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the current date as a dictionary of keys: [code]year[/code], [code]month[/code], [code]day[/code], [code]weekday[/code], [code]hour[/code], [code]minute[/code], [code]second[/code], and [code]dst[/code] (Daylight Savings Time).
*/
//go:nosplit
func (self class) GetDatetimeDictFromSystem(utc bool) gd.Dictionary {
	var frame = callframe.New()
	callframe.Arg(frame, utc)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Time.Bind_get_datetime_dict_from_system, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Dictionary](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the current date as a dictionary of keys: [code]year[/code], [code]month[/code], [code]day[/code], and [code]weekday[/code].
The returned values are in the system's local time when [param utc] is [code]false[/code], otherwise they are in UTC.
*/
//go:nosplit
func (self class) GetDateDictFromSystem(utc bool) gd.Dictionary {
	var frame = callframe.New()
	callframe.Arg(frame, utc)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Time.Bind_get_date_dict_from_system, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Dictionary](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the current time as a dictionary of keys: [code]hour[/code], [code]minute[/code], and [code]second[/code].
The returned values are in the system's local time when [param utc] is [code]false[/code], otherwise they are in UTC.
*/
//go:nosplit
func (self class) GetTimeDictFromSystem(utc bool) gd.Dictionary {
	var frame = callframe.New()
	callframe.Arg(frame, utc)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Time.Bind_get_time_dict_from_system, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Dictionary](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the current date and time as an ISO 8601 date and time string (YYYY-MM-DDTHH:MM:SS).
The returned values are in the system's local time when [param utc] is [code]false[/code], otherwise they are in UTC.
If [param use_space] is [code]true[/code], the date and time bits are separated by an empty space character instead of the letter T.
*/
//go:nosplit
func (self class) GetDatetimeStringFromSystem(utc bool, use_space bool) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, utc)
	callframe.Arg(frame, use_space)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Time.Bind_get_datetime_string_from_system, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the current date as an ISO 8601 date string (YYYY-MM-DD).
The returned values are in the system's local time when [param utc] is [code]false[/code], otherwise they are in UTC.
*/
//go:nosplit
func (self class) GetDateStringFromSystem(utc bool) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, utc)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Time.Bind_get_date_string_from_system, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the current time as an ISO 8601 time string (HH:MM:SS).
The returned values are in the system's local time when [param utc] is [code]false[/code], otherwise they are in UTC.
*/
//go:nosplit
func (self class) GetTimeStringFromSystem(utc bool) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, utc)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Time.Bind_get_time_string_from_system, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the current time zone as a dictionary of keys: [code]bias[/code] and [code]name[/code].
- [code]bias[/code] is the offset from UTC in minutes, since not all time zones are multiples of an hour from UTC.
- [code]name[/code] is the localized name of the time zone, according to the OS locale settings of the current user.
*/
//go:nosplit
func (self class) GetTimeZoneFromSystem() gd.Dictionary {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Time.Bind_get_time_zone_from_system, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Dictionary](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the current Unix timestamp in seconds based on the system time in UTC. This method is implemented by the operating system and always returns the time in UTC. The Unix timestamp is the number of seconds passed since 1970-01-01 at 00:00:00, the [url=https://en.wikipedia.org/wiki/Unix_time]Unix epoch[/url].
[b]Note:[/b] Unlike other methods that use integer timestamps, this method returns the timestamp as a [float] for sub-second precision.
*/
//go:nosplit
func (self class) GetUnixTimeFromSystem() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Time.Bind_get_unix_time_from_system, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the amount of time passed in milliseconds since the engine started.
Will always be positive or 0 and uses a 64-bit value (it will wrap after roughly 500 million years).
*/
//go:nosplit
func (self class) GetTicksMsec() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Time.Bind_get_ticks_msec, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the amount of time passed in microseconds since the engine started.
Will always be positive or 0 and uses a 64-bit value (it will wrap after roughly half a million years).
*/
//go:nosplit
func (self class) GetTicksUsec() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Time.Bind_get_ticks_usec, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsObject(), name)
	}
}
func init() {
	classdb.Register("Time", func(ptr gd.Object) any { return [1]classdb.Time{*(*classdb.Time)(unsafe.Pointer(&ptr))} })
}

type Month = classdb.TimeMonth

const (
	/*The month of January, represented numerically as [code]01[/code].*/
	MonthJanuary Month = 1
	/*The month of February, represented numerically as [code]02[/code].*/
	MonthFebruary Month = 2
	/*The month of March, represented numerically as [code]03[/code].*/
	MonthMarch Month = 3
	/*The month of April, represented numerically as [code]04[/code].*/
	MonthApril Month = 4
	/*The month of May, represented numerically as [code]05[/code].*/
	MonthMay Month = 5
	/*The month of June, represented numerically as [code]06[/code].*/
	MonthJune Month = 6
	/*The month of July, represented numerically as [code]07[/code].*/
	MonthJuly Month = 7
	/*The month of August, represented numerically as [code]08[/code].*/
	MonthAugust Month = 8
	/*The month of September, represented numerically as [code]09[/code].*/
	MonthSeptember Month = 9
	/*The month of October, represented numerically as [code]10[/code].*/
	MonthOctober Month = 10
	/*The month of November, represented numerically as [code]11[/code].*/
	MonthNovember Month = 11
	/*The month of December, represented numerically as [code]12[/code].*/
	MonthDecember Month = 12
)

type Weekday = classdb.TimeWeekday

const (
	/*The day of the week Sunday, represented numerically as [code]0[/code].*/
	WeekdaySunday Weekday = 0
	/*The day of the week Monday, represented numerically as [code]1[/code].*/
	WeekdayMonday Weekday = 1
	/*The day of the week Tuesday, represented numerically as [code]2[/code].*/
	WeekdayTuesday Weekday = 2
	/*The day of the week Wednesday, represented numerically as [code]3[/code].*/
	WeekdayWednesday Weekday = 3
	/*The day of the week Thursday, represented numerically as [code]4[/code].*/
	WeekdayThursday Weekday = 4
	/*The day of the week Friday, represented numerically as [code]5[/code].*/
	WeekdayFriday Weekday = 5
	/*The day of the week Saturday, represented numerically as [code]6[/code].*/
	WeekdaySaturday Weekday = 6
)
