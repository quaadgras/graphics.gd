#ifndef _SYS_CDEFS_H
#define _SYS_CDEFS_H

#define __BEGIN_DECLS extern "C" {
#define __END_DECLS }

#ifndef __GNUC_PREREQ
#define __GNUC_PREREQ(maj, min) ((__GNUC__ << 16) + __GNUC_MINOR__ >= ((maj) << 16) + (min))
#endif

#ifndef __printflike
#define __printflike(x, y) __attribute__((__format__(__printf__, x, y)))
#endif

#ifndef __scanflike
#define __scanflike(x, y) __attribute__((__format__(__scanf__, x, y)))
#endif

#ifndef __BIONIC__
#define __INTRODUCED_IN(x)
#define __INTRODUCED_IN_32(x)
#define __INTRODUCED_IN_64(x)
#define __DEPRECATED_IN(x)
#define __REMOVED_IN(x)
#endif

#ifndef __attribute_pure__
#define __attribute_pure__ __attribute__((__pure__))
#endif

#ifndef __attribute_const__
#define __attribute_const__ __attribute__((__const__))
#endif

#ifndef __wur
#define __wur __attribute__((__warn_unused_result__))
#endif

#ifndef __CONCAT
#define __CONCAT(x,y) x ## y
#endif

#ifndef __STRING
#define __STRING(x) #x
#endif

#endif
