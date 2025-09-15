package gdjson

import (
	"net/netip"
	"reflect"

	"graphics.gd/variant/Callable"
	"graphics.gd/variant/Color"
	"graphics.gd/variant/Error"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Object"
	"graphics.gd/variant/RID"
	"graphics.gd/variant/Rect2"
	"graphics.gd/variant/Transform3D"
	"graphics.gd/variant/Vector2"
	"graphics.gd/variant/Vector2i"
	"graphics.gd/variant/Vector3"
	"graphics.gd/variant/Vector3i"
)

type SignalInfo struct {
	Name        string                `gd:"name"`
	Flags       int                   `gd:"flags"`
	ID          int                   `gd:"id"`
	DefaultArgs []any                 `gd:"default_args"`
	Args        []Object.PropertyInfo `gd:"args"`
}

type CompletionInfo struct {
	Kind         any        `gd:"kind" type:"CodeCompletionKind"`
	DisplayText  string     `gd:"display_text"`
	InsertText   string     `gd:"insert_text"`
	FontColor    Color.RGBA `gd:"font_color"`
	Icon         string     `gd:"icon"`
	DefaultValue string     `gd:"default_value"`
}

type Completion struct {
	Kind         any        `gd:"kind" type:"CodeCompletionKind"`
	Display      string     `gd:"display"`
	InsertText   string     `gd:"insert_text"`
	FontColor    Color.RGBA `gd:"font_color"`
	Icon         string     `gd:"icon"`
	DefaultValue string     `gd:"default_value"`
	Location     string     `gd:"location"`
	Matches      []int32    `gd:"matches"`
	Force        bool       `gd:"force"`
	CallHint     string     `gd:"call_hint"`
	Result       Error.Code `gd:"result"`
}

type DiffLine map[any]any
type DiffHunk map[any]any
type DiffFile map[any]any
type Commit map[any]any
type StatusFile map[any]any

type VersionInfo struct {
	Major     int    `gd:"major"`
	Minor     int    `gd:"minor"`
	Patch     int    `gd:"patch"`
	Hex       int    `gd:"hex"`
	Status    string `gd:"status"`
	Build     string `gd:"build"`
	Hash      string `gd:"hash"`
	Timestamp int    `gd:"timestamp"`
	String    string `gd:"string"`
}

type AuthorInfo struct {
	LeadDevelopers  []string `gd:"lead_developers"`
	Founders        []string `gd:"founders"`
	ProjectManagers []string `gd:"project_managers"`
	Developers      []string `gd:"developers"`
}

type DonorInfo struct {
	PlatinumSponsors []string `gd:"platinum_sponsors"`
	GoldSponsors     []string `gd:"gold_sponsors"`
	SilverSponsors   []string `gd:"silver_sponsors"`
	BronzeSponsors   []string `gd:"bronze_sponsors"`
	MiniSponsors     []string `gd:"mini_sponsors"`
	GoldDonors       []string `gd:"gold_donors"`
	SilverDonors     []string `gd:"silver_donors"`
	BronzeDonors     []string `gd:"bronze_donors"`
}

type Structure map[any]any

type Atlas struct {
	Points []Vector2.XY `gd:"points"`
	Size   Vector2i.XY  `gd:"size"`
}

type Metrics struct {
	Max             Float.X `gd:"max"`
	Mean            Float.X `gd:"mean"`
	MeanSquared     Float.X `gd:"mean_squared"`
	RootMeanSquared Float.X `gd:"root_mean_squared"`
	PeakSNR         Float.X `gd:"peak_snr"`
}

type JoyInfo struct {
	XinputIndex       int    `gd:"xinput_index"`
	RawName           string `gd:"raw_name"`
	VendorID          string `gd:"vendor_id"`
	ProductID         string `gd:"product_id"`
	SteamInputPresent int    `gd:"steam_input_present"`
}

type Request struct {
	Method string `gd:"method"`
	Params any    `gd:"params"`
	ID     string `gd:"id"`
}

type Response struct {
	Result any    `gd:"result"`
	ID     string `gd:"id"`
}

type Notification struct {
	Method string `gd:"method"`
	Params any    `gd:"params"`
}

type ResponseError struct {
	Code    int    `gd:"code"`
	Message string `gd:"message"`
	ID      int    `gd:"id"`
}

type Pipe struct {
	Stdio  any `gd:"stdio" type:"[1]gdclass.FileAccess"`
	Stderr any `gd:"stderr" type:"[1]gdclass.FileAccess"`
	PID    int `gd:"pid"`
}

type MemoryInfo struct {
	Physical  int `gd:"physical"`
	Free      int `gd:"free"`
	Available int `gd:"available"`
	Stack     int `gd:"stack"`
}

type PhysicsDirectSpaceState2D_Intersection struct {
	Collider   any        `gd:"collider" type:"Object.Instance"`
	ColliderID int64      `gd:"collider_id" type:"Object.ID"`
	Normal     Vector2.XY `gd:"normal"`
	Position   Vector2.XY `gd:"position"`
	RID        RID.Any    `gd:"rid"`
	Shape      int        `gd:"shape"`
}

type PhysicsDirectSpaceState2D_RestInfo struct {
	ColliderID     int64      `gd:"collider_id" type:"Object.ID"`
	LinearVelocity Vector2.XY `gd:"linear_velocity"`
	Normal         Vector2.XY `gd:"normal"`
	Point          Vector2.XY `gd:"point"`
	RID            RID.Any    `gd:"rid"`
	Shape          int        `gd:"shape"`
}

type PhysicsDirectSpaceState3D_Intersection struct {
	Collider   any         `gd:"collider" type:"Object.Instance"`
	ColliderID int64       `gd:"collider_id" type:"Object.ID"`
	Normal     Vector3.XYZ `gd:"normal"`
	Position   Vector3.XYZ `gd:"position"`
	FaceIndex  int         `gd:"face_index"`
	RID        RID.Any     `gd:"rid"`
	Shape      int         `gd:"shape"`
}

type PhysicsDirectSpaceState3D_RestInfo struct {
	ColliderID     int64       `gd:"collider_id" type:"Object.ID"`
	LinearVelocity Vector3.XYZ `gd:"linear_velocity"`
	Normal         Vector3.XYZ `gd:"normal"`
	Point          Vector3.XYZ `gd:"point"`
	RID            RID.Any     `gd:"rid"`
	Shape          int         `gd:"shape"`
}

type Surface map[any]any

type Entry struct {
	Color Color.RGBA `gd:"color"`
}

type DateOnly struct {
	Year    int `gd:"year"`
	Month   int `gd:"month"`
	Day     int `gd:"day"`
	Weekday int `gd:"weekday"`
}

type Date struct {
	Year    int `gd:"year"`
	Month   int `gd:"month"`
	Day     int `gd:"day"`
	Weekday int `gd:"weekday"`
	Hour    int `gd:"hour"`
	Minute  int `gd:"minute"`
	Second  int `gd:"second"`
}

type OnTheClock struct {
	Hour   int `gd:"hour"`
	Minute int `gd:"minute"`
	Second int `gd:"second"`
}

type RangeConfig struct {
	Min  Float.X `gd:"min"`
	Max  Float.X `gd:"max"`
	Step Float.X `gd:"step"`
	Expr string  `gd:"expr"`
}

type Conn struct {
	Connection any  `gd:"connection" type:"[1]gdclass.WebRTCPeerConnection"`
	Channels   any  `gd:"channels" type:"[][1]gdclass.WebRTCDataChannel"`
	Connected  bool `gd:"connected"`
}

