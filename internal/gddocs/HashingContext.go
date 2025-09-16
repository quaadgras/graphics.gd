/*
[gdscript]
const CHUNK_SIZE = 1024

func hash_file(path):
	# Check that file exists.
	if not FileAccess.file_exists(path):
		return
	# Start an SHA-256 context.
	var ctx = HashingContext.new()
	ctx.start(HashingContext.HASH_SHA256)
	# Open the file to hash.
	var file = FileAccess.open(path, FileAccess.READ)
	# Update the context after reading each chunk.
	while file.get_position() < file.get_length():
		var remaining = file.get_length() - file.get_position()
		ctx.update(file.get_buffer(min(remaining, CHUNK_SIZE)))
	# Get the computed hash.
	var res = ctx.finish()
	# Print the result as hex string and array.
	printt(res.hex_encode(), Array(res))
[/gdscript]
[csharp]
public const int ChunkSize = 1024;

public void HashFile(string path)
{
	// Check that file exists.
	if (!FileAccess.FileExists(path))
	{
		return;
	}
	// Start an SHA-256 context.
	var ctx = new HashingContext();
	ctx.Start(HashingContext.HashType.Sha256);
	// Open the file to hash.
	using var file = FileAccess.Open(path, FileAccess.ModeFlags.Read);
	// Update the context after reading each chunk.
	while (file.GetPosition() < file.GetLength())
	{
		int remaining = (int)(file.GetLength() - file.GetPosition());
		ctx.Update(file.GetBuffer(Mathf.Min(remaining, ChunkSize)));
	}
	// Get the computed hash.
	byte[] res = ctx.Finish();
	// Print the result as hex string and array.
	GD.PrintT(res.HexEncode(), (Variant)res);
}
[/csharp]
*/

package main

import (
	"encoding/hex"
	"fmt"

	"graphics.gd/classdb/FileAccess"
	"graphics.gd/classdb/HashingContext"
)

const ChunkSize = 1024

func HashFile(path string) {
	if !FileAccess.FileExists(path) {
		return
	}
	var ctx = HashingContext.New()
	ctx.Start(HashingContext.HashSha256)
	var file = FileAccess.Open(path, FileAccess.Read)
	for file.GetPosition() < file.GetLength() {
		remaining := int(file.GetLength() - file.GetPosition())
		ctx.Update(file.GetBuffer(min(remaining, ChunkSize)))
	}
	res := ctx.Finish()
	fmt.Println(hex.EncodeToString(res), res)
}
