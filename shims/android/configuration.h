#ifndef ANDROID_CONFIGURATION_H
#define ANDROID_CONFIGURATION_H

#include <stdint.h>
#include <sys/cdefs.h>

#ifdef __cplusplus
extern "C" {
#endif

struct AConfiguration;
typedef struct AConfiguration AConfiguration;
struct AAssetManager;

enum {
    ACONFIGURATION_ORIENTATION_ANY = 0x0000,
    ACONFIGURATION_ORIENTATION_PORT = 0x0001,
    ACONFIGURATION_ORIENTATION_LAND = 0x0002,
    ACONFIGURATION_ORIENTATION_SQUARE = 0x0003,

    ACONFIGURATION_TOUCHSCREEN_ANY = 0x0000,
    ACONFIGURATION_TOUCHSCREEN_NOTOUCH = 0x0001,
    ACONFIGURATION_TOUCHSCREEN_STYLUS = 0x0002,
    ACONFIGURATION_TOUCHSCREEN_FINGER = 0x0003,

    ACONFIGURATION_DENSITY_DEFAULT = 0,
    ACONFIGURATION_DENSITY_LOW = 120,
    ACONFIGURATION_DENSITY_MEDIUM = 160,
    ACONFIGURATION_DENSITY_TV = 213,
    ACONFIGURATION_DENSITY_HIGH = 240,
    ACONFIGURATION_DENSITY_XHIGH = 320,
    ACONFIGURATION_DENSITY_XXHIGH = 480,
    ACONFIGURATION_DENSITY_XXXHIGH = 640,
    ACONFIGURATION_DENSITY_ANY = 0xfffe,
    ACONFIGURATION_DENSITY_NONE = 0xffff,

    ACONFIGURATION_KEYBOARD_ANY = 0x0000,
    ACONFIGURATION_KEYBOARD_NOKEYS = 0x0001,
    ACONFIGURATION_KEYBOARD_QWERTY = 0x0002,
    ACONFIGURATION_KEYBOARD_12KEY = 0x0003,

    ACONFIGURATION_NAVIGATION_ANY = 0x0000,
    ACONFIGURATION_NAVIGATION_NONAV = 0x0001,
    ACONFIGURATION_NAVIGATION_DPAD = 0x0002,
    ACONFIGURATION_NAVIGATION_TRACKBALL = 0x0003,
    ACONFIGURATION_NAVIGATION_WHEEL = 0x0004,

    ACONFIGURATION_KEYSHIDDEN_ANY = 0x0000,
    ACONFIGURATION_KEYSHIDDEN_NO = 0x0001,
    ACONFIGURATION_KEYSHIDDEN_YES = 0x0002,
    ACONFIGURATION_KEYSHIDDEN_SOFT = 0x0003,

    ACONFIGURATION_NAVHIDDEN_ANY = 0x0000,
    ACONFIGURATION_NAVHIDDEN_NO = 0x0001,
    ACONFIGURATION_NAVHIDDEN_YES = 0x0002,

    ACONFIGURATION_SCREENSIZE_ANY = 0x00,
    ACONFIGURATION_SCREENSIZE_SMALL = 0x01,
    ACONFIGURATION_SCREENSIZE_NORMAL = 0x02,
    ACONFIGURATION_SCREENSIZE_LARGE = 0x03,
    ACONFIGURATION_SCREENSIZE_XLARGE = 0x04,

    ACONFIGURATION_SCREENLONG_ANY = 0x00,
    ACONFIGURATION_SCREENLONG_NO = 0x1,
    ACONFIGURATION_SCREENLONG_YES = 0x2,

    ACONFIGURATION_SCREENROUND_ANY = 0x00,
    ACONFIGURATION_SCREENROUND_NO = 0x1,
    ACONFIGURATION_SCREENROUND_YES = 0x2,

    ACONFIGURATION_WIDE_COLOR_GAMUT_ANY = 0x00,
    ACONFIGURATION_WIDE_COLOR_GAMUT_NO = 0x1,
    ACONFIGURATION_WIDE_COLOR_GAMUT_YES = 0x2,

    ACONFIGURATION_HDR_ANY = 0x00,
    ACONFIGURATION_HDR_NO = 0x1,
    ACONFIGURATION_HDR_YES = 0x2,

    ACONFIGURATION_UI_MODE_TYPE_ANY = 0x00,
    ACONFIGURATION_UI_MODE_TYPE_NORMAL = 0x01,
    ACONFIGURATION_UI_MODE_TYPE_DESK = 0x02,
    ACONFIGURATION_UI_MODE_TYPE_CAR = 0x03,
    ACONFIGURATION_UI_MODE_TYPE_TELEVISION = 0x04,
    ACONFIGURATION_UI_MODE_TYPE_APPLIANCE = 0x05,
    ACONFIGURATION_UI_MODE_TYPE_WATCH = 0x06,
    ACONFIGURATION_UI_MODE_TYPE_VR_HEADSET = 0x07,

    ACONFIGURATION_UI_MODE_NIGHT_ANY = 0x00,
    ACONFIGURATION_UI_MODE_NIGHT_NO = 0x1,
    ACONFIGURATION_UI_MODE_NIGHT_YES = 0x2,

    ACONFIGURATION_SCREEN_WIDTH_DP_ANY = 0x0000,
    ACONFIGURATION_SCREEN_HEIGHT_DP_ANY = 0x0000,
    ACONFIGURATION_SMALLEST_SCREEN_WIDTH_DP_ANY = 0x0000,

    ACONFIGURATION_LAYOUTDIR_ANY = 0x00,
    ACONFIGURATION_LAYOUTDIR_LTR = 0x1,
    ACONFIGURATION_LAYOUTDIR_RTL = 0x2,

    ACONFIGURATION_MCC = 0x0001,
    ACONFIGURATION_MNC = 0x0002,
    ACONFIGURATION_LOCALE = 0x0004,
    ACONFIGURATION_TOUCHSCREEN = 0x0008,
    ACONFIGURATION_KEYBOARD = 0x0010,
    ACONFIGURATION_KEYBOARD_HIDDEN = 0x0020,
    ACONFIGURATION_NAVIGATION = 0x0040,
    ACONFIGURATION_ORIENTATION = 0x0080,
    ACONFIGURATION_DENSITY = 0x0100,
    ACONFIGURATION_SCREEN_SIZE = 0x0200,
    ACONFIGURATION_VERSION = 0x0400,
    ACONFIGURATION_SCREEN_LAYOUT = 0x0800,
    ACONFIGURATION_UI_MODE = 0x1000,
    ACONFIGURATION_SMALLEST_SCREEN_SIZE = 0x2000,
    ACONFIGURATION_LAYOUTDIR = 0x4000,
    ACONFIGURATION_SCREEN_ROUND = 0x8000,
    ACONFIGURATION_COLOR_MODE = 0x10000,
    ACONFIGURATION_GRAMMATICAL_GENDER = 0x20000
};

#define ACONFIGURATION_MNC_ZERO 0xffff

AConfiguration* AConfiguration_new(void) __INTRODUCED_IN(9);
void AConfiguration_delete(AConfiguration* config) __INTRODUCED_IN(9);
void AConfiguration_fromAssetManager(AConfiguration* out, AAssetManager* am) __INTRODUCED_IN(9);
void AConfiguration_copy(AConfiguration* dest, AConfiguration* src) __INTRODUCED_IN(9);
int32_t AConfiguration_getMcc(AConfiguration* config) __INTRODUCED_IN(9);
void AConfiguration_setMcc(AConfiguration* config, int32_t mcc) __INTRODUCED_IN(9);
int32_t AConfiguration_getMnc(AConfiguration* config) __INTRODUCED_IN(9);
void AConfiguration_setMnc(AConfiguration* config, int32_t mnc) __INTRODUCED_IN(9);
void AConfiguration_getLanguage(AConfiguration* config, char* outLanguage) __INTRODUCED_IN(9);
void AConfiguration_setLanguage(AConfiguration* config, const char* language) __INTRODUCED_IN(9);
void AConfiguration_getCountry(AConfiguration* config, char* outCountry) __INTRODUCED_IN(9);
void AConfiguration_setCountry(AConfiguration* config, const char* country) __INTRODUCED_IN(9);
int32_t AConfiguration_getOrientation(AConfiguration* config) __INTRODUCED_IN(9);
void AConfiguration_setOrientation(AConfiguration* config, int32_t orientation) __INTRODUCED_IN(9);
int32_t AConfiguration_getTouchscreen(AConfiguration* config) __INTRODUCED_IN(9);
void AConfiguration_setTouchscreen(AConfiguration* config, int32_t touchscreen) __INTRODUCED_IN(9);
int32_t AConfiguration_getDensity(AConfiguration* config) __INTRODUCED_IN(9);
void AConfiguration_setDensity(AConfiguration* config, int32_t density) __INTRODUCED_IN(9);
int32_t AConfiguration_getKeyboard(AConfiguration* config) __INTRODUCED_IN(9);
void AConfiguration_setKeyboard(AConfiguration* config, int32_t keyboard) __INTRODUCED_IN(9);
int32_t AConfiguration_getNavigation(AConfiguration* config) __INTRODUCED_IN(9);
void AConfiguration_setNavigation(AConfiguration* config, int32_t navigation) __INTRODUCED_IN(9);
int32_t AConfiguration_getKeysHidden(AConfiguration* config) __INTRODUCED_IN(9);
void AConfiguration_setKeysHidden(AConfiguration* config, int32_t keysHidden) __INTRODUCED_IN(9);
int32_t AConfiguration_getNavHidden(AConfiguration* config) __INTRODUCED_IN(9);
void AConfiguration_setNavHidden(AConfiguration* config, int32_t navHidden) __INTRODUCED_IN(9);
int32_t AConfiguration_getSdkVersion(AConfiguration* config) __INTRODUCED_IN(9);
void AConfiguration_setSdkVersion(AConfiguration* config, int32_t sdkVersion) __INTRODUCED_IN(9);
int32_t AConfiguration_getScreenSize(AConfiguration* config) __INTRODUCED_IN(9);
void AConfiguration_setScreenSize(AConfiguration* config, int32_t screenSize) __INTRODUCED_IN(9);
int32_t AConfiguration_getScreenLong(AConfiguration* config) __INTRODUCED_IN(9);
void AConfiguration_setScreenLong(AConfiguration* config, int32_t screenLong) __INTRODUCED_IN(9);
int32_t AConfiguration_getUiModeType(AConfiguration* config) __INTRODUCED_IN(9);
void AConfiguration_setUiModeType(AConfiguration* config, int32_t uiModeType) __INTRODUCED_IN(9);
int32_t AConfiguration_getUiModeNight(AConfiguration* config) __INTRODUCED_IN(9);
void AConfiguration_setUiModeNight(AConfiguration* config, int32_t uiModeNight) __INTRODUCED_IN(9);
int32_t AConfiguration_getScreenWidthDp(AConfiguration* config) __INTRODUCED_IN(13);
void AConfiguration_setScreenWidthDp(AConfiguration* config, int32_t value) __INTRODUCED_IN(13);
int32_t AConfiguration_getScreenHeightDp(AConfiguration* config) __INTRODUCED_IN(13);
void AConfiguration_setScreenHeightDp(AConfiguration* config, int32_t value) __INTRODUCED_IN(13);
int32_t AConfiguration_getSmallestScreenWidthDp(AConfiguration* config) __INTRODUCED_IN(13);
void AConfiguration_setSmallestScreenWidthDp(AConfiguration* config, int32_t value) __INTRODUCED_IN(13);
int32_t AConfiguration_diff(AConfiguration* config1, AConfiguration* config2) __INTRODUCED_IN(9);
int32_t AConfiguration_match(AConfiguration* base, AConfiguration* requested) __INTRODUCED_IN(9);
int32_t AConfiguration_isBetterThan(AConfiguration* base, AConfiguration* test, AConfiguration* requested) __INTRODUCED_IN(9);
void AConfiguration_getLocaleScript(AConfiguration* config, char* outLocaleScript) __INTRODUCED_IN(24);
void AConfiguration_setLocaleScript(AConfiguration* config, const char* localeScript) __INTRODUCED_IN(24);
void AConfiguration_getLocaleVariant(AConfiguration* config, char* outLocaleVariant) __INTRODUCED_IN(24);
void AConfiguration_setLocaleVariant(AConfiguration* config, const char* localeVariant) __INTRODUCED_IN(24);
int32_t AConfiguration_getLayoutDirection(AConfiguration* config) __INTRODUCED_IN(17);
void AConfiguration_setLayoutDirection(AConfiguration* config, int32_t value) __INTRODUCED_IN(17);
int32_t AConfiguration_getScreenRound(AConfiguration* config) __INTRODUCED_IN(23);
void AConfiguration_setScreenRound(AConfiguration* config, int32_t screenRound) __INTRODUCED_IN(23);
int32_t AConfiguration_getColorMode(AConfiguration* config) __INTRODUCED_IN(26);
void AConfiguration_setColorMode(AConfiguration* config, int32_t colorMode) __INTRODUCED_IN(26);
int32_t AConfiguration_getWideColorGamut(AConfiguration* config) __INTRODUCED_IN(26);
void AConfiguration_setWideColorGamut(AConfiguration* config, int32_t wideColorGamut) __INTRODUCED_IN(26);
int32_t AConfiguration_getHdr(AConfiguration* config) __INTRODUCED_IN(26);
void AConfiguration_setHdr(AConfiguration* config, int32_t hdr) __INTRODUCED_IN(26);
int32_t AConfiguration_getGrammaticalGender(AConfiguration* config) __INTRODUCED_IN(34);
void AConfiguration_setGrammaticalGender(AConfiguration* config, int32_t grammaticalGender) __INTRODUCED_IN(34);

#ifdef __cplusplus
}
#endif

#endif // ANDROID_CONFIGURATION_H
