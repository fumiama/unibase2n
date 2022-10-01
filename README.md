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

### Base256
```
12345678 -> ᄱᄲᄳᄴᄵᄶᄷᄸ
```

### Base128
```
1234567 -> ⑸⒬⒦⒓⒁⒴Ⓦ⒗
```

### Base64Gua
```
123456 -> ䷌䷓䷈䷳䷍䷃䷔䷶
```

### Base32
```
12345 -> ▆▄▙▃▆▍▁▕
```

### Base16
```
1234 -> ㆓㆑㆓㆒㆓㆓㆓㆔
```

### Base8Gua
```
123456 -> ☱☴☲☳☱☰☶☳☱☵☰☳☲☴☶☶
```
