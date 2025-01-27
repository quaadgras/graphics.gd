package gdjson

import (
	"net/netip"
	"reflect"

	"graphics.gd/variant/Callable"
	"graphics.gd/variant/Color"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Path"
	"graphics.gd/variant/RID"
	"graphics.gd/variant/Vector2"
	"graphics.gd/variant/Vector2i"
	"graphics.gd/variant/Vector3"
)

type PropertyInfo struct {
	ClassName  string       `gd:"class_name"`
	Name       string       `gd:"name"`
	Hint       int          `gd:"hint"`
	HintString string       `gd:"hint_string"`
	Type       reflect.Type `gd:"type"`
	Usage      int          `gd:"usage"`
}

type SignalInfo struct {
	Name        string         `gd:"name"`
	Flags       int            `gd:"flags"`
	ID          int            `gd:"id"`
	DefaultArgs []any          `gd:"default_args"`
	Args        []PropertyInfo `gd:"args"`
}

type CompletionInfo struct {
	Kind         any             `gd:"kind" type:"[1]gdclass.CodeEditCodeCompletionKind"`
	DisplayText  string          `gd:"display_text"`
	InsertText   string          `gd:"insert_text"`
	FontColor    Color.RGBA      `gd:"font_color"`
	Icon         Path.ToResource `gd:"icon"`
	DefaultValue string          `gd:"default_value"`
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

type Connection struct {
	FromPort int    `gd:"from_port"`
	FromNode string `gd:"from_node"`
	ToPort   int    `gd:"to_port"`
	ToNode   string `gd:"to_node"`
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

var Structables = map[string]reflect.Type{
	"ArrayMesh.add_surface_from_arrays.lods":                            reflect.TypeFor[map[Float.X][]int32](),
	"CharFXTransform.get_environment.":                                  reflect.TypeFor[map[string]any](),
	"CharFXTransform.set_environment.environment":                       reflect.TypeFor[map[string]any](),
	"ClassDB.class_get_signal.":                                         reflect.TypeFor[SignalInfo](),
	"ClassDB.class_get_signal_list.":                                    reflect.TypeFor[SignalInfo](),
	"ClassDB.class_get_property_list.":                                  reflect.TypeFor[PropertyInfo](),
	"ClassDB.class_get_method_list.":                                    reflect.TypeFor[PropertyInfo](),
	"CodeEdit.set_auto_brace_completion_pairs.pairs":                    reflect.TypeFor[map[string]string](),
	"CodeEdit.get_auto_brace_completion_pairs.":                         reflect.TypeFor[map[string]string](),
	"CodeEdit.get_code_completion_option.":                              reflect.TypeFor[CompletionInfo](),
	"CodeEdit.get_code_completion_options.":                             reflect.TypeFor[CompletionInfo](),
	"CodeHighlighter.set_keyword_colors.keywords":                       reflect.TypeFor[map[string]Color.RGBA](),
	"CodeHighlighter.get_keyword_colors.":                               reflect.TypeFor[map[string]Color.RGBA](),
	"CodeHighlighter.set_member_keyword_colors.member_keyword":          reflect.TypeFor[map[string]Color.RGBA](),
	"CodeHighlighter.get_member_keyword_colors.":                        reflect.TypeFor[map[string]Color.RGBA](),
	"CodeHighlighter.set_color_regions.color_regions":                   reflect.TypeFor[map[string]Color.RGBA](),
	"CodeHighlighter.get_color_regions.":                                reflect.TypeFor[map[string]Color.RGBA](),
	"DisplayServer.global_menu_get_system_menu_roots.":                  reflect.TypeFor[map[string]string](),
	"DisplayServer.tts_get_voices.":                                     reflect.TypeFor[TextToSpeechVoice](),
	"DisplayServer.file_dialog_with_options_show.options":               reflect.TypeFor[FileDialogOption](),
	"EditorFileDialog.get_selected_options.":                            reflect.TypeFor[map[string]int](),
	"EditorImportPlugin.append_import_external_resource.custom_options": reflect.TypeFor[map[string]any](),
	"EditorSettings.add_property_info.info":                             reflect.TypeFor[PropertyInfo](),
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
	"Engine.get_version_info.":                                          reflect.TypeFor[VersionInfo](),
	"Engine.get_author_info.":                                           reflect.TypeFor[AuthorInfo](),
	"Engine.get_donor_info.":                                            reflect.TypeFor[DonorInfo](),
	"Engine.get_license_info.":                                          reflect.TypeFor[map[string]string](),
	"Engine.get_copyright_info.":                                        reflect.TypeFor[Copyright](),
	"FileDialog.get_selected_options.":                                  reflect.TypeFor[map[string]int](),
	"Font.find_variation.variation_coordinates":                         reflect.TypeFor[map[string]Float.X](),
	"Font.get_ot_name_strings.":                                         reflect.TypeFor[map[string]string](), // FIXME
	"Font.get_opentype_features.":                                       reflect.TypeFor[map[string]string](), // FIXME
	"Font.get_supported_feature_list.":                                  reflect.TypeFor[map[string]string](), // FIXME
	"Font.get_supported_variation_list.":                                reflect.TypeFor[map[string]string](), // FIXME
	"FontFile.set_variation_coordinates.variation_coordinates":          reflect.TypeFor[map[string]Float.X](),
	"FontFile.get_variation_coordinates.":                               reflect.TypeFor[map[string]Float.X](),
	"FontFile.set_opentype_feature_overrides.overrides":                 reflect.TypeFor[map[string]string](), // FIXME
	"FontFile.get_opentype_feature_overrides.":                          reflect.TypeFor[map[string]string](), // FIXME
	"FontVariation.set_variation_opentype.coords":                       reflect.TypeFor[map[any]Float.X](),   // FIXME
	"FontVariation.get_variation_opentype.":                             reflect.TypeFor[map[any]Float.X](),   // FIXME
	"FontVariation.set_opentype_features.features":                      reflect.TypeFor[map[string]string](), // FIXME
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
	"Object.get_property_list.":                                         reflect.TypeFor[PropertyInfo](),
	"Object.get_method_list.":                                           reflect.TypeFor[PropertyInfo](),
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
	"ProjectSettings.add_property_info.hint":                            reflect.TypeFor[PropertyInfo](),
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
	"Script.get_script_property_list.":                                  reflect.TypeFor[PropertyInfo](),
	"Script.get_script_method_list.":                                    reflect.TypeFor[PropertyInfo](),
	"Script.get_script_signal_list.":                                    reflect.TypeFor[SignalInfo](),
	"SyntaxHighlighter.get_line_syntax_highlighting.":                   reflect.TypeFor[map[int]Entry](),
	"TextServer.font_get_ot_name_strings.":                              reflect.TypeFor[map[string]string](), // FIXME
	"TextServer.font_set_variation_coordinates.variation_coordinates":   reflect.TypeFor[map[string]Float.X](),
	"TextServer.font_get_variation_coordinates.":                        reflect.TypeFor[map[string]Float.X](),
	"TextServer.font_get_glyph_contours.":                               reflect.TypeFor[map[string]any](),     // FIXME
	"TextServer.font_set_opentype_feature_overrides.overrides":          reflect.TypeFor[map[string]string](),  // FIXME
	"TextServer.font_get_opentype_feature_overrides.":                   reflect.TypeFor[map[string]string](),  // FIXME
	"TextServer.font_supported_feature_list.":                           reflect.TypeFor[map[string]string](),  // FIXME
	"TextServer.font_supported_variation_list.":                         reflect.TypeFor[map[string]string](),  // FIXME
	"TextServer.shaped_text_add_string.opentype_features":               reflect.TypeFor[map[string]string](),  // FIXME
	"TextServer.shaped_set_span_update_font.opentype_features":          reflect.TypeFor[map[string]string](),  // FIXME
	"TextServer.shaped_text_get_glyphs.":                                reflect.TypeFor[map[int]Vector2.XY](), // FIXME
	"TextServer.shaped_text_sort_logical.":                              reflect.TypeFor[map[int]int](),        // FIXME
	"TextServer.shaped_text_get_ellipsis_glyphs.":                       reflect.TypeFor[map[int]Vector2.XY](), // FIXME
	"TextServer.shaped_text_get_carets.":                                reflect.TypeFor[map[int]Vector2.XY](), // FIXME
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
	"Time.get_time_zone_from_system.":                                   reflect.TypeFor[map[string]string](),
	"TreeItem.get_range_config.":                                        reflect.TypeFor[RangeConfig](),
	"WebRTCMultiplayerPeer.get_peer.":                                   reflect.TypeFor[Conn](),
	"VisualShader.get_node_connections.":                                reflect.TypeFor[map[string]any](),
	"WebRTCMultiplayerPeer.get_peers.":                                  reflect.TypeFor[map[int]Conn](),
	"WebRTCPeerConnection.initialize.configuration":                     reflect.TypeFor[Configuration](),
	"WebRTCPeerConnection.create_data_channel.options":                  reflect.TypeFor[Options](),
	"XRInterface.get_system_info.":                                      reflect.TypeFor[map[string]any](),
	"XRServer.get_trackers.":                                            reflect.TypeFor[map[any]any](),
	"XRServer.get_interfaces.":                                          reflect.TypeFor[map[int]string](),
}
