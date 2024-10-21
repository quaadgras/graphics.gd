package EditorVCSInterface

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Defines the API that the editor uses to extract information from the underlying VCS. The implementation of this API is included in VCS plugins, which are GDExtension plugins that inherit [EditorVCSInterface] and are attached (on demand) to the singleton instance of [EditorVCSInterface]. Instead of performing the task themselves, all the virtual functions listed below are calling the internally overridden functions in the VCS plugins to provide a plug-n-play experience. A custom VCS plugin is supposed to inherit from [EditorVCSInterface] and override each of these virtual functions.
	// EditorVCSInterface methods that can be overridden by a [Class] that extends it.
	type EditorVCSInterface interface {
		//Initializes the VCS plugin when called from the editor. Returns whether or not the plugin was successfully initialized. A VCS project is initialized at [param project_path].
		Initialize(project_path gd.String) bool
		//Set user credentials in the underlying VCS. [param username] and [param password] are used only during HTTPS authentication unless not already mentioned in the remote URL. [param ssh_public_key_path], [param ssh_private_key_path], and [param ssh_passphrase] are only used during SSH authentication.
		SetCredentials(username gd.String, password gd.String, ssh_public_key_path gd.String, ssh_private_key_path gd.String, ssh_passphrase gd.String) 
		//Returns an [Array] of [Dictionary] items (see [method create_status_file]), each containing the status data of every modified file in the project folder.
		GetModifiedFilesData() gd.ArrayOf[gd.Dictionary]
		//Stages the file present at [param file_path] to the staged area.
		StageFile(file_path gd.String) 
		//Unstages the file present at [param file_path] from the staged area to the unstaged area.
		UnstageFile(file_path gd.String) 
		//Discards the changes made in a file present at [param file_path].
		DiscardFile(file_path gd.String) 
		//Commits the currently staged changes and applies the commit [param msg] to the resulting commit.
		Commit(msg gd.String) 
		//Returns an array of [Dictionary] items (see [method create_diff_file], [method create_diff_hunk], [method create_diff_line], [method add_line_diffs_into_diff_hunk] and [method add_diff_hunks_into_diff_file]), each containing information about a diff. If [param identifier] is a file path, returns a file diff, and if it is a commit identifier, then returns a commit diff.
		GetDiff(identifier gd.String, area gd.Int) gd.ArrayOf[gd.Dictionary]
		//Shuts down VCS plugin instance. Called when the user either closes the editor or shuts down the VCS plugin through the editor UI.
		ShutDown() bool
		//Returns the name of the underlying VCS provider.
		GetVcsName() gd.String
		//Returns an [Array] of [Dictionary] items (see [method create_commit]), each containing the data for a past commit.
		GetPreviousCommits(max_commits gd.Int) gd.ArrayOf[gd.Dictionary]
		//Gets an instance of an [Array] of [String]s containing available branch names in the VCS.
		GetBranchList() gd.ArrayOf[gd.String]
		//Returns an [Array] of [String]s, each containing the name of a remote configured in the VCS.
		GetRemotes() gd.ArrayOf[gd.String]
		//Creates a new branch named [param branch_name] in the VCS.
		CreateBranch(branch_name gd.String) 
		//Remove a branch from the local VCS.
		RemoveBranch(branch_name gd.String) 
		//Creates a new remote destination with name [param remote_name] and points it to [param remote_url]. This can be an HTTPS remote or an SSH remote.
		CreateRemote(remote_name gd.String, remote_url gd.String) 
		//Remove a remote from the local VCS.
		RemoveRemote(remote_name gd.String) 
		//Gets the current branch name defined in the VCS.
		GetCurrentBranchName() gd.String
		//Checks out a [param branch_name] in the VCS.
		CheckoutBranch(branch_name gd.String) bool
		//Pulls changes from the remote. This can give rise to merge conflicts.
		Pull(remote gd.String) 
		//Pushes changes to the [param remote]. If [param force] is [code]true[/code], a force push will override the change history already present on the remote.
		Push(remote gd.String, force bool) 
		//Fetches new changes from the [param remote], but doesn't write changes to the current working directory. Equivalent to [code]git fetch[/code].
		Fetch(remote gd.String) 
		//Returns an [Array] of [Dictionary] items (see [method create_diff_hunk]), each containing a line diff between a file at [param file_path] and the [param text] which is passed in.
		GetLineDiff(file_path gd.String, text gd.String) gd.ArrayOf[gd.Dictionary]
	}

