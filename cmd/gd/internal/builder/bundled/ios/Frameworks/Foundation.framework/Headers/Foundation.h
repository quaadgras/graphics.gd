// Foundation declarations for graphics.gd's iOS SDK.
//
// Written from the public documentation of the API, covering what Godot and
// SDL make use of. Not derived from Apple's SDK headers.
#ifndef GD_FOUNDATION_H
#define GD_FOUNDATION_H

#include <CoreFoundation/CoreFoundation.h>
#include <CoreGraphics/CoreGraphics.h>
#include <os/availability.h>
#include <dispatch/dispatch.h>
#include <stdarg.h>

#if defined(__OBJC__)
#import <objc/NSObject.h>
#import <objc/objc.h>
#import <objc/runtime.h>

#if defined(__cplusplus)
#define FOUNDATION_EXTERN extern "C"
#else
#define FOUNDATION_EXTERN extern
#endif
#define FOUNDATION_EXPORT FOUNDATION_EXTERN
#define NS_INLINE static inline
#define NS_ENUM(type, name) CF_ENUM(type, name)
#define NS_OPTIONS(type, name) CF_OPTIONS(type, name)
#define NS_CLOSED_ENUM(type, name) CF_ENUM(type, name)
#define NS_ERROR_ENUM(domain, name) CF_ENUM(NSInteger, name)
#define NS_TYPED_ENUM
#define NS_TYPED_EXTENSIBLE_ENUM
#define NS_STRING_ENUM
#define NS_EXTENSIBLE_STRING_ENUM
#define NS_ASSUME_NONNULL_BEGIN _Pragma("clang assume_nonnull begin")
#define NS_ASSUME_NONNULL_END _Pragma("clang assume_nonnull end")
#define NS_DESIGNATED_INITIALIZER __attribute__((objc_designated_initializer))
#define NS_UNAVAILABLE __attribute__((unavailable))
#define NS_REQUIRES_SUPER __attribute__((objc_requires_super))
#define NS_RETURNS_RETAINED __attribute__((ns_returns_retained))
#define NS_RETURNS_INNER_POINTER __attribute__((objc_returns_inner_pointer))
#define NS_FORMAT_FUNCTION(F, A) __attribute__((format(__NSString__, F, A)))
#define NS_REQUIRES_NIL_TERMINATION __attribute__((sentinel(0, 1)))
#define NS_NOESCAPE __attribute__((noescape))
#define NS_SWIFT_NAME(name)
#define NS_SWIFT_UNAVAILABLE(message)
#define NS_REFINED_FOR_SWIFT
#define NS_SWIFT_UI_ACTOR
#define NS_SWIFT_SENDABLE
#define NS_SWIFT_NONISOLATED
#define NS_AVAILABLE_IOS(version)
#define NS_DEPRECATED_IOS(...)
#define NS_EXTENSION_UNAVAILABLE_IOS(message)

NS_ASSUME_NONNULL_BEGIN

typedef long NSInteger;
typedef unsigned long NSUInteger;
typedef double NSTimeInterval;
typedef struct _NSZone NSZone;
#define NSIntegerMax LONG_MAX
#define NSIntegerMin LONG_MIN
#define NSUIntegerMax ULONG_MAX
static const NSInteger NSNotFound = 0x7fffffffffffffffL;

typedef struct _NSRange {
	NSUInteger location;
	NSUInteger length;
} NSRange;
typedef NSRange *NSRangePointer;
NS_INLINE NSRange NSMakeRange(NSUInteger loc, NSUInteger len) {
	NSRange range;
	range.location = loc;
	range.length = len;
	return range;
}
NS_INLINE NSUInteger NSMaxRange(NSRange range) {
	return range.location + range.length;
}
NS_INLINE BOOL NSLocationInRange(NSUInteger loc, NSRange range) {
	return loc >= range.location && loc - range.location < range.length;
}

typedef NS_CLOSED_ENUM(NSInteger, NSComparisonResult) {
	NSOrderedAscending = -1L,
	NSOrderedSame,
	NSOrderedDescending,
};
typedef NSComparisonResult (^NSComparator)(id obj1, id obj2);

typedef struct {
	NSInteger majorVersion;
	NSInteger minorVersion;
	NSInteger patchVersion;
} NSOperatingSystemVersion;

@class NSString, NSArray<ObjectType>, NSDictionary<KeyType, ObjectType>, NSSet<ObjectType>, NSData, NSURL, NSError, NSDate, NSNumber, NSCoder, NSLocale, NSBundle, NSEnumerator<ObjectType>, NSMutableArray<ObjectType>, NSRunLoop, NSTimer, NSThread, NSOperationQueue, NSNotification, NSAttributedString, NSUserActivity, NSIndexSet, NSUUID;

typedef NSString *NSExceptionName;
typedef NSString *NSRunLoopMode;
typedef NSString *NSNotificationName;
typedef NSString *NSErrorDomain;
typedef NSString *NSErrorUserInfoKey;
typedef NSString *NSLocaleKey;
typedef NSString *NSAttributedStringKey;
typedef NSString *NSURLResourceKey;
typedef NSString *NSFileAttributeKey;

@protocol NSCopying
- (id)copyWithZone:(nullable NSZone *)zone;
@end
@protocol NSMutableCopying
- (id)mutableCopyWithZone:(nullable NSZone *)zone;
@end
@protocol NSCoding
- (void)encodeWithCoder:(NSCoder *)coder;
- (nullable instancetype)initWithCoder:(NSCoder *)coder;
@end
@protocol NSSecureCoding <NSCoding>
@property(class, readonly) BOOL supportsSecureCoding;
@end

typedef struct {
	unsigned long state;
	id __unsafe_unretained _Nullable *_Nullable itemsPtr;
	unsigned long *_Nullable mutationsPtr;
	unsigned long extra[5];
} NSFastEnumerationState;
@protocol NSFastEnumeration
- (NSUInteger)countByEnumeratingWithState:(NSFastEnumerationState *)state objects:(id __unsafe_unretained _Nullable[_Nonnull])buffer count:(NSUInteger)len;
@end

