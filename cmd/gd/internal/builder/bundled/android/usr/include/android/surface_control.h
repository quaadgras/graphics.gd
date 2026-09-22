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

/**
 * @addtogroup NativeActivity Native Activity
 * @{
 */

/**
 * @file surface_control.h
 */

#ifndef ANDROID_SURFACE_CONTROL_H
#define ANDROID_SURFACE_CONTROL_H

#include <sys/cdefs.h>

#include <android/display_luts.h>
#include <android/choreographer.h>
#include <android/data_space.h>
#include <android/hardware_buffer.h>
#include <android/hdr_metadata.h>
#include <android/native_window.h>

__BEGIN_DECLS

struct ASurfaceControl;

/**
 * The SurfaceControl API can be used to provide a hierarchy of surfaces for
 * composition to the system compositor. ASurfaceControl represents a content node in
 * this hierarchy.
 */
typedef struct ASurfaceControl ASurfaceControl;

/**
 * Creates an ASurfaceControl with either ANativeWindow or an ASurfaceControl as its parent.
 * \a debug_name is a debug name associated with this surface. It can be used to
 * identify this surface in the SurfaceFlinger's layer tree. It must not be
 * null.
 *
 * The caller takes ownership of the ASurfaceControl returned and must release it
 * using ASurfaceControl_release below.
 *
 * By default the \a ASurfaceControl will be visible and display any buffer submitted. In
 * addition, the default buffer submission control may release and not display all buffers
 * that are submitted before receiving a callback for the previous buffer. See
 * \a ASurfaceTransaction_setVisibility and \a ASurfaceTransaction_setEnableBackPressure to
 * change the default behaviors after creation.
 *
 * Available since API level 29.
 */
ASurfaceControl* _Nullable ASurfaceControl_createFromWindow(ANativeWindow* _Nonnull parent,
                                                            const char* _Nonnull debug_name)
        __INTRODUCED_IN(29);

/**
 * See ASurfaceControl_createFromWindow.
 *
 * Available since API level 29.
 */
ASurfaceControl* _Nullable ASurfaceControl_create(ASurfaceControl* _Nonnull parent,
                                                  const char* _Nonnull debug_name)
        __INTRODUCED_IN(29);

/**
 * Acquires a reference on the given ASurfaceControl object.  This prevents the object
 * from being deleted until the reference is removed.
 *
 * To release the reference, use the ASurfaceControl_release function.
 *
 * Available since API level 31.
 */
void ASurfaceControl_acquire(ASurfaceControl* _Nonnull surface_control) __INTRODUCED_IN(31);

/**
 * Removes a reference that was previously acquired with one of the following functions:
 *   ASurfaceControl_createFromWindow
 *   ASurfaceControl_create
 *   ANativeWindow_acquire
 * The surface and its children may remain on display as long as their parent remains on display.
 *
 * Available since API level 29.
 */
void ASurfaceControl_release(ASurfaceControl* _Nonnull surface_control) __INTRODUCED_IN(29);

struct ASurfaceTransaction;

/**
 * ASurfaceTransaction is a collection of updates to the surface tree that must
 * be applied atomically.
 */
typedef struct ASurfaceTransaction ASurfaceTransaction;

/**
 * The caller takes ownership of the transaction and must release it using
 * ASurfaceTransaction_delete() below.
 *
 * Available since API level 29.
 */
ASurfaceTransaction* _Nonnull ASurfaceTransaction_create() __INTRODUCED_IN(29);

/**
 * Destroys the \a transaction object.
 *
 * Available since API level 29.
 */
void ASurfaceTransaction_delete(ASurfaceTransaction* _Nullable transaction) __INTRODUCED_IN(29);

/**
 * Applies the updates accumulated in \a transaction.
 *
 * Note that the transaction is guaranteed to be applied atomically. The
 * transactions which are applied on the same thread are also guaranteed to be
 * applied in order.
 *
 * Available since API level 29.
 */
void ASurfaceTransaction_apply(ASurfaceTransaction* _Nonnull transaction) __INTRODUCED_IN(29);

/**
 * An opaque handle returned during a callback that can be used to query general stats and stats for
 * surfaces which were either removed or for which buffers were updated after this transaction was
 * applied.
 */
typedef struct ASurfaceTransactionStats ASurfaceTransactionStats;

/**
 * Since the transactions are applied asynchronously, the
 * ASurfaceTransaction_OnComplete callback can be used to be notified when a frame
 * including the updates in a transaction was presented.
 *
 * Buffers which are replaced or removed from the scene in the transaction invoking
 * this callback may be reused after this point.
 *
 * Starting with API level 36, prefer using \a ASurfaceTransaction_OnBufferRelease to listen
 * to when a buffer is ready to be reused.
 *
 * \param context Optional context provided by the client that is passed into
 * the callback.
 *
 * \param stats Opaque handle that can be passed to ASurfaceTransactionStats functions to query
 * information about the transaction. The handle is only valid during the callback.
 *
 * THREADING
 * The transaction completed callback can be invoked on any thread.
 */
