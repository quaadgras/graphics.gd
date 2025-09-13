/*
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
*/

package main

import (
	"graphics.gd/classdb/Node"
	"graphics.gd/classdb/WorkerThreadPool"
)

var enemies []Node.Instance

func process_enemy_ai(enemy_index int) {
	var processed_enemy = enemies[enemy_index]
	// Expensive logic...
	_ = processed_enemy
}

func process() {
	var task_id = WorkerThreadPool.AddGroupTask(process_enemy_ai, len(enemies), false, "")
	// Other code...
	WorkerThreadPool.WaitForGroupTaskCompletion(task_id)
	// Other code that depends on the enemy AI already being processed.
}