@interface NSEnumerator<ObjectType> : NSObject <NSFastEnumeration>
- (nullable ObjectType)nextObject;
@property(readonly, copy) NSArray<ObjectType> *allObjects;
@end

// NSString

typedef NSUInteger NSStringEncoding;
enum : NSStringEncoding {
	NSASCIIStringEncoding = 1,
	NSNEXTSTEPStringEncoding = 2,
	NSJapaneseEUCStringEncoding = 3,
	NSUTF8StringEncoding = 4,
	NSISOLatin1StringEncoding = 5,
	NSSymbolStringEncoding = 6,
	NSNonLossyASCIIStringEncoding = 7,
	NSShiftJISStringEncoding = 8,
	NSISOLatin2StringEncoding = 9,
	NSUnicodeStringEncoding = 10,
	NSWindowsCP1251StringEncoding = 11,
	NSWindowsCP1252StringEncoding = 12,
	NSWindowsCP1253StringEncoding = 13,
	NSWindowsCP1254StringEncoding = 14,
	NSWindowsCP1250StringEncoding = 15,
	NSISO2022JPStringEncoding = 21,
	NSMacOSRomanStringEncoding = 30,
	NSUTF16StringEncoding = NSUnicodeStringEncoding,
	NSUTF16BigEndianStringEncoding = 0x90000100,
	NSUTF16LittleEndianStringEncoding = 0x94000100,
	NSUTF32StringEncoding = 0x8c000100,
	NSUTF32BigEndianStringEncoding = 0x98000100,
	NSUTF32LittleEndianStringEncoding = 0x9c000100,
};
typedef NS_OPTIONS(NSUInteger, NSStringCompareOptions) {
	NSCaseInsensitiveSearch = 1,
	NSLiteralSearch = 2,
	NSBackwardsSearch = 4,
	NSAnchoredSearch = 8,
	NSNumericSearch = 64,
};
typedef unsigned short unichar;

@interface NSString : NSObject <NSCopying, NSMutableCopying, NSSecureCoding>
@property(readonly) NSUInteger length;
- (unichar)characterAtIndex:(NSUInteger)index;
- (instancetype)init;
- (nullable instancetype)initWithUTF8String:(const char *)nullTerminatedCString;
- (nullable instancetype)initWithCString:(const char *)nullTerminatedCString encoding:(NSStringEncoding)encoding;
- (nullable instancetype)initWithBytes:(const void *)bytes length:(NSUInteger)len encoding:(NSStringEncoding)encoding;
- (nullable instancetype)initWithData:(NSData *)data encoding:(NSStringEncoding)encoding;
- (instancetype)initWithString:(NSString *)aString;
- (instancetype)initWithCharacters:(const unichar *)characters length:(NSUInteger)length;
- (instancetype)initWithFormat:(NSString *)format, ... NS_FORMAT_FUNCTION(1, 2);
+ (instancetype)string;
+ (instancetype)stringWithString:(NSString *)string;
+ (nullable instancetype)stringWithUTF8String:(const char *)nullTerminatedCString;
+ (nullable instancetype)stringWithCString:(const char *)cString encoding:(NSStringEncoding)enc;
+ (instancetype)stringWithCharacters:(const unichar *)characters length:(NSUInteger)length;
+ (instancetype)stringWithFormat:(NSString *)format, ... NS_FORMAT_FUNCTION(1, 2);
+ (nullable instancetype)stringWithContentsOfFile:(NSString *)path encoding:(NSStringEncoding)enc error:(NSError **)error;
@property(nullable, readonly) const char *UTF8String NS_RETURNS_INNER_POINTER;
- (nullable const char *)cStringUsingEncoding:(NSStringEncoding)encoding;
- (BOOL)getCString:(char *)buffer maxLength:(NSUInteger)maxBufferCount encoding:(NSStringEncoding)encoding;
- (void)getCharacters:(unichar *)buffer range:(NSRange)range;
- (BOOL)getBytes:(nullable void *)buffer maxLength:(NSUInteger)maxBufferCount usedLength:(nullable NSUInteger *)usedBufferCount encoding:(NSStringEncoding)encoding options:(NSUInteger)options range:(NSRange)range remainingRange:(nullable NSRangePointer)leftover;
- (nullable NSData *)dataUsingEncoding:(NSStringEncoding)encoding;
- (NSUInteger)lengthOfBytesUsingEncoding:(NSStringEncoding)enc;
- (NSUInteger)maximumLengthOfBytesUsingEncoding:(NSStringEncoding)enc;
- (BOOL)isEqualToString:(NSString *)aString;
- (NSComparisonResult)compare:(NSString *)string;
- (NSComparisonResult)compare:(NSString *)string options:(NSStringCompareOptions)mask;
- (NSComparisonResult)caseInsensitiveCompare:(NSString *)string;
- (BOOL)hasPrefix:(NSString *)str;
- (BOOL)hasSuffix:(NSString *)str;
- (BOOL)containsString:(NSString *)str;
- (NSRange)rangeOfString:(NSString *)searchString;
- (NSRange)rangeOfString:(NSString *)searchString options:(NSStringCompareOptions)mask;
- (NSString *)substringFromIndex:(NSUInteger)from;
- (NSString *)substringToIndex:(NSUInteger)to;
- (NSString *)substringWithRange:(NSRange)range;
- (NSString *)stringByAppendingString:(NSString *)aString;
- (NSString *)stringByAppendingFormat:(NSString *)format, ... NS_FORMAT_FUNCTION(1, 2);
- (NSString *)stringByReplacingOccurrencesOfString:(NSString *)target withString:(NSString *)replacement;
- (NSString *)stringByReplacingCharactersInRange:(NSRange)range withString:(NSString *)replacement;
- (NSArray<NSString *> *)componentsSeparatedByString:(NSString *)separator;
@property(readonly, copy) NSString *uppercaseString;
@property(readonly, copy) NSString *lowercaseString;
@property(readonly, copy) NSString *description;
@property(readonly) double doubleValue;
@property(readonly) float floatValue;
@property(readonly) int intValue;
@property(readonly) NSInteger integerValue;
@property(readonly) long long longLongValue;
@property(readonly) BOOL boolValue;
// NSPathUtilities
@property(readonly, copy) NSString *lastPathComponent;
@property(readonly, copy) NSString *pathExtension;
@property(readonly, copy) NSString *stringByDeletingLastPathComponent;
@property(readonly, copy) NSString *stringByDeletingPathExtension;
@property(readonly, copy) NSString *stringByExpandingTildeInPath;
@property(readonly, copy) NSString *stringByStandardizingPath;
@property(readonly, copy) NSArray<NSString *> *pathComponents;
@property(readonly) const char *fileSystemRepresentation NS_RETURNS_INNER_POINTER;
- (NSString *)stringByAppendingPathComponent:(NSString *)str;
- (nullable NSString *)stringByAppendingPathExtension:(NSString *)str;
- (nullable NSString *)stringByAddingPercentEncodingWithAllowedCharacters:(id)allowedCharacters;
@property(nullable, readonly, copy) NSString *stringByRemovingPercentEncoding;
@end