typedef void (*ASurfaceTransaction_OnComplete)(void* _Null_unspecified context,
                                               ASurfaceTransactionStats* _Nonnull stats);

/**
 * The ASurfaceTransaction_OnCommit callback is invoked when transaction is applied and the updates
 * are ready to be presented. This callback will be invoked before the
 * ASurfaceTransaction_OnComplete callback.
 *
 * This callback does not mean buffers have been released! It simply means that any new
 * transactions applied will not overwrite the transaction for which we are receiving
 * a callback and instead will be included in the next frame. If you are trying to avoid
 * dropping frames (overwriting transactions), and unable to use timestamps (Which provide
 * a more efficient solution), then this method provides a method to pace your transaction
 * application.
 *
 * \param context Optional context provided by the client that is passed into the callback.
 *
 * \param stats Opaque handle that can be passed to ASurfaceTransactionStats functions to query
 * information about the transaction. The handle is only valid during the callback.
 * Present and release fences are not available for this callback. Querying them using
 * ASurfaceTransactionStats_getPresentFenceFd and ASurfaceTransactionStats_getPreviousReleaseFenceFd
 * will result in failure.
 *
 * THREADING
 * The transaction committed callback can be invoked on any thread.
 */
typedef void (*ASurfaceTransaction_OnCommit)(void* _Null_unspecified context,
                                             ASurfaceTransactionStats* _Nonnull stats);

/**
 * The ASurfaceTransaction_OnBufferRelease callback is invoked when a buffer that was passed in
 * ASurfaceTransaction_setBuffer is ready to be reused.
 *
 * This callback is guaranteed to be invoked if ASurfaceTransaction_setBuffer is called with a non
 * null buffer. If the buffer in the transaction is replaced via another call to
 * ASurfaceTransaction_setBuffer, the callback will be invoked immediately. Otherwise the callback
 * will be invoked before the ASurfaceTransaction_OnComplete callback after the buffer was
 * presented.
 *
 * If this callback is set, caller should not release the buffer using the
 * ASurfaceTransaction_OnComplete.
 *
 * \param context Optional context provided by the client that is passed into the callback.
 *
 * \param release_fence_fd Returns the fence file descriptor used to signal the release of buffer
 * associated with this callback. If this fence is valid (>=0), the buffer has not yet been released
 * and the fence will signal when the buffer has been released. If the fence is -1 , the buffer is
 * already released. The recipient of the callback takes ownership of the fence fd and is
 * responsible for closing it.
 *
 * THREADING
 * The callback can be invoked on any thread.
 */
typedef void (*ASurfaceTransaction_OnBufferRelease)(void* _Null_unspecified context,
                                                    int release_fence_fd);

/**
 * Returns the timestamp of when the frame was latched by the framework. Once a frame is
 * latched by the framework, it is presented at the following hardware vsync.
 *
 * Available since API level 29.
 */
int64_t ASurfaceTransactionStats_getLatchTime(
        ASurfaceTransactionStats* _Nonnull surface_transaction_stats) __INTRODUCED_IN(29);

/**
 * Returns a sync fence that signals when the transaction has been presented.
 * The recipient of the callback takes ownership of the fence and is responsible for closing
 * it. If a device does not support present fences, a -1 will be returned.
 *
 * This query is not valid for ASurfaceTransaction_OnCommit callback.
 *
 * Available since API level 29.
 */
int ASurfaceTransactionStats_getPresentFenceFd(
        ASurfaceTransactionStats* _Nonnull surface_transaction_stats) __INTRODUCED_IN(29);

/**
 * \a outASurfaceControls returns an array of ASurfaceControl pointers that were updated during the
 * transaction. Stats for the surfaces can be queried through ASurfaceTransactionStats functions.
 * When the client is done using the array, it must release it by calling
 * ASurfaceTransactionStats_releaseASurfaceControls.
 *
 * Available since API level 29.
 *
 * \a outASurfaceControlsSize returns the size of the ASurfaceControls array.
 */
void ASurfaceTransactionStats_getASurfaceControls(
        ASurfaceTransactionStats* _Nonnull surface_transaction_stats,
        ASurfaceControl* _Nullable* _Nullable* _Nonnull outASurfaceControls,
        size_t* _Nonnull outASurfaceControlsSize) __INTRODUCED_IN(29);
/**
 * Releases the array of ASurfaceControls that were returned by
 * ASurfaceTransactionStats_getASurfaceControls().
 *
 * Available since API level 29.
 */
