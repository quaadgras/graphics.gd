#include <sys/mman.h>
#include <unistd.h>
#include <cstring>

namespace android::incfs {

class IncFsFileMap {
public:
    IncFsFileMap() = default;
    ~IncFsFileMap() {
        if (data_ && length_ > 0) {
            munmap(const_cast<void*>(data_), length_);
        }
    }

    IncFsFileMap(IncFsFileMap&& other) noexcept
        : data_(other.data_), length_(other.length_), offset_(other.offset_) {
        other.data_ = nullptr;
        other.length_ = 0;
    }

    IncFsFileMap& operator=(IncFsFileMap&& other) noexcept {
        if (this != &other) {
            if (data_ && length_ > 0) munmap(const_cast<void*>(data_), length_);
            data_ = other.data_;
            length_ = other.length_;
            offset_ = other.offset_;
            other.data_ = nullptr;
            other.length_ = 0;
        }
        return *this;
    }

    static IncFsFileMap Create(int fd, off_t offset, size_t length, const char* name, bool verify = false) {
        IncFsFileMap map;
        void* ptr = mmap(nullptr, length, PROT_READ, MAP_PRIVATE, fd, offset);
        if (ptr != MAP_FAILED) {
            map.data_ = ptr;
            map.length_ = length;
            map.offset_ = offset;
        }
        return map;
    }

    const void* unsafe_data() const { return data_; }
    size_t length() const { return length_; }
    off_t offset() const { return offset_; }
    const char* file_name() const { return ""; }

private:
    const void* data_ = nullptr;
    size_t length_ = 0;
    off_t offset_ = 0;
};

} // namespace android::incfs