*/
type Simple [1]classdb.EditorVCSInterface
func (Simple) _initialize(impl func(ptr unsafe.Pointer, project_path string) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var project_path = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,0))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, project_path.String())
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _set_credentials(impl func(ptr unsafe.Pointer, username string, password string, ssh_public_key_path string, ssh_private_key_path string, ssh_passphrase string) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var username = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,0))
		var password = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,1))
		var ssh_public_key_path = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,2))
		var ssh_private_key_path = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,3))
		var ssh_passphrase = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,4))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, username.String(), password.String(), ssh_public_key_path.String(), ssh_private_key_path.String(), ssh_passphrase.String())
		gc.End()
	}
}
func (Simple) _get_modified_files_data(impl func(ptr unsafe.Pointer) gd.ArrayOf[gd.Dictionary], api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(ret.Array()))
		gc.End()
	}
}
func (Simple) _stage_file(impl func(ptr unsafe.Pointer, file_path string) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var file_path = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,0))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, file_path.String())
		gc.End()
	}
}
func (Simple) _unstage_file(impl func(ptr unsafe.Pointer, file_path string) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var file_path = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,0))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, file_path.String())
		gc.End()
	}
}
func (Simple) _discard_file(impl func(ptr unsafe.Pointer, file_path string) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var file_path = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,0))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, file_path.String())
		gc.End()
	}
}
func (Simple) _commit(impl func(ptr unsafe.Pointer, msg string) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var msg = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,0))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, msg.String())
		gc.End()
	}
}
func (Simple) _get_diff(impl func(ptr unsafe.Pointer, identifier string, area int) gd.ArrayOf[gd.Dictionary], api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var identifier = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,0))
		var area = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, identifier.String(), int(area))
		gd.UnsafeSet(p_back, mmm.End(ret.Array()))
		gc.End()
	}
}
func (Simple) _shut_down(impl func(ptr unsafe.Pointer) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _get_vcs_name(impl func(ptr unsafe.Pointer) string, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(gc.String(ret)))
		gc.End()
	}
}
func (Simple) _get_previous_commits(impl func(ptr unsafe.Pointer, max_commits int) gd.ArrayOf[gd.Dictionary], api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var max_commits = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(max_commits))
		gd.UnsafeSet(p_back, mmm.End(ret.Array()))
		gc.End()
	}
}
func (Simple) _get_branch_list(impl func(ptr unsafe.Pointer) gd.ArrayOf[gd.String], api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(ret.Array()))
		gc.End()
	}
}
func (Simple) _get_remotes(impl func(ptr unsafe.Pointer) gd.ArrayOf[gd.String], api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(ret.Array()))
		gc.End()
	}
}
func (Simple) _create_branch(impl func(ptr unsafe.Pointer, branch_name string) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var branch_name = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,0))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, branch_name.String())
		gc.End()
	}
}
func (Simple) _remove_branch(impl func(ptr unsafe.Pointer, branch_name string) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var branch_name = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,0))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, branch_name.String())
		gc.End()
	}
}
func (Simple) _create_remote(impl func(ptr unsafe.Pointer, remote_name string, remote_url string) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var remote_name = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,0))
		var remote_url = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, remote_name.String(), remote_url.String())
		gc.End()
	}
}
func (Simple) _remove_remote(impl func(ptr unsafe.Pointer, remote_name string) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var remote_name = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,0))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, remote_name.String())
		gc.End()
	}
}
func (Simple) _get_current_branch_name(impl func(ptr unsafe.Pointer) string, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(gc.String(ret)))
		gc.End()
	}
}
func (Simple) _checkout_branch(impl func(ptr unsafe.Pointer, branch_name string) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var branch_name = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,0))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, branch_name.String())
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _pull(impl func(ptr unsafe.Pointer, remote string) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var remote = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,0))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, remote.String())
		gc.End()
	}
}
func (Simple) _push(impl func(ptr unsafe.Pointer, remote string, force bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var remote = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,0))
		var force = gd.UnsafeGet[bool](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, remote.String(), force)
		gc.End()
	}
}
func (Simple) _fetch(impl func(ptr unsafe.Pointer, remote string) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var remote = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,0))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, remote.String())
		gc.End()
	}
}
func (Simple) _get_line_diff(impl func(ptr unsafe.Pointer, file_path string, text string) gd.ArrayOf[gd.Dictionary], api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var file_path = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,0))
		var text = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, file_path.String(), text.String())
		gd.UnsafeSet(p_back, mmm.End(ret.Array()))
		gc.End()
	}
}
func (self Simple) CreateDiffLine(new_line_no int, old_line_no int, content string, status string) gd.Dictionary {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Dictionary(Expert(self).CreateDiffLine(gc, gd.Int(new_line_no), gd.Int(old_line_no), gc.String(content), gc.String(status)))
}
func (self Simple) CreateDiffHunk(old_start int, new_start int, old_lines int, new_lines int) gd.Dictionary {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Dictionary(Expert(self).CreateDiffHunk(gc, gd.Int(old_start), gd.Int(new_start), gd.Int(old_lines), gd.Int(new_lines)))
}
func (self Simple) CreateDiffFile(new_file string, old_file string) gd.Dictionary {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Dictionary(Expert(self).CreateDiffFile(gc, gc.String(new_file), gc.String(old_file)))
}
func (self Simple) CreateCommit(msg string, author string, id string, unix_timestamp int, offset_minutes int) gd.Dictionary {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Dictionary(Expert(self).CreateCommit(gc, gc.String(msg), gc.String(author), gc.String(id), gd.Int(unix_timestamp), gd.Int(offset_minutes)))
}
func (self Simple) CreateStatusFile(file_path string, change_type classdb.EditorVCSInterfaceChangeType, area classdb.EditorVCSInterfaceTreeArea) gd.Dictionary {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Dictionary(Expert(self).CreateStatusFile(gc, gc.String(file_path), change_type, area))
}
func (self Simple) AddDiffHunksIntoDiffFile(diff_file gd.Dictionary, diff_hunks gd.ArrayOf[gd.Dictionary]) gd.Dictionary {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Dictionary(Expert(self).AddDiffHunksIntoDiffFile(gc, diff_file, diff_hunks))
}
func (self Simple) AddLineDiffsIntoDiffHunk(diff_hunk gd.Dictionary, line_diffs gd.ArrayOf[gd.Dictionary]) gd.Dictionary {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Dictionary(Expert(self).AddLineDiffsIntoDiffHunk(gc, diff_hunk, line_diffs))
}
func (self Simple) PopupError(msg string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).PopupError(gc.String(msg))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.EditorVCSInterface
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Initializes the VCS plugin when called from the editor. Returns whether or not the plugin was successfully initialized. A VCS project is initialized at [param project_path].
*/
func (class) _initialize(impl func(ptr unsafe.Pointer, project_path gd.String) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var project_path = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,0))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, project_path)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Set user credentials in the underlying VCS. [param username] and [param password] are used only during HTTPS authentication unless not already mentioned in the remote URL. [param ssh_public_key_path], [param ssh_private_key_path], and [param ssh_passphrase] are only used during SSH authentication.
*/
func (class) _set_credentials(impl func(ptr unsafe.Pointer, username gd.String, password gd.String, ssh_public_key_path gd.String, ssh_private_key_path gd.String, ssh_passphrase gd.String) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var username = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,0))
		var password = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,1))
		var ssh_public_key_path = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,2))
		var ssh_private_key_path = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,3))
		var ssh_passphrase = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,4))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, username, password, ssh_public_key_path, ssh_private_key_path, ssh_passphrase)
		ctx.End()
	}
}

