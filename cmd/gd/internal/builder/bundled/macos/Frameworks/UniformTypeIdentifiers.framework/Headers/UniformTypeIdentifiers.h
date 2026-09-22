// UniformTypeIdentifiers declarations for graphics.gd's macOS SDK.
//
// Written from the public documentation of the API, covering what Godot and
// SDL make use of. Not derived from Apple's SDK headers.
#ifndef GD_UNIFORMTYPEIDENTIFIERS_H
#define GD_UNIFORMTYPEIDENTIFIERS_H

#import <Foundation/Foundation.h>

NS_ASSUME_NONNULL_BEGIN

API_AVAILABLE(macos(11.0), ios(14.0))
@interface UTType : NSObject <NSCopying, NSSecureCoding>
+ (nullable UTType *)typeWithIdentifier:(NSString *)identifier;
+ (nullable UTType *)typeWithFilenameExtension:(NSString *)filenameExtension;
+ (nullable UTType *)typeWithFilenameExtension:(NSString *)filenameExtension conformingToType:(UTType *)supertype;
+ (nullable UTType *)typeWithMIMEType:(NSString *)mimeType;
+ (nullable UTType *)typeWithMIMEType:(NSString *)mimeType conformingToType:(UTType *)supertype;
@property(readonly) NSString *identifier;
@property(nullable, readonly) NSString *preferredFilenameExtension;
@property(nullable, readonly) NSString *preferredMIMEType;
@property(nullable, readonly) NSString *localizedDescription;
@property(readonly, getter=isDynamic) BOOL dynamic;
@property(readonly, getter=isDeclared) BOOL declared;
@property(readonly, getter=isPublicType) BOOL publicType;
- (BOOL)conformsToType:(UTType *)type;
@end

FOUNDATION_EXTERN UTType *const UTTypeItem;
FOUNDATION_EXTERN UTType *const UTTypeContent;
FOUNDATION_EXTERN UTType *const UTTypeData;
FOUNDATION_EXTERN UTType *const UTTypeDirectory;
FOUNDATION_EXTERN UTType *const UTTypeFolder;
FOUNDATION_EXTERN UTType *const UTTypeText;
FOUNDATION_EXTERN UTType *const UTTypePlainText;
FOUNDATION_EXTERN UTType *const UTTypeImage;
FOUNDATION_EXTERN UTType *const UTTypePNG;
FOUNDATION_EXTERN UTType *const UTTypeJPEG;
FOUNDATION_EXTERN UTType *const UTTypeURL;
FOUNDATION_EXTERN UTType *const UTTypeFileURL;

NS_ASSUME_NONNULL_END

#endif
