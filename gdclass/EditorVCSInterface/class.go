package EditorVCSInterface

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

/*
Defines the API that the editor uses to extract information from the underlying VCS. The implementation of this API is included in VCS plugins, which are GDExtension plugins that inherit [EditorVCSInterface] and are attached (on demand) to the singleton instance of [EditorVCSInterface]. Instead of performing the task themselves, all the virtual functions listed below are calling the internally overridden functions in the VCS plugins to provide a plug-n-play experience. A custom VCS plugin is supposed to inherit from [EditorVCSInterface] and override each of these virtual functions.

	// EditorVCSInterface methods that can be overridden by a [Class] that extends it.
	type EditorVCSInterface interface {
		//Initializes the VCS plugin when called from the editor. Returns whether or not the plugin was successfully initialized. A VCS project is initialized at [param project_path].
		Initialize(project_path string) bool
		//Set user credentials in the underlying VCS. [param username] and [param password] are used only during HTTPS authentication unless not already mentioned in the remote URL. [param ssh_public_key_path], [param ssh_private_key_path], and [param ssh_passphrase] are only used during SSH authentication.
		SetCredentials(username string, password string, ssh_public_key_path string, ssh_private_key_path string, ssh_passphrase string)
		//Returns an [Array] of [Dictionary] items (see [method create_status_file]), each containing the status data of every modified file in the project folder.
		GetModifiedFilesData() gd.Array
		//Stages the file present at [param file_path] to the staged area.
		StageFile(file_path string)
		//Unstages the file present at [param file_path] from the staged area to the unstaged area.
		UnstageFile(file_path string)
		//Discards the changes made in a file present at [param file_path].
		DiscardFile(file_path string)
		//Commits the currently staged changes and applies the commit [param msg] to the resulting commit.
		Commit(msg string)
		//Returns an array of [Dictionary] items (see [method create_diff_file], [method create_diff_hunk], [method create_diff_line], [method add_line_diffs_into_diff_hunk] and [method add_diff_hunks_into_diff_file]), each containing information about a diff. If [param identifier] is a file path, returns a file diff, and if it is a commit identifier, then returns a commit diff.
		GetDiff(identifier string, area int) gd.Array
		//Shuts down VCS plugin instance. Called when the user either closes the editor or shuts down the VCS plugin through the editor UI.
		ShutDown() bool
		//Returns the name of the underlying VCS provider.
		GetVcsName() string
		//Returns an [Array] of [Dictionary] items (see [method create_commit]), each containing the data for a past commit.
		GetPreviousCommits(max_commits int) gd.Array
		//Gets an instance of an [Array] of [String]s containing available branch names in the VCS.
		GetBranchList() gd.Array
		//Returns an [Array] of [String]s, each containing the name of a remote configured in the VCS.
		GetRemotes() gd.Array
		//Creates a new branch named [param branch_name] in the VCS.
		CreateBranch(branch_name string)
		//Remove a branch from the local VCS.
		RemoveBranch(branch_name string)
		//Creates a new remote destination with name [param remote_name] and points it to [param remote_url]. This can be an HTTPS remote or an SSH remote.
		CreateRemote(remote_name string, remote_url string)
		//Remove a remote from the local VCS.
		RemoveRemote(remote_name string)
		//Gets the current branch name defined in the VCS.
		GetCurrentBranchName() string
		//Checks out a [param branch_name] in the VCS.
		CheckoutBranch(branch_name string) bool
		//Pulls changes from the remote. This can give rise to merge conflicts.
		Pull(remote string)
		//Pushes changes to the [param remote]. If [param force] is [code]true[/code], a force push will override the change history already present on the remote.
		Push(remote string, force bool)
		//Fetches new changes from the [param remote], but doesn't write changes to the current working directory. Equivalent to [code]git fetch[/code].
		Fetch(remote string)
		//Returns an [Array] of [Dictionary] items (see [method create_diff_hunk]), each containing a line diff between a file at [param file_path] and the [param text] which is passed in.
		GetLineDiff(file_path string, text string) gd.Array
	}
*/
type Instance [1]classdb.EditorVCSInterface

