/*
 * Copyright 2021 The Android Open Source Project
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

#include "system_utils.h"

#include <android/api-level.h>
#include <stdlib.h>
#include <sys/system_properties.h>

namespace gamesdk {

#if __ANDROID_API__ >= 26
std::string getSystemPropViaCallback(const char* key,
                                     const char* default_value = "") {
    const prop_info* prop = __system_property_find(key);
    if (prop == nullptr) {
        return default_value;
    }
    std::string return_value;
    auto thunk = [](void* cookie, const char* /*name*/, const char* value,
                    uint32_t /*serial*/) {
        if (value != nullptr) {
            std::string* r = static_cast<std::string*>(cookie);
            *r = value;
        }
    };
    __system_property_read_callback(prop, thunk, &return_value);
    return return_value;
}
#else
std::string getSystemPropViaGet(const char* key,
                                const char* default_value = "") {
    char buffer[PROP_VALUE_MAX + 1] = "";  // +1 for terminator
    int bufferLen = __system_property_get(key, buffer);
    if (bufferLen > 0)
        return buffer;
    else
        return "";
}
#endif

std::string GetSystemProp(const char* key, const char* default_value) {
#if __ANDROID_API__ >= 26
    return getSystemPropViaCallback(key, default_value);
#else
    return getSystemPropViaGet(key, default_value);
#endif
}

int GetSystemPropAsInt(const char* key, int default_value) {
    std::string prop = GetSystemProp(key);
    return prop == "" ? default_value : strtoll(prop.c_str(), nullptr, 10);
}

bool GetSystemPropAsBool(const char* key, bool default_value) {
    return GetSystemPropAsInt(key, default_value) != 0;
}

}  // namespace gamesdk