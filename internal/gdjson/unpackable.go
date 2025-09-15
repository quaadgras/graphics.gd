package gdjson

import (
	"reflect"
)

var Unpackables = map[string][]reflect.Type{
	"ENetConnection.service": {
		TypeFromString("", "EventType"),
		TypeFromString("ENetPacketPeer", "Instance"),
		reflect.TypeFor[int](),
		reflect.TypeFor[int](),
	},
	"CSGShape3D.get_meshes": {
		TypeFromString("", "Transform3D.BasisOrigin"),
		TypeFromString("Mesh", "Instance"),
	},
	"Node.get_node_and_resource": {
		TypeFromString("", "Instance"),
		TypeFromString("Resource", "Instance"),
		reflect.TypeFor[string](),
	},
	"TileSet.get_coords_level_tile_proxy": {
		reflect.TypeFor[int](),
		TypeFromString("", "Vector2i.XY"),
	},
	"TileSet.get_alternative_level_tile_proxy": {
		reflect.TypeFor[int](),
		TypeFromString("", "Vector2i.XY"),
		reflect.TypeFor[int](),
	},
	"TileSet.map_tile_proxy": {
		reflect.TypeFor[int](),
		TypeFromString("", "Vector2i.XY"),
		reflect.TypeFor[int](),
	},
	"StreamPeer.put_partial_data": {
		reflect.TypeFor[error](),
		reflect.TypeFor[int](),
	},
	"StreamPeer.get_data": {
		reflect.TypeFor[error](),
		reflect.TypeFor[[]byte](),
	},
	"StreamPeer.get_partial_data": {
		reflect.TypeFor[error](),
		reflect.TypeFor[[]byte](),
	},
}