@interface NSMutableString : NSString
- (void)appendString:(NSString *)aString;
- (void)appendFormat:(NSString *)format, ... NS_FORMAT_FUNCTION(1, 2);
- (void)setString:(NSString *)aString;
- (void)replaceCharactersInRange:(NSRange)range withString:(NSString *)aString;
- (void)deleteCharactersInRange:(NSRange)range;
- (void)insertString:(NSString *)aString atIndex:(NSUInteger)loc;
+ (instancetype)stringWithCapacity:(NSUInteger)capacity;
@end

FOUNDATION_EXPORT NSAttributedStringKey NSFontAttributeName;
FOUNDATION_EXPORT NSAttributedStringKey NSForegroundColorAttributeName;
FOUNDATION_EXPORT NSAttributedStringKey NSParagraphStyleAttributeName;

@interface NSAttributedString : NSObject <NSCopying, NSMutableCopying, NSSecureCoding>
@property(readonly, copy) NSString *string;
@property(readonly) NSUInteger length;
- (instancetype)initWithString:(NSString *)str;
- (instancetype)initWithString:(NSString *)str attributes:(nullable NSDictionary<NSAttributedStringKey, id> *)attrs;
- (instancetype)initWithAttributedString:(NSAttributedString *)attrStr;
- (NSDictionary<NSAttributedStringKey, id> *)attributesAtIndex:(NSUInteger)location effectiveRange:(nullable NSRangePointer)range;
@end

@interface NSMutableAttributedString : NSAttributedString
@property(readonly, retain) NSMutableString *mutableString;
- (void)replaceCharactersInRange:(NSRange)range withString:(NSString *)str;
- (void)setAttributes:(nullable NSDictionary<NSAttributedStringKey, id> *)attrs range:(NSRange)range;
- (void)addAttribute:(NSAttributedStringKey)name value:(id)value range:(NSRange)range;
- (void)addAttributes:(NSDictionary<NSAttributedStringKey, id> *)attrs range:(NSRange)range;
- (void)removeAttribute:(NSAttributedStringKey)name range:(NSRange)range;
- (void)appendAttributedString:(NSAttributedString *)attrString;
- (void)setAttributedString:(NSAttributedString *)attrString;
- (void)beginEditing;
- (void)endEditing;
@end

// NSValue & NSNumber

@interface NSValue : NSObject <NSCopying, NSSecureCoding>
- (void)getValue:(void *)value size:(NSUInteger)size;
@property(readonly) const char *objCType NS_RETURNS_INNER_POINTER;
+ (NSValue *)valueWithBytes:(const void *)value objCType:(const char *)type;
+ (NSValue *)valueWithPointer:(nullable const void *)pointer;
+ (NSValue *)valueWithNonretainedObject:(nullable id)anObject;
@property(nullable, readonly) void *pointerValue;
@property(nullable, readonly) id nonretainedObjectValue;
+ (NSValue *)valueWithRange:(NSRange)range;
@property(readonly) NSRange rangeValue;
@end

@interface NSNumber : NSValue
+ (NSNumber *)numberWithChar:(char)value;
+ (NSNumber *)numberWithUnsignedChar:(unsigned char)value;
+ (NSNumber *)numberWithShort:(short)value;
+ (NSNumber *)numberWithUnsignedShort:(unsigned short)value;
+ (NSNumber *)numberWithInt:(int)value;
+ (NSNumber *)numberWithUnsignedInt:(unsigned int)value;
+ (NSNumber *)numberWithLong:(long)value;
+ (NSNumber *)numberWithUnsignedLong:(unsigned long)value;
+ (NSNumber *)numberWithLongLong:(long long)value;
+ (NSNumber *)numberWithUnsignedLongLong:(unsigned long long)value;
+ (NSNumber *)numberWithFloat:(float)value;
+ (NSNumber *)numberWithDouble:(double)value;
+ (NSNumber *)numberWithBool:(BOOL)value;
+ (NSNumber *)numberWithInteger:(NSInteger)value;
+ (NSNumber *)numberWithUnsignedInteger:(NSUInteger)value;
@property(readonly) char charValue;
@property(readonly) unsigned char unsignedCharValue;
@property(readonly) short shortValue;
@property(readonly) unsigned short unsignedShortValue;
@property(readonly) int intValue;
@property(readonly) unsigned int unsignedIntValue;
@property(readonly) long longValue;
@property(readonly) unsigned long unsignedLongValue;
@property(readonly) long long longLongValue;
@property(readonly) unsigned long long unsignedLongLongValue;
@property(readonly) float floatValue;
@property(readonly) double doubleValue;
@property(readonly) BOOL boolValue;
@property(readonly) NSInteger integerValue;
@property(readonly) NSUInteger unsignedIntegerValue;
@property(readonly, copy) NSString *stringValue;
- (NSComparisonResult)compare:(NSNumber *)otherNumber;
- (BOOL)isEqualToNumber:(NSNumber *)number;
@end

