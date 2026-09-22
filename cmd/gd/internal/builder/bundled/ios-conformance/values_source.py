import os,re,sys
# usage: values_source.py [ios|macos]   (macos covers the macos bundle and the ios frameworks it shares)
SDK=sys.argv[1] if len(sys.argv)>1 else 'ios'
here=os.path.dirname(os.path.abspath(__file__))
B=os.path.join(here, '..', SDK)
files=[]
roots=[B] if SDK=='ios' else [B, os.path.join(here, '..', 'ios')]
for root in roots:
    for d,_,fs in os.walk(root):
        if 'OpenGLES.framework/Headers/ES' in d or '/KHR' in d or 'include/mach-o' in d or 'include/CoreFoundation' in d: continue
        if SDK=='macos' and root.endswith('ios') and ('UIKit' in d or 'OpenGLES' in d or 'CoreMotion' in d): continue
        for f in fs:
            if f.endswith('.h'): files.append(os.path.join(d,f))
consts=[]; structs={}
enum_start=re.compile(r'(NS_ENUM|NS_OPTIONS|NS_CLOSED_ENUM|NS_ERROR_ENUM|CF_ENUM|CF_OPTIONS)\(|^\s*enum\b[^;]*\{')
enumerator=re.compile(r'^\s*([A-Za-z_][A-Za-z0-9_]*)\s*(=\s*[^,]+)?,?\s*(//.*)?$')
struct_start=re.compile(r'^\s*(?:typedef\s+)?struct\s*([A-Za-z_][A-Za-z0-9_]*)?\s*\{')
field=re.compile(r'([A-Za-z_][A-Za-z0-9_]*)\s*(?:\[[^\]]*\])?\s*(?=[,;])')
define=re.compile(r'^#define\s+([A-Za-z_][A-Za-z0-9_]*)\s+\(?\(?[A-Za-z_ ]*\)?\s*(-?(?:0x[0-9A-Fa-f]+|\d+)U?L?L?)\)?\s*$')
for path in sorted(files):
    mode=None; cur=None; fields=[]; iphone=0; depth=[]
    for line in open(path):
        # declarations under #if TARGET_OS_IPHONE are not the macOS SDK's.
        st=line.strip()
        if st.startswith('#if'):
            depth.append('TARGET_OS_IPHONE' in st)
            if depth[-1]: iphone+=1
        elif st.startswith('#endif') and depth:
            if depth.pop(): iphone-=1
        if iphone and SDK=='macos': continue
        if mode is None:
            m=define.match(line)
            if m and not m.group(1).startswith(('GD_','_')) and m.group(1) not in ('TRUE','FALSE','CGFLOAT_IS_DOUBLE','CGFLOAT_DEFINED','GL_GLEXT_PROTOTYPES'): consts.append(m.group(1))
            if enum_start.search(line) and '{' in line: mode='enum'; continue
            m=struct_start.match(line)
            if m and '@' not in line: mode='struct'; cur=m.group(1); fields=[]; continue
        elif mode=='enum':
            if '}' in line: mode=None; continue
            m=enumerator.match(line)
            if m: consts.append(m.group(1))
        elif mode=='struct':
            if '}' in line:
                m=re.search(r'\}\s*([A-Za-z_][A-Za-z0-9_]*)\s*;',line)
                name=(m.group(1) if m else None) or (('struct '+cur) if cur else None)
                if name and not name.startswith('struct _NS') and '(' not in ''.join(fields): structs[name]=fields
                mode=None; continue
            if '(' in line or line.strip().startswith('//'): continue
            decl=line.split('//')[0]
            if ';' in decl:
                names=field.findall(decl)
                fields += [n for n in names if not n.isupper()]
consts=sorted(set(consts)-{'FFERR_DEVICENOTREG','REGDB_E_CLASSNOTREG','kIOMasterPortDefault','kIOMainPortDefault'}) # which Apple's header defines in terms of a macro it does not declare.
macos_imports=['// Generated: every constant, struct size and field offset graphics.gd\'s macOS SDK declares.',
 '#import <Cocoa/Cocoa.h>','#import <Carbon/Carbon.h>','#import <IOKit/IOKitLib.h>','#import <IOKit/hid/IOHIDLib.h>','#import <IOKit/pwr_mgt/IOPMLib.h>','#import <IOKit/usb/USBSpec.h>','#import <IOKit/hidsystem/ev_keymap.h>','#import <IOKit/IOCFPlugIn.h>','#import <ForceFeedback/ForceFeedback.h>','#import <ForceFeedback/ForceFeedbackConstants.h>','#import <AVFoundation/AVFoundation.h>','#import <CoreAudio/AudioHardware.h>','#import <CoreMIDI/CoreMIDI.h>','#import <CoreVideo/CoreVideo.h>','#import <QuartzCore/CAMetalLayer.h>','#import <Metal/Metal.h>','#import <AVFoundation/AVFoundation.h>','#import <CoreMedia/CoreMedia.h>','#import <CoreHaptics/CoreHaptics.h> // after AVFoundation, which Apple\'s CoreHaptics expects.','#import <GameController/GameController.h>','#import <CoreText/CoreText.h>','#import <AudioToolbox/AudioToolbox.h>','#import <IOSurface/IOSurfaceRef.h>','#import <Security/Security.h>','#import <UniformTypeIdentifiers/UniformTypeIdentifiers.h>','#include <os/log.h>','#include <os/signpost.h>','#include <sys/xattr.h>','#include <stddef.h>','']
out=['// Generated: every constant, struct size and field offset graphics.gd\'s iOS SDK declares.',
 '#import <CoreFoundation/CoreFoundation.h>','#import <CoreGraphics/CoreGraphics.h>','#import <Foundation/Foundation.h>','#import <QuartzCore/QuartzCore.h>','#import <QuartzCore/CAMetalLayer.h>','#import <Metal/Metal.h>','#import <UIKit/UIKit.h>','#import <AVFoundation/AVFoundation.h>','#import <CoreMedia/CoreMedia.h>','#import <CoreVideo/CoreVideo.h>','#import <CoreMotion/CoreMotion.h>','#import <CoreHaptics/CoreHaptics.h>','#import <GameController/GameController.h>','#import <CoreText/CoreText.h>','#import <AudioToolbox/AudioToolbox.h>','#import <IOSurface/IOSurfaceRef.h>','#import <OpenGLES/EAGL.h>','#import <OpenGLES/EAGLDrawable.h>','#include <os/log.h>','#include <os/signpost.h>','#include <sys/xattr.h>','#include <stddef.h>','']
if SDK=='macos': out=macos_imports
for c in consts: out.append(f'const long long gdv_{c} = (long long)({c});')
for s,fs in sorted(structs.items()):
    tag=s.replace('struct ','')
    out.append(f'const long long gds_{tag} = (long long)sizeof({s});')
    for f in fs: out.append(f'const long long gdo_{tag}__{f} = (long long)offsetof({s}, {f});')
open('conformance.m','w').write('\n'.join(out)+'\n')
print(len(consts),'constants',len(structs),'structs',sum(len(v) for v in structs.values()),'fields')
