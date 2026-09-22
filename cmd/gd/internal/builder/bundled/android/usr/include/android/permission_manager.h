/*
 * Copyright (C) 2020 The Android Open Source Project
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

/**
 * Structures and functions related to permission checks in native code.
 *
 * @addtogroup Permission
 * @{
 */

/**
 * @file permission_manager.h
 */

#ifndef ANDROID_PERMISSION_MANAGER_H
#define ANDROID_PERMISSION_MANAGER_H

#include <sys/cdefs.h>
#include <sys/types.h>

__BEGIN_DECLS

/**
 * Permission check results.
 *
 * Introduced in API 31.
 */
enum {
    /**
     * This is returned by APermissionManager_checkPermission()
     * if the permission has been granted to the given package.
     */
    PERMISSION_MANAGER_PERMISSION_GRANTED = 0,
    /**
     * This is returned by APermissionManager_checkPermission()
     * if the permission has not been granted to the given package.
     */
    PERMISSION_MANAGER_PERMISSION_DENIED = -1,
};

/**
 * Permission check return status values.
 *
 * Introduced in API 31.
 */
enum {
    /**
     * This is returned if the permission check completed without errors.
     * The output result is valid and contains one of {::PERMISSION_MANAGER_PERMISSION_GRANTED,
     * ::PERMISSION_MANAGER_PERMISSION_DENIED}.
     */
    PERMISSION_MANAGER_STATUS_OK = 0,
    /**
     * This is returned if the permission check encountered an unspecified error.
     * The output result is unmodified.
     */
    PERMISSION_MANAGER_STATUS_ERROR_UNKNOWN = -1,
    /**
     * This is returned if the permission check failed because the service is
     * unavailable. The output result is unmodified.
     */
    PERMISSION_MANAGER_STATUS_SERVICE_UNAVAILABLE = -2,
};

/**
 * Checks whether the package with the given pid/uid has been granted a permission.
 *
 * Note that the Java API of Context#checkPermission() is usually faster due to caching,
 * thus is preferred over this API wherever possible.
 *
 * @param permission the permission to be checked.
 * @param pid the process id of the package to be checked.
 * @param uid the uid of the package to be checked.
 * @param outResult output of the permission check result.
 *
 * @return error codes if any error happened during the check.
 */
int32_t APermissionManager_checkPermission(const char* permission,
                                           pid_t pid,
                                           uid_t uid,
                                           int32_t* outResult) __INTRODUCED_IN(31);

__END_DECLS

#endif  // ANDROID_PERMISSION_MANAGER_H

/** @} */