@interface NSNull : NSObject <NSCopying, NSSecureCoding>
+ (NSNull *)null;
@end

// NSData

@interface NSData : NSObject <NSCopying, NSMutableCopying, NSSecureCoding>
@property(readonly) NSUInteger length;
@property(readonly) const void *bytes NS_RETURNS_INNER_POINTER;
- (void)getBytes:(void *)buffer length:(NSUInteger)length;
- (void)getBytes:(void *)buffer range:(NSRange)range;
+ (instancetype)data;
+ (instancetype)dataWithBytes:(nullable const void *)bytes length:(NSUInteger)length;
+ (instancetype)dataWithBytesNoCopy:(void *)bytes length:(NSUInteger)length freeWhenDone:(BOOL)b;
+ (nullable instancetype)dataWithContentsOfFile:(NSString *)path;
+ (nullable instancetype)dataWithContentsOfURL:(NSURL *)url;
- (instancetype)initWithBytes:(nullable const void *)bytes length:(NSUInteger)length;
- (BOOL)writeToFile:(NSString *)path atomically:(BOOL)useAuxiliaryFile;
- (BOOL)isEqualToData:(NSData *)other;
@end

@interface NSMutableData : NSData
@property(readonly) void *mutableBytes NS_RETURNS_INNER_POINTER;
@property NSUInteger length;
- (void)appendBytes:(const void *)bytes length:(NSUInteger)length;
- (void)appendData:(NSData *)other;
+ (nullable instancetype)dataWithCapacity:(NSUInteger)aNumItems;
+ (nullable instancetype)dataWithLength:(NSUInteger)length;
@end

// Collections

@interface NSArray<__covariant ObjectType> : NSObject <NSCopying, NSMutableCopying, NSSecureCoding, NSFastEnumeration>
@property(readonly) NSUInteger count;
- (ObjectType)objectAtIndex:(NSUInteger)index;
- (ObjectType)objectAtIndexedSubscript:(NSUInteger)idx;
- (instancetype)init;
- (instancetype)initWithObjects:(const ObjectType _Nonnull[_Nullable])objects count:(NSUInteger)cnt;
- (instancetype)initWithArray:(NSArray<ObjectType> *)array;
+ (instancetype)array;
+ (instancetype)arrayWithObject:(ObjectType)anObject;
+ (instancetype)arrayWithObjects:(const ObjectType _Nonnull[_Nonnull])objects count:(NSUInteger)cnt;
+ (instancetype)arrayWithObjects:(ObjectType)firstObj, ... NS_REQUIRES_NIL_TERMINATION;
+ (instancetype)arrayWithArray:(NSArray<ObjectType> *)array;
- (NSArray<ObjectType> *)arrayByAddingObject:(ObjectType)anObject;
- (NSArray<ObjectType> *)arrayByAddingObjectsFromArray:(NSArray<ObjectType> *)otherArray;
- (NSString *)componentsJoinedByString:(NSString *)separator;
- (BOOL)containsObject:(ObjectType)anObject;
- (NSUInteger)indexOfObject:(ObjectType)anObject;
@property(nullable, nonatomic, readonly) ObjectType firstObject;
@property(nullable, nonatomic, readonly) ObjectType lastObject;
- (NSEnumerator<ObjectType> *)objectEnumerator;
- (NSArray<ObjectType> *)sortedArrayUsingComparator:(NSComparator NS_NOESCAPE)cmptr;
- (NSArray<ObjectType> *)sortedArrayUsingSelector:(SEL)comparator;
- (NSArray<ObjectType> *)subarrayWithRange:(NSRange)range;
- (void)enumerateObjectsUsingBlock:(void (NS_NOESCAPE ^)(ObjectType obj, NSUInteger idx, BOOL *stop))block;
- (BOOL)isEqualToArray:(NSArray<ObjectType> *)otherArray;
@end

@interface NSPredicate : NSObject <NSSecureCoding, NSCopying>
+ (NSPredicate *)predicateWithFormat:(NSString *)predicateFormat, ...;
+ (NSPredicate *)predicateWithBlock:(BOOL (^)(id _Nullable evaluatedObject, NSDictionary<NSString *, id> *_Nullable bindings))block;
- (BOOL)evaluateWithObject:(nullable id)object;
@end

@interface NSArray <ObjectType>(NSPredicateSupport)
- (NSArray<ObjectType> *)filteredArrayUsingPredicate:(NSPredicate *)predicate;
@end

@interface NSMutableArray<ObjectType> : NSArray <ObjectType>
- (void)addObject:(ObjectType)anObject;
- (void)addObjectsFromArray:(NSArray<ObjectType> *)otherArray;
- (void)insertObject:(ObjectType)anObject atIndex:(NSUInteger)index;
- (void)removeLastObject;
- (void)removeObjectAtIndex:(NSUInteger)index;
- (void)removeObject:(ObjectType)anObject;
- (void)removeAllObjects;
- (void)replaceObjectAtIndex:(NSUInteger)index withObject:(ObjectType)anObject;
- (void)setObject:(ObjectType)obj atIndexedSubscript:(NSUInteger)idx;
- (instancetype)initWithCapacity:(NSUInteger)numItems;
+ (instancetype)arrayWithCapacity:(NSUInteger)numItems;
- (void)sortUsingComparator:(NSComparator NS_NOESCAPE)cmptr;
@end

