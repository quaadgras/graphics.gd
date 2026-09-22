import re,sys
cur=None; vals={}
for line in open(sys.argv[1]):
    m=re.match(r'^_(gd[vso]_[A-Za-z0-9_]+):',line)
    if m: cur=m.group(1); continue
    m=re.match(r'^\s*\.(quad|long|space|zero)\s+(\S+)',line)
    if cur and m:
        v=m.group(2)
        vals[cur]= 0 if m.group(1) in ('space','zero') else int(v,0)
        cur=None
for k in sorted(vals): print(k, vals[k] & 0xFFFFFFFFFFFFFFFF)
