ifeq ($(OS),Windows_NT)
   TARGET = example.dll
else
	UNAME_S := $(shell uname -s)
	ifeq ($(UNAME_S),Linux)
		TARGET = example.so
	endif
	ifeq ($(UNAME_S),Darwin)
		TARGET = example.framework
	endif
   
endif

all:
	go build -o ../bin/$(TARGET) -buildmode=c-shared