@interface NSDictionary<__covariant KeyType, __covariant ObjectType> : NSObject <NSCopying, NSMutableCopying, NSSecureCoding, NSFastEnumeration>
@property(readonly) NSUInteger count;
- (nullable ObjectType)objectForKey:(KeyType)aKey;
- (nullable ObjectType)objectForKeyedSubscript:(KeyType)key;
- (nullable ObjectType)valueForKey:(NSString *)key;
- (NSEnumerator<KeyType> *)keyEnumerator;
- (instancetype)init;
- (instancetype)initWithObjects:(const ObjectType _Nonnull[_Nullable])objects forKeys:(const KeyType<NSCopying> _Nonnull[_Nullable])keys count:(NSUInteger)cnt;
- (instancetype)initWithDictionary:(NSDictionary<KeyType, ObjectType> *)otherDictionary;
@property(readonly, copy) NSArray<KeyType> *allKeys;
@property(readonly, copy) NSArray<ObjectType> *allValues;
+ (instancetype)dictionary;
+ (instancetype)dictionaryWithObject:(ObjectType)object forKey:(KeyType<NSCopying>)key;
+ (instancetype)dictionaryWithObjects:(const ObjectType _Nonnull[_Nullable])objects forKeys:(const KeyType<NSCopying> _Nonnull[_Nullable])keys count:(NSUInteger)cnt;
+ (instancetype)dictionaryWithObjectsAndKeys:(id)firstObject, ... NS_REQUIRES_NIL_TERMINATION;
+ (instancetype)dictionaryWithDictionary:(NSDictionary<KeyType, ObjectType> *)dict;
+ (nullable NSDictionary<KeyType, ObjectType> *)dictionaryWithContentsOfFile:(NSString *)path;
- (void)enumerateKeysAndObjectsUsingBlock:(void (NS_NOESCAPE ^)(KeyType key, ObjectType obj, BOOL *stop))block;
@end

@interface NSMutableDictionary<KeyType, ObjectType> : NSDictionary <KeyType, ObjectType>
- (void)removeObjectForKey:(KeyType)aKey;
- (void)setObject:(ObjectType)anObject forKey:(KeyType<NSCopying>)aKey;
- (void)setObject:(nullable ObjectType)obj forKeyedSubscript:(KeyType<NSCopying>)key;
- (void)setValue:(nullable ObjectType)value forKey:(NSString *)key;
- (void)removeAllObjects;
- (void)addEntriesFromDictionary:(NSDictionary<KeyType, ObjectType> *)otherDictionary;
- (instancetype)initWithCapacity:(NSUInteger)numItems;
+ (instancetype)dictionaryWithCapacity:(NSUInteger)numItems;
@end

@interface NSSet<__covariant ObjectType> : NSObject <NSCopying, NSMutableCopying, NSSecureCoding, NSFastEnumeration>
@property(readonly) NSUInteger count;
- (nullable ObjectType)member:(ObjectType)object;
- (NSEnumerator<ObjectType> *)objectEnumerator;
@property(readonly, copy) NSArray<ObjectType> *allObjects;
- (nullable ObjectType)anyObject;
- (BOOL)containsObject:(ObjectType)anObject;
+ (instancetype)set;
+ (instancetype)setWithObject:(ObjectType)object;
+ (instancetype)setWithObjects:(ObjectType)firstObj, ... NS_REQUIRES_NIL_TERMINATION;
+ (instancetype)setWithArray:(NSArray<ObjectType> *)array;
- (void)enumerateObjectsUsingBlock:(void (NS_NOESCAPE ^)(ObjectType obj, BOOL *stop))block;
@end

@interface NSMutableSet<ObjectType> : NSSet <ObjectType>
- (void)addObject:(ObjectType)object;
- (void)removeObject:(ObjectType)object;
- (void)removeAllObjects;
+ (instancetype)setWithCapacity:(NSUInteger)numItems;
@end

@interface NSIndexSet : NSObject <NSCopying, NSMutableCopying, NSSecureCoding>
@property(readonly) NSUInteger count;
@end

@interface NSUUID : NSObject <NSCopying, NSSecureCoding>
+ (instancetype)UUID;
@property(readonly, copy) NSString *UUIDString;
@end

// NSDate

@interface NSDate : NSObject <NSCopying, NSSecureCoding>
@property(readonly) NSTimeInterval timeIntervalSinceReferenceDate;
@property(readonly) NSTimeInterval timeIntervalSince1970;
@property(readonly) NSTimeInterval timeIntervalSinceNow;
- (NSTimeInterval)timeIntervalSinceDate:(NSDate *)anotherDate;
+ (instancetype)date;
+ (instancetype)dateWithTimeIntervalSinceNow:(NSTimeInterval)secs;
+ (instancetype)dateWithTimeIntervalSince1970:(NSTimeInterval)secs;
@property(class, readonly, copy) NSDate *distantFuture;
@property(class, readonly, copy) NSDate *distantPast;
@end

// NSError & NSException

FOUNDATION_EXPORT NSErrorDomain const NSCocoaErrorDomain;
FOUNDATION_EXPORT NSErrorDomain const NSPOSIXErrorDomain;
FOUNDATION_EXPORT NSErrorDomain const NSOSStatusErrorDomain;
FOUNDATION_EXPORT NSErrorUserInfoKey const NSLocalizedDescriptionKey;
FOUNDATION_EXPORT NSErrorUserInfoKey const NSUnderlyingErrorKey;

@interface NSError : NSObject <NSCopying, NSSecureCoding>
- (instancetype)initWithDomain:(NSErrorDomain)domain code:(NSInteger)code userInfo:(nullable NSDictionary<NSErrorUserInfoKey, id> *)dict;
+ (instancetype)errorWithDomain:(NSErrorDomain)domain code:(NSInteger)code userInfo:(nullable NSDictionary<NSErrorUserInfoKey, id> *)dict;
@property(readonly, copy) NSErrorDomain domain;
@property(readonly) NSInteger code;
@property(readonly, copy) NSDictionary<NSErrorUserInfoKey, id> *userInfo;
@property(readonly, copy) NSString *localizedDescription;
@property(nullable, readonly, copy) NSString *localizedFailureReason;
@end

@interface NSException : NSObject <NSCopying, NSSecureCoding>
@property(readonly, copy) NSExceptionName name;
@property(nullable, readonly, copy) NSString *reason;
@property(nullable, readonly, copy) NSDictionary *userInfo;
+ (void)raise:(NSExceptionName)name format:(NSString *)format, ... NS_FORMAT_FUNCTION(2, 3);
@end

// NSURL

