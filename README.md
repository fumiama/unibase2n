# unibase2n
generate base2^n starting from any position in unicode table

## Interface
```go
// NewBase generates a new base2n config
func NewBase(off, til uint16, bit uint8) (bs Base, err error)
```
- `off` is the starting offset in unicode table
- `til` is the starting offset of the tail character, which could not cross with off area
- `bit` is the number `n` of `2^n`, for example bit `6` means `2^6=64`, thus `base64`

## Supported Base2n
> see more in define.go

### Base16384
> see https://github.com/fumiama/base16384
```
1234567 -> 婌焳廔萷
```
### Base8192
```
12345678 -> 눦듌옚뽣며찈
```
### Base4096
```
123456 -> 㜓㘳㝃㤶
```

### Base512
```
123456789 -> ᑢᓈᖙᕃᒦᖍᖜᐹ
```

### Base256
```
12345678 -> ᄱᄲᄳᄴᄵᄶᄷᄸ
```

### BaseTanWi
```
12345678 -> ㌱㌲㌳㌴㌵㌶㌷㌸
```

### Base128
```
1234567 -> ⑸⒬⒦⒓⒁⒴Ⓦ⒗
```
### BaseDevanagari
```
1234567 -> घौॆळड॔६ष
```

### Base64Gua
```
123456 -> ䷌䷓䷈䷳䷍䷃䷔䷶
```

### BaseRune
```
123456 -> ᚬᚳᚨᛓᚭᚣᚴᛖ
```

### BaseMongolian
```
123456 -> ᠬᠳᠨᡓᠭᠣᠴᡖ
```

### Base32
```
12345 -> ▆▄▙▃▆▍▁▕
```

### BaseTibetan
```
12345 -> ཏཌྷརཌཏབཊཞ
```

### Base16
```
1234 -> ㆓㆑㆓㆒㆓㆓㆓㆔
```

### BaseBuginese
```
1234 -> ᨃᨁᨃᨂᨃᨃᨃᨄ
```

### Base8Gua
```
123456 -> ☱☴☲☳☱☰☶☳☱☵☰☳☲☴☶☶
```
