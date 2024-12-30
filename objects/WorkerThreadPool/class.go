package WorkerThreadPool

import "unsafe"
import "sync"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
The [WorkerThreadPool] singleton allocates a set of [Thread]s (called worker threads) on project startup and provides methods for offloading tasks to them. This can be used for simple multithreading without having to create [Thread]s.
Tasks hold the [Callable] to be run by the threads. [WorkerThreadPool] can be used to create regular tasks, which will be taken by one worker thread, or group tasks, which can be distributed between multiple worker threads. Group tasks execute the [Callable] multiple times, which makes them useful for iterating over a lot of elements, such as the enemies in an arena.
Here's a sample on how to offload an expensive function to worker threads:
[codeblocks]
[gdscript]
var enemies = [] # An array to be filled with enemies.

func process_enemy_ai(enemy_index):

	var processed_enemy = enemies[enemy_index]
	# Expensive logic...

func _process(delta):

	var task_id = WorkerThreadPool.add_group_task(process_enemy_ai, enemies.size())
	# Other code...
	WorkerThreadPool.wait_for_group_task_completion(task_id)
	# Other code that depends on the enemy AI already being processed.

[/gdscript]
[csharp]
private List<Node> _enemies = new List<Node>(); // A list to be filled with enemies.

private void ProcessEnemyAI(int enemyIndex)

	{
	    Node processedEnemy = _enemies[enemyIndex];
	    // Expensive logic here.
	}

public override void _Process(double delta)

	{
	    long taskId = WorkerThreadPool.AddGroupTask(Callable.From<int>(ProcessEnemyAI), _enemies.Count);
	    // Other code...
	    WorkerThreadPool.WaitForGroupTaskCompletion(taskId);
	    // Other code that depends on the enemy AI already being processed.
	}

[/csharp]
[/codeblocks]
The above code relies on the number of elements in the [code]enemies[/code] array remaining constant during the multithreaded part.
[b]Note:[/b] Using this singleton could affect performance negatively if the task being distributed between threads is not computationally expensive.
*/
var self objects.WorkerThreadPool
var once sync.Once

func singleton() {
	obj := gd.Global.Object.GetSingleton(gd.Global.Singletons.WorkerThreadPool)
	self = *(*objects.WorkerThreadPool)(unsafe.Pointer(&obj))
}

/*
Adds [param action] as a task to be executed by a worker thread. [param high_priority] determines if the task has a high priority or a low priority (default). You can optionally provide a [param description] to help with debugging.
Returns a task ID that can be used by other methods.
[b]Warning:[/b] Every task must be waited for completion using [method wait_for_task_completion] or [method wait_for_group_task_completion] at some point so that any allocated resources inside the task can be cleaned up.
*/
func AddTask(action func()) int {
	once.Do(singleton)
	return int(int(class(self).AddTask(gd.NewCallable(action), false, gd.NewString(""))))
}

/*
Returns [code]true[/code] if the task with the given ID is completed.
[b]Note:[/b] You should only call this method between adding the task and awaiting its completion.
*/
func IsTaskCompleted(task_id int) bool {
	once.Do(singleton)
	return bool(class(self).IsTaskCompleted(gd.Int(task_id)))
}

/*
Pauses the thread that calls this method until the task with the given ID is completed.
Returns [constant @GlobalScope.OK] if the task could be successfully awaited.
Returns [constant @GlobalScope.ERR_INVALID_PARAMETER] if a task with the passed ID does not exist (maybe because it was already awaited and disposed of).
Returns [constant @GlobalScope.ERR_BUSY] if the call is made from another running task and, due to task scheduling, there's potential for deadlocking (e.g., the task to await may be at a lower level in the call stack and therefore can't progress). This is an advanced situation that should only matter when some tasks depend on others (in the current implementation, the tricky case is a task trying to wait on an older one).
*/
func WaitForTaskCompletion(task_id int) error {
	once.Do(singleton)
	return error(class(self).WaitForTaskCompletion(gd.Int(task_id)))
}

/*
Adds [param action] as a group task to be executed by the worker threads. The [Callable] will be called a number of times based on [param elements], with the first thread calling it with the value [code]0[/code] as a parameter, and each consecutive execution incrementing this value by 1 until it reaches [code]element - 1[/code].
The number of threads the task is distributed to is defined by [param tasks_needed], where the default value [code]-1[/code] means it is distributed to all worker threads. [param high_priority] determines if the task has a high priority or a low priority (default). You can optionally provide a [param description] to help with debugging.
Returns a group task ID that can be used by other methods.
[b]Warning:[/b] Every task must be waited for completion using [method wait_for_task_completion] or [method wait_for_group_task_completion] at some point so that any allocated resources inside the task can be cleaned up.
*/
func AddGroupTask(action func(), elements int) int {
	once.Do(singleton)
	return int(int(class(self).AddGroupTask(gd.NewCallable(action), gd.Int(elements), gd.Int(-1), false, gd.NewString(""))))
}