/*
Initializes the VCS plugin when called from the editor. Returns whether or not the plugin was successfully initialized. A VCS project is initialized at [param project_path].
*/
func (Instance) _initialize(impl func(ptr unsafe.Pointer, project_path string) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var project_path = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 0))
		defer pointers.End(project_path)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, project_path.String())
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Set user credentials in the underlying VCS. [param username] and [param password] are used only during HTTPS authentication unless not already mentioned in the remote URL. [param ssh_public_key_path], [param ssh_private_key_path], and [param ssh_passphrase] are only used during SSH authentication.
*/
func (Instance) _set_credentials(impl func(ptr unsafe.Pointer, username string, password string, ssh_public_key_path string, ssh_private_key_path string, ssh_passphrase string)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var username = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 0))
		defer pointers.End(username)
		var password = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 1))
		defer pointers.End(password)
		var ssh_public_key_path = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 2))
		defer pointers.End(ssh_public_key_path)
		var ssh_private_key_path = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 3))
		defer pointers.End(ssh_private_key_path)
		var ssh_passphrase = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 4))
		defer pointers.End(ssh_passphrase)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, username.String(), password.String(), ssh_public_key_path.String(), ssh_private_key_path.String(), ssh_passphrase.String())
	}
}

/*
Returns an [Array] of [Dictionary] items (see [method create_status_file]), each containing the status data of every modified file in the project folder.
*/
func (Instance) _get_modified_files_data(impl func(ptr unsafe.Pointer) gd.Array) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Stages the file present at [param file_path] to the staged area.
*/
func (Instance) _stage_file(impl func(ptr unsafe.Pointer, file_path string)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var file_path = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 0))
		defer pointers.End(file_path)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, file_path.String())
	}
}

/*
Unstages the file present at [param file_path] from the staged area to the unstaged area.
*/
func (Instance) _unstage_file(impl func(ptr unsafe.Pointer, file_path string)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var file_path = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 0))
		defer pointers.End(file_path)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, file_path.String())
	}
}

/*
Discards the changes made in a file present at [param file_path].
*/
func (Instance) _discard_file(impl func(ptr unsafe.Pointer, file_path string)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var file_path = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 0))
		defer pointers.End(file_path)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, file_path.String())
	}
}

/*
Commits the currently staged changes and applies the commit [param msg] to the resulting commit.
*/
func (Instance) _commit(impl func(ptr unsafe.Pointer, msg string)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var msg = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 0))
		defer pointers.End(msg)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, msg.String())
	}
}

/*
Returns an array of [Dictionary] items (see [method create_diff_file], [method create_diff_hunk], [method create_diff_line], [method add_line_diffs_into_diff_hunk] and [method add_diff_hunks_into_diff_file]), each containing information about a diff. If [param identifier] is a file path, returns a file diff, and if it is a commit identifier, then returns a commit diff.
*/
func (Instance) _get_diff(impl func(ptr unsafe.Pointer, identifier string, area int) gd.Array) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var identifier = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 0))
		defer pointers.End(identifier)
		var area = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, identifier.String(), int(area))
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Shuts down VCS plugin instance. Called when the user either closes the editor or shuts down the VCS plugin through the editor UI.
*/
func (Instance) _shut_down(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Returns the name of the underlying VCS provider.
*/
func (Instance) _get_vcs_name(impl func(ptr unsafe.Pointer) string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.NewString(ret))
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Returns an [Array] of [Dictionary] items (see [method create_commit]), each containing the data for a past commit.
*/
func (Instance) _get_previous_commits(impl func(ptr unsafe.Pointer, max_commits int) gd.Array) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var max_commits = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(max_commits))
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Gets an instance of an [Array] of [String]s containing available branch names in the VCS.
*/
func (Instance) _get_branch_list(impl func(ptr unsafe.Pointer) gd.Array) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Returns an [Array] of [String]s, each containing the name of a remote configured in the VCS.
*/
func (Instance) _get_remotes(impl func(ptr unsafe.Pointer) gd.Array) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Creates a new branch named [param branch_name] in the VCS.
*/
func (Instance) _create_branch(impl func(ptr unsafe.Pointer, branch_name string)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var branch_name = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 0))
		defer pointers.End(branch_name)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, branch_name.String())
	}
}