void ASurfaceTransactionStats_releaseASurfaceControls(
        ASurfaceControl* _Nonnull* _Nonnull surface_controls) __INTRODUCED_IN(29);

/**
 * Returns the timestamp of when the CURRENT buffer was acquired. A buffer is considered
 * acquired when its acquire_fence_fd has signaled. A buffer cannot be latched or presented until
 * it is acquired. If no acquire_fence_fd was provided, this timestamp will be set to -1.
 *
 * Available since API level 29.
 *
 * @deprecated This may return SIGNAL_PENDING because the stats can arrive before the acquire
 * fence has signaled, depending on internal timing differences. Therefore the caller should
 * use the acquire fence passed in to setBuffer and query the signal time.
 */
int64_t ASurfaceTransactionStats_getAcquireTime(
        ASurfaceTransactionStats* _Nonnull surface_transaction_stats,
        ASurfaceControl* _Nonnull surface_control) __INTRODUCED_IN(29);

/**
 * The returns the fence used to signal the release of the PREVIOUS buffer set on
 * this surface. If this fence is valid (>=0), the PREVIOUS buffer has not yet been released and the
 * fence will signal when the PREVIOUS buffer has been released. If the fence is -1, the PREVIOUS
 * buffer is already released. The recipient of the callback takes ownership of the
 * previousReleaseFenceFd and is responsible for closing it.
 *
 * Each time a buffer is set through ASurfaceTransaction_setBuffer() on a transaction
 * which is applied, the framework takes a ref on this buffer. The framework treats the
 * addition of a buffer to a particular surface as a unique ref. When a transaction updates or
 * removes a buffer from a surface, or removes the surface itself from the tree, this ref is
 * guaranteed to be released in the OnComplete callback for this transaction. The
 * ASurfaceControlStats provided in the callback for this surface may contain an optional fence
 * which must be signaled before the ref is assumed to be released.
 *
 * The client must ensure that all pending refs on a buffer are released before attempting to reuse
 * this buffer, otherwise synchronization errors may occur.
 *
 * This query is not valid for ASurfaceTransaction_OnCommit callback.
 *
 * Available since API level 29.
 */
int ASurfaceTransactionStats_getPreviousReleaseFenceFd(
        ASurfaceTransactionStats* _Nonnull surface_transaction_stats,
        ASurfaceControl* _Nonnull surface_control) __INTRODUCED_IN(29);

/**
 * Sets the callback that will be invoked when the updates from this transaction
 * are presented. For details on the callback semantics and data, see the
 * comments on the ASurfaceTransaction_OnComplete declaration above.
 *
 * Available since API level 29.
 */
void ASurfaceTransaction_setOnComplete(ASurfaceTransaction* _Nonnull transaction,
                                       void* _Null_unspecified context,
                                       ASurfaceTransaction_OnComplete _Nonnull func)
        __INTRODUCED_IN(29);

/**
 * Sets the callback that will be invoked when the updates from this transaction are applied and are
 * ready to be presented. This callback will be invoked before the ASurfaceTransaction_OnComplete
 * callback.
 *
 * Available since API level 31.
 */
void ASurfaceTransaction_setOnCommit(ASurfaceTransaction* _Nonnull transaction,
                                     void* _Null_unspecified context,
                                     ASurfaceTransaction_OnCommit _Nonnull func)
        __INTRODUCED_IN(31);

/**
 * Reparents the \a surface_control from its old parent to the \a new_parent surface control.
 * Any children of the reparented \a surface_control will remain children of the \a surface_control.
 *
 * The \a new_parent can be null. Surface controls with a null parent do not appear on the display.
 *
 * Available since API level 29.
 */
void ASurfaceTransaction_reparent(ASurfaceTransaction* _Nonnull transaction,
                                  ASurfaceControl* _Nonnull surface_control,
                                  ASurfaceControl* _Nullable new_parent) __INTRODUCED_IN(29);

/**
 * Parameter for ASurfaceTransaction_setVisibility().
 */
enum ASurfaceTransactionVisibility : int8_t {
    ASURFACE_TRANSACTION_VISIBILITY_HIDE = 0,
    ASURFACE_TRANSACTION_VISIBILITY_SHOW = 1,
};
/**
 * Updates the visibility of \a surface_control. If show is set to
 * ASURFACE_TRANSACTION_VISIBILITY_HIDE, the \a surface_control and all surfaces in its subtree will
 * be hidden.
 *
 * Available since API level 29.
 */
void ASurfaceTransaction_setVisibility(ASurfaceTransaction* _Nonnull transaction,
                                       ASurfaceControl* _Nonnull surface_control,
                                       enum ASurfaceTransactionVisibility visibility)
        __INTRODUCED_IN(29);

