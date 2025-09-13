/*
# Apply IK effect automatically on every new frame (not the current)
skeleton_ik_node.start()

# Apply IK effect only on the current frame
skeleton_ik_node.start(true)

# Stop IK effect and reset bones_global_pose_override on Skeleton
skeleton_ik_node.stop()

# Apply full IK effect
skeleton_ik_node.set_influence(1.0)

# Apply half IK effect
skeleton_ik_node.set_influence(0.5)

# Apply zero IK effect (a value at or below 0.01 also removes bones_global_pose_override on Skeleton)
skeleton_ik_node.set_influence(0.0)
*/

package main

import "graphics.gd/classdb/SkeletonIK3D"

func ExampleSkeletonIK3D(skeleton_ik_node SkeletonIK3D.Instance) {
	// Apply IK effect automatically on every new frame (not the current)
	skeleton_ik_node.Start()

	// Apply IK effect only on the current frame
	SkeletonIK3D.Expanded(skeleton_ik_node).Start(true)

	// Stop IK effect and reset bones_global_pose_override on Skeleton
	skeleton_ik_node.Stop()

	// Apply full IK effect
	skeleton_ik_node.AsSkeletonModifier3D().SetInfluence(1.0)

	// Apply half IK effect
	skeleton_ik_node.AsSkeletonModifier3D().SetInfluence(0.5)

	// Apply zero IK effect (a value at or below 0.01 also removes bones_global_pose_override on Skeleton)
	skeleton_ik_node.AsSkeletonModifier3D().SetInfluence(0.0)
}
