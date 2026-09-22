/*
 * Copyright (C) 2015 The Android Open Source Project
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
 * @addtogroup Networking
 * @{
 */

/**
 * @file multinetwork.h
 */

#ifndef ANDROID_MULTINETWORK_H
#define ANDROID_MULTINETWORK_H

#include <netdb.h>
#include <stdlib.h>
#include <sys/cdefs.h>

__BEGIN_DECLS

/**
 * The corresponding C type for android.net.Network#getNetworkHandle() return
 * values.  The Java signed long value can be safely cast to a net_handle_t:
 *
 *     [C]    ((net_handle_t) java_long_network_handle)
 *     [C++]  static_cast<net_handle_t>(java_long_network_handle)
 *
 * as appropriate.
 */
typedef uint64_t net_handle_t;

/**
 * The value NETWORK_UNSPECIFIED indicates no specific network.
 *
 * For some functions (documented below), a previous binding may be cleared
 * by an invocation with NETWORK_UNSPECIFIED.
 *
 * Depending on the context it may indicate an error.  It is expressly
 * not used to indicate some notion of the "current default network".
 */
#define NETWORK_UNSPECIFIED  ((net_handle_t)0)


/**
 * All functions below that return an int return 0 on success or -1
 * on failure with an appropriate errno value set.
 */

/**
 * Set the network to be used by the given socket file descriptor.
 *
 * To clear a previous socket binding, invoke with NETWORK_UNSPECIFIED.
 *
 * This is the equivalent of: [android.net.Network#bindSocket()](https://developer.android.com/reference/android/net/Network.html#bindSocket(java.net.Socket))
 *
 * Available since API level 23.
 */
int android_setsocknetwork(net_handle_t network, int fd) __INTRODUCED_IN(23);


/**
 * Binds the current process to |network|.  All sockets created in the future
 * (and not explicitly bound via android_setsocknetwork()) will be bound to
 * |network|.  All host name resolutions will be limited to |network| as well.
 * Note that if the network identified by |network| ever disconnects, all
 * sockets created in this way will cease to work and all host name
 * resolutions will fail.  This is by design so an application doesn't
 * accidentally use sockets it thinks are still bound to a particular network.
 *
 * To clear a previous process binding, invoke with NETWORK_UNSPECIFIED.
 *
 * This is the equivalent of: [android.net.ConnectivityManager#bindProcessToNetwork()](https://developer.android.com/reference/android/net/ConnectivityManager.html#bindProcessToNetwork(android.net.Network))
 *
 * Available since API level 23.
 */
int android_setprocnetwork(net_handle_t network) __INTRODUCED_IN(23);


/**
 * Gets the |network| bound to the current process, as per android_setprocnetwork.
 *
 * This is the equivalent of: [android.net.ConnectivityManager#getBoundNetworkForProcess()](https://developer.android.com/reference/android/net/ConnectivityManager.html#getBoundNetworkForProcess(android.net.Network))
 * Returns 0 on success, or -1 setting errno to EINVAL if a null pointer is
 * passed in.
 *
 *
 * Available since API level 31.
 */
int android_getprocnetwork(net_handle_t *network) __INTRODUCED_IN(31);

/**
 * Binds domain name resolutions performed by this process to |network|.
 * android_setprocnetwork takes precedence over this setting.
 *
 * To clear a previous process binding, invoke with NETWORK_UNSPECIFIED.
 * On success 0 is returned. On error -1 is returned, and errno is set.
 *
 * Available since API level 31.
 */
int android_setprocdns(net_handle_t network) __INTRODUCED_IN(31);

/**
 * Gets the |network| to which domain name resolutions are bound on the
 * current process.
 *
 * Returns 0 on success, or -1 setting errno to EINVAL if a null pointer is
 * passed in.
 *
 * Available since API level 31.
 */
int android_getprocdns(net_handle_t *network) __INTRODUCED_IN(31);


/**
 * Perform hostname resolution via the DNS servers associated with |network|.
 *
 * All arguments (apart from |network|) are used identically as those passed
 * to getaddrinfo(3).  Return and error values are identical to those of
 * getaddrinfo(3), and in particular gai_strerror(3) can be used as expected.
 * Similar to getaddrinfo(3):
 *     - |hints| may be NULL (in which case man page documented defaults apply)
 *     - either |node| or |service| may be NULL, but not both
 *     - |res| must not be NULL
 *
 * This is the equivalent of: [android.net.Network#getAllByName()](https://developer.android.com/reference/android/net/Network.html#getAllByName(java.lang.String))
 *
 * Available since API level 23.
 */
int android_getaddrinfofornetwork(net_handle_t network,
        const char *node, const char *service,
        const struct addrinfo *hints, struct addrinfo **res) __INTRODUCED_IN(23);

/**
 * Possible values of the flags argument to android_res_nsend and android_res_nquery.
 * Values are ORed together.
 */
enum ResNsendFlags : uint32_t {
    /**
     * Send a single request to a single resolver and fail on timeout or network errors
     */
    ANDROID_RESOLV_NO_RETRY = 1 << 0,

    /**
     * Don't lookup this request in the cache, and don't cache the result of the lookup.
     * This flag implies {@link #ANDROID_RESOLV_NO_CACHE_LOOKUP}.
     */
    ANDROID_RESOLV_NO_CACHE_STORE = 1 << 1,

