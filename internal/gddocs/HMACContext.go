/*
[gdscript]
extends Node
var ctx = HMACContext.new()

func _ready():
	var key = "supersecret".to_utf8_buffer()
	var err = ctx.start(HashingContext.HASH_SHA256, key)
	assert(err == OK)
	var msg1 = "this is ".to_utf8_buffer()
	var msg2 = "super duper secret".to_utf8_buffer()
	err = ctx.update(msg1)
	assert(err == OK)
	err = ctx.update(msg2)
	assert(err == OK)
	var hmac = ctx.finish()
	print(hmac.hex_encode())

[/gdscript]
[csharp]
using Godot;
using System.Diagnostics;

public partial class MyNode : Node
{
	private HmacContext _ctx = new HmacContext();

	public override void _Ready()
	{
		byte[] key = "supersecret".ToUtf8Buffer();
		Error err = _ctx.Start(HashingContext.HashType.Sha256, key);
		Debug.Assert(err == Error.Ok);
		byte[] msg1 = "this is ".ToUtf8Buffer();
		byte[] msg2 = "super duper secret".ToUtf8Buffer();
		err = _ctx.Update(msg1);
		Debug.Assert(err == Error.Ok);
		err = _ctx.Update(msg2);
		Debug.Assert(err == Error.Ok);
		byte[] hmac = _ctx.Finish();
		GD.Print(hmac.HexEncode());
	}
}
[/csharp]
*/

package main

import (
	"encoding/hex"
	"fmt"

	"graphics.gd/classdb/HMACContext"
	"graphics.gd/classdb/HashingContext"
	"graphics.gd/classdb/Node"
)

type MyHMAC struct {
	Node.Extension[MyHMAC]

	ctx HMACContext.Instance
}

func (n *MyHMAC) Ready() {
	n.ctx = HMACContext.New()

	var key = []byte("supersecret")
	var err = n.ctx.Start(HashingContext.HashSha256, key)
	if err != nil {
		panic(err)
	}
	var msg1 = []byte("this is ")
	var msg2 = []byte("super duper secret")
	err = n.ctx.Update(msg1)
	if err != nil {
		panic(err)
	}
	err = n.ctx.Update(msg2)
	if err != nil {
		panic(err)
	}
	var hmac = n.ctx.Finish()
	fmt.Println(hex.EncodeToString(hmac))
}