/**
 * Updates the z order index for \a surface_control. Note that the z order for a surface
 * is relative to other surfaces which are siblings of this surface. The behavior of sibilings with
 * the same z order is undefined.
 *
 * Z orders may be from MIN_INT32 to MAX_INT32. A layer's default z order index is 0.
 *
 * Available since API level 29.
 */
void ASurfaceTransaction_setZOrder(ASurfaceTransaction* _Nonnull transaction,
                                   ASurfaceControl* _Nonnull surface_control, int32_t z_order)
        __INTRODUCED_IN(29);

/**
 * Updates the AHardwareBuffer displayed for \a surface_control. If not -1, the
 * acquire_fence_fd should be a file descriptor that is signaled when all pending work
 * for the buffer is complete and the buffer can be safely read.
 *
 * The frameworks takes ownership of the \a acquire_fence_fd passed and is responsible
 * for closing it.
 *
 * Note that the buffer must be allocated with AHARDWAREBUFFER_USAGE_GPU_SAMPLED_IMAGE
 * as the surface control might be composited using the GPU.
 *
 * Starting with API level 36, prefer using \a ASurfaceTransaction_setBufferWithRelease to
 * set a buffer and a callback which will be invoked when the buffer is ready to be reused.
 *
 * Available since API level 29.
 */
void ASurfaceTransaction_setBuffer(ASurfaceTransaction* _Nonnull transaction,
                                   ASurfaceControl* _Nonnull surface_control,
                                   AHardwareBuffer* _Nonnull buffer, int acquire_fence_fd)
        __INTRODUCED_IN(29);

/**
 * Updates the AHardwareBuffer displayed for \a surface_control. If not -1, the
 * acquire_fence_fd should be a file descriptor that is signaled when all pending work
 * for the buffer is complete and the buffer can be safely read.
 *
 * The frameworks takes ownership of the \a acquire_fence_fd passed and is responsible
 * for closing it.
 *
 * Note that the buffer must be allocated with AHARDWAREBUFFER_USAGE_GPU_SAMPLED_IMAGE
 * as the surface control might be composited using the GPU.
 *
 * When the buffer is ready to be reused, the ASurfaceTransaction_OnBufferRelease
 * callback will be invoked. If the buffer is null, the callback will not be invoked.
 *
 * Available since API level 36.
 */
void ASurfaceTransaction_setBufferWithRelease(ASurfaceTransaction* _Nonnull transaction,
                                              ASurfaceControl* _Nonnull surface_control,
                                              AHardwareBuffer* _Nonnull buffer,
                                              int acquire_fence_fd, void* _Null_unspecified context,
                                              ASurfaceTransaction_OnBufferRelease _Nonnull func)
        __INTRODUCED_IN(36);

/**
 * Updates the color for \a surface_control.  This will make the background color for the
 * ASurfaceControl visible in transparent regions of the surface.  Colors \a r, \a g,
 * and \a b must be within the range that is valid for \a dataspace.  \a dataspace and \a alpha
 * will be the dataspace and alpha set for the background color layer.
 *
 * Available since API level 29.
 */
void ASurfaceTransaction_setColor(ASurfaceTransaction* _Nonnull transaction,
                                  ASurfaceControl* _Nonnull surface_control, float r, float g,
                                  float b, float alpha, enum ADataSpace dataspace)
        __INTRODUCED_IN(29);

// These APIs (setGeometry and setCrop) were originally written in a
// C-incompatible form using references instead of pointers, and the OS shipped
// that version for years before it was noticed. Fortunately the compiled code
// for callers is the same regardless of whether it's a pointer or a reference,
// so we can declare this as a nonnull pointer for C and keep the existing C++
// decl and definition.
//
// We could alternatively change the decl and the definition to both be a
// pointer (with an inline definition using references to preserve source compat
// for existing C++ callers), but that requires changing the definition of an
// API that has been in the OS for years. It's theoretically a safe change, but
// without being able to prove it that's a very big risk to take. By keeping the
// C-compatibility hack in the header, we can be sure that we haven't changed
// anything for existing callers. By definition there were no C users of the
// reference-based decl; if there were any C callers of the API at all, they were
// using the same workaround that is now used below.
//
// Even if this workaround turns out to not work for C, there's no permanent
// damage done to the platform (unlike if we were to change the definition). At
// worst it continues to work for C++ (since the preprocessed header as seen by
// C++ hasn't changed, nor has the definition) and continues to not work for C.