    /**
     * Don't lookup the request in cache.
     */
    ANDROID_RESOLV_NO_CACHE_LOOKUP = 1 << 2,
};

/**
 * Look up the {|ns_class|, |ns_type|} Resource Record (RR) associated
 * with Domain Name |dname| on the given |network|.
 * The typical value for |ns_class| is ns_c_in, while |type| can be any
 * record type (for instance, ns_t_aaaa or ns_t_txt).
 * |flags| is a additional config to control actual querying behavior, see
 * ResNsendFlags for detail.
 *
 * Returns a file descriptor to watch for read events, or a negative
 * POSIX error code (see errno.h) if an immediate error occurs.
 *
 * Available since API level 29.
 */
int android_res_nquery(net_handle_t network,
        const char *dname, int ns_class, int ns_type, uint32_t flags) __INTRODUCED_IN(29);

/**
 * Issue the query |msg| on the given |network|.
 * |flags| is a additional config to control actual querying behavior, see
 * ResNsendFlags for detail.
 *
 * Returns a file descriptor to watch for read events, or a negative
 * POSIX error code (see errno.h) if an immediate error occurs.
 *
 * Available since API level 29.
 */
int android_res_nsend(net_handle_t network,
        const uint8_t *msg, size_t msglen, uint32_t flags) __INTRODUCED_IN(29);

/**
 * Read a result for the query associated with the |fd| descriptor.
 * Closes |fd| before returning.
 *
 * Available since 29.
 *
 * Returns:
 *     < 0: negative POSIX error code (see errno.h for possible values). |rcode| is not set.
 *     >= 0: length of |answer|. |rcode| is the resolver return code (e.g., ns_r_nxdomain)
 */
int android_res_nresult(int fd,
        int *rcode, uint8_t *answer, size_t anslen) __INTRODUCED_IN(29);

/**
 * Attempts to cancel the in-progress query associated with the |nsend_fd|
 * descriptor.
 *
 * Available since API level 29.
 */
void android_res_cancel(int nsend_fd) __INTRODUCED_IN(29);

/*
 * Set the socket tag and owning UID for traffic statistics on the specified
 * socket.
 *
 * Subsequent calls always replace any existing parameters. The socket tag and
 * uid (if set) are kept when the socket is sent to another process using binder
 * IPCs or other mechanisms such as UNIX socket fd passing. Any app can accept
 * blame for future traffic performed on a socket originally created by another
 * app by calling this method with its own UID (or calling
 * android_tag_socket(int sockfd, int tag)). However, only apps holding the
 * android.Manifest.permission#UPDATE_DEVICE_STATS permission may assign blame
 * to another UIDs. If unset (default) the socket tag is 0, and the uid is the
 * socket creator's uid.
 *
 * Returns 0 on success, or a negative POSIX error code (see errno.h) on
 * failure.
 *
 * Available since API level 33.
 */
int android_tag_socket_with_uid(int sockfd, uint32_t tag, uid_t uid) __INTRODUCED_IN(33);

/*
 * Set the socket tag for traffic statistics on the specified socket.
 *
 * This function tags the socket with the caller's UID (accepting blame for
 * future traffic performed on this socket) even if the socket was originally
 * opened by another UID or was previously tagged by another UID. Subsequent
 * calls always replace any existing parameters. The socket tag is kept when the
 * socket is sent to another process using binder IPCs or other mechanisms such
 * as UNIX socket fd passing. The tag is a value defined by the caller and used
 * together with uid for data traffic accounting, so that the function callers
 * can account different types of data usage for a uid.
 *
 * Returns 0 on success, or a negative POSIX error code (see errno.h) on
 * failure.
 *
 * Some possible error codes:
 * -EBADF           Bad socketfd.
 * -EPERM           No permission.
 * -EAFNOSUPPORT    Socket family is neither AF_INET nor AF_INET6.
 * -EPROTONOSUPPORT Socket protocol is neither IPPROTO_UDP nor IPPROTO_TCP.
 * -EMFILE          Too many stats entries.
 * There are still other error codes that may provided by -errno of
 * [getsockopt()](https://man7.org/linux/man-pages/man2/getsockopt.2.html) or by
 * BPF maps read/write sys calls, which are set appropriately.
 *
 * Available since API level 33.
 */
int android_tag_socket(int sockfd, uint32_t tag) __INTRODUCED_IN(33);

/*
 * Untag a network socket.
 *
 * Future traffic on this socket will no longer be associated with any
 * previously configured tag and uid. If the socket was created by another UID
 * or was previously tagged by another UID, calling this function will clear the
 * statistics parameters, and thus the UID blamed for traffic on the socket will
 * be the UID that originally created the socket, even if the socket was
 * subsequently tagged by a different UID.
 *
 * Returns 0 on success, or a negative POSIX error code (see errno.h) on
 * failure.
 *
 * One of possible error code:
 * -EBADF           Bad socketfd.
 * Other error codes are either provided by -errno of
 * [getsockopt()](https://man7.org/linux/man-pages/man2/getsockopt.2.html) or by
 * BPF map element deletion sys call, which are set appropriately.
 *
 * Available since API level 33.
 */
int android_untag_socket(int sockfd) __INTRODUCED_IN(33);

__END_DECLS

#endif  // ANDROID_MULTINETWORK_H

/** @} */