/*
Remove a branch from the local VCS.
*/
func (Instance) _remove_branch(impl func(ptr unsafe.Pointer, branch_name string)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var branch_name = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 0))
		defer pointers.End(branch_name)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, branch_name.String())
	}
}

/*
Creates a new remote destination with name [param remote_name] and points it to [param remote_url]. This can be an HTTPS remote or an SSH remote.
*/
func (Instance) _create_remote(impl func(ptr unsafe.Pointer, remote_name string, remote_url string)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var remote_name = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 0))
		defer pointers.End(remote_name)
		var remote_url = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 1))
		defer pointers.End(remote_url)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, remote_name.String(), remote_url.String())
	}
}

/*
Remove a remote from the local VCS.
*/
func (Instance) _remove_remote(impl func(ptr unsafe.Pointer, remote_name string)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var remote_name = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 0))
		defer pointers.End(remote_name)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, remote_name.String())
	}
}

/*
Gets the current branch name defined in the VCS.
*/
func (Instance) _get_current_branch_name(impl func(ptr unsafe.Pointer) string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.NewString(ret))
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Checks out a [param branch_name] in the VCS.
*/
func (Instance) _checkout_branch(impl func(ptr unsafe.Pointer, branch_name string) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var branch_name = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 0))
		defer pointers.End(branch_name)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, branch_name.String())
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Pulls changes from the remote. This can give rise to merge conflicts.
*/
func (Instance) _pull(impl func(ptr unsafe.Pointer, remote string)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var remote = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 0))
		defer pointers.End(remote)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, remote.String())
	}
}

/*
Pushes changes to the [param remote]. If [param force] is [code]true[/code], a force push will override the change history already present on the remote.
*/
func (Instance) _push(impl func(ptr unsafe.Pointer, remote string, force bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var remote = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 0))
		defer pointers.End(remote)
		var force = gd.UnsafeGet[bool](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, remote.String(), force)
	}
}

/*
Fetches new changes from the [param remote], but doesn't write changes to the current working directory. Equivalent to [code]git fetch[/code].
*/
func (Instance) _fetch(impl func(ptr unsafe.Pointer, remote string)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var remote = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 0))
		defer pointers.End(remote)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, remote.String())
	}
}

/*
Returns an [Array] of [Dictionary] items (see [method create_diff_hunk]), each containing a line diff between a file at [param file_path] and the [param text] which is passed in.
*/
func (Instance) _get_line_diff(impl func(ptr unsafe.Pointer, file_path string, text string) gd.Array) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var file_path = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 0))
		defer pointers.End(file_path)
		var text = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 1))
		defer pointers.End(text)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, file_path.String(), text.String())
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Helper function to create a [Dictionary] for storing a line diff. [param new_line_no] is the line number in the new file (can be [code]-1[/code] if the line is deleted). [param old_line_no] is the line number in the old file (can be [code]-1[/code] if the line is added). [param content] is the diff text. [param status] is a single character string which stores the line origin.
*/
func (self Instance) CreateDiffLine(new_line_no int, old_line_no int, content string, status string) gd.Dictionary {
	return gd.Dictionary(class(self).CreateDiffLine(gd.Int(new_line_no), gd.Int(old_line_no), gd.NewString(content), gd.NewString(status)))
}

/*
Helper function to create a [Dictionary] for storing diff hunk data. [param old_start] is the starting line number in old file. [param new_start] is the starting line number in new file. [param old_lines] is the number of lines in the old file. [param new_lines] is the number of lines in the new file.
*/
func (self Instance) CreateDiffHunk(old_start int, new_start int, old_lines int, new_lines int) gd.Dictionary {
	return gd.Dictionary(class(self).CreateDiffHunk(gd.Int(old_start), gd.Int(new_start), gd.Int(old_lines), gd.Int(new_lines)))
}

/*
Helper function to create a [Dictionary] for storing old and new diff file paths.
*/
func (self Instance) CreateDiffFile(new_file string, old_file string) gd.Dictionary {
	return gd.Dictionary(class(self).CreateDiffFile(gd.NewString(new_file), gd.NewString(old_file)))
}