/*
Returns an [Array] of [Dictionary] items (see [method create_status_file]), each containing the status data of every modified file in the project folder.
*/
func (class) _get_modified_files_data(impl func(ptr unsafe.Pointer) gd.ArrayOf[gd.Dictionary], api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(ret.Array()))
		ctx.End()
	}
}

/*
Stages the file present at [param file_path] to the staged area.
*/
func (class) _stage_file(impl func(ptr unsafe.Pointer, file_path gd.String) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var file_path = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,0))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, file_path)
		ctx.End()
	}
}

/*
Unstages the file present at [param file_path] from the staged area to the unstaged area.
*/
func (class) _unstage_file(impl func(ptr unsafe.Pointer, file_path gd.String) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var file_path = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,0))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, file_path)
		ctx.End()
	}
}

/*
Discards the changes made in a file present at [param file_path].
*/
func (class) _discard_file(impl func(ptr unsafe.Pointer, file_path gd.String) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var file_path = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,0))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, file_path)
		ctx.End()
	}
}

/*
Commits the currently staged changes and applies the commit [param msg] to the resulting commit.
*/
func (class) _commit(impl func(ptr unsafe.Pointer, msg gd.String) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var msg = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,0))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, msg)
		ctx.End()
	}
}

/*
Returns an array of [Dictionary] items (see [method create_diff_file], [method create_diff_hunk], [method create_diff_line], [method add_line_diffs_into_diff_hunk] and [method add_diff_hunks_into_diff_file]), each containing information about a diff. If [param identifier] is a file path, returns a file diff, and if it is a commit identifier, then returns a commit diff.
*/
func (class) _get_diff(impl func(ptr unsafe.Pointer, identifier gd.String, area gd.Int) gd.ArrayOf[gd.Dictionary], api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var identifier = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,0))
		var area = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, identifier, area)
		gd.UnsafeSet(p_back, mmm.End(ret.Array()))
		ctx.End()
	}
}