/**
 * \param source The sub-rect within the buffer's content to be rendered inside the surface's area
 * The surface's source rect is clipped by the bounds of its current buffer. The source rect's width
 * and height must be > 0.
 *
 * \param destination Specifies the rect in the parent's space where this surface will be drawn. The
 * post source rect bounds are scaled to fit the destination rect. The surface's destination rect is
 * clipped by the bounds of its parent. The destination rect's width and height must be > 0.
 *
 * \param transform The transform applied after the source rect is applied to the buffer. This
 * parameter should be set to 0 for no transform. To specify a transform use the
 * NATIVE_WINDOW_TRANSFORM_* enum.
 *
 * Available since API level 29.
 *
 * @deprecated Use setCrop, setPosition, setBufferTransform, and setScale instead. Those functions
 * provide well defined behavior and allows for more control by the apps. It also allows the caller
 * to set different properties at different times, instead of having to specify all the desired
 * properties at once.
 */
void ASurfaceTransaction_setGeometry(ASurfaceTransaction* _Nonnull transaction,
                                     ASurfaceControl* _Nonnull surface_control,
#if defined(__cplusplus)
                                     const ARect& source, const ARect& destination,
#else
                                     const ARect* _Nonnull source,
                                     const ARect* _Nonnull destination,
#endif
                                     int32_t transform) __INTRODUCED_IN(29);

/**
 * Bounds the surface and its children to the bounds specified. The crop and buffer size will be
 * used to determine the bounds of the surface. If no crop is specified and the surface has no
 * buffer, the surface bounds is only constrained by the size of its parent bounds.
 *
 * \param crop The bounds of the crop to apply.
 *
 * Available since API level 31.
 */
void ASurfaceTransaction_setCrop(ASurfaceTransaction* _Nonnull transaction,
                                 ASurfaceControl* _Nonnull surface_control,
#if defined(__cplusplus)
                                 const ARect& crop)
#else
                                 const ARect* _Nonnull crop)
#endif
        __INTRODUCED_IN(31);

/**
 * Specifies the position in the parent's space where the surface will be drawn.
 *
 * \param x The x position to render the surface.
 * \param y The y position to render the surface.
 *
 * Available since API level 31.
 */
void ASurfaceTransaction_setPosition(ASurfaceTransaction* _Nonnull transaction,
                                     ASurfaceControl* _Nonnull surface_control, int32_t x,
                                     int32_t y) __INTRODUCED_IN(31);

/**
 * \param transform The transform applied after the source rect is applied to the buffer. This
 * parameter should be set to 0 for no transform. To specify a transform use the
 * NATIVE_WINDOW_TRANSFORM_* enum.
 *
 * Available since API level 31.
 */
void ASurfaceTransaction_setBufferTransform(ASurfaceTransaction* _Nonnull transaction,
                                            ASurfaceControl* _Nonnull surface_control,
                                            int32_t transform) __INTRODUCED_IN(31);

/**
 * Sets an x and y scale of a surface with (0, 0) as the centerpoint of the scale.
 *
 * \param xScale The scale in the x direction. Must be greater than 0.
 * \param yScale The scale in the y direction. Must be greater than 0.
 *
 * Available since API level 31.
 */
void ASurfaceTransaction_setScale(ASurfaceTransaction* _Nonnull transaction,
                                  ASurfaceControl* _Nonnull surface_control, float xScale,
                                  float yScale) __INTRODUCED_IN(31);
/**
 * Parameter for ASurfaceTransaction_setBufferTransparency().
 */
enum ASurfaceTransactionTransparency : int8_t {
    ASURFACE_TRANSACTION_TRANSPARENCY_TRANSPARENT = 0,
    ASURFACE_TRANSACTION_TRANSPARENCY_TRANSLUCENT = 1,
    ASURFACE_TRANSACTION_TRANSPARENCY_OPAQUE = 2,
};
/**
 * Updates whether the content for the buffer associated with this surface is
 * completely opaque. If true, every pixel of content inside the buffer must be
 * opaque or visual errors can occur.
 *
 * Available since API level 29.
 */
void ASurfaceTransaction_setBufferTransparency(ASurfaceTransaction* _Nonnull transaction,
                                               ASurfaceControl* _Nonnull surface_control,
                                               enum ASurfaceTransactionTransparency transparency)
                                               __INTRODUCED_IN(29);

/**
 * Updates the region for the content on this surface updated in this
 * transaction. If unspecified, the complete surface is assumed to be damaged.
 *
 * Available since API level 29.
 */
void ASurfaceTransaction_setDamageRegion(ASurfaceTransaction* _Nonnull transaction,
                                         ASurfaceControl* _Nonnull surface_control,
                                         const ARect* _Nullable rects, uint32_t count)
                                         __INTRODUCED_IN(29);