/*
Returns [code]true[/code] if the group task with the given ID is completed.
[b]Note:[/b] You should only call this method between adding the group task and awaiting its completion.
*/
func IsGroupTaskCompleted(group_id int) bool {
	once.Do(singleton)
	return bool(class(self).IsGroupTaskCompleted(gd.Int(group_id)))
}

/*
Returns how many times the [Callable] of the group task with the given ID has already been executed by the worker threads.
[b]Note:[/b] If a thread has started executing the [Callable] but is yet to finish, it won't be counted.
*/
func GetGroupProcessedElementCount(group_id int) int {
	once.Do(singleton)
	return int(int(class(self).GetGroupProcessedElementCount(gd.Int(group_id))))
}

/*
Pauses the thread that calls this method until the group task with the given ID is completed.
*/
func WaitForGroupTaskCompletion(group_id int) {
	once.Do(singleton)
	class(self).WaitForGroupTaskCompletion(gd.Int(group_id))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
func Advanced() class { once.Do(singleton); return self }

type class [1]classdb.WorkerThreadPool

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }

/*
Adds [param action] as a task to be executed by a worker thread. [param high_priority] determines if the task has a high priority or a low priority (default). You can optionally provide a [param description] to help with debugging.
Returns a task ID that can be used by other methods.
[b]Warning:[/b] Every task must be waited for completion using [method wait_for_task_completion] or [method wait_for_group_task_completion] at some point so that any allocated resources inside the task can be cleaned up.
*/
//go:nosplit
func (self class) AddTask(action gd.Callable, high_priority bool, description gd.String) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(action))
	callframe.Arg(frame, high_priority)
	callframe.Arg(frame, pointers.Get(description))
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WorkerThreadPool.Bind_add_task, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the task with the given ID is completed.
[b]Note:[/b] You should only call this method between adding the task and awaiting its completion.
*/
//go:nosplit
func (self class) IsTaskCompleted(task_id gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, task_id)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WorkerThreadPool.Bind_is_task_completed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Pauses the thread that calls this method until the task with the given ID is completed.
Returns [constant @GlobalScope.OK] if the task could be successfully awaited.
Returns [constant @GlobalScope.ERR_INVALID_PARAMETER] if a task with the passed ID does not exist (maybe because it was already awaited and disposed of).
Returns [constant @GlobalScope.ERR_BUSY] if the call is made from another running task and, due to task scheduling, there's potential for deadlocking (e.g., the task to await may be at a lower level in the call stack and therefore can't progress). This is an advanced situation that should only matter when some tasks depend on others (in the current implementation, the tricky case is a task trying to wait on an older one).
*/
//go:nosplit
func (self class) WaitForTaskCompletion(task_id gd.Int) error {
	var frame = callframe.New()
	callframe.Arg(frame, task_id)
	var r_ret = callframe.Ret[error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WorkerThreadPool.Bind_wait_for_task_completion, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Adds [param action] as a group task to be executed by the worker threads. The [Callable] will be called a number of times based on [param elements], with the first thread calling it with the value [code]0[/code] as a parameter, and each consecutive execution incrementing this value by 1 until it reaches [code]element - 1[/code].
The number of threads the task is distributed to is defined by [param tasks_needed], where the default value [code]-1[/code] means it is distributed to all worker threads. [param high_priority] determines if the task has a high priority or a low priority (default). You can optionally provide a [param description] to help with debugging.
Returns a group task ID that can be used by other methods.
[b]Warning:[/b] Every task must be waited for completion using [method wait_for_task_completion] or [method wait_for_group_task_completion] at some point so that any allocated resources inside the task can be cleaned up.
*/
//go:nosplit
func (self class) AddGroupTask(action gd.Callable, elements gd.Int, tasks_needed gd.Int, high_priority bool, description gd.String) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(action))
	callframe.Arg(frame, elements)
	callframe.Arg(frame, tasks_needed)
	callframe.Arg(frame, high_priority)
	callframe.Arg(frame, pointers.Get(description))
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WorkerThreadPool.Bind_add_group_task, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the group task with the given ID is completed.
[b]Note:[/b] You should only call this method between adding the group task and awaiting its completion.
*/
//go:nosplit
func (self class) IsGroupTaskCompleted(group_id gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, group_id)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WorkerThreadPool.Bind_is_group_task_completed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns how many times the [Callable] of the group task with the given ID has already been executed by the worker threads.
[b]Note:[/b] If a thread has started executing the [Callable] but is yet to finish, it won't be counted.
*/
//go:nosplit
func (self class) GetGroupProcessedElementCount(group_id gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, group_id)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WorkerThreadPool.Bind_get_group_processed_element_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Pauses the thread that calls this method until the group task with the given ID is completed.
*/
//go:nosplit
func (self class) WaitForGroupTaskCompletion(group_id gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, group_id)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WorkerThreadPool.Bind_wait_for_group_task_completion, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsObject(), name)
	}
}
func init() {
	classdb.Register("WorkerThreadPool", func(ptr gd.Object) any { return classdb.WorkerThreadPool(ptr) })
}

