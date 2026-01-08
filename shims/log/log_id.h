#ifndef _LOG_LOG_ID_H
#define _LOG_LOG_ID_H

#include <android/log.h>

#ifdef __cplusplus
extern "C" {
#endif

static inline log_id_t android_name_to_log_id(const char* logName) {
    return LOG_ID_MAIN;
}

static inline const char* android_log_id_to_name(log_id_t log_id) {
    return "main";
}

#ifdef __cplusplus
}
#endif

#endif
