// CoreText declarations for graphics.gd's iOS SDK.
//
// Written from the public documentation of the API, covering what Godot and
// SDL make use of. Not derived from Apple's SDK headers.
#ifndef GD_CORETEXT_H
#define GD_CORETEXT_H

#include <CoreFoundation/CoreFoundation.h>
#include <CoreGraphics/CoreGraphics.h>

CF_EXTERN_C_BEGIN

typedef const struct CF_BRIDGED_TYPE(id) __CTFont *CTFontRef;
typedef const struct CF_BRIDGED_TYPE(id) __CTFontDescriptor *CTFontDescriptorRef;

typedef CF_OPTIONS(uint32_t, CTFontSymbolicTraits) {
	kCTFontTraitItalic = (1 << 0),
	kCTFontTraitBold = (1 << 1),
	kCTFontTraitExpanded = (1 << 5),
	kCTFontTraitCondensed = (1 << 6),
	kCTFontTraitMonoSpace = (1 << 10),
	kCTFontTraitVertical = (1 << 11),
	kCTFontTraitUIOptimized = (1 << 12),
	kCTFontTraitColorGlyphs = (1 << 13),
	kCTFontTraitComposite = (1 << 14),
	kCTFontItalicTrait = kCTFontTraitItalic,
	kCTFontBoldTrait = kCTFontTraitBold,
	kCTFontExpandedTrait = kCTFontTraitExpanded,
	kCTFontCondensedTrait = kCTFontTraitCondensed,
	kCTFontMonoSpaceTrait = kCTFontTraitMonoSpace,
};

CF_EXPORT const CFStringRef kCTFontURLAttribute;
CF_EXPORT const CFStringRef kCTFontNameAttribute;
CF_EXPORT const CFStringRef kCTFontDisplayNameAttribute;
CF_EXPORT const CFStringRef kCTFontFamilyNameAttribute;
CF_EXPORT const CFStringRef kCTFontStyleNameAttribute;
CF_EXPORT const CFStringRef kCTFontTraitsAttribute;
CF_EXPORT const CFStringRef kCTFontSymbolicTrait;
CF_EXPORT const CFStringRef kCTFontWeightTrait;
CF_EXPORT const CFStringRef kCTFontWidthTrait;
CF_EXPORT const CFStringRef kCTFontSlantTrait;

CF_EXPORT CTFontDescriptorRef CTFontDescriptorCreateWithAttributes(CFDictionaryRef attributes) CF_RETURNS_RETAINED;
CF_EXPORT CTFontDescriptorRef CTFontDescriptorCreateWithNameAndSize(CFStringRef name, CGFloat size) CF_RETURNS_RETAINED;
CF_EXPORT CFTypeRef CTFontDescriptorCopyAttribute(CTFontDescriptorRef descriptor, CFStringRef attribute) CF_RETURNS_RETAINED;
CF_EXPORT CFArrayRef CTFontDescriptorCreateMatchingFontDescriptors(CTFontDescriptorRef descriptor, CFSetRef mandatoryAttributes) CF_RETURNS_RETAINED;
CF_EXPORT CTFontRef CTFontCreateWithName(CFStringRef name, CGFloat size, const CGAffineTransform *matrix) CF_RETURNS_RETAINED;
CF_EXPORT CTFontRef CTFontCreateWithFontDescriptor(CTFontDescriptorRef descriptor, CGFloat size, const CGAffineTransform *matrix) CF_RETURNS_RETAINED;
CF_EXPORT CTFontRef CTFontCreateForString(CTFontRef currentFont, CFStringRef string, CFRange range) CF_RETURNS_RETAINED;
CF_EXPORT CTFontDescriptorRef CTFontCopyFontDescriptor(CTFontRef font) CF_RETURNS_RETAINED;
CF_EXPORT CFStringRef CTFontCopyFamilyName(CTFontRef font) CF_RETURNS_RETAINED;
CF_EXPORT CFArrayRef CTFontManagerCopyAvailableFontFamilyNames(void) CF_RETURNS_RETAINED;
CF_EXPORT CFArrayRef CTFontManagerCopyAvailablePostScriptNames(void) CF_RETURNS_RETAINED;

CF_EXTERN_C_END

#endif