/**
 * Specifies a desiredPresentTime for the transaction. The framework will try to present
 * the transaction at or after the time specified.
 *
 * Transactions will not be presented until all of their acquire fences have signaled even if the
 * app requests an earlier present time.
 *
 * If an earlier transaction has a desired present time of x, and a later transaction has a desired
 * present time that is before x, the later transaction will not preempt the earlier transaction.
 *
 * Available since API level 29.
 */
void ASurfaceTransaction_setDesiredPresentTime(ASurfaceTransaction* _Nonnull transaction,
                                               int64_t desiredPresentTime) __INTRODUCED_IN(29);

/**
 * Sets the alpha for the buffer. It uses a premultiplied blending.
 *
 * The \a alpha must be between 0.0 and 1.0.
 *
 * Available since API level 29.
 */
void ASurfaceTransaction_setBufferAlpha(ASurfaceTransaction* _Nonnull transaction,
                                        ASurfaceControl* _Nonnull surface_control, float alpha)
                                        __INTRODUCED_IN(29);

/**
 * Sets the data space of the surface_control's buffers.
 *
 * If no data space is set, the surface control defaults to ADATASPACE_SRGB.
 *
 * Available since API level 29.
 */
void ASurfaceTransaction_setBufferDataSpace(ASurfaceTransaction* _Nonnull transaction,
                                            ASurfaceControl* _Nonnull surface_control,
                                            enum ADataSpace data_space) __INTRODUCED_IN(29);

/**
 * SMPTE ST 2086 "Mastering Display Color Volume" static metadata
 *
 * When \a metadata is set to null, the framework does not use any smpte2086 metadata when rendering
 * the surface's buffer.
 *
 * Available since API level 29.
 */
void ASurfaceTransaction_setHdrMetadata_smpte2086(ASurfaceTransaction* _Nonnull transaction,
                                                  ASurfaceControl* _Nonnull surface_control,
                                                  struct AHdrMetadata_smpte2086* _Nullable metadata)
                                                  __INTRODUCED_IN(29);

/**
 * Sets the CTA 861.3 "HDR Static Metadata Extension" static metadata on a surface.
 *
 * When \a metadata is set to null, the framework does not use any cta861.3 metadata when rendering
 * the surface's buffer.
 *
 * Available since API level 29.
 */
void ASurfaceTransaction_setHdrMetadata_cta861_3(ASurfaceTransaction* _Nonnull transaction,
                                                 ASurfaceControl* _Nonnull surface_control,
                                                 struct AHdrMetadata_cta861_3* _Nullable metadata)
                                                 __INTRODUCED_IN(29);

/**
 * Sets the desired extended range brightness for the layer. This only applies for layers whose
 * dataspace has RANGE_EXTENDED set on it. See: ASurfaceTransaction_setDesiredHdrHeadroom, prefer
 * using this API for formats that encode an HDR/SDR ratio as part of generating the buffer.
 *
 * @param surface_control The layer whose extended range brightness is being specified
 * @param currentBufferRatio The current HDR/SDR ratio of the current buffer as represented as
 *                           peakHdrBrightnessInNits / targetSdrWhitePointInNits. For example if the
 *                           buffer was rendered with a target SDR whitepoint of 100nits and a max
 *                           display brightness of 200nits, this should be set to 2.0f.
 *
 *                           Default value is 1.0f.
 *
 *                           Transfer functions that encode their own brightness ranges, such as
 *                           HLG or PQ, should also set this to 1.0f and instead communicate
 *                           extended content brightness information via metadata such as CTA861_3
 *                           or SMPTE2086.
 *
 *                           Must be finite && >= 1.0f
 *
 * @param desiredRatio The desired HDR/SDR ratio as represented as peakHdrBrightnessInNits /
 *                     targetSdrWhitePointInNits. This can be used to communicate the max desired
 *                     brightness range. This is similar to the "max luminance" value in other
 *                     HDR metadata formats, but represented as a ratio of the target SDR whitepoint
 *                     to the max display brightness. The system may not be able to, or may choose
 *                     not to, deliver the requested range.
 *
 *                     While requesting a large desired ratio will result in the most
 *                     dynamic range, voluntarily reducing the requested range can help
 *                     improve battery life as well as can improve quality by ensuring
 *                     greater bit depth is allocated to the luminance range in use.
 *
 *                     Default value is 1.0f and indicates that extended range brightness
 *                     is not being used, so the resulting SDR or HDR behavior will be
 *                     determined entirely by the dataspace being used (ie, typically SDR
 *                     however PQ or HLG transfer functions will still result in HDR)
 *
 *                     When called after ASurfaceTransaction_setDesiredHdrHeadroom, the
 *                     desiredRatio will override the desiredHeadroom provided by
 *                     ASurfaceTransaction_setDesiredHdrHeadroom. Conversely, when called before
 *                     ASurfaceTransaction_setDesiredHdrHeadroom, the desiredHeadroom provided by
 *.                    ASurfaceTransaction_setDesiredHdrHeadroom will override the desiredRatio.
 *
 *                     Must be finite && >= 1.0f
 *
 * Available since API level 34.
 */