@interface NSURL : NSObject <NSCopying, NSSecureCoding>
- (nullable instancetype)initWithString:(NSString *)URLString;
+ (nullable instancetype)URLWithString:(NSString *)URLString;
+ (NSURL *)fileURLWithPath:(NSString *)path;
+ (NSURL *)fileURLWithPath:(NSString *)path isDirectory:(BOOL)isDir;
@property(nullable, readonly, copy) NSString *absoluteString;
@property(nullable, readonly, copy) NSString *path;
@property(nullable, readonly, copy) NSString *scheme;
@property(nullable, readonly, copy) NSString *host;
@property(nullable, readonly, copy) NSString *query;
@property(nullable, readonly, copy) NSString *lastPathComponent;
@property(nullable, readonly, copy) NSString *pathExtension;
@property(readonly, getter=isFileURL) BOOL fileURL;
@property(readonly) const char *fileSystemRepresentation NS_RETURNS_INNER_POINTER;
- (nullable NSURL *)URLByAppendingPathComponent:(NSString *)pathComponent;
- (BOOL)startAccessingSecurityScopedResource;
- (void)stopAccessingSecurityScopedResource;
@end

// NSBundle

@interface NSBundle : NSObject
@property(class, readonly, strong) NSBundle *mainBundle;
+ (nullable instancetype)bundleWithPath:(NSString *)path;
+ (nullable instancetype)bundleWithURL:(NSURL *)url;
+ (nullable NSBundle *)bundleWithIdentifier:(NSString *)identifier;
+ (NSBundle *)bundleForClass:(Class)aClass;
@property(readonly, copy) NSString *bundlePath;
@property(nullable, readonly, copy) NSString *resourcePath;
@property(nullable, readonly, copy) NSString *executablePath;
@property(nullable, readonly, copy) NSString *privateFrameworksPath;
@property(nullable, readonly, copy) NSString *bundleIdentifier;
@property(readonly, copy) NSURL *bundleURL;
@property(nullable, readonly, copy) NSURL *resourceURL;
@property(nullable, readonly, copy) NSDictionary<NSString *, id> *infoDictionary;
@property(nullable, readonly, copy) NSDictionary<NSString *, id> *localizedInfoDictionary;
@property(readonly, copy) NSArray<NSString *> *preferredLocalizations;
- (nullable id)objectForInfoDictionaryKey:(NSString *)key;
- (nullable NSString *)pathForResource:(nullable NSString *)name ofType:(nullable NSString *)ext;
- (nullable NSURL *)URLForResource:(nullable NSString *)name withExtension:(nullable NSString *)ext;
- (NSString *)localizedStringForKey:(NSString *)key value:(nullable NSString *)value table:(nullable NSString *)tableName;
- (BOOL)load;
@end

// NSLocale

FOUNDATION_EXPORT NSLocaleKey const NSLocaleIdentifier;
FOUNDATION_EXPORT NSLocaleKey const NSLocaleLanguageCode;
FOUNDATION_EXPORT NSLocaleKey const NSLocaleCountryCode;
FOUNDATION_EXPORT NSLocaleKey const NSLocaleScriptCode;

@interface NSLocale : NSObject <NSCopying, NSSecureCoding>
- (nullable id)objectForKey:(NSLocaleKey)key;
- (instancetype)initWithLocaleIdentifier:(NSString *)string;
+ (instancetype)localeWithLocaleIdentifier:(NSString *)ident;
@property(class, readonly, copy) NSLocale *currentLocale;
@property(class, readonly, strong) NSLocale *autoupdatingCurrentLocale;
@property(class, readonly, copy) NSLocale *systemLocale;
@property(class, readonly, copy) NSArray<NSString *> *preferredLanguages;
@property(readonly, copy) NSString *localeIdentifier;
@property(readonly, copy) NSString *languageCode;
@property(nullable, readonly, copy) NSString *countryCode;
@property(nullable, readonly, copy) NSString *scriptCode;
@end

// NSNotification

@interface NSNotification : NSObject <NSCopying, NSCoding>
@property(readonly, copy) NSNotificationName name;
@property(nullable, readonly, retain) id object;
@property(nullable, readonly, copy) NSDictionary *userInfo;
+ (instancetype)notificationWithName:(NSNotificationName)aName object:(nullable id)anObject;
@end

@interface NSNotificationCenter : NSObject
@property(class, readonly, strong) NSNotificationCenter *defaultCenter;
- (void)addObserver:(id)observer selector:(SEL)aSelector name:(nullable NSNotificationName)aName object:(nullable id)anObject;
- (void)postNotification:(NSNotification *)notification;
- (void)postNotificationName:(NSNotificationName)aName object:(nullable id)anObject;
- (void)postNotificationName:(NSNotificationName)aName object:(nullable id)anObject userInfo:(nullable NSDictionary *)aUserInfo;
- (void)removeObserver:(id)observer;
- (void)removeObserver:(id)observer name:(nullable NSNotificationName)aName object:(nullable id)anObject;
- (id<NSObject>)addObserverForName:(nullable NSNotificationName)name object:(nullable id)obj queue:(nullable NSOperationQueue *)queue usingBlock:(void (^)(NSNotification *notification))block;
@end

// NSUserDefaults

@interface NSUserDefaults : NSObject
@property(class, readonly, strong) NSUserDefaults *standardUserDefaults;
- (nullable id)objectForKey:(NSString *)defaultName;
- (void)setObject:(nullable id)value forKey:(NSString *)defaultName;
- (void)removeObjectForKey:(NSString *)defaultName;
- (nullable NSString *)stringForKey:(NSString *)defaultName;
- (nullable NSArray *)arrayForKey:(NSString *)defaultName;
- (NSInteger)integerForKey:(NSString *)defaultName;
- (float)floatForKey:(NSString *)defaultName;
- (double)doubleForKey:(NSString *)defaultName;
- (BOOL)boolForKey:(NSString *)defaultName;
- (void)setInteger:(NSInteger)value forKey:(NSString *)defaultName;
- (void)setBool:(BOOL)value forKey:(NSString *)defaultName;
- (BOOL)synchronize;
@end

// NSProcessInfo