/*
Helper function to create a commit [Dictionary] item. [param msg] is the commit message of the commit. [param author] is a single human-readable string containing all the author's details, e.g. the email and name configured in the VCS. [param id] is the identifier of the commit, in whichever format your VCS may provide an identifier to commits. [param unix_timestamp] is the UTC Unix timestamp of when the commit was created. [param offset_minutes] is the timezone offset in minutes, recorded from the system timezone where the commit was created.
*/
func (self Instance) CreateCommit(msg string, author string, id string, unix_timestamp int, offset_minutes int) gd.Dictionary {
	return gd.Dictionary(class(self).CreateCommit(gd.NewString(msg), gd.NewString(author), gd.NewString(id), gd.Int(unix_timestamp), gd.Int(offset_minutes)))
}

/*
Helper function to create a [Dictionary] used by editor to read the status of a file.
*/
func (self Instance) CreateStatusFile(file_path string, change_type classdb.EditorVCSInterfaceChangeType, area classdb.EditorVCSInterfaceTreeArea) gd.Dictionary {
	return gd.Dictionary(class(self).CreateStatusFile(gd.NewString(file_path), change_type, area))
}

/*
Helper function to add an array of [param diff_hunks] into a [param diff_file].
*/
func (self Instance) AddDiffHunksIntoDiffFile(diff_file gd.Dictionary, diff_hunks gd.Array) gd.Dictionary {
	return gd.Dictionary(class(self).AddDiffHunksIntoDiffFile(diff_file, diff_hunks))
}

/*
Helper function to add an array of [param line_diffs] into a [param diff_hunk].
*/
func (self Instance) AddLineDiffsIntoDiffHunk(diff_hunk gd.Dictionary, line_diffs gd.Array) gd.Dictionary {
	return gd.Dictionary(class(self).AddLineDiffsIntoDiffHunk(diff_hunk, line_diffs))
}

/*
Pops up an error message in the editor which is shown as coming from the underlying VCS. Use this to show VCS specific error messages.
*/
func (self Instance) PopupError(msg string) {
	class(self).PopupError(gd.NewString(msg))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.EditorVCSInterface

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("EditorVCSInterface"))
	return Instance{classdb.EditorVCSInterface(object)}
}

/*
Initializes the VCS plugin when called from the editor. Returns whether or not the plugin was successfully initialized. A VCS project is initialized at [param project_path].
*/
func (class) _initialize(impl func(ptr unsafe.Pointer, project_path gd.String) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var project_path = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 0))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, project_path)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Set user credentials in the underlying VCS. [param username] and [param password] are used only during HTTPS authentication unless not already mentioned in the remote URL. [param ssh_public_key_path], [param ssh_private_key_path], and [param ssh_passphrase] are only used during SSH authentication.
*/
func (class) _set_credentials(impl func(ptr unsafe.Pointer, username gd.String, password gd.String, ssh_public_key_path gd.String, ssh_private_key_path gd.String, ssh_passphrase gd.String)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var username = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 0))
		var password = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 1))
		var ssh_public_key_path = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 2))
		var ssh_private_key_path = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 3))
		var ssh_passphrase = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 4))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, username, password, ssh_public_key_path, ssh_private_key_path, ssh_passphrase)
	}
}

/*
Returns an [Array] of [Dictionary] items (see [method create_status_file]), each containing the status data of every modified file in the project folder.
*/
func (class) _get_modified_files_data(impl func(ptr unsafe.Pointer) gd.Array) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Stages the file present at [param file_path] to the staged area.
*/
func (class) _stage_file(impl func(ptr unsafe.Pointer, file_path gd.String)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var file_path = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 0))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, file_path)
	}
}

/*
Unstages the file present at [param file_path] from the staged area to the unstaged area.
*/
func (class) _unstage_file(impl func(ptr unsafe.Pointer, file_path gd.String)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var file_path = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 0))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, file_path)
	}
}

/*
Discards the changes made in a file present at [param file_path].
*/
func (class) _discard_file(impl func(ptr unsafe.Pointer, file_path gd.String)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var file_path = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 0))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, file_path)
	}
}

/*
Commits the currently staged changes and applies the commit [param msg] to the resulting commit.
*/
func (class) _commit(impl func(ptr unsafe.Pointer, msg gd.String)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var msg = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 0))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, msg)
	}
}

