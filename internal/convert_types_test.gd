extends Converter

func _init():
    if HelloWorld() != "Hello, World!": fail("HelloWorld()")
    if Echo("1234") != "1234":          fail("Echo()")
    if Int() != 22:                     fail("Int()")
    if Float() != 2.2:                  fail("Float()")
    if Bool() != true:                  fail("Bool()")
    if Int8() != 8:                     fail("Int8()")
    if Int16() != 16:                   fail("Int16()")
    if Int32() != 32:                   fail("Int32()")
    if Int64() != 64:                   fail("Int64()")
    if Uint() != 22:                    fail("UInt()")
    if Uint8() != 8:                    fail("UInt8()")
    if Uint16() != 16:                  fail("UInt16()")
    if Uint32() != 32:                  fail("UInt32()")
    Done()

func fail(name: String):
    push_error(name + " failed")
    Fail()
