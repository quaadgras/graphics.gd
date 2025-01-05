extends Converter

func _init():
    if HelloWorld() != "Hello, World!":
        push_error("HelloWorld() failed")
        Fail()
    if Echo("1234") != "1234":
        push_error("Echo() failed")
        Fail()
    if Int() != 22:
        push_error("Int() failed")
        Fail()
    if Float() != 2.2:
        push_error("Float() failed")
        Fail()
    if Bool() != true:
        push_error("Bool() failed")
        Fail()
    Done()