/*
Returns an array of [Dictionary] items (see [method create_diff_file], [method create_diff_hunk], [method create_diff_line], [method add_line_diffs_into_diff_hunk] and [method add_diff_hunks_into_diff_file]), each containing information about a diff. If [param identifier] is a file path, returns a file diff, and if it is a commit identifier, then returns a commit diff.
*/
func (class) _get_diff(impl func(ptr unsafe.Pointer, identifier gd.String, area gd.Int) gd.Array) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var identifier = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 0))
		var area = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, identifier, area)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Shuts down VCS plugin instance. Called when the user either closes the editor or shuts down the VCS plugin through the editor UI.
*/
func (class) _shut_down(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Returns the name of the underlying VCS provider.
*/
func (class) _get_vcs_name(impl func(ptr unsafe.Pointer) gd.String) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Returns an [Array] of [Dictionary] items (see [method create_commit]), each containing the data for a past commit.
*/
func (class) _get_previous_commits(impl func(ptr unsafe.Pointer, max_commits gd.Int) gd.Array) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var max_commits = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, max_commits)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Gets an instance of an [Array] of [String]s containing available branch names in the VCS.
*/
func (class) _get_branch_list(impl func(ptr unsafe.Pointer) gd.Array) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Returns an [Array] of [String]s, each containing the name of a remote configured in the VCS.
*/
func (class) _get_remotes(impl func(ptr unsafe.Pointer) gd.Array) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Creates a new branch named [param branch_name] in the VCS.
*/
func (class) _create_branch(impl func(ptr unsafe.Pointer, branch_name gd.String)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var branch_name = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 0))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, branch_name)
	}
}

/*
Remove a branch from the local VCS.
*/
func (class) _remove_branch(impl func(ptr unsafe.Pointer, branch_name gd.String)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var branch_name = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 0))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, branch_name)
	}
}

/*
Creates a new remote destination with name [param remote_name] and points it to [param remote_url]. This can be an HTTPS remote or an SSH remote.
*/
func (class) _create_remote(impl func(ptr unsafe.Pointer, remote_name gd.String, remote_url gd.String)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var remote_name = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 0))
		var remote_url = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 1))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, remote_name, remote_url)
	}
}

/*
Remove a remote from the local VCS.
*/
func (class) _remove_remote(impl func(ptr unsafe.Pointer, remote_name gd.String)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var remote_name = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 0))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, remote_name)
	}
}

/*
Gets the current branch name defined in the VCS.
*/
func (class) _get_current_branch_name(impl func(ptr unsafe.Pointer) gd.String) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Checks out a [param branch_name] in the VCS.
*/
func (class) _checkout_branch(impl func(ptr unsafe.Pointer, branch_name gd.String) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var branch_name = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 0))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, branch_name)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Pulls changes from the remote. This can give rise to merge conflicts.
*/
func (class) _pull(impl func(ptr unsafe.Pointer, remote gd.String)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var remote = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 0))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, remote)
	}
}

/*
Pushes changes to the [param remote]. If [param force] is [code]true[/code], a force push will override the change history already present on the remote.
*/
func (class) _push(impl func(ptr unsafe.Pointer, remote gd.String, force bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var remote = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 0))
		var force = gd.UnsafeGet[bool](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, remote, force)
	}
}

/*
Fetches new changes from the [param remote], but doesn't write changes to the current working directory. Equivalent to [code]git fetch[/code].
*/
func (class) _fetch(impl func(ptr unsafe.Pointer, remote gd.String)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var remote = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 0))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, remote)
	}
}

/*
Returns an [Array] of [Dictionary] items (see [method create_diff_hunk]), each containing a line diff between a file at [param file_path] and the [param text] which is passed in.
*/
func (class) _get_line_diff(impl func(ptr unsafe.Pointer, file_path gd.String, text gd.String) gd.Array) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var file_path = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 0))
		var text = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 1))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, file_path, text)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Helper function to create a [Dictionary] for storing a line diff. [param new_line_no] is the line number in the new file (can be [code]-1[/code] if the line is deleted). [param old_line_no] is the line number in the old file (can be [code]-1[/code] if the line is added). [param content] is the diff text. [param status] is a single character string which stores the line origin.
