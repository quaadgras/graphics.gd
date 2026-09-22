/*
 * Copyright 2018 The Android Open Source Project
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

// API entry points for both OpenGL and Vulkan

#include "swappy/swappy_common.h"

extern "C" {

void SWAPPY_VERSION_SYMBOL() {
    // Intentionally empty: this function is needed in order to interrogate
    // shared libraries for the Swappy version.
}

uint32_t Swappy_version() { return SWAPPY_PACKED_VERSION; }

const char* Swappy_versionString() {
  static const char version[] = AGDK_STRING_VERSION(
      SWAPPY_MAJOR_VERSION, SWAPPY_MINOR_VERSION, SWAPPY_BUGFIX_VERSION);
  return version;
}

}  // extern "C"
