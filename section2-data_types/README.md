# Data Types :
## Introduction of memory :
- Stack memory 
  - last in, first out. use for primitive types of data
  - this called __value type__
- Heap memory 
  - use for objects, their reference stores in stack.
  - `gorbage collector` free the unused data
  - this called __reference type__

## Basic data types : 
```
var number int8 
fmt.Printf("size of my variable %d", unsafe.Sizeof())
```
> This shows us the number of bytes (a byte is 8 bit)
```
var a = bits.UintSize
fmt.Println(a)
```
> This code shows us the size of bits

- Numbers : 
  - Integer
    - signed
      - int 
      - int8
      - int16
      - int32
      - int64
    - unsigned
      - uint
      - uint8
      - uint16
      - uint32
      - uint64
  - Float
    - float32
    - float64
  - Complex Numbers
    - an actual part and an unknown part
    - complexNumber64
    - complexNumber128
- Strings 
- Booleans

## Composite data types :
- Arrays 
  - the size of that is static
  - if we copy that with `:=`, that don't change but if we pointed that with `:= &`, that change.
- Slice 
  - that are dynamic size
  - that stores in heap and point a reference in start member
  - we can use `len` for number of full member and `cap` for the size of start index until end index of origin list.
- Map 
  - key value sets
  - like dict and HashTable
  - it is very faster
- Struct
  - make a custom data type 
  - fields and properties are store back to back
- Pointer
  - it uses to call by reference a value
  - use `*` to show the amount of a variable or being pointer that
  - use `&` to access the address of that
- Interface
  - this is a contract and you should obey it.
  - you should implement its methods
- Channels
  - send and receive values between `Goroutine` 
  - the connection between `Goroutines`

## Reference vs. Value types :
- Value types
  - int, float, string, boolean, struct
- Reference types
  - slices, maps, channels, pointers, functions

## Enum and Const introduction
- Const types
  - you can't change that value
- Enum types
  - is a variable that value is a const set
  - ```
    type status int
    const (
      status1   status = 1
      pay       status = 2
      strongest status = 3
      status4   status = 4
      status5   status = 5
    )
    ```
> __Tip1 :__
>> __You should define your function names as Camelcase.__
> 
> __Tip2 :__
>> __`json.Marshal()` printer returns error or that value__
>>
>> ```
>> state1 := state{number: 1, presentState: pay}
>> stateJson1, _ := json.Marshal(state1)
>> fmt.Println(string(stateJson1))
>> ```

## Pointer introduction
- the difference between cal by value and call by reference
> __Tip1 :__
>> __use `:=` if you don't define before and want to assess a value for that__
>>
>> __use `=` if you define it before and want to assess a value for that__
> ```
> i, j := 1, 2
> var pi, pj *int = &i, &j
> ```

## Rune introduction
- to store the characters that contain 1 byte to 4 byte.
- that is a `4 byte int variable`
- the common characters store 1 byte (ASCII characters)
- the special character store 4 bite (another language)
  - we can't use `len in string` to move on it.
> __Tip1 :__
>> __the application of `%T` is returning the type a variable
