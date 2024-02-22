# Golang

**This repository contains various tutorials and learning materials for Golang, For quick overlook you can look into [cheatsheet](Golang%20cheatsheet.pdf)**

## Table of Contents

- [01HelloWorld](01HelloWorld)
- [02Variables](02Variables)
- [03Input](03Input)
- [04Conversions](04Conversions)
- [05DateandTime](05DateandTime)
- [07mypointers](07mypointers)
- [08Array](08Array)
- [09Slices](09Slices)
- [10map](10map)
- [11Structs](11Structs)
- [12IfElse](12IfElse)
- [13Functions](13Functions)
- [14Defer](14Defer)
- [15Files](15Files)
- [16WebRequests](16WebRequests)
- [17HandlingURLS](17HandlingURLS)
- [18Server](18Server)
- [20BitMoreJSON](20BitMoreJSON)
- [21BuildAPI](21BuildAPI)
- [22modImportance](22modImportance)
- [23Mongo](23Mongo)
- [24GoRoutines](24GoRoutines)
- [25RaceCondition](25RaceCondition)
- [26Deadlock](26Deadlock)
- [27Maths](27Maths)
- [Web Server](Webserver)


## Important Things to remember

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