/*
Shuts down VCS plugin instance. Called when the user either closes the editor or shuts down the VCS plugin through the editor UI.
*/
func (class) _shut_down(impl func(ptr unsafe.Pointer) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Returns the name of the underlying VCS provider.
*/
func (class) _get_vcs_name(impl func(ptr unsafe.Pointer) gd.String, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
Returns an [Array] of [Dictionary] items (see [method create_commit]), each containing the data for a past commit.
*/
func (class) _get_previous_commits(impl func(ptr unsafe.Pointer, max_commits gd.Int) gd.ArrayOf[gd.Dictionary], api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var max_commits = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, max_commits)
		gd.UnsafeSet(p_back, mmm.End(ret.Array()))
		ctx.End()
	}
}

/*
Gets an instance of an [Array] of [String]s containing available branch names in the VCS.
*/
func (class) _get_branch_list(impl func(ptr unsafe.Pointer) gd.ArrayOf[gd.String], api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(ret.Array()))
		ctx.End()
	}
}

/*
Returns an [Array] of [String]s, each containing the name of a remote configured in the VCS.
*/
func (class) _get_remotes(impl func(ptr unsafe.Pointer) gd.ArrayOf[gd.String], api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(ret.Array()))
		ctx.End()
	}
}

/*
Creates a new branch named [param branch_name] in the VCS.
*/
func (class) _create_branch(impl func(ptr unsafe.Pointer, branch_name gd.String) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var branch_name = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,0))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, branch_name)
		ctx.End()
	}
}

