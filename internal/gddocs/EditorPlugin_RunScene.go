/*
func _run_scene(scene, args):
	args.append("--an-extra-argument")
	return args
*/

package main

func (MyEditorPlugin) RunScene(scene string, args []string) []string {
	args = append(args, "--an-extra-argument")
	return args
}