type Configuration struct {
	IceServers []IceServer `gd:"ice_servers"`
}

type IceServer struct {
	URLs       []string `gd:"urls"`
	Username   string   `gd:"username"`
	Credential string   `gd:"credential"`
}

type Options struct {
	Negotiated        bool   `gd:"negotiated"`
	ID                int    `gd:"id"`
	MaxRetransmits    int    `gd:"max_retransmits"`
	MaxPacketLifeTime int    `gd:"max_packet_life_time"`
	Ordered           bool   `gd:"ordered"`
	Protocol          string `gd:"protocol"`
}

type TextToSpeechVoice struct {
	Name     string `gd:"name"`
	ID       string `gd:"id"`
	Language string `gd:"language"`
}

type FileDialogOption struct {
	Name    string   `gd:"name"`
	Values  []string `gd:"values"`
	Default int      `gd:"default"`
}

type Copyright struct {
	Name  string `gd:"name"`
	Parts []Part `gd:"parts"`
}

type Part struct {
	Files     []string `gd:"files"`
	Copyright []string `gd:"copyright"`
	License   string   `gd:"license"`
}

type LocalInterface struct {
	Index     string       `gd:"index"`
	Name      string       `gd:"name"`
	Friendly  string       `gd:"friendly"`
	Addresses []netip.Addr `gd:"addresses"`
}

type SignalConnection struct {
	Signal   any               `gd:"signal" type:"SignalType.Any"`
	Callable Callable.Function `gd:"callable"`
	Flags    int               `gd:"Signal.ConnectFlags"`
}

type GlobalClass struct {
	Base     string `gd:"base"`
	Class    string `gd:"class"`
	Icon     string `gd:"icon"`
	Language string `gd:"language"`
	Path     string `gd:"path"`
}

type PointData struct {
	ID          Vector2i.XY `gd:"id"`
	Position    Vector2.XY  `gd:"position"`
	Solid       bool        `gd:"solid"`
	WeightScale Float.X     `gd:"weight_scale"`
}

type Connection struct {
	FromNode  string `gd:"from_node"`
	FromPort  int    `gd:"from_port"`
	ToNode    string `gd:"to_node"`
	ToPort    int    `gd:"to_port"`
	KeepAlive bool   `gd:"keep_alive"`
}

var WavOptions = Named[struct {
	CompressionMode int     `gd:"compress/mode"`
	LoopBegin       int     `gd:"edit/loop_begin"`
	LoopEnd         int     `gd:"edit/loop_end"`
	LoopMode        int     `gd:"edit/loop_mode"`
	Normalize       int     `gd:"edit/normalize"`
	Trim            bool    `gd:"edit/trim"`
	ForceMinRate    bool    `gd:"force/8_bit"`
	ForceMaxRate    bool    `gd:"force/max_rate"`
	ForceMaxRateHz  Float.X `gd:"force/max_rate_hz"`
	ForceMono       bool    `gd:"force/mono"`
}]("Options")

type Zone struct {
	Name string `gd:"name"` // is the localized name of the time zone, according to the OS locale settings of the current user
	Bias int64  `gd:"bias"` // is the offset from UTC in minutes
}

func Named[T any](name string) reflect.Type {
	return namedType{reflect.TypeFor[T](), name}
}

type namedType struct {
	reflect.Type
	name string
}

func (nt namedType) PkgPath() string {
	return "graphics.gd/internal/gdjson"
}

func (nt namedType) Name() string { return nt.name }

type writtenType struct {
	reflect.Type
	written string
	pkgtype string
}

func (wt writtenType) PkgPath() string {
	return wt.pkgtype
}

func (wt writtenType) Name() string { return wt.written }

func (wt writtenType) String() string {
	if wt.pkgtype != "" {
		return wt.pkgtype + "." + wt.written
	}
	return wt.written
}

func TypeFromString(pkgtype, written string) reflect.Type {
	return writtenType{reflect.TypeFor[struct{}](), written, pkgtype}
}

type sliceType struct {
	reflect.Type
	elem reflect.Type
}

func SliceOf(elem reflect.Type) reflect.Type {
	return sliceType{reflect.SliceOf(elem), elem}
}

func (st sliceType) PkgPath() string {
	return ""
}

func (st sliceType) String() string {
	return "[]" + st.elem.String()
}

type FormatParameters struct {
	Output string `gd:"output"`
}

type Report struct {
	Error         Error.Code `gd:"result"`
	Files         []File     `gd:"so_files"`
	EmbeddedStart int        `gd:"embedded_start"`
	EmbeddedSize  int        `gd:"embedded_size"`
}

type File struct {
	Path         string   `gd:"path"`
	Tags         []string `gd:"tags"`
	TargetFolder string   `gd:"target_folder"`
}

type TutorialDoc struct {
	Link  string `gd:"link"`
	Title string `gd:"title"`
}

type MethodDoc struct {
	Name                string        `gd:"name"`
	ReturnType          string        `gd:"return_type"`
	ReturnEnum          string        `gd:"return_enum"`
	ReturnIsBitfield    bool          `gd:"return_is_bitfield"`
	Qualifiers          string        `gd:"qualifiers"`
	Description         string        `gd:"description"`
	IsDeprecated        bool          `gd:"is_deprecated"`
	DeprecatedMessage   string        `gd:"deprecated_message"`
	IsExperimental      bool          `gd:"is_experimental"`
	ExperimentalMessage string        `gd:"experimental_message"`
	Arguments           []ArgumentDoc `gd:"arguments"`
	ErrorsReturned      []Error.Code  `gd:"errors_returned"`
	Keywords            string        `gd:"keywords"`
}

type ArgumentDoc struct {
	Name         string `gd:"name"`
	Type         string `gd:"type"`
	Enumeration  string `gd:"enumeration"`
	IsBitfield   bool   `gd:"is_bitfield"`
	DefaultValue string `gd:"default_value"`
}

type ConstantDoc struct {
	Name                string `gd:"name"`
	Value               string `gd:"value"`
	IsValueValid        bool   `gd:"is_value_valid"`
	Type                string `gd:"type"`
	Enumeration         string `gd:"enumeration"`
	IsBitfield          bool   `gd:"is_bitfield"`
	Description         string `gd:"description"`
	IsDeprecated        bool   `gd:"is_deprecated"`
	DeprecatedMessage   string `gd:"deprecated_message"`
	IsExperimental      bool   `gd:"is_experimental"`
	ExperimentalMessage string `gd:"experimental_message"`
	Keywords            string `gd:"keywords"`
}

type EnumDoc struct {
	Description         string `gd:"description"`
	IsDeprecated        bool   `gd:"is_deprecated"`
	DeprecatedMessage   string `gd:"deprecated_message"`
	IsExperimental      bool   `gd:"is_experimental"`
	ExperimentalMessage string `gd:"experimental_message"`
}

type PropertyDoc struct {
	Name                string `gd:"name"`
	Type                string `gd:"type"`
	Enumeration         string `gd:"enumeration"`
	IsBitfield          bool   `gd:"is_bitfield"`
	Description         string `gd:"description"`
	Setter              string `gd:"setter"`
	Getter              string `gd:"getter"`
	DefaultValue        string `gd:"default_value"`
	Overridden          bool   `gd:"overridden"`
	Overrides           bool   `gd:"overrides"`
	IsDeprecated        bool   `gd:"is_deprecated"`
	DeprecatedMessage   string `gd:"deprecated_message"`
	IsExperimental      bool   `gd:"is_experimental"`
	ExperimentalMessage string `gd:"experimental_message"`
	Keywords            string `gd:"keywords"`
}