type Error int

const (
	/*Methods that return [enum Error] return [constant OK] when no error occurred.
	  Since [constant OK] has value 0, and all other error constants are positive integers, it can also be used in boolean checks.
	  [b]Example:[/b]
	  [codeblock]
	  var error = method_that_returns_error()
	  if error != OK:
	      printerr("Failure!")

	  # Or, alternatively:
	  if error:
	      printerr("Still failing!")
	  [/codeblock]
	  [b]Note:[/b] Many functions do not return an error code, but will print error messages to standard output.*/
	Ok Error = 0
	/*Generic error.*/
	Failed Error = 1
	/*Unavailable error.*/
	ErrUnavailable Error = 2
	/*Unconfigured error.*/
	ErrUnconfigured Error = 3
	/*Unauthorized error.*/
	ErrUnauthorized Error = 4
	/*Parameter range error.*/
	ErrParameterRangeError Error = 5
	/*Out of memory (OOM) error.*/
	ErrOutOfMemory Error = 6
	/*File: Not found error.*/
	ErrFileNotFound Error = 7
	/*File: Bad drive error.*/
	ErrFileBadDrive Error = 8
	/*File: Bad path error.*/
	ErrFileBadPath Error = 9
	/*File: No permission error.*/
	ErrFileNoPermission Error = 10
	/*File: Already in use error.*/
	ErrFileAlreadyInUse Error = 11
	/*File: Can't open error.*/
	ErrFileCantOpen Error = 12
	/*File: Can't write error.*/
	ErrFileCantWrite Error = 13
	/*File: Can't read error.*/
	ErrFileCantRead Error = 14
	/*File: Unrecognized error.*/
	ErrFileUnrecognized Error = 15
	/*File: Corrupt error.*/
	ErrFileCorrupt Error = 16
	/*File: Missing dependencies error.*/
	ErrFileMissingDependencies Error = 17
	/*File: End of file (EOF) error.*/
	ErrFileEof Error = 18
	/*Can't open error.*/
	ErrCantOpen Error = 19
	/*Can't create error.*/
	ErrCantCreate Error = 20
	/*Query failed error.*/
	ErrQueryFailed Error = 21
	/*Already in use error.*/
	ErrAlreadyInUse Error = 22
	/*Locked error.*/
	ErrLocked Error = 23
	/*Timeout error.*/
	ErrTimeout Error = 24
	/*Can't connect error.*/
	ErrCantConnect Error = 25
	/*Can't resolve error.*/
	ErrCantResolve Error = 26
	/*Connection error.*/
	ErrConnectionError Error = 27
	/*Can't acquire resource error.*/
	ErrCantAcquireResource Error = 28
	/*Can't fork process error.*/
	ErrCantFork Error = 29
	/*Invalid data error.*/
	ErrInvalidData Error = 30
	/*Invalid parameter error.*/
	ErrInvalidParameter Error = 31
	/*Already exists error.*/
	ErrAlreadyExists Error = 32
	/*Does not exist error.*/
	ErrDoesNotExist Error = 33
	/*Database: Read error.*/
	ErrDatabaseCantRead Error = 34
	/*Database: Write error.*/
	ErrDatabaseCantWrite Error = 35
	/*Compilation failed error.*/
	ErrCompilationFailed Error = 36
	/*Method not found error.*/
	ErrMethodNotFound Error = 37
	/*Linking failed error.*/
	ErrLinkFailed Error = 38
	/*Script failed error.*/
	ErrScriptFailed Error = 39
	/*Cycling link (import cycle) error.*/
	ErrCyclicLink Error = 40
	/*Invalid declaration error.*/
	ErrInvalidDeclaration Error = 41
	/*Duplicate symbol error.*/
	ErrDuplicateSymbol Error = 42
	/*Parse error.*/
	ErrParseError Error = 43
	/*Busy error.*/
	ErrBusy Error = 44
	/*Skip error.*/
	ErrSkip Error = 45
	/*Help error. Used internally when passing [code]--version[/code] or [code]--help[/code] as executable options.*/
	ErrHelp Error = 46
	/*Bug error, caused by an implementation issue in the method.
	  [b]Note:[/b] If a built-in method returns this code, please open an issue on [url=https://github.com/godotengine/godot/issues]the GitHub Issue Tracker[/url].*/
	ErrBug Error = 47
	/*Printer on fire error (This is an easter egg, no built-in methods return this error code).*/
	ErrPrinterOnFire Error = 48
)
