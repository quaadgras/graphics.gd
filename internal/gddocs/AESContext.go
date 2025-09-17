/*
[gdscript]
extends Node

var aes = AESContext.new()

func _ready():
	var key = "My secret key!!!" # Key must be either 16 or 32 bytes.
	var data = "My secret text!!" # Data size must be multiple of 16 bytes, apply padding if needed.
	# Encrypt ECB
	aes.start(AESContext.MODE_ECB_ENCRYPT, key.to_utf8_buffer())
	var encrypted = aes.update(data.to_utf8_buffer())
	aes.finish()
	# Decrypt ECB
	aes.start(AESContext.MODE_ECB_DECRYPT, key.to_utf8_buffer())
	var decrypted = aes.update(encrypted)
	aes.finish()
	# Check ECB
	assert(decrypted == data.to_utf8_buffer())

	var iv = "My secret iv!!!!" # IV must be of exactly 16 bytes.
	# Encrypt CBC
	aes.start(AESContext.MODE_CBC_ENCRYPT, key.to_utf8_buffer(), iv.to_utf8_buffer())
	encrypted = aes.update(data.to_utf8_buffer())
	aes.finish()
	# Decrypt CBC
	aes.start(AESContext.MODE_CBC_DECRYPT, key.to_utf8_buffer(), iv.to_utf8_buffer())
	decrypted = aes.update(encrypted)
	aes.finish()
	# Check CBC
	assert(decrypted == data.to_utf8_buffer())
[/gdscript]
[csharp]
using Godot;
using System.Diagnostics;

public partial class MyNode : Node
{
	private AesContext _aes = new AesContext();

	public override void _Ready()
	{
		string key = "My secret key!!!"; // Key must be either 16 or 32 bytes.
		string data = "My secret text!!"; // Data size must be multiple of 16 bytes, apply padding if needed.
		// Encrypt ECB
		_aes.Start(AesContext.Mode.EcbEncrypt, key.ToUtf8Buffer());
		byte[] encrypted = _aes.Update(data.ToUtf8Buffer());
		_aes.Finish();
		// Decrypt ECB
		_aes.Start(AesContext.Mode.EcbDecrypt, key.ToUtf8Buffer());
		byte[] decrypted = _aes.Update(encrypted);
		_aes.Finish();
		// Check ECB
		Debug.Assert(decrypted == data.ToUtf8Buffer());

		string iv = "My secret iv!!!!"; // IV must be of exactly 16 bytes.
		// Encrypt CBC
		_aes.Start(AesContext.Mode.EcbEncrypt, key.ToUtf8Buffer(), iv.ToUtf8Buffer());
		encrypted = _aes.Update(data.ToUtf8Buffer());
		_aes.Finish();
		// Decrypt CBC
		_aes.Start(AesContext.Mode.EcbDecrypt, key.ToUtf8Buffer(), iv.ToUtf8Buffer());
		decrypted = _aes.Update(encrypted);
		_aes.Finish();
		// Check CBC
		Debug.Assert(decrypted == data.ToUtf8Buffer());
	}
}
[/csharp]
*/

package main

import (
	"graphics.gd/classdb/AESContext"
	"graphics.gd/classdb/Node"
)

type MyEncryption struct {
	Node.Extension[MyEncryption]

	AES AESContext.Instance
}

func NewEncryptionNode() *MyEncryption {
	return &MyEncryption{
		AES: AESContext.New(),
	}
}

func (node *MyEncryption) Ready() {
	var key = "My secret key!!!"  // Key must be either 16 or 32 bytes.
	var data = "My secret text!!" // Data size must be a multiple of 16 bytes, apply padding if needed.
	// Encrypt ECB
	node.AES.Start(AESContext.ModeEcbEncrypt, []byte(key))
	var encrypted = node.AES.Update([]byte(data))
	node.AES.Finish()
	// Decrypt ECB
	node.AES.Start(AESContext.ModeEcbDecrypt, []byte(key))
	var decrypted = node.AES.Update(encrypted)
	node.AES.Finish()
	if string(decrypted) != data {
		panic("decrypted data does not match original")
	}

	var iv = "My secret iv!!!!" // IV must be exactly 16 bytes.
	// Encrypt CBC
	node.AES.MoreArgs().Start(AESContext.ModeCbcEncrypt, []byte(key), []byte(iv))
	encrypted = node.AES.Update([]byte(data))
	node.AES.Finish()
	// Decrypt CBC
	node.AES.MoreArgs().Start(AESContext.ModeCbcDecrypt, []byte(key), []byte(iv))
	decrypted = node.AES.Update(encrypted)
	node.AES.Finish()
	if string(decrypted) != data {
		panic("decrypted data does not match original")
	}
}
