The files in this directory are copied from the Android Open Source Project:

	https://android.googlesource.com/platform/bionic @ 731631f300090436d7f5df80d50b6275c8c60a93
	https://android.googlesource.com/platform/frameworks/native @ 4f463a6b1de9198963dc6aff74154a504ba3f8f6
	https://android.googlesource.com/platform/frameworks/base @ 1cdfff555f4a21f71ccc978290e2e212e2f8b168
	https://android.googlesource.com/platform/frameworks/av @ e2f098935447ca4945946de5cb69db843fe3f003
	https://android.googlesource.com/platform/frameworks/wilhelm @ 5674f27e4c8333495518d07a676d0664d92fd74d
	https://android.googlesource.com/platform/libnativehelper @ aef2939781fc0b57b4477df7160935cdf5697919
	https://android.googlesource.com/platform/system/logging @ bcac7c30d88a3773a7c0bc9f5617a23a886331fd
	https://android.googlesource.com/platform/external/zlib @ 46c6da99965067627e6c078e197106988d57d4ff

regenerate them with: go run ./cmd/gd/internal/builder/bundled/gen-android
