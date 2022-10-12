extends Node2D

var hello: HelloWorld = HelloWorld.new()

# Called when the node enters the scene tree for the first time.
func _ready():
	hello.print()
	hello.echo("Hello from GDScript")
	print(Engine.get_singleton("DisplayServer"))
	pass # Replace with function body.

func _notification(what):
	if what == NOTIFICATION_PREDELETE:
		hello.free()
# Called every frame. 'delta' is the elapsed time since the previous frame.
func _process(delta):
	pass
