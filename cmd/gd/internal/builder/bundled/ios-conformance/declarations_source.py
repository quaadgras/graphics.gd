import os,re,sys
# usage: declarations_source.py [ios|macos]  (after values_source.py, whose imports it reuses)
SDK=sys.argv[1] if len(sys.argv)>1 else 'ios'
here=os.path.dirname(os.path.abspath(__file__))
B=os.path.join(here, '..', SDK)
files=[]
roots=[B] if SDK=='ios' else [B, os.path.join(here, '..', 'ios')]
for root in roots:
  for d,_,fs in os.walk(root):
    if SDK=='macos' and root.endswith('ios') and ('UIKit' in d or 'OpenGLES' in d or 'CoreMotion' in d): continue
    if 'include/mach-o' in d or 'include/CoreFoundation' in d: continue
    if 'OpenGLES.framework/Headers/ES' in d or '/KHR' in d: continue
    files += [os.path.join(d,f) for f in fs if f.endswith('.h')]
sels=set(); protos=[]
def strip_parens(s):
    prev=None
    while prev!=s:
        prev=s; s=re.sub(r'\([^()]*\)','',s)
    return s
prefix=re.compile(r'^(CF_EXPORT|FOUNDATION_EXPORT|FOUNDATION_EXTERN|UIKIT_EXTERN|GAMECONTROLLER_EXPORT|GAMECONTROLLER_EXTERN)\s+')
for path in sorted(files):
    for line in open(path):
        l=line.strip()
        if re.match(r'^[-+]\s*\(',l):
            body=strip_parens(l[1:]).split(';')[0]
            body=re.sub(r'\b(NS_[A-Z_]+|API_AVAILABLE|__attribute__)\b.*$','',body).strip()
            parts=re.findall(r'([A-Za-z_][A-Za-z0-9_]*)\s*:',body)
            if parts: sels.add(''.join(p+':' for p in parts))
            else:
                m=re.match(r'([A-Za-z_][A-Za-z0-9_]*)',body)
                if m: sels.add(m.group(1))
        elif l.startswith('@property'):
            attrs=re.match(r'@property\s*(\(([^)]*)\))?',l).group(2) or ''
            decl=strip_parens(l[len('@property'):]).split(';')[0]
            decl=re.sub(r'\b(NS_[A-Z_]+|API_AVAILABLE)\b.*$','',decl)
            names=re.findall(r'([A-Za-z_][A-Za-z0-9_]*)\s*$',decl.strip())
            if not names: continue
            name=names[0]
            g=re.search(r'getter\s*=\s*([A-Za-z_][A-Za-z0-9_]*)',attrs)
            sels.add(g.group(1) if g else name)
            if 'readonly' not in attrs: sels.add('set'+name[0].upper()+name[1:]+':')
        elif prefix.match(l) and l.endswith(';'):
            p=prefix.sub('extern ',l)
            p=re.sub(r'\s*(CF_RETURNS_RETAINED|NS_RETURNS_RETAINED|NS_FORMAT_FUNCTION\([^)]*\)|API_AVAILABLE\((?:[^()]|\([^()]*\))*\))','',p)
            protos.append(p)
imports=[l for l in open('conformance.m') if l.startswith('#i')] # written by values_source.py
out=imports+['','// C functions & constants, redeclared as graphics.gd declares them: a conflicting type is an error.','_Pragma("clang assume_nonnull begin")']
out+=[p+'\n' for p in protos]+['_Pragma("clang assume_nonnull end")','','// Every selector graphics.gd declares, -Wundeclared-selector reports the ones nothing in the SDK does.','void gd_selectors(void) {']
out+=[f'\t(void)@selector({s});\n' for s in sorted(sels)]+['}','']
open('conformance2.m','w').write(''.join(x if x.endswith('\n') else x+'\n' for x in out))
print(len(sels),'selectors',len(protos),'prototypes')