/*
Remove a branch from the local VCS.
*/
func (class) _remove_branch(impl func(ptr unsafe.Pointer, branch_name gd.String) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var branch_name = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,0))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, branch_name)
		ctx.End()
	}
}

/*
Creates a new remote destination with name [param remote_name] and points it to [param remote_url]. This can be an HTTPS remote or an SSH remote.
*/
func (class) _create_remote(impl func(ptr unsafe.Pointer, remote_name gd.String, remote_url gd.String) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var remote_name = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,0))
		var remote_url = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, remote_name, remote_url)
		ctx.End()
	}
}

/*
Remove a remote from the local VCS.
*/
func (class) _remove_remote(impl func(ptr unsafe.Pointer, remote_name gd.String) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var remote_name = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,0))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, remote_name)
		ctx.End()
	}
}

/*
Gets the current branch name defined in the VCS.
*/
func (class) _get_current_branch_name(impl func(ptr unsafe.Pointer) gd.String, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
Checks out a [param branch_name] in the VCS.
*/
func (class) _checkout_branch(impl func(ptr unsafe.Pointer, branch_name gd.String) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var branch_name = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,0))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, branch_name)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Pulls changes from the remote. This can give rise to merge conflicts.
*/
func (class) _pull(impl func(ptr unsafe.Pointer, remote gd.String) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var remote = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,0))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, remote)
		ctx.End()
	}
}

/*
Pushes changes to the [param remote]. If [param force] is [code]true[/code], a force push will override the change history already present on the remote.
*/
func (class) _push(impl func(ptr unsafe.Pointer, remote gd.String, force bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var remote = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,0))
		var force = gd.UnsafeGet[bool](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, remote, force)
		ctx.End()
	}
}

/*
Fetches new changes from the [param remote], but doesn't write changes to the current working directory. Equivalent to [code]git fetch[/code].
*/
func (class) _fetch(impl func(ptr unsafe.Pointer, remote gd.String) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var remote = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,0))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, remote)
		ctx.End()
	}
}

/*
Returns an [Array] of [Dictionary] items (see [method create_diff_hunk]), each containing a line diff between a file at [param file_path] and the [param text] which is passed in.
*/
func (class) _get_line_diff(impl func(ptr unsafe.Pointer, file_path gd.String, text gd.String) gd.ArrayOf[gd.Dictionary], api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var file_path = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,0))
		var text = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, file_path, text)
		gd.UnsafeSet(p_back, mmm.End(ret.Array()))
		ctx.End()
	}
}

