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

#pragma once

#include "string"

namespace gamesdk {

// Get the value of the given system property
std::string GetSystemProp(const char* key, const char* default_value = "");

// Get the value of the given system property as an integer
int GetSystemPropAsInt(const char* key, int default_value = 0);

// Get the value of the given system property as a bool
bool GetSystemPropAsBool(const char* key, bool default_value = false);

}  // namespace gamesdk