void ASurfaceTransaction_setExtendedRangeBrightness(ASurfaceTransaction* _Nonnull transaction,
                                                    ASurfaceControl* _Nonnull surface_control,
                                                    float currentBufferRatio, float desiredRatio)
                                                    __INTRODUCED_IN(__ANDROID_API_U__);

/**
 * Sets the desired HDR headroom for the layer. See: ASurfaceTransaction_setExtendedRangeBrightness,
 * prefer using this API for formats that conform to HDR standards like HLG or HDR10, that do not
 * communicate a HDR/SDR ratio as part of generating the buffer.
 *
 * @param surface_control The layer whose desired HDR headroom is being specified
 *
 * @param desiredHeadroom The desired HDR/SDR ratio as represented as peakHdrBrightnessInNits /
 *                        targetSdrWhitePointInNits. This can be used to communicate the max
 *                        desired brightness range of the panel. The system may not be able to, or
 *                        may choose not to, deliver the requested range.
 *
 *                        While requesting a large desired ratio will result in the most
 *                        dynamic range, voluntarily reducing the requested range can help
 *                        improve battery life as well as can improve quality by ensuring
 *                        greater bit depth is allocated to the luminance range in use.
 *
 *                        Default value is 0.0f and indicates that the system will choose the best
 *                        headroom for this surface control's content. Typically, this means that
 *                        HLG/PQ encoded content will be displayed with some HDR headroom greater
 *                        than 1.0.
 *
 *                        When called after ASurfaceTransaction_setExtendedRangeBrightness, the
 *                        desiredHeadroom will override the desiredRatio provided by
 *                        ASurfaceTransaction_setExtendedRangeBrightness. Conversely, when called
 *                        before ASurfaceTransaction_setExtendedRangeBrightness, the desiredRatio
 *                        provided by ASurfaceTransaction_setExtendedRangeBrightness will override
 *                        the desiredHeadroom.
 *
 *                        Must be finite && >= 1.0f or 0.0f to indicate there is no desired
 *                        headroom.
 *
 * Available since API level 35.
 */
void ASurfaceTransaction_setDesiredHdrHeadroom(ASurfaceTransaction* _Nonnull transaction,
                                               ASurfaceControl* _Nonnull surface_control,
                                               float desiredHeadroom)
        __INTRODUCED_IN(__ANDROID_API_V__);

/**
 * Sets the Lut(s) to be applied for the layer.
 *
 * The function makes a deep copy of the provided `luts`.
 * Any modifications made to the `luts` object after calling this function
 * will not affect the Lut(s) applied to the layer.
 *
 * @param surface_control The layer where Lut(s) is being applied
 * @param luts The Lut(s) to be applied
 *
 * Available since API level 36.
 */
void ASurfaceTransaction_setLuts(ASurfaceTransaction* _Nonnull transaction,
                                 ASurfaceControl* _Nonnull surface_control,
                                 const struct ADisplayLuts* _Nullable luts)
        __INTRODUCED_IN(36);

/**
 * Same as ASurfaceTransaction_setFrameRateWithChangeStrategy(transaction, surface_control,
 * frameRate, compatibility, ANATIVEWINDOW_CHANGE_FRAME_RATE_ONLY_IF_SEAMLESS).
 *
 * See ASurfaceTransaction_setFrameRateWithChangeStrategy().
 *
 * Available since API level 30.
 */
void ASurfaceTransaction_setFrameRate(ASurfaceTransaction* _Nonnull transaction,
                                      ASurfaceControl* _Nonnull surface_control, float frameRate,
                                      int8_t compatibility) __INTRODUCED_IN(30);

/**
 * Sets the intended frame rate for \a surface_control.
 *
 * On devices that are capable of running the display at different refresh rates, the system may
 * choose a display refresh rate to better match this surface's frame rate. Usage of this API won't
 * directly affect the application's frame production pipeline. However, because the system may
 * change the display refresh rate, calls to this function may result in changes to Choreographer
 * callback timings, and changes to the time interval at which the system releases buffers back to
 * the application.
 *
 * You can register for changes in the refresh rate using
 * \a AChoreographer_registerRefreshRateCallback.
 *
 * See ASurfaceTransaction_clearFrameRate().
 *
 * \param frameRate is the intended frame rate of this surface, in frames per second. 0 is a special
 * value that indicates the app will accept the system's choice for the display frame rate, which is
 * the default behavior if this function isn't called. The frameRate param does <em>not</em> need to
 * be a valid refresh rate for this device's display - e.g., it's fine to pass 30fps to a device
 * that can only run the display at 60fps.
 *
 * \param compatibility The frame rate compatibility of this surface. The compatibility value may
 * influence the system's choice of display frame rate. To specify a compatibility use the
 * ANATIVEWINDOW_FRAME_RATE_COMPATIBILITY_* enum. This parameter is ignored when frameRate is 0.
 *
 * \param changeFrameRateStrategy Whether display refresh rate transitions caused by this
 * surface should be seamless. A seamless transition is one that doesn't have any visual
 * interruptions, such as a black screen for a second or two. See the
 * ANATIVEWINDOW_CHANGE_FRAME_RATE_* values. This parameter is ignored when frameRate is 0.
 *
 * Available since API level 31.
 */
