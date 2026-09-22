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

#pragma once

#include <jni.h>

#include <cstdio>
#include <fstream>
#include <sstream>
#include <string>

#include "Log.h"
#include "system_utils.h"

// The code in this file will dynamically load Java classes from a binary
// resource linked to the library. The binary data is in DEX format, accessible
// using the _binary_classes_dex_start and _binary_classes_dex_end linker
// symbols.

// If you make AGDK classes available on the Java classpath, e.g. by adding them
// to your game/engine's own Java component, you do not need to add the binary
// resource and can instead define ANDROIDGAMESDK_NO_BINARY_DEX_LINKAGE which
// will avoid the linker requiring these symbols.

#ifndef ANDROIDGAMESDK_NO_BINARY_DEX_LINKAGE
extern const char _binary_classes_dex_start;
extern const char _binary_classes_dex_end;
#endif

namespace gamesdk {

#ifndef ANDROIDGAMESDK_NO_BINARY_DEX_LINKAGE

static bool saveBytesToFile(std::string fileName, const char* bytes,
                            size_t size) {
  std::ofstream save_file(fileName, std::ios::binary);
  if (save_file.good()) {
    save_file.write(bytes, size);
    return true;
  }
  return false;
}

static bool deleteFile(std::string fileName) {
  if (remove(fileName.c_str()) != 0)
    return false;
  else
    return true;
}

static bool createTempFile(JNIEnv* env, jobject activity, const char* ext,
                           std::string& tempFileName) {
  bool result = false;
  jclass activityClass = env->GetObjectClass(activity);
  jmethodID getCacheDir =
      env->GetMethodID(activityClass, "getCacheDir", "()Ljava/io/File;");
  jobject cacheDir = env->CallObjectMethod(activity, getCacheDir);
  if (env->ExceptionCheck()) {
    env->ExceptionDescribe();
    env->ExceptionClear();
  } else {
    jclass fileClass = env->FindClass("java/io/File");
    jmethodID createTempFile =
        env->GetStaticMethodID(fileClass, "createTempFile",
                               "(Ljava/lang/String;Ljava/lang/String;Ljava/"
                               "io/File;)Ljava/io/File;");
    jstring prefix = env->NewStringUTF("ags");
    jstring suffix = env->NewStringUTF(ext);
    jobject tempFile = env->CallStaticObjectMethod(fileClass, createTempFile,
                                                   prefix, suffix, cacheDir);
    if (env->ExceptionCheck()) {
      env->ExceptionDescribe();
      env->ExceptionClear();
    } else {
      jmethodID getPath =
          env->GetMethodID(fileClass, "getPath", "()Ljava/lang/String;");
      jstring pathString = (jstring)env->CallObjectMethod(tempFile, getPath);
      if (env->ExceptionCheck()) {
        env->ExceptionDescribe();
        env->ExceptionClear();
      } else {
        const char* path = env->GetStringUTFChars(pathString, NULL);
        tempFileName.assign(path);
        env->ReleaseStringUTFChars(pathString, path);
        result = true;
      }
    }
    env->DeleteLocalRef(prefix);
    env->DeleteLocalRef(suffix);
  }
  return result;
}

#endif  // #ifndef ANDROIDGAMESDK_NO_BINARY_DEX_LINKAGE

static jclass loadClass(JNIEnv* env, jobject activity, const char* name,
                        JNINativeMethod* nativeMethods,
                        size_t nativeMethodsSize) {
  /*
   *   1. Get a classloader from actvity
   *   2. Try to create the requested class from the activity classloader
   *   3. If step 2 not successful then get a classloder for dex bytes (in
   * memory or file)
   *   4. If step 3 is  successful then register native methods
   */
  if (!env || !activity || !name) {
    return nullptr;
  }
  jclass activityClass = env->GetObjectClass(activity);
  jclass classLoaderClass = env->FindClass("java/lang/ClassLoader");
  jmethodID getClassLoader = env->GetMethodID(activityClass, "getClassLoader",
                                              "()Ljava/lang/ClassLoader;");
  jobject classLoaderObj = env->CallObjectMethod(activity, getClassLoader);
  jmethodID loadClass = env->GetMethodID(
      classLoaderClass, "loadClass", "(Ljava/lang/String;)Ljava/lang/Class;");
  jstring className = env->NewStringUTF(name);
  jclass targetClass = static_cast<jclass>(
      env->CallObjectMethod(classLoaderObj, loadClass, className));
  if (env->ExceptionCheck()) {
    env->ExceptionClear();

#ifdef ANDROIDGAMESDK_NO_BINARY_DEX_LINKAGE
    ALOGE(
        "Couldn't find class %s and ANDROIDGAMESDK_NO_BINARY_DEX_LINKAGE "
        "is defined",
        name);
#else
    jstring dexLoaderClassName =
        env->NewStringUTF("dalvik/system/InMemoryDexClassLoader");
    jclass imclassloaderClass = static_cast<jclass>(
        env->CallObjectMethod(classLoaderObj, loadClass, dexLoaderClassName));
    env->DeleteLocalRef(dexLoaderClassName);

    int sdkVersion = gamesdk::GetSystemPropAsInt("ro.build.version.sdk");

    if (env->ExceptionCheck() || !imclassloaderClass || sdkVersion < 26) {
      env->ExceptionClear();
      // For older SDK versions <26, where InMemoryDexClassLoader is not
      // available
      dexLoaderClassName = env->NewStringUTF("dalvik/system/PathClassLoader");
      imclassloaderClass = static_cast<jclass>(
          env->CallObjectMethod(classLoaderObj, loadClass, dexLoaderClassName));
      env->DeleteLocalRef(dexLoaderClassName);
      if (env->ExceptionCheck() || !imclassloaderClass) {
        env->ExceptionDescribe();
        env->ExceptionClear();
        ALOGE("Unable to find dalvik/system/PathClassLoader.");
        targetClass = nullptr;
      } else {
        jmethodID constructor =
            env->GetMethodID(imclassloaderClass, "<init>",
                             "(Ljava/lang/String;Ljava/lang/ClassLoader;)V");
        std::string tempPath;
        if (!createTempFile(env, activity, ".dex", tempPath)) {
          ALOGE(
              "Unable to create a temporary file to store DEX with "
              "Java classes.");
        } else {
          size_t dex_file_size =
              (size_t)(&_binary_classes_dex_end - &_binary_classes_dex_start);
          if (!saveBytesToFile(tempPath, &_binary_classes_dex_start,
                               dex_file_size)) {
            ALOGE("Unable to write to %s file.", tempPath.c_str());
          } else {
            jstring dexPathString = env->NewStringUTF(tempPath.c_str());
            jobject imclassloaderObj = env->NewObject(
                imclassloaderClass, constructor, dexPathString, classLoaderObj);
            env->DeleteLocalRef(dexPathString);
            targetClass = static_cast<jclass>(
                env->CallObjectMethod(imclassloaderObj, loadClass, className));
            if (env->ExceptionCheck()) {
              env->ExceptionDescribe();
              env->ExceptionClear();
              ALOGE("Unable to find %s class", name);
            } else {
              env->RegisterNatives(targetClass, nativeMethods,
                                   nativeMethodsSize);
              ALOGI("Using internal %s class from dex bytes.", name);
            }
            if (imclassloaderObj) {
              env->DeleteLocalRef(imclassloaderObj);
            }
          }
          deleteFile(tempPath);
        }
      }
    } else {
      jmethodID constructor =
          env->GetMethodID(imclassloaderClass, "<init>",
                           "(Ljava/nio/ByteBuffer;Ljava/lang/ClassLoader;)V");

      size_t dex_file_size =
          (size_t)(&_binary_classes_dex_end - &_binary_classes_dex_start);
      auto byteBuffer = env->NewDirectByteBuffer(
          (void*)&_binary_classes_dex_start, dex_file_size);
      jobject imclassloaderObj = env->NewObject(imclassloaderClass, constructor,
                                                byteBuffer, classLoaderObj);

      targetClass = static_cast<jclass>(
          env->CallObjectMethod(imclassloaderObj, loadClass, className));
      if (env->ExceptionCheck()) {
        env->ExceptionDescribe();
        env->ExceptionClear();
        ALOGE("Unable to find %s class", name);
      } else {
        env->RegisterNatives(targetClass, nativeMethods, nativeMethodsSize);
        ALOGI("Using internal %s class from dex bytes.", name);
      }
      if (imclassloaderObj) {
        env->DeleteLocalRef(imclassloaderObj);
      }
    }
    if (imclassloaderClass) {
      env->DeleteLocalRef(imclassloaderClass);
    }
#endif  // #ifdef ANDROIDGAMESDK_NO_BINARY_DEX_LINKAGE
  } else {
    // Register the native methods loaded by the classLoaderClass
    env->RegisterNatives(targetClass, nativeMethods, nativeMethodsSize);
  }
  env->DeleteLocalRef(className);
  return targetClass;
}

}  // namespace gamesdk