/*
Helper function to create a [Dictionary] for storing a line diff. [param new_line_no] is the line number in the new file (can be [code]-1[/code] if the line is deleted). [param old_line_no] is the line number in the old file (can be [code]-1[/code] if the line is added). [param content] is the diff text. [param status] is a single character string which stores the line origin.
*/
//go:nosplit
func (self class) CreateDiffLine(ctx gd.Lifetime, new_line_no gd.Int, old_line_no gd.Int, content gd.String, status gd.String) gd.Dictionary {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, new_line_no)
	callframe.Arg(frame, old_line_no)
	callframe.Arg(frame, mmm.Get(content))
	callframe.Arg(frame, mmm.Get(status))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorVCSInterface.Bind_create_diff_line, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Dictionary](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Helper function to create a [Dictionary] for storing diff hunk data. [param old_start] is the starting line number in old file. [param new_start] is the starting line number in new file. [param old_lines] is the number of lines in the old file. [param new_lines] is the number of lines in the new file.
*/
//go:nosplit
func (self class) CreateDiffHunk(ctx gd.Lifetime, old_start gd.Int, new_start gd.Int, old_lines gd.Int, new_lines gd.Int) gd.Dictionary {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, old_start)
	callframe.Arg(frame, new_start)
	callframe.Arg(frame, old_lines)
	callframe.Arg(frame, new_lines)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorVCSInterface.Bind_create_diff_hunk, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Dictionary](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Helper function to create a [Dictionary] for storing old and new diff file paths.
*/
//go:nosplit
func (self class) CreateDiffFile(ctx gd.Lifetime, new_file gd.String, old_file gd.String) gd.Dictionary {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(new_file))
	callframe.Arg(frame, mmm.Get(old_file))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorVCSInterface.Bind_create_diff_file, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Dictionary](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Helper function to create a commit [Dictionary] item. [param msg] is the commit message of the commit. [param author] is a single human-readable string containing all the author's details, e.g. the email and name configured in the VCS. [param id] is the identifier of the commit, in whichever format your VCS may provide an identifier to commits. [param unix_timestamp] is the UTC Unix timestamp of when the commit was created. [param offset_minutes] is the timezone offset in minutes, recorded from the system timezone where the commit was created.
*/
//go:nosplit
func (self class) CreateCommit(ctx gd.Lifetime, msg gd.String, author gd.String, id gd.String, unix_timestamp gd.Int, offset_minutes gd.Int) gd.Dictionary {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(msg))
	callframe.Arg(frame, mmm.Get(author))
	callframe.Arg(frame, mmm.Get(id))
	callframe.Arg(frame, unix_timestamp)
	callframe.Arg(frame, offset_minutes)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorVCSInterface.Bind_create_commit, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Dictionary](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Helper function to create a [Dictionary] used by editor to read the status of a file.
*/
//go:nosplit
func (self class) CreateStatusFile(ctx gd.Lifetime, file_path gd.String, change_type classdb.EditorVCSInterfaceChangeType, area classdb.EditorVCSInterfaceTreeArea) gd.Dictionary {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(file_path))
	callframe.Arg(frame, change_type)
	callframe.Arg(frame, area)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorVCSInterface.Bind_create_status_file, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Dictionary](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Helper function to add an array of [param diff_hunks] into a [param diff_file].
*/
//go:nosplit
func (self class) AddDiffHunksIntoDiffFile(ctx gd.Lifetime, diff_file gd.Dictionary, diff_hunks gd.ArrayOf[gd.Dictionary]) gd.Dictionary {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(diff_file))
	callframe.Arg(frame, mmm.Get(diff_hunks))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorVCSInterface.Bind_add_diff_hunks_into_diff_file, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Dictionary](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Helper function to add an array of [param line_diffs] into a [param diff_hunk].
*/
//go:nosplit
func (self class) AddLineDiffsIntoDiffHunk(ctx gd.Lifetime, diff_hunk gd.Dictionary, line_diffs gd.ArrayOf[gd.Dictionary]) gd.Dictionary {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(diff_hunk))
	callframe.Arg(frame, mmm.Get(line_diffs))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorVCSInterface.Bind_add_line_diffs_into_diff_hunk, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Dictionary](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Pops up an error message in the editor which is shown as coming from the underlying VCS. Use this to show VCS specific error messages.
*/
//go:nosplit
func (self class) PopupError(msg gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(msg))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorVCSInterface.Bind_popup_error, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) AsEditorVCSInterface() Expert { return self[0].AsEditorVCSInterface() }