void ASurfaceTransaction_setFrameRateWithChangeStrategy(ASurfaceTransaction* _Nonnull transaction,
                                                        ASurfaceControl* _Nonnull surface_control,
                                                        float frameRate, int8_t compatibility,
                                                        int8_t changeFrameRateStrategy)
                                                        __INTRODUCED_IN(31);

/**
 * Clears the frame rate which is set for \a surface_control.
 *
 * This is equivalent to calling
 * ASurfaceTransaction_setFrameRateWithChangeStrategy(
 * transaction, 0, compatibility, changeFrameRateStrategy).
 *
 * Usage of this API won't directly affect the application's frame production pipeline. However,
 * because the system may change the display refresh rate, calls to this function may result in
 * changes to Choreographer callback timings, and changes to the time interval at which the system
 * releases buffers back to the application.
 *
 * See ASurfaceTransaction_setFrameRateWithChangeStrategy()
 *
 * You can register for changes in the refresh rate using
 * \a AChoreographer_registerRefreshRateCallback.
 *
 * See ASurfaceTransaction_setFrameRateWithChangeStrategy().
 *
 * Available since API level 34.
 */
void ASurfaceTransaction_clearFrameRate(ASurfaceTransaction* _Nonnull transaction,
                                        ASurfaceControl* _Nonnull surface_control)
                                        __INTRODUCED_IN(__ANDROID_API_U__);

/**
 * Indicate whether to enable backpressure for buffer submission to a given SurfaceControl.
 *
 * By default backpressure is disabled, which means submitting a buffer prior to receiving
 * a callback for the previous buffer could lead to that buffer being "dropped". In cases
 * where you are selecting for latency, this may be a desirable behavior! We had a new buffer
 * ready, why shouldn't we show it?
 *
 * When back pressure is enabled, each buffer will be required to be presented
 * before it is released and the callback delivered
 * (absent the whole SurfaceControl being removed).
 *
 * Most apps are likely to have some sort of backpressure internally, e.g. you are
 * waiting on the callback from frame N-2 before starting frame N. In high refresh
 * rate scenarios there may not be much time between SurfaceFlinger completing frame
 * N-1 (and therefore releasing buffer N-2) and beginning frame N. This means
 * your app may not have enough time to respond in the callback. Using this flag
 * and pushing buffers earlier for server side queuing will be advantageous
 * in such cases.
 *
 * Available since API level 31.
 *
 * \param transaction The transaction in which to make the change.
 * \param surface_control The ASurfaceControl on which to control buffer backpressure behavior.
 * \param enableBackPressure Whether to enable back pressure.
 */
void ASurfaceTransaction_setEnableBackPressure(ASurfaceTransaction* _Nonnull transaction,
                                               ASurfaceControl* _Nonnull surface_control,
                                               bool enableBackPressure) __INTRODUCED_IN(31);

/**
 * Sets the frame timeline to use in SurfaceFlinger.
 *
 * A frame timeline should be chosen based on the frame deadline the application
 * can meet when rendering the frame and the application's desired presentation time.
 * By setting a frame timeline, SurfaceFlinger tries to present the frame at the corresponding
 * expected presentation time.
 *
 * To receive frame timelines, a callback must be posted to Choreographer using
 * AChoreographer_postVsyncCallback(). The \c vsyncId can then be extracted from the
 * callback payload using AChoreographerFrameCallbackData_getFrameTimelineVsyncId().
 *
 * Available since API level 33.
 *
 * \param vsyncId The vsync ID received from AChoreographer, setting the frame's presentation target
 * to the corresponding expected presentation time and deadline from the frame to be rendered. A
 * stale or invalid value will be ignored.
 */
void ASurfaceTransaction_setFrameTimeline(ASurfaceTransaction* _Nonnull transaction,
                                          AVsyncId vsyncId) __INTRODUCED_IN(33);

__END_DECLS

#endif // ANDROID_SURFACE_CONTROL_H

/** @} */