*/
//go:nosplit
func (self class) CreateDiffLine(new_line_no gd.Int, old_line_no gd.Int, content gd.String, status gd.String) gd.Dictionary {
	var frame = callframe.New()
	callframe.Arg(frame, new_line_no)
	callframe.Arg(frame, old_line_no)
	callframe.Arg(frame, pointers.Get(content))
	callframe.Arg(frame, pointers.Get(status))
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorVCSInterface.Bind_create_diff_line, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Dictionary](r_ret.Get())
	frame.Free()
	return ret
}

/*
Helper function to create a [Dictionary] for storing diff hunk data. [param old_start] is the starting line number in old file. [param new_start] is the starting line number in new file. [param old_lines] is the number of lines in the old file. [param new_lines] is the number of lines in the new file.
*/
//go:nosplit
func (self class) CreateDiffHunk(old_start gd.Int, new_start gd.Int, old_lines gd.Int, new_lines gd.Int) gd.Dictionary {
	var frame = callframe.New()
	callframe.Arg(frame, old_start)
	callframe.Arg(frame, new_start)
	callframe.Arg(frame, old_lines)
	callframe.Arg(frame, new_lines)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorVCSInterface.Bind_create_diff_hunk, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Dictionary](r_ret.Get())
	frame.Free()
	return ret
}

/*
Helper function to create a [Dictionary] for storing old and new diff file paths.
*/
//go:nosplit
func (self class) CreateDiffFile(new_file gd.String, old_file gd.String) gd.Dictionary {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(new_file))
	callframe.Arg(frame, pointers.Get(old_file))
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorVCSInterface.Bind_create_diff_file, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Dictionary](r_ret.Get())
	frame.Free()
	return ret
}

/*
Helper function to create a commit [Dictionary] item. [param msg] is the commit message of the commit. [param author] is a single human-readable string containing all the author's details, e.g. the email and name configured in the VCS. [param id] is the identifier of the commit, in whichever format your VCS may provide an identifier to commits. [param unix_timestamp] is the UTC Unix timestamp of when the commit was created. [param offset_minutes] is the timezone offset in minutes, recorded from the system timezone where the commit was created.
*/
//go:nosplit
func (self class) CreateCommit(msg gd.String, author gd.String, id gd.String, unix_timestamp gd.Int, offset_minutes gd.Int) gd.Dictionary {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(msg))
	callframe.Arg(frame, pointers.Get(author))
	callframe.Arg(frame, pointers.Get(id))
	callframe.Arg(frame, unix_timestamp)
	callframe.Arg(frame, offset_minutes)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorVCSInterface.Bind_create_commit, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Dictionary](r_ret.Get())
	frame.Free()
	return ret
}

/*
Helper function to create a [Dictionary] used by editor to read the status of a file.
*/
//go:nosplit
func (self class) CreateStatusFile(file_path gd.String, change_type classdb.EditorVCSInterfaceChangeType, area classdb.EditorVCSInterfaceTreeArea) gd.Dictionary {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(file_path))
	callframe.Arg(frame, change_type)
	callframe.Arg(frame, area)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorVCSInterface.Bind_create_status_file, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Dictionary](r_ret.Get())
	frame.Free()
	return ret
}

/*
Helper function to add an array of [param diff_hunks] into a [param diff_file].
*/
//go:nosplit
func (self class) AddDiffHunksIntoDiffFile(diff_file gd.Dictionary, diff_hunks gd.Array) gd.Dictionary {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(diff_file))
	callframe.Arg(frame, pointers.Get(diff_hunks))
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorVCSInterface.Bind_add_diff_hunks_into_diff_file, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Dictionary](r_ret.Get())
	frame.Free()
	return ret
}

/*
Helper function to add an array of [param line_diffs] into a [param diff_hunk].
*/
//go:nosplit
func (self class) AddLineDiffsIntoDiffHunk(diff_hunk gd.Dictionary, line_diffs gd.Array) gd.Dictionary {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(diff_hunk))
	callframe.Arg(frame, pointers.Get(line_diffs))
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorVCSInterface.Bind_add_line_diffs_into_diff_hunk, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Dictionary](r_ret.Get())
	frame.Free()
	return ret
}