type ThemeItemDoc struct {
	Name                string `gd:"name"`
	Type                string `gd:"type"`
	DataType            string `gd:"data_type"`
	Description         string `gd:"description"`
	IsDeprecated        bool   `gd:"is_deprecated"`
	DeprecatedMessage   string `gd:"deprecated_message"`
	IsExperimental      bool   `gd:"is_experimental"`
	ExperimentalMessage string `gd:"experimental_message"`
	DefaultValue        string `gd:"default_value"`
	Keywords            string `gd:"keywords"`
}

type ClassDoc struct {
	Name                string             `gd:"name"`
	Inherits            string             `gd:"inherits"`
	BriefDescription    string             `gd:"brief_description"`
	Description         string             `gd:"description"`
	Keywords            string             `gd:"keywords"`
	Tutorials           []TutorialDoc      `gd:"tutorials"`
	Constructors        []MethodDoc        `gd:"constructors"`
	Methods             []MethodDoc        `gd:"methods"`
	Operators           []MethodDoc        `gd:"operators"`
	Signals             []MethodDoc        `gd:"signals"`
	Constants           []ConstantDoc      `gd:"constants"`
	Enums               map[string]EnumDoc `gd:"enums"`
	Properties          []PropertyDoc      `gd:"properties"`
	Annotations         []MethodDoc        `gd:"annotations"`
	ThemeProperties     []ThemeItemDoc     `gd:"theme_properties"`
	IsDeprecated        bool               `gd:"is_deprecated"`
	DeprecatedMessage   string             `gd:"deprecated_message"`
	IsExperimental      bool               `gd:"is_experimental"`
	ExperimentalMessage string             `gd:"experimental_message"`
	IsScriptDoc         bool               `gd:"is_script_doc"`
	ScriptPath          string             `gd:"script_path"`
}

type Template struct {
	Inherit     string           `gd:"inherit"`
	Name        string           `gd:"name"`
	Description string           `gd:"description"`
	Content     string           `gd:"content"`
	ID          int32            `gd:"id"`
	Origin      TemplateLocation `gd:"origin"`
}

type TemplateLocation int

type Validation struct {
	Functions []string          `gd:"functions"`
	Errors    []ValidationError `gd:"errors"`
	Warnings  []Warning         `gd:"warnings"`
	SafeLines []int32           `gd:"safe_lines"`
	Valid     bool              `gd:"valid"`
}

type ValidationError struct {
	Line    int    `gd:"line"`
	Column  int    `gd:"column"`
	Message string `gd:"message"`
	Path    string `gd:"path"`
}

type Warning struct {
	StartLine       int    `gd:"start_line"`
	EndLine         int    `gd:"end_line"`
	LeftMostColumn  int    `gd:"left_most_column"`
	RightMostColumn int    `gd:"right_most_column"`
	Code            int    `gd:"code"`
	StringCode      string `gd:"string_code"`
	Message         string `gd:"message"`
}

type Code struct {
	Result              Error.Code `gd:"result"`
	Type                int        `gd:"type"`
	ClassName           string     `gd:"class_name"`
	ClassMember         string     `gd:"class_member"`
	Description         string     `gd:"description"`
	IsDeprecated        bool       `gd:"is_deprecated"`
	DeprecatedMessage   string     `gd:"deprecated_message"`
	IsExperimental      bool       `gd:"is_experimental"`
	ExperimentalMessage string     `gd:"experimental_message"`
	DocType             string     `gd:"doc_type"`
	Enumeration         string     `gd:"enumeration"`
	IsBitfield          bool       `gd:"is_bitfield"`
	Value               string     `gd:"value"`
	Script              any        `gd:"script" type:"Script.Instance"`
	ScriptPath          string     `gd:"script_path"`
	Location            int        `gd:"location"`
}

type StackLevelLocals struct {
	Locals []string `gd:"locals"`
	Values []any    `gd:"values"`
}

type StackLevelMembers struct {
	Members []string `gd:"members"`
	Values  []any    `gd:"values"`
}

type Globals struct {
	Globals []string `gd:"globals"`
	Values  []any    `gd:"values"`
}

type StackInfo struct {
	File string `gd:"file"`
	Func string `gd:"func"`
	Line int    `gd:"line"`
}

type Constant struct {
	Name  string `gd:"name"`
	Value any    `gd:"value"`
}

type ClassName struct {
	Name       string `gd:"name"`
	BaseType   string `gd:"base_type"`
	IconPath   string `gd:"icon_path"`
	IsAbstract bool   `gd:"is_abstract"`
	IsTool     bool   `gd:"is_tool"`
}

type GlyphContours struct {
	Points      []Vector3.XYZ `gd:"points"`
	Contours    []int32       `gd:"contours"`
	Orientation bool          `gd:"orientation"`
}

type OpenTypeFeature struct {
	Label  string       `gd:"label"`
	Type   reflect.Type `gd:"type"`
	Hidden bool         `gd:"hidden"`
}

type Glyph struct {
	Start    int32      `gd:"start"`
	End      int32      `gd:"end"`
	Repeat   uint8      `gd:"repeat"`
	Count    uint8      `gd:"count"`
	Flags    uint16     `gd:"flags"`
	Offset   Vector2.XY `gd:"offset"`
	Advance  Float.X    `gd:"advance"`
	FontRID  RID.Font   `gd:"font_rid"`
	FontSize int32      `gd:"font_size"`
	Index    int32      `gd:"index"`
}

type Carets struct {
	LeadingRect       Rect2.PositionSize `gd:"leading_rect"`
	LeadingDirection  int                `gd:"leading_direction" type:"Direction"`
	TrailingRect      Rect2.PositionSize `gd:"trailing_rect"`
	TrailingDirection int                `gd:"trailing_direction" type:"Direction"`
}

type ProjectedObstruction2D struct {
	Vertices []float32 `gd:"vertices"`
	Carve    bool      `gd:"carve"`
}

type ProjectedObstruction3D struct {
	Vertices  []float32 `gd:"vertices"`
	Elevation Float.X   `gd:"elevation"`
	Height    Float.X   `gd:"height"`
	Carve     bool      `gd:"carve"`
}

type Format struct {
	Width            int    `gd:"width"`
	Height           int    `gd:"height"`
	Format           string `gd:"format"`
	FrameNumerator   int    `gd:"frame_numerator"`
	FrameDenominator int    `gd:"frame_denominator"`
}