//go:nosplit
func (self Simple) AsEditorVCSInterface() Simple { return self[0].AsEditorVCSInterface() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_initialize": return reflect.ValueOf(self._initialize);
	case "_set_credentials": return reflect.ValueOf(self._set_credentials);
	case "_get_modified_files_data": return reflect.ValueOf(self._get_modified_files_data);
	case "_stage_file": return reflect.ValueOf(self._stage_file);
	case "_unstage_file": return reflect.ValueOf(self._unstage_file);
	case "_discard_file": return reflect.ValueOf(self._discard_file);
	case "_commit": return reflect.ValueOf(self._commit);
	case "_get_diff": return reflect.ValueOf(self._get_diff);
	case "_shut_down": return reflect.ValueOf(self._shut_down);
	case "_get_vcs_name": return reflect.ValueOf(self._get_vcs_name);
	case "_get_previous_commits": return reflect.ValueOf(self._get_previous_commits);
	case "_get_branch_list": return reflect.ValueOf(self._get_branch_list);
	case "_get_remotes": return reflect.ValueOf(self._get_remotes);
	case "_create_branch": return reflect.ValueOf(self._create_branch);
	case "_remove_branch": return reflect.ValueOf(self._remove_branch);
	case "_create_remote": return reflect.ValueOf(self._create_remote);
	case "_remove_remote": return reflect.ValueOf(self._remove_remote);
	case "_get_current_branch_name": return reflect.ValueOf(self._get_current_branch_name);
	case "_checkout_branch": return reflect.ValueOf(self._checkout_branch);
	case "_pull": return reflect.ValueOf(self._pull);
	case "_push": return reflect.ValueOf(self._push);
	case "_fetch": return reflect.ValueOf(self._fetch);
	case "_get_line_diff": return reflect.ValueOf(self._get_line_diff);
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	case "_initialize": return reflect.ValueOf(self._initialize);
	case "_set_credentials": return reflect.ValueOf(self._set_credentials);
	case "_get_modified_files_data": return reflect.ValueOf(self._get_modified_files_data);
	case "_stage_file": return reflect.ValueOf(self._stage_file);
	case "_unstage_file": return reflect.ValueOf(self._unstage_file);
	case "_discard_file": return reflect.ValueOf(self._discard_file);
	case "_commit": return reflect.ValueOf(self._commit);
	case "_get_diff": return reflect.ValueOf(self._get_diff);
	case "_shut_down": return reflect.ValueOf(self._shut_down);
	case "_get_vcs_name": return reflect.ValueOf(self._get_vcs_name);
	case "_get_previous_commits": return reflect.ValueOf(self._get_previous_commits);
	case "_get_branch_list": return reflect.ValueOf(self._get_branch_list);
	case "_get_remotes": return reflect.ValueOf(self._get_remotes);
	case "_create_branch": return reflect.ValueOf(self._create_branch);
	case "_remove_branch": return reflect.ValueOf(self._remove_branch);
	case "_create_remote": return reflect.ValueOf(self._create_remote);
	case "_remove_remote": return reflect.ValueOf(self._remove_remote);
	case "_get_current_branch_name": return reflect.ValueOf(self._get_current_branch_name);
	case "_checkout_branch": return reflect.ValueOf(self._checkout_branch);
	case "_pull": return reflect.ValueOf(self._pull);
	case "_push": return reflect.ValueOf(self._push);
	case "_fetch": return reflect.ValueOf(self._fetch);
	case "_get_line_diff": return reflect.ValueOf(self._get_line_diff);
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("EditorVCSInterface", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type ChangeType = classdb.EditorVCSInterfaceChangeType

const (
/*A new file has been added.*/
	ChangeTypeNew ChangeType = 0
/*An earlier added file has been modified.*/
	ChangeTypeModified ChangeType = 1
/*An earlier added file has been renamed.*/
	ChangeTypeRenamed ChangeType = 2
/*An earlier added file has been deleted.*/
	ChangeTypeDeleted ChangeType = 3
/*An earlier added file has been typechanged.*/
	ChangeTypeTypechange ChangeType = 4
/*A file is left unmerged.*/
	ChangeTypeUnmerged ChangeType = 5
)
type TreeArea = classdb.EditorVCSInterfaceTreeArea

const (
/*A commit is encountered from the commit area.*/
	TreeAreaCommit TreeArea = 0
/*A file is encountered from the staged area.*/
	TreeAreaStaged TreeArea = 1
/*A file is encountered from the unstaged area.*/
	TreeAreaUnstaged TreeArea = 2
)