typedef NS_ENUM(NSInteger, NSProcessInfoThermalState) {
	NSProcessInfoThermalStateNominal,
	NSProcessInfoThermalStateFair,
	NSProcessInfoThermalStateSerious,
	NSProcessInfoThermalStateCritical,
};

@interface NSProcessInfo : NSObject
@property(class, readonly, strong) NSProcessInfo *processInfo;
@property(readonly, copy) NSDictionary<NSString *, NSString *> *environment;
@property(readonly, copy) NSArray<NSString *> *arguments;
@property(readonly, copy) NSString *hostName;
@property(copy) NSString *processName;
@property(readonly) int processIdentifier;
@property(readonly, copy) NSString *globallyUniqueString;
@property(readonly, copy) NSString *operatingSystemVersionString;
@property(readonly) NSOperatingSystemVersion operatingSystemVersion;
@property(readonly) NSUInteger processorCount;
@property(readonly) NSUInteger activeProcessorCount;
@property(readonly) unsigned long long physicalMemory;
@property(readonly) NSTimeInterval systemUptime;
@property(readonly) NSProcessInfoThermalState thermalState;
@property(readonly, getter=isLowPowerModeEnabled) BOOL lowPowerModeEnabled;
@property(readonly, getter=isiOSAppOnMac) BOOL iOSAppOnMac;
@property(readonly, getter=isMacCatalystApp) BOOL macCatalystApp;
- (BOOL)isOperatingSystemAtLeastVersion:(NSOperatingSystemVersion)version;
@end

// NSRunLoop, NSTimer, NSThread & NSOperationQueue

FOUNDATION_EXPORT NSRunLoopMode const NSDefaultRunLoopMode;
FOUNDATION_EXPORT NSRunLoopMode const NSRunLoopCommonModes;

@interface NSRunLoop : NSObject
@property(class, readonly, strong) NSRunLoop *currentRunLoop;
@property(class, readonly, strong) NSRunLoop *mainRunLoop;
@property(nullable, readonly, copy) NSRunLoopMode currentMode;
- (CFRunLoopRef)getCFRunLoop;
- (void)addTimer:(NSTimer *)timer forMode:(NSRunLoopMode)mode;
- (void)run;
- (void)runUntilDate:(NSDate *)limitDate;
- (BOOL)runMode:(NSRunLoopMode)mode beforeDate:(NSDate *)limitDate;
@end

@interface NSTimer : NSObject
+ (NSTimer *)timerWithTimeInterval:(NSTimeInterval)ti target:(id)aTarget selector:(SEL)aSelector userInfo:(nullable id)userInfo repeats:(BOOL)yesOrNo;
+ (NSTimer *)scheduledTimerWithTimeInterval:(NSTimeInterval)ti target:(id)aTarget selector:(SEL)aSelector userInfo:(nullable id)userInfo repeats:(BOOL)yesOrNo;
+ (NSTimer *)scheduledTimerWithTimeInterval:(NSTimeInterval)interval repeats:(BOOL)repeats block:(void (^)(NSTimer *timer))block;
- (void)fire;
- (void)invalidate;
@property(readonly, getter=isValid) BOOL valid;
@property(nullable, readonly, retain) id userInfo;
@end

@interface NSThread : NSObject
@property(class, readonly, strong) NSThread *currentThread;
@property(class, readonly, strong) NSThread *mainThread;
@property(class, readonly) BOOL isMainThread;
@property(readonly) BOOL isMainThread;
@property(nullable, copy) NSString *name;
+ (void)sleepForTimeInterval:(NSTimeInterval)ti;
+ (BOOL)isMultiThreaded;
+ (void)detachNewThreadSelector:(SEL)selector toTarget:(id)target withObject:(nullable id)argument;
@end

@interface NSOperationQueue : NSObject
@property(class, readonly, strong) NSOperationQueue *mainQueue;
@property(class, readonly, strong, nullable) NSOperationQueue *currentQueue;
- (void)addOperationWithBlock:(void (^)(void))block;
@end

@interface NSObject (NSThreadPerformAdditions)
- (void)performSelectorOnMainThread:(SEL)aSelector withObject:(nullable id)arg waitUntilDone:(BOOL)wait;
- (void)performSelectorInBackground:(SEL)aSelector withObject:(nullable id)arg;
@end

@interface NSObject (NSDelayedPerforming)
- (void)performSelector:(SEL)aSelector withObject:(nullable id)anArgument afterDelay:(NSTimeInterval)delay;
+ (void)cancelPreviousPerformRequestsWithTarget:(id)aTarget;
@end

@interface NSObject (NSKeyValueCoding)
- (nullable id)valueForKey:(NSString *)key;
- (void)setValue:(nullable id)value forKey:(NSString *)key;
- (nullable id)valueForKeyPath:(NSString *)keyPath;
@end

typedef NS_OPTIONS(NSUInteger, NSKeyValueObservingOptions) {
	NSKeyValueObservingOptionNew = 0x01,
	NSKeyValueObservingOptionOld = 0x02,
	NSKeyValueObservingOptionInitial = 0x04,
	NSKeyValueObservingOptionPrior = 0x08,
};
typedef NSString *NSKeyValueChangeKey;
FOUNDATION_EXPORT NSKeyValueChangeKey const NSKeyValueChangeNewKey;
FOUNDATION_EXPORT NSKeyValueChangeKey const NSKeyValueChangeOldKey;

@interface NSObject (NSKeyValueObserving)
- (void)observeValueForKeyPath:(nullable NSString *)keyPath ofObject:(nullable id)object change:(nullable NSDictionary<NSKeyValueChangeKey, id> *)change context:(nullable void *)context;
- (void)addObserver:(NSObject *)observer forKeyPath:(NSString *)keyPath options:(NSKeyValueObservingOptions)options context:(nullable void *)context;
- (void)removeObserver:(NSObject *)observer forKeyPath:(NSString *)keyPath;
- (void)removeObserver:(NSObject *)observer forKeyPath:(NSString *)keyPath context:(nullable void *)context;
@end

// NSCoder & NSUserActivity

