LLVM_CONFIG := /usr/lib/llvm/21/bin/llvm-config

CXX      := clang++
CXXFLAGS := -std=c++17 -fno-rtti -O2 $(shell $(LLVM_CONFIG) --cxxflags)
LDFLAGS  := $(shell $(LLVM_CONFIG) --ldflags) -lclang-cpp -lLLVM-21

TARGET := trivial_methods

all: $(TARGET)

$(TARGET): trivial_methods.cpp
	$(CXX) $(CXXFLAGS) -o $@ $< $(LDFLAGS)

clean:
	rm -f $(TARGET)

.PHONY: all clean