/*
Pops up an error message in the editor which is shown as coming from the underlying VCS. Use this to show VCS specific error messages.
*/
//go:nosplit
func (self class) PopupError(msg gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(msg))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorVCSInterface.Bind_popup_error, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsEditorVCSInterface() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsEditorVCSInterface() Instance { return *((*Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_initialize":
		return reflect.ValueOf(self._initialize)
	case "_set_credentials":
		return reflect.ValueOf(self._set_credentials)
	case "_get_modified_files_data":
		return reflect.ValueOf(self._get_modified_files_data)
	case "_stage_file":
		return reflect.ValueOf(self._stage_file)
	case "_unstage_file":
		return reflect.ValueOf(self._unstage_file)
	case "_discard_file":
		return reflect.ValueOf(self._discard_file)
	case "_commit":
		return reflect.ValueOf(self._commit)
	case "_get_diff":
		return reflect.ValueOf(self._get_diff)
	case "_shut_down":
		return reflect.ValueOf(self._shut_down)
	case "_get_vcs_name":
		return reflect.ValueOf(self._get_vcs_name)
	case "_get_previous_commits":
		return reflect.ValueOf(self._get_previous_commits)
	case "_get_branch_list":
		return reflect.ValueOf(self._get_branch_list)
	case "_get_remotes":
		return reflect.ValueOf(self._get_remotes)
	case "_create_branch":
		return reflect.ValueOf(self._create_branch)
	case "_remove_branch":
		return reflect.ValueOf(self._remove_branch)
	case "_create_remote":
		return reflect.ValueOf(self._create_remote)
	case "_remove_remote":
		return reflect.ValueOf(self._remove_remote)
	case "_get_current_branch_name":
		return reflect.ValueOf(self._get_current_branch_name)
	case "_checkout_branch":
		return reflect.ValueOf(self._checkout_branch)
	case "_pull":
		return reflect.ValueOf(self._pull)
	case "_push":
		return reflect.ValueOf(self._push)
	case "_fetch":
		return reflect.ValueOf(self._fetch)
	case "_get_line_diff":
		return reflect.ValueOf(self._get_line_diff)
	default:
		return gd.VirtualByName(self.AsObject(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_initialize":
		return reflect.ValueOf(self._initialize)
	case "_set_credentials":
		return reflect.ValueOf(self._set_credentials)
	case "_get_modified_files_data":
		return reflect.ValueOf(self._get_modified_files_data)
	case "_stage_file":
		return reflect.ValueOf(self._stage_file)
	case "_unstage_file":
		return reflect.ValueOf(self._unstage_file)
	case "_discard_file":
		return reflect.ValueOf(self._discard_file)
	case "_commit":
		return reflect.ValueOf(self._commit)
	case "_get_diff":
		return reflect.ValueOf(self._get_diff)
	case "_shut_down":
		return reflect.ValueOf(self._shut_down)
	case "_get_vcs_name":
		return reflect.ValueOf(self._get_vcs_name)
	case "_get_previous_commits":
		return reflect.ValueOf(self._get_previous_commits)
	case "_get_branch_list":
		return reflect.ValueOf(self._get_branch_list)
	case "_get_remotes":
		return reflect.ValueOf(self._get_remotes)
	case "_create_branch":
		return reflect.ValueOf(self._create_branch)
	case "_remove_branch":
		return reflect.ValueOf(self._remove_branch)
	case "_create_remote":
		return reflect.ValueOf(self._create_remote)
	case "_remove_remote":
		return reflect.ValueOf(self._remove_remote)
	case "_get_current_branch_name":
		return reflect.ValueOf(self._get_current_branch_name)
	case "_checkout_branch":
		return reflect.ValueOf(self._checkout_branch)
	case "_pull":
		return reflect.ValueOf(self._pull)
	case "_push":
		return reflect.ValueOf(self._push)
	case "_fetch":
		return reflect.ValueOf(self._fetch)
	case "_get_line_diff":
		return reflect.ValueOf(self._get_line_diff)
	default:
		return gd.VirtualByName(self.AsObject(), name)
	}
}
func init() {
	classdb.Register("EditorVCSInterface", func(ptr gd.Object) any { return classdb.EditorVCSInterface(ptr) })
}

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