@interface NSCoder : NSObject
- (void)encodeObject:(nullable id)object forKey:(NSString *)key;
- (nullable id)decodeObjectForKey:(NSString *)key;
- (void)encodeBool:(BOOL)value forKey:(NSString *)key;
- (BOOL)decodeBoolForKey:(NSString *)key;
- (void)encodeInteger:(NSInteger)value forKey:(NSString *)key;
- (NSInteger)decodeIntegerForKey:(NSString *)key;
@end

FOUNDATION_EXPORT NSString *const NSUserActivityTypeBrowsingWeb;

@interface NSUserActivity : NSObject
- (instancetype)initWithActivityType:(NSString *)activityType;
@property(readonly, copy) NSString *activityType;
@property(nullable, copy) NSString *title;
@property(nullable, copy) NSDictionary *userInfo;
@property(nullable, copy) NSURL *webpageURL;
@end

// NSFileManager & paths

typedef NS_ENUM(NSUInteger, NSSearchPathDirectory) {
	NSApplicationDirectory = 1,
	NSLibraryDirectory = 5,
	NSUserDirectory = 7,
	NSDocumentDirectory = 9,
	NSDesktopDirectory = 12,
	NSCachesDirectory = 13,
	NSApplicationSupportDirectory = 14,
	NSDownloadsDirectory = 15,
	NSMoviesDirectory = 17,
	NSMusicDirectory = 18,
	NSPicturesDirectory = 19,
};
typedef NS_OPTIONS(NSUInteger, NSSearchPathDomainMask) {
	NSUserDomainMask = 1,
	NSLocalDomainMask = 2,
	NSNetworkDomainMask = 4,
	NSSystemDomainMask = 8,
	NSAllDomainsMask = 0x0ffff,
};

FOUNDATION_EXPORT NSArray<NSString *> *NSSearchPathForDirectoriesInDomains(NSSearchPathDirectory directory, NSSearchPathDomainMask domainMask, BOOL expandTilde);
FOUNDATION_EXPORT NSString *NSTemporaryDirectory(void);
FOUNDATION_EXPORT NSString *NSHomeDirectory(void);
FOUNDATION_EXPORT NSString *NSUserName(void);

FOUNDATION_EXPORT NSFileAttributeKey const NSFileSize;
FOUNDATION_EXPORT NSFileAttributeKey const NSFileModificationDate;
FOUNDATION_EXPORT NSFileAttributeKey const NSFileSystemFreeSize;
FOUNDATION_EXPORT NSFileAttributeKey const NSFileSystemSize;
FOUNDATION_EXPORT NSURLResourceKey const NSURLIsExcludedFromBackupKey;

@interface NSFileManager : NSObject
@property(class, readonly, strong) NSFileManager *defaultManager;
@property(readonly, copy) NSString *currentDirectoryPath;
@property(readonly, copy) NSURL *temporaryDirectory;
- (BOOL)fileExistsAtPath:(NSString *)path;
- (BOOL)fileExistsAtPath:(NSString *)path isDirectory:(nullable BOOL *)isDirectory;
- (BOOL)createDirectoryAtPath:(NSString *)path withIntermediateDirectories:(BOOL)createIntermediates attributes:(nullable NSDictionary<NSFileAttributeKey, id> *)attributes error:(NSError **)error;
- (BOOL)removeItemAtPath:(NSString *)path error:(NSError **)error;
- (BOOL)removeItemAtURL:(NSURL *)URL error:(NSError **)error;
- (BOOL)copyItemAtPath:(NSString *)srcPath toPath:(NSString *)dstPath error:(NSError **)error;
- (BOOL)moveItemAtPath:(NSString *)srcPath toPath:(NSString *)dstPath error:(NSError **)error;
- (BOOL)changeCurrentDirectoryPath:(NSString *)path;
- (nullable NSArray<NSString *> *)contentsOfDirectoryAtPath:(NSString *)path error:(NSError **)error;
- (nullable NSDictionary<NSFileAttributeKey, id> *)attributesOfItemAtPath:(NSString *)path error:(NSError **)error;
- (nullable NSDictionary<NSFileAttributeKey, id> *)attributesOfFileSystemForPath:(NSString *)path error:(NSError **)error;
- (NSArray<NSURL *> *)URLsForDirectory:(NSSearchPathDirectory)directory inDomains:(NSSearchPathDomainMask)domainMask;
- (nullable NSURL *)URLForDirectory:(NSSearchPathDirectory)directory inDomain:(NSSearchPathDomainMask)domain appropriateForURL:(nullable NSURL *)url create:(BOOL)shouldCreate error:(NSError **)error;
- (nullable NSURL *)containerURLForSecurityApplicationGroupIdentifier:(NSString *)groupIdentifier;
@end

@interface NSURL (NSURLResourceValues)
- (BOOL)setResourceValue:(nullable id)value forKey:(NSURLResourceKey)key error:(NSError **)error;
@end

// Functions

NS_INLINE CF_RETURNS_RETAINED CFTypeRef _Nullable CFBridgingRetain(id _Nullable X) {
	return (__bridge_retained CFTypeRef)X;
}
NS_INLINE id _Nullable CFBridgingRelease(CFTypeRef CF_CONSUMED _Nullable X) {
	return (__bridge_transfer id)X;
}

FOUNDATION_EXPORT void NSLog(NSString *format, ...) NS_FORMAT_FUNCTION(1, 2);
FOUNDATION_EXPORT void NSLogv(NSString *format, va_list args) NS_FORMAT_FUNCTION(1, 0);
FOUNDATION_EXPORT NSString *NSStringFromSelector(SEL aSelector);
FOUNDATION_EXPORT SEL NSSelectorFromString(NSString *aSelectorName);
FOUNDATION_EXPORT NSString *NSStringFromClass(Class aClass);
FOUNDATION_EXPORT Class _Nullable NSClassFromString(NSString *aClassName);
FOUNDATION_EXPORT NSString *NSStringFromRange(NSRange range);

NS_ASSUME_NONNULL_END

#endif // __OBJC__

#endif
