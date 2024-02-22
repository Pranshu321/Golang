# Golang

- No switch case
- No function overloading , overriding
- No generics
- It is a compiled language
- `lexer` does a lot of work
- It is written we have to write semi-colon but can skip that.

## Types in Golang

- First Capital letter means it is public.
- Everything is a `Type`
- Variable type should be known in advance
  ### TYPES
  - `Bool`
  - `floating` : float32 , float64
  - `string`
  - `Integer` : uint8 , uint64 , int8 , int64 , uintptr
  - `Complex`
  - `Array`
  - `Slices`
  - `Struct`
  - `Interface`
  - `Map`
  - `Pointers`

**Alomost everything is a FUNCTION OR CHANNELS**

## Build

`go build`

### We can create build for every system

```bash
GOOS="darwin" go build

OR

GOOS="linux" go build
```

## Memory Management

- **new() :** Allocate memory but no INIT , Zeroed storage
- **make() :** Allocate memory and INIT , Non-zeroed storage

- Garbage Collector automatically
- Out of scope or nil