var Structables = map[string]reflect.Type{
	"AStarGrid2D.get_point_data_in_region.":                           reflect.TypeFor[PointData](),
	"AudioStreamWAV.load_from_buffer.options":                         WavOptions,
	"AudioStreamWAV.load_from_file.options":                           WavOptions,
	"CameraFeed.set_format.parameters":                                reflect.TypeFor[FormatParameters](),
	"ArrayMesh.add_surface_from_arrays.lods":                          reflect.TypeFor[map[Float.X][]int32](),
	"CharFXTransform.get_environment.":                                reflect.TypeFor[map[string]any](),
	"CharFXTransform.set_environment.environment":                     reflect.TypeFor[map[string]any](),
	"ClassDB.class_get_signal.":                                       reflect.TypeFor[SignalInfo](),
	"ClassDB.class_get_signal_list.":                                  reflect.TypeFor[SignalInfo](),
	"ClassDB.class_get_property_list.":                                reflect.TypeFor[Object.PropertyInfo](),
	"ClassDB.class_get_method_list.":                                  reflect.TypeFor[Object.PropertyInfo](),
	"JavaClass.get_java_method_list.":                                 reflect.TypeFor[Object.PropertyInfo](),
	"RenderingServer.canvas_item_get_instance_shader_parameter_list.": reflect.TypeFor[Object.PropertyInfo](),
	"CodeEdit.set_auto_brace_completion_pairs.pairs":                  reflect.TypeFor[map[string]string](),
	"CodeEdit.get_auto_brace_completion_pairs.":                       reflect.TypeFor[map[string]string](),
	"CodeEdit.get_code_completion_option.":                            reflect.TypeFor[CompletionInfo](),
	"CodeEdit.get_code_completion_options.":                           reflect.TypeFor[CompletionInfo](),
	"CodeHighlighter.set_keyword_colors.keywords":                     reflect.TypeFor[map[string]Color.RGBA](),
	"CodeHighlighter.get_keyword_colors.":                             reflect.TypeFor[map[string]Color.RGBA](),
	"CodeHighlighter.set_member_keyword_colors.member_keyword":        reflect.TypeFor[map[string]Color.RGBA](),
	"CodeHighlighter.get_member_keyword_colors.":                      reflect.TypeFor[map[string]Color.RGBA](),
	"CodeHighlighter.set_color_regions.color_regions":                 reflect.TypeFor[map[string]Color.RGBA](),
	"CodeHighlighter.get_color_regions.":                              reflect.TypeFor[map[string]Color.RGBA](),
	"DisplayServer.global_menu_get_system_menu_roots.":                reflect.TypeFor[map[string]string](),
	"DisplayServer.tts_get_voices.":                                   reflect.TypeFor[TextToSpeechVoice](),
	"DisplayServer.file_dialog_with_options_show.options":             reflect.TypeFor[FileDialogOption](),
	"EditorExportPlatform.find_export_template.": Named[struct {
		Path  string `gd:"path"`
		Error string `gd:"error"`
	}]("Template"),
	"EditorExportPlatform.save_pack.":                                   reflect.TypeFor[Report](),
	"EditorExportPlatform.save_zip.":                                    reflect.TypeFor[Report](),
	"EditorExportPlatform.save_pack_patch.":                             reflect.TypeFor[Report](),
	"EditorExportPlatform.save_zip_patch.":                              reflect.TypeFor[Report](),
	"EditorExportPlatform.get_internal_export_files.":                   reflect.TypeFor[map[string][]byte](),
	"EditorExportPreset.get_customized_files.":                          reflect.TypeFor[map[string]string](),
	"EditorFileDialog.get_selected_options.":                            reflect.TypeFor[map[string]int](),
	"EditorImportPlugin.append_import_external_resource.custom_options": reflect.TypeFor[map[string]any](),
	"EditorSettings.add_property_info.info":                             reflect.TypeFor[Object.PropertyInfo](),
	"EditorVCSInterface.create_diff_line.":                              reflect.TypeFor[DiffLine](),
	"EditorVCSInterface.create_diff_hunk.":                              reflect.TypeFor[DiffHunk](),
	"EditorVCSInterface.create_diff_file.":                              reflect.TypeFor[DiffFile](),
	"EditorVCSInterface.create_commit.":                                 reflect.TypeFor[Commit](),
	"EditorVCSInterface.create_status_file.":                            reflect.TypeFor[StatusFile](),
	"EditorVCSInterface.add_diff_hunks_into_diff_file.diff_file":        reflect.TypeFor[DiffFile](),
	"EditorVCSInterface.add_diff_hunks_into_diff_file.":                 reflect.TypeFor[DiffFile](),
	"EditorVCSInterface.add_line_diffs_into_diff_hunk.diff_hunk":        reflect.TypeFor[DiffHunk](),
	"EditorVCSInterface.add_line_diffs_into_diff_hunk.":                 reflect.TypeFor[DiffHunk](),
	"EditorVCSInterface.add_diff_hunks_into_diff_file.diff_hunks":       reflect.TypeFor[DiffHunk](),
	"EditorVCSInterface.add_line_diffs_into_diff_hunk.line_diffs":       reflect.TypeFor[DiffLine](),
	"EditorExportPlugin._get_export_options_overrides.":                 reflect.TypeFor[map[string]any](),
	"Engine.get_version_info.":                                          reflect.TypeFor[VersionInfo](),
	"Engine.get_author_info.":                                           reflect.TypeFor[AuthorInfo](),
	"Engine.get_donor_info.":                                            reflect.TypeFor[DonorInfo](),
	"Engine.get_license_info.":                                          reflect.TypeFor[map[string]string](),
	"Engine.get_copyright_info.":                                        reflect.TypeFor[Copyright](),
	"FileDialog.get_selected_options.":                                  reflect.TypeFor[map[string]int](),
	"Font.find_variation.variation_coordinates":                         reflect.TypeFor[map[string]Float.X](),
	"Font.get_ot_name_strings.":                                         reflect.TypeFor[map[string]map[string]string](),
	"Font.get_opentype_features.":                                       reflect.TypeFor[map[string][2]string](),
	"Font.get_supported_feature_list.":                                  reflect.TypeFor[map[string]OpenTypeFeature](),
	"Font.get_supported_variation_list.":                                reflect.TypeFor[map[int]Vector3i.XYZ](),
	"FontFile.set_variation_coordinates.variation_coordinates":          reflect.TypeFor[map[string]Float.X](),
	"FontFile.get_variation_coordinates.":                               reflect.TypeFor[map[string]Float.X](),
	"FontFile.set_opentype_feature_overrides.overrides":                 reflect.TypeFor[map[string][2]string](),
	"FontFile.get_opentype_feature_overrides.":                          reflect.TypeFor[map[string][2]string](),
	"FontVariation.set_variation_opentype.coords":                       reflect.TypeFor[map[any]Float.X](),
	"FontVariation.get_variation_opentype.":                             reflect.TypeFor[map[any]Float.X](),
	"FontVariation.set_opentype_features.features":                      reflect.TypeFor[map[string]uint32](),
	"GLTFCamera.from_dictionary.dictionary":                             reflect.TypeFor[Structure](),
	"GLTFCamera.to_dictionary.":                                         reflect.TypeFor[Structure](),
	"GLTFLight.from_dictionary.dictionary":                              reflect.TypeFor[Structure](),
	"GLTFLight.to_dictionary.":                                          reflect.TypeFor[Structure](),
	"GLTFPhysicsBody.from_dictionary.dictionary":                        reflect.TypeFor[Structure](),
	"GLTFPhysicsBody.to_dictionary.":                                    reflect.TypeFor[Structure](),
	"GLTFPhysicsShape.from_dictionary.dictionary":                       reflect.TypeFor[Structure](),
	"GLTFPhysicsShape.to_dictionary.":                                   reflect.TypeFor[Structure](),
	"GLTFSkeleton.get_godot_bone_node.":                                 reflect.TypeFor[map[int]int](),
	"GLTFSkeleton.set_godot_bone_node.godot_bone_node":                  reflect.TypeFor[map[int]int](),
	"GLTFSkin.get_joint_i_to_bone_i.":                                   reflect.TypeFor[map[int]int](),
	"GLTFSkin.set_joint_i_to_bone_i.joint_i_to_bone_i":                  reflect.TypeFor[map[int]int](),
	"GLTFSkin.get_joint_i_to_name.":                                     reflect.TypeFor[map[int]string](),
	"GLTFSkin.set_joint_i_to_name.joint_i_to_name":                      reflect.TypeFor[map[int]string](),
	"GLTFState.get_json.":                                               reflect.TypeFor[map[string]any](),
	"GLTFState.set_json.json":                                           reflect.TypeFor[map[string]any](),
	"Geometry2D.make_atlas.":                                            reflect.TypeFor[Atlas](),
	"GraphEdit.get_closest_connection_at_point.":                        reflect.TypeFor[Connection](),
	"GraphEdit.get_connection_list.":                                    reflect.TypeFor[Connection](),
	"GraphEdit.get_connections_intersecting_with_rect.":                 reflect.TypeFor[Connection](),
	"GraphEdit.set_connections.connections":                             reflect.TypeFor[Connection](),
	"HTTPClient.get_response_headers_as_dictionary.":                    reflect.TypeFor[map[string]string](),
	"HTTPClient.query_string_from_dict.fields":                          reflect.TypeFor[map[string]string](),
	"Image.compute_image_metrics.":                                      reflect.TypeFor[Metrics](),
	"ImporterMesh.add_surface.lods":                                     reflect.TypeFor[map[Float.X][]int32](),
	"Input.get_joy_info.":                                               reflect.TypeFor[JoyInfo](),
	"InstancePlaceholder.get_stored_values.":                            reflect.TypeFor[map[string]any](),
	"IP.get_local_interfaces.":                                          reflect.TypeFor[LocalInterface](),
	"JSONRPC.make_request.":                                             reflect.TypeFor[Request](),
	"JSONRPC.make_response.":                                            reflect.TypeFor[Response](),
	"JSONRPC.make_notification.":                                        reflect.TypeFor[Notification](),
	"JSONRPC.make_response_error.":                                      reflect.TypeFor[ResponseError](),
	"Object.get_property_list.":                                         reflect.TypeFor[Object.PropertyInfo](),
	"Object.get_method_list.":                                           reflect.TypeFor[Object.PropertyInfo](),
	"Object.get_signal_list.":                                           reflect.TypeFor[SignalInfo](),
	"Object.get_signal_connection_list.":                                reflect.TypeFor[SignalConnection](),
	"Object.get_incoming_connections.":                                  reflect.TypeFor[SignalConnection](),
	"OS.execute_with_pipe.":                                             reflect.TypeFor[Pipe](),
	"OS.get_memory_info.":                                               reflect.TypeFor[MemoryInfo](),
	"PhysicsDirectSpaceState2D.intersect_ray.":                          reflect.TypeFor[PhysicsDirectSpaceState2D_Intersection](),
	"PhysicsDirectSpaceState2D.intersect_point.":                        reflect.TypeFor[PhysicsDirectSpaceState2D_Intersection](),
	"PhysicsDirectSpaceState2D.intersect_shape.":                        reflect.TypeFor[PhysicsDirectSpaceState2D_Intersection](),
	"PhysicsDirectSpaceState2D.get_rest_info.":                          reflect.TypeFor[PhysicsDirectSpaceState2D_RestInfo](),
	"PhysicsDirectSpaceState3D.intersect_ray.":                          reflect.TypeFor[PhysicsDirectSpaceState3D_Intersection](),
	"PhysicsDirectSpaceState3D.intersect_point.":                        reflect.TypeFor[PhysicsDirectSpaceState3D_Intersection](),
	"PhysicsDirectSpaceState3D.intersect_shape.":                        reflect.TypeFor[PhysicsDirectSpaceState3D_Intersection](),
	"PhysicsDirectSpaceState3D.get_rest_info.":                          reflect.TypeFor[PhysicsDirectSpaceState3D_RestInfo](),
	"ProjectSettings.add_property_info.hint":                            reflect.TypeFor[Object.PropertyInfo](),
	"ProjectSettings.get_global_class_list.":                            reflect.TypeFor[GlobalClass](),
	"RegExMatch.get_names.":                                             reflect.TypeFor[map[string]int](),
	"RenderingServer.mesh_add_surface.surface":                          reflect.TypeFor[Surface](),
	"RenderingServer.mesh_add_surface_from_arrays.lods":                 reflect.TypeFor[map[Float.X][]int32](),
	"RenderingServer.mesh_get_surface.":                                 reflect.TypeFor[Surface](),
	"RenderingServer.get_shader_parameter_list.":                        reflect.TypeFor[map[string]any](),
	"RenderingServer.mesh_create_from_surfaces.surfaces":                reflect.TypeFor[Surface](),
	"RenderingServer.instance_geometry_get_shader_parameter_list.":      reflect.TypeFor[map[string]any](),
	"RichTextLabel.push_customfx.env":                                   reflect.TypeFor[map[string]any](),
	"RichTextLabel.parse_expressions_for_values.":                       reflect.TypeFor[map[string]any](),
	"Script.get_script_constant_map.":                                   reflect.TypeFor[map[string]any](),
	"Script.get_script_property_list.":                                  reflect.TypeFor[Object.PropertyInfo](),
	"Script.get_script_method_list.":                                    reflect.TypeFor[Object.PropertyInfo](),
	"Script.get_script_signal_list.":                                    reflect.TypeFor[SignalInfo](),
	"ShapeCast3D.collision_result":                                      reflect.TypeFor[[]PhysicsDirectSpaceState3D_RestInfo](),
	"SyntaxHighlighter.get_line_syntax_highlighting.":                   reflect.TypeFor[map[int]Entry](),
	"TextServer.font_get_ot_name_strings.":                              reflect.TypeFor[map[string]map[string]string](),
	"TextServer.font_set_variation_coordinates.variation_coordinates":   reflect.TypeFor[map[string]Float.X](),
	"TextServer.font_get_variation_coordinates.":                        reflect.TypeFor[map[string]Float.X](),
	"TextServer.font_get_glyph_contours.":                               reflect.TypeFor[GlyphContours](),
	"TextServer.font_set_opentype_feature_overrides.overrides":          reflect.TypeFor[map[string][2]string](),
	"TextServer.font_get_opentype_feature_overrides.":                   reflect.TypeFor[map[string][2]string](),
	"TextServer.font_supported_feature_list.":                           reflect.TypeFor[map[string]OpenTypeFeature](),
	"TextServer.font_supported_variation_list.":                         reflect.TypeFor[map[int]Vector3i.XYZ](),
	"TextServer.shaped_text_add_string.opentype_features":               reflect.TypeFor[map[string]uint32](),
	"TextServer.shaped_set_span_update_font.opentype_features":          reflect.TypeFor[map[string]uint32](),
	"TextServer.shaped_text_get_glyphs.":                                reflect.TypeFor[[]Glyph](),
	"TextServer.shaped_text_sort_logical.":                              reflect.TypeFor[[]Glyph](),
	"TextServer.shaped_text_get_ellipsis_glyphs.":                       reflect.TypeFor[[]Glyph](),
	"TextServer.shaped_text_get_carets.":                                reflect.TypeFor[Carets](),
	"TextServerManager.get_interfaces.":                                 reflect.TypeFor[map[int]string](),
	"Time.get_datetime_dict_from_unix_time.":                            reflect.TypeFor[Date](),
	"Time.get_date_dict_from_unix_time.":                                reflect.TypeFor[DateOnly](),
	"Time.get_time_dict_from_unix_time.":                                reflect.TypeFor[OnTheClock](),
	"Time.get_datetime_dict_from_datetime_string.":                      reflect.TypeFor[Date](),
	"Time.get_datetime_string_from_datetime_dict.datetime":              reflect.TypeFor[Date](),
	"Time.get_unix_time_from_datetime_dict.datetime":                    reflect.TypeFor[Date](),
	"Time.get_datetime_dict_from_system.":                               reflect.TypeFor[Date](),
	"Time.get_date_dict_from_system.":                                   reflect.TypeFor[DateOnly](),
	"Time.get_time_dict_from_system.":                                   reflect.TypeFor[OnTheClock](),
	"Time.get_time_zone_from_system.":                                   reflect.TypeFor[Zone](),
	"TreeItem.get_range_config.":                                        reflect.TypeFor[RangeConfig](),
	"WebRTCMultiplayerPeer.get_peer.":                                   reflect.TypeFor[Conn](),
	"VisualShader.get_node_connections.":                                reflect.TypeFor[map[string]any](),
	"WebRTCMultiplayerPeer.get_peers.":                                  reflect.TypeFor[map[int]Conn](),
	"WebRTCPeerConnection.initialize.configuration":                     reflect.TypeFor[Configuration](),
	"WebRTCPeerConnection.create_data_channel.options":                  reflect.TypeFor[Options](),
	"XRInterface.get_system_info.":                                      reflect.TypeFor[map[string]any](),
	"XRServer.get_trackers.":                                            reflect.TypeFor[map[any]any](),
	"XRServer.get_interfaces.":                                          reflect.TypeFor[map[int]string](),

	"AnimationNode._get_child_nodes.":                        reflect.MapOf(reflect.TypeFor[string](), TypeFromString("graphics.gd/classdb/Node", "Instance")),
	"AudioStream._get_parameter_list.":                       reflect.TypeFor[[]Object.PropertyInfo](),
	"CodeEdit._filter_code_completion_candidates.candidates": reflect.TypeFor[[]CompletionInfo](),
	"CodeEdit._filter_code_completion_candidates.":           reflect.TypeFor[[]CompletionInfo](),

	"EditorExportPlatformExtension._get_export_options.": SliceOf(Named[struct {
		Hint             int    `gd:"hint"`
		HintString       string `gd:"hint_string"`
		Usage            int    `gd:"usage"`
		ClassName        string `gd:"class_name"`
		DefaultValue     any    `gd:"default_value"`
		UpdateVisibility bool   `gd:"update_visibility"`
		Required         bool   `gd:"required"`
	}]("Option")),
	"EditorExportPlugin._get_export_options.": SliceOf(Named[struct {
		Option           Object.PropertyInfo `gd:"option"`
		DefaultValue     any                 `gd:"default_value"`
		UpdateVisibility bool                `gd:"update_visibility"`
	}]("Option")),
	"EditorImportPlugin._get_import_options.": SliceOf(Named[struct {
		Name         string `gd:"name"`
		DefaultValue any    `gd:"default_value"`
		PropertyHint int    `gd:"property_hint"`
		HintString   string `gd:"hint_string"`
		Usage        int    `gd:"usage"`
	}]("Option")),
	"EditorImportPlugin._get_option_visibility.options":                                                               reflect.TypeFor[map[string]any](),
	"EditorImportPlugin._import.options":                                                                              reflect.TypeFor[map[string]any](),
	"EditorPlugin._get_state.":                                                                                        reflect.TypeFor[map[any]any](),
	"EditorPlugin._set_state.state":                                                                                   reflect.TypeFor[map[any]any](),
	"EditorResourcePreviewGenerator._generate.metadata":                                                               reflect.TypeFor[map[string]any](),
	"EditorResourcePreviewGenerator._generate_from_path.metadata":                                                     reflect.TypeFor[map[string]any](),
	"EditorResourceTooltipPlugin._make_tooltip_for_path.metadata":                                                     reflect.TypeFor[map[string]any](),
	"EditorSceneFormatImporter._import_scene.options":                                                                 reflect.TypeFor[map[string]any](),
	"EditorVCSInterface._get_modified_files_data.":                                                                    reflect.TypeFor[[]StatusFile](),
	"EditorVCSInterface._get_diff.":                                                                                   reflect.TypeFor[[]DiffFile](),
	"EditorVCSInterface._get_previous_commits.":                                                                       reflect.TypeFor[[]Commit](),
	"EditorVCSInterface._get_line_diff.":                                                                              reflect.TypeFor[[]DiffLine](),
	"GLTFDocumentExtension._parse_node_extensions.extensions":                                                         reflect.TypeFor[map[string]any](),
	"GLTFDocumentExtension._parse_texture_json.texture_json":                                                          reflect.TypeFor[map[string]any](),
	"GLTFDocumentExtension._import_node.json":                                                                         reflect.TypeFor[map[string]any](),
	"GLTFDocumentExtension._serialize_image_to_bytes.image_dict":                                                      reflect.TypeFor[map[string][]byte](),
	"GLTFDocumentExtension._serialize_texture_json.texture_json":                                                      reflect.TypeFor[map[string]any](),
	"GLTFDocumentExtension._export_node.json":                                                                         reflect.TypeFor[map[string]any](),
	"Mesh._surface_get_lods.":                                                                                         reflect.TypeFor[map[Float.X][]int32](),
	"OpenXRExtensionWrapperExtension._get_requested_extensions.":                                                      reflect.TypeFor[map[string]*bool](),
	"OpenXRExtensionWrapperExtension._set_viewport_composition_layer_and_get_next_pointer.property_values":            reflect.TypeFor[map[string]Object.PropertyInfo](),
	"OpenXRExtensionWrapperExtension._get_viewport_composition_layer_extension_properties.":                           reflect.TypeFor[[]Object.PropertyInfo](),
	"OpenXRExtensionWrapperExtension._set_android_surface_swapchain_create_info_and_get_next_pointer.property_values": reflect.TypeFor[map[string]Object.PropertyInfo](),
	"OpenXRExtensionWrapperExtension._get_viewport_composition_layer_extension_property_defaults.":                    reflect.TypeFor[map[string]any](),
	"ResourceFormatLoader._rename_dependencies.renames":                                                               reflect.TypeFor[map[string]string](),

	"ScriptExtension._get_documentation.":        reflect.TypeFor[[]ClassDoc](),
	"ScriptExtension._get_method_info.":          reflect.TypeFor[Object.MethodInfo](),
	"ScriptExtension._get_script_signal_list.":   reflect.TypeFor[[]Object.MethodInfo](),
	"ScriptExtension._get_script_method_list.":   reflect.TypeFor[[]Object.MethodInfo](),
	"ScriptExtension._get_script_property_list.": reflect.TypeFor[[]Object.PropertyInfo](),
	"ScriptExtension._get_constants.":            reflect.TypeFor[map[string]any](),

	"ScriptLanguageExtension._get_built_in_templates.":        reflect.TypeFor[[]Template](),
	"ScriptLanguageExtension._validate.":                      reflect.TypeFor[Validation](),
	"ScriptLanguageExtension._complete_code.":                 reflect.TypeFor[Completion](),
	"ScriptLanguageExtension._lookup_code.":                   reflect.TypeFor[Code](),
	"ScriptLanguageExtension._debug_get_stack_level_locals.":  reflect.TypeFor[StackLevelLocals](),
	"ScriptLanguageExtension._debug_get_stack_level_members.": reflect.TypeFor[StackLevelMembers](),
	"ScriptLanguageExtension._debug_get_globals.":             reflect.TypeFor[Globals](),
	"ScriptLanguageExtension._debug_get_current_stack_info.":  reflect.TypeFor[StackInfo](),
	"ScriptLanguageExtension._get_public_functions.":          reflect.TypeFor[[]Object.MethodInfo](),
	"ScriptLanguageExtension._get_public_constants.":          reflect.TypeFor[[]Constant](),
	"ScriptLanguageExtension._get_public_annotations.":        reflect.TypeFor[[]Object.MethodInfo](),
	"ScriptLanguageExtension._get_global_class_name.":         reflect.TypeFor[ClassName](),

	"SyntaxHighlighter._get_line_syntax_highlighting.": reflect.TypeFor[map[int]Entry](),

	"TextServerExtension._font_get_ot_name_strings.":                            reflect.TypeFor[map[string]map[string]string](),
	"TextServerExtension._font_set_variation_coordinates.variation_coordinates": reflect.TypeFor[map[string]Float.X](),
	"TextServerExtension._font_get_variation_coordinates.":                      reflect.TypeFor[map[string]Float.X](),
	"TextServerExtension._font_get_glyph_contours.":                             reflect.TypeFor[GlyphContours](),
	"TextServerExtension._font_set_opentype_feature_overrides.overrides":        reflect.TypeFor[map[string][2]string](),
	"TextServerExtension._font_get_opentype_feature_overrides.":                 reflect.TypeFor[map[string][2]string](),
	"TextServerExtension._font_supported_feature_list.":                         reflect.TypeFor[map[string]OpenTypeFeature](),
	"TextServerExtension._font_supported_variation_list.":                       reflect.TypeFor[map[int]Vector3i.XYZ](),
	"TextServerExtension._shaped_text_add_string.opentype_features":             reflect.TypeFor[map[string]uint32](),
	"TextServerExtension._shaped_set_span_update_font.opentype_features":        reflect.TypeFor[map[string]uint32](),

	"WebRTCPeerConnectionExtension._initialize.p_config":          reflect.TypeFor[Configuration](),
	"WebRTCPeerConnectionExtension._create_data_channel.p_config": reflect.TypeFor[Configuration](),

	"XRInterfaceExtension._get_system_info.": reflect.TypeFor[map[string]any](),

	"Animation.method_track_get_params.":                                                   reflect.TypeFor[[]any](),
	"AnimationNode._get_parameter_list.":                                                   reflect.TypeFor[[]Object.PropertyInfo](),
	"ArrayMesh.add_surface_from_arrays.arrays":                                             reflect.TypeFor[[]any](),
	"ArrayMesh.add_surface_from_arrays.blend_shapes":                                       reflect.TypeFor[[][]any](),
	"CameraFeed.get_formats.":                                                              reflect.TypeFor[[]Format](),
	"Control._structured_text_parser.args":                                                 reflect.TypeFor[[]any](),
	"EditorDebuggerPlugin._capture.data":                                                   reflect.TypeFor[[]any](),
	"EditorDebuggerPlugin.get_sessions.":                                                   SliceOf(TypeFromString("EditorDebuggerSession", "Instance")),
	"EditorDebuggerSession.send_message.data":                                              reflect.TypeFor[[]any](),
	"EditorDebuggerSession.toggle_profiler.data":                                           reflect.TypeFor[[]any](),
	"EditorExportPlatform.get_current_presets.":                                            SliceOf(TypeFromString("EditorExportPreset", "Instance")),
	"EditorExportPlatform.ssh_run_on_remote.output":                                        reflect.TypeFor[[]string](), // FIXME
	"EngineDebugger.profiler_add_frame_data.data":                                          reflect.TypeFor[[]any](),
	"EngineDebugger.profiler_enable.arguments":                                             reflect.TypeFor[[]any](),
	"EngineDebugger.send_message.data":                                                     reflect.TypeFor[[]any](),
	"EngineProfiler._toggle.options":                                                       reflect.TypeFor[[]any](),
	"EngineProfiler._add_frame.data":                                                       reflect.TypeFor[[]any](),
	"Expression.execute.inputs":                                                            reflect.TypeFor[[]any](),
	"GridMap.get_meshes.":                                                                  reflect.TypeFor[[]any](),
	"GridMap.get_bake_meshes.":                                                             reflect.TypeFor[[]any](),
	"GridMapEditorPlugin.get_selected_cells.":                                              reflect.TypeFor[[]Vector3i.XYZ](),
	"IP.get_resolve_item_addresses.":                                                       reflect.TypeFor[[]string](),
	"ImporterMesh.add_surface.arrays":                                                      reflect.TypeFor[[]any](),
	"ImporterMesh.add_surface.blend_shapes":                                                reflect.TypeFor[[][]any](),
	"ImporterMesh.get_surface_arrays.":                                                     reflect.TypeFor[[]any](),
	"ImporterMesh.get_surface_blend_shape_arrays.":                                         reflect.TypeFor[[][]any](),
	"ImporterMesh.generate_lods.bone_transform_array":                                      reflect.TypeFor[[]Transform3D.BasisOrigin](),
	"Label.set_structured_text_bidi_override_options.args":                                 reflect.TypeFor[[]any](),
	"Label.get_structured_text_bidi_override_options.":                                     reflect.TypeFor[[]any](),
	"Label3D.set_structured_text_bidi_override_options.args":                               reflect.TypeFor[[]any](),
	"Label3D.get_structured_text_bidi_override_options.":                                   reflect.TypeFor[[]any](),
	"LineEdit.set_structured_text_bidi_override_options.args":                              reflect.TypeFor[[]any](),
	"LineEdit.get_structured_text_bidi_override_options.":                                  reflect.TypeFor[[]any](),
	"LinkButton.set_structured_text_bidi_override_options.args":                            reflect.TypeFor[[]any](),
	"LinkButton.get_structured_text_bidi_override_options.":                                reflect.TypeFor[[]any](),
	"Mesh._surface_get_arrays.":                                                            reflect.TypeFor[[]any](),
	"Mesh._surface_get_blend_shape_arrays.":                                                reflect.TypeFor[[][]any](),
	"Mesh.surface_get_arrays.":                                                             reflect.TypeFor[[]any](),
	"Mesh.surface_get_blend_shape_arrays.":                                                 reflect.TypeFor[[][]any](),
	"MeshLibrary.set_item_shapes.shapes":                                                   SliceOf(TypeFromString("Shape3D", "Instance")),
	"MeshLibrary.get_item_shapes.":                                                         SliceOf(TypeFromString("Shape3D", "Instance")),
	"MultiplayerAPI.rpc.arguments":                                                         reflect.TypeFor[[]any](),
	"MultiplayerAPIExtension._rpc.args":                                                    reflect.TypeFor[[]any](),
	"NavigationMeshSourceGeometryData2D.set_projected_obstructions.projected_obstructions": reflect.TypeFor[[]ProjectedObstruction2D](),
	"NavigationMeshSourceGeometryData2D.get_projected_obstructions.":                       reflect.TypeFor[[]ProjectedObstruction2D](),
	"NavigationMeshSourceGeometryData3D.set_projected_obstructions.projected_obstructions": reflect.TypeFor[[]ProjectedObstruction3D](),
	"NavigationMeshSourceGeometryData3D.add_mesh_array.mesh_array":                         reflect.TypeFor[[]any](),
	"NavigationMeshSourceGeometryData3D.get_projected_obstructions.":                       reflect.TypeFor[[]ProjectedObstruction3D](),
	"Node.propagate_call.args":                                                             reflect.TypeFor[[]any](),
	"OS.execute.output":                                                                    reflect.TypeFor[[]string](), // FIXME
	"Object.add_user_signal.arguments":                                                     reflect.TypeFor[[]any](),
	"Object.callv.arg_array":                                                               reflect.TypeFor[[]any](),
	"OggPacketSequence.set_packet_data.packet_data":                                        reflect.TypeFor[[][]any](),
	"OggPacketSequence.get_packet_data.":                                                   reflect.TypeFor[[][]any](),
	"OpenXRAPIExtension.xr_result.args":                                                    reflect.TypeFor[[]any](),
	"OpenXRActionMap.set_action_sets.action_sets":                                          SliceOf(TypeFromString("OpenXRActionSet", "Instance")),
	"OpenXRActionMap.get_action_sets.":                                                     SliceOf(TypeFromString("OpenXRActionSet", "Instance")),
	"OpenXRActionMap.set_interaction_profiles.interaction_profiles":                        SliceOf(TypeFromString("OpenXRInteractionProfile", "Instance")),
	"OpenXRActionMap.get_interaction_profiles.":                                            SliceOf(TypeFromString("OpenXRInteractionProfile", "Instance")),
	"OpenXRActionSet.set_actions.actions":                                                  SliceOf(TypeFromString("OpenXRAction", "Instance")),
	"OpenXRActionSet.get_actions.":                                                         SliceOf(TypeFromString("OpenXRAction", "Instance")),
	"OpenXRIPBinding.set_binding_modifiers.binding_modifiers":                              SliceOf(TypeFromString("OpenXRActionBindingModifier", "Instance")),
	"OpenXRIPBinding.get_binding_modifiers.":                                               SliceOf(TypeFromString("OpenXRActionBindingModifier", "Instance")),
	"OpenXRInteractionProfile.set_bindings.bindings":                                       SliceOf(TypeFromString("OpenXRIPBinding", "Instance")),
	"OpenXRInteractionProfile.get_bindings.":                                               SliceOf(TypeFromString("OpenXRIPBinding", "Instance")),
	"OpenXRInteractionProfile.set_binding_modifiers.binding_modifiers":                     SliceOf(TypeFromString("OpenXRIPBindingModifier", "Instance")),
	"OpenXRInteractionProfile.get_binding_modifiers.":                                      SliceOf(TypeFromString("OpenXRIPBindingModifier", "Instance")),
	"OpenXRInterface.get_action_sets.":                                                     SliceOf(TypeFromString("OpenXRActionSet", "Instance")),
	"OpenXRInterface.get_available_display_refresh_rates.":                                 reflect.TypeFor[[]Float.X](),
	"Performance.add_custom_monitor.arguments":                                             reflect.TypeFor[[]any](),
	"Polygon2D.set_polygons.polygons":                                                      reflect.TypeFor[[][]int32](),
	"Polygon2D.get_polygons.":                                                              reflect.TypeFor[[][]int32](),
	"PrimitiveMesh._create_mesh_array.":                                                    reflect.TypeFor[[][]any](),
	"PrimitiveMesh.get_mesh_arrays.":                                                       reflect.TypeFor[[]any](),
	"RenderingServer.mesh_add_surface_from_arrays.arrays":                                  reflect.TypeFor[[]any](),
	"RenderingServer.mesh_add_surface_from_arrays.blend_shapes":                            reflect.TypeFor[[][]any](),
	"RenderingServer.mesh_surface_get_arrays.":                                             reflect.TypeFor[[]any](),
	"RenderingServer.mesh_surface_get_blend_shape_arrays.":                                 reflect.TypeFor[[][]any](),
	"ResourceLoader.load_threaded_get_status.progress":                                     reflect.TypeFor[[]Float.X](), // FIXME
	"RichTextLabel.set_structured_text_bidi_override_options.args":                         reflect.TypeFor[[]any](),
	"RichTextLabel.get_structured_text_bidi_override_options.":                             reflect.TypeFor[[]any](),
	"RichTextLabel.set_effects.effects":                                                    SliceOf(TypeFromString("RichTextEffect", "Instance")),
	"RichTextLabel.get_effects.":                                                           SliceOf(TypeFromString("RichTextEffect", "Instance")),
	"SceneState.get_connection_binds.":                                                     reflect.TypeFor[[]any](),
	"ScriptLanguageExtension._reload_scripts.scripts":                                      SliceOf(TypeFromString("Script", "Instance")),
	"Shader.get_shader_uniform_list.":                                                      reflect.TypeFor[[]Object.PropertyInfo](),
	"ShapeCast2D.get_collision_result.":                                                    reflect.TypeFor[[]PhysicsDirectSpaceState2D_RestInfo](),
	"ShapeCast3D.get_collision_result.":                                                    reflect.TypeFor[[]PhysicsDirectSpaceState3D_RestInfo](),
	"Shortcut.set_events.events":                                                           SliceOf(TypeFromString("InputEvent", "Instance")),
	"Shortcut.get_events.":                                                                 SliceOf(TypeFromString("InputEvent", "Instance")),
	"SurfaceTool.create_from_arrays.arrays":                                                reflect.TypeFor[[]any](),
	"SurfaceTool.commit_to_arrays.":                                                        reflect.TypeFor[[]any](),
	"TextEdit.set_structured_text_bidi_override_options.args":                              reflect.TypeFor[[]any](),
	"TextEdit.get_structured_text_bidi_override_options.":                                  reflect.TypeFor[[]any](),
	"TextLine.set_bidi_override.override":                                                  reflect.TypeFor[[]any](),
	"TextLine.get_objects.":                                                                reflect.TypeFor[[]any](),
	"TextMesh.set_structured_text_bidi_override_options.args":                              reflect.TypeFor[[]any](),
	"TextMesh.get_structured_text_bidi_override_options.":                                  reflect.TypeFor[[]any](),
	"TextParagraph.set_bidi_override.override":                                             reflect.TypeFor[[]any](),
	"TextParagraph.get_line_objects.":                                                      reflect.TypeFor[[]any](),
	"TextServer.shaped_text_set_bidi_override.override":                                    reflect.TypeFor[[]any](),
	"TextServer.shaped_text_get_objects.":                                                  reflect.TypeFor[[]any](),
	"TextServer.parse_structured_text.args":                                                reflect.TypeFor[[]any](),
	"TextServerExtension._shaped_text_set_bidi_override.override":                          reflect.TypeFor[[]any](),
	"TextServerExtension._shaped_text_get_objects.":                                        reflect.TypeFor[[]any](),
	"TextServerExtension._parse_structured_text.args":                                      reflect.TypeFor[[]any](),
	"TreeItem.set_structured_text_bidi_override_options.args":                              reflect.TypeFor[[]any](),
	"TreeItem.get_structured_text_bidi_override_options.":                                  reflect.TypeFor[[]any](),
	"VisualShaderNode.set_default_input_values.values":                                     reflect.TypeFor[[]any](),
	"VisualShaderNode.get_default_input_values.":                                           reflect.TypeFor[[]any](),
	"WebRTCMultiplayerPeer.create_server.channels_config":                                  SliceOf(TypeFromString("MultiplayerPeer", "TransferMode")),
	"WebRTCMultiplayerPeer.create_client.channels_config":                                  SliceOf(TypeFromString("MultiplayerPeer", "TransferMode")),
	"WebRTCMultiplayerPeer.create_mesh.channels_config":                                    SliceOf(TypeFromString("MultiplayerPeer", "TransferMode")),
	"WebXRInterface.get_available_display_refresh_rates.":                                  reflect.TypeFor[[]Float.X](),
	"XRInterface.get_supported_environment_blend_modes.":                                   SliceOf(TypeFromString("", "EnvironmentBlendMode")),
}
