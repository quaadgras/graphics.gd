/*
[gdscript]
# Create the objects.
var node = Node2D.new()
var body = RigidBody2D.new()
var collision = CollisionShape2D.new()

# Create the object hierarchy.
body.add_child(collision)
node.add_child(body)

# Change owner of `body`, but not of `collision`.
body.owner = node
var scene = PackedScene.new()

# Only `node` and `body` are now packed.
var result = scene.pack(node)
if result == OK:
    var error = ResourceSaver.save(scene, "res://path/name.tscn")  # Or "user://..."
    if error != OK:
        push_error("An error occurred while saving the scene to disk.")
[/gdscript]
[csharp]
// Create the objects.
var node = new Node2D();
var body = new RigidBody2D();
var collision = new CollisionShape2D();

// Create the object hierarchy.
body.AddChild(collision);
node.AddChild(body);

// Change owner of `body`, but not of `collision`.
body.Owner = node;
var scene = new PackedScene();

// Only `node` and `body` are now packed.
Error result = scene.Pack(node);
if (result == Error.Ok)
{
    Error error = ResourceSaver.Save(scene, "res://path/name.tscn"); // Or "user://..."
    if (error != Error.Ok)
    {
        GD.PushError("An error occurred while saving the scene to disk.");
    }
}
[/csharp]
*/

package main

import (
	"graphics.gd/classdb/CollisionShape2D"
	"graphics.gd/classdb/Engine"
	"graphics.gd/classdb/Node2D"
	"graphics.gd/classdb/PackedScene"
	"graphics.gd/classdb/ResourceSaver"
	"graphics.gd/classdb/RigidBody2D"
)

func ExamplePackedSceneSave() {
	var node = Node2D.New()
	var body = RigidBody2D.New()
	var collision = CollisionShape2D.New()
	body.AsNode().AddChild(collision.AsNode())
	node.AsNode().AddChild(body.AsNode())
	body.AsNode().SetOwner(node.AsNode()) // Change owner of `body`, but not of `collision`.
	var scene = PackedScene.New()
	var result = scene.Pack(node.AsNode()) // Only `node` and `body` are now packed.
	if result == nil {
		var err = ResourceSaver.Save(scene.AsResource(), "res://path/name.tscn", 0) // Or "user://..."
		if err != nil {
			Engine.Raise(err)
		}
	}
}
