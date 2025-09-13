/*
[gdscript]
class_name MyAStar3D
extends AStar3D

func _compute_cost(u, v):
    var u_pos = get_point_position(u)
    var v_pos = get_point_position(v)
    return abs(u_pos.x - v_pos.x) + abs(u_pos.y - v_pos.y) + abs(u_pos.z - v_pos.z)

func _estimate_cost(u, v):
    var u_pos = get_point_position(u)
    var v_pos = get_point_position(v)
    return abs(u_pos.x - v_pos.x) + abs(u_pos.y - v_pos.y) + abs(u_pos.z - v_pos.z)
[/gdscript]
[csharp]
using Godot;

[GlobalClass]
public partial class MyAStar3D : AStar3D
{
    public override float _ComputeCost(long fromId, long toId)
    {
        Vector3 fromPoint = GetPointPosition(fromId);
        Vector3 toPoint = GetPointPosition(toId);

        return Mathf.Abs(fromPoint.X - toPoint.X) + Mathf.Abs(fromPoint.Y - toPoint.Y) + Mathf.Abs(fromPoint.Z - toPoint.Z);
    }

    public override float _EstimateCost(long fromId, long toId)
    {
        Vector3 fromPoint = GetPointPosition(fromId);
        Vector3 toPoint = GetPointPosition(toId);
        return Mathf.Abs(fromPoint.X - toPoint.X) + Mathf.Abs(fromPoint.Y - toPoint.Y) + Mathf.Abs(fromPoint.Z - toPoint.Z);
    }
}
[/csharp]
*/

package main

import (
	"graphics.gd/classdb/AStar3D"
	"graphics.gd/variant/Float"
)

type MyAStar3D struct {
	AStar3D.Extension[MyAStar3D]
}

func (astar *MyAStar3D) ComputeCost(u, v AStar3D.Point) Float.X {
	var u_pos = astar.AsAStar3D().GetPointPosition(u)
	var v_pos = astar.AsAStar3D().GetPointPosition(v)
	return Float.Abs(u_pos.X-v_pos.X) + Float.Abs(u_pos.Y-v_pos.Y) + Float.Abs(u_pos.Z-v_pos.Z)
}

func (astar *MyAStar3D) EstimateCost(u, v AStar3D.Point) Float.X {
	var u_pos = astar.AsAStar3D().GetPointPosition(u)
	var v_pos = astar.AsAStar3D().GetPointPosition(v)
	return Float.Abs(u_pos.X-v_pos.X) + Float.Abs(u_pos.Y-v_pos.Y) + Float.Abs(u_pos.Z-v_pos.Z)
}
