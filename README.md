# Go Language Beginner readme.

Go is a statically and Strongly typed language and also its a compiled language unlike JS, Python(dynamically typed). It also supports
concurrency model, It also supports type system out of the box(like typescript for javascript). Since it is compiled language
it turns our code into machine code quickly which makes the application faster on the run time. It also has
built in garbage collection.

## Go is Compiled Language

Unlike python, javascript etc Go is a compiled language, meaning it will compile the human readable code and converts into machine code
where that output machine code alone can be transfered to other people or put in servers to execute.
Unlike hpython or js, other people or servers dont need js runtime(node) or python runtime, since in go we are transferring the direct machine code
the OS(platform specific) can able to execute directly.

## Memory management in Go

Python, javascript etc are managed by garbage collector(Automatic memory management), and the code(Garbage collector code) that needs to mange
this memory will be in the python or js runtime.
But languages like C, C++ these are manual memory management, and needs to handled manually in our program.
Go is also managed by garbage collector(Automatic memory management).
Now there arise a question, since go is a compiled language like C, C++ etc, where does the Garabge collector code will be present 
The answer is `In every Go program when we are compiling, the garbage collector code(Go Runtime) will also be compiled along with it`


Example for Statically and strongly typed language
`const str: string; const num: number; console.log(str + ":" + num) => this will work in typescript(only statically typed)`
`var str string; var num int; println(str + ":" + num) => will not work in go(since it is both statically and strong type)`


To summarize:

    - Fast.
    - Statically type.
    - Concurrency.
    - type safety.
    - Compiled Language.

## Introduction to Go

To start with the files that are ending .go is considered as go lang file, similar to package.json
go.mod file situated in the root of working directory is the mod file which has the module name.
eg: 
`module "MODULE_NAME`

    - This MODULE_NAME is used when this file is being used as an external module by some other go code.
    - Similar to npm publish package

`package` is another important concept in go, any file with package main is kinda root file on that working directory
also we can have a main function which will be defaultly called when running the go run FILE_NAME/regexpath
command.

We can also use some inbuilt standard library functions like fmt for printing, getting some inputs from 
command line.

EG: go lang folder structure(From noob's perspective :( )

    $HOME/
        go.mod => where it contains the module name of this folder(may or may not be used as external module by others)
        filename.go => main go file
    running this file => either we can use go run filename.go / go run .(from that folder, which can trigger the main function)

Some initial observations
    
    - Package name can start with either main / working_directory_folder name
    - Each package should have a single name
    Eg:
        golang
            -> go.mod => module "golang/test"
            -> initial.go => can have package has main / golang
                NOTE: package nested can be imported here if it packaged has nested,
            nested
                -> nested.go => can have package has main / nested

Go Variables

    - bool => boolean type
    - string => string type
    - int(int, int8, int16, int32, int64) => integer(positive/negative whole numbers) type with various bits
    - uint(unit, unint8, unit16, unit32, unit64) => unsigned(positive whole numbers) integer type with various bits
    - float(float32, float64) => fraction numbers with various bits
    - complex => imaginary numbers with various bits(very rare)
    - byte => alias for unit8(basically one byte = 8bits)
    - rune => alias for int32
    - Go variables are pass by value, not by reference

    Unless any application needs extreme performance care, otherwise its ok to use the below variables
    - int, uint, float64,complex128
    
    Variables can also be typecasted for EG:
    - age := 12; ageFloat := float64(age);

Go also supports `const` for immutable type of variables, constant variables doesn't support `:=`syntax

Go Print Statement
    Go follows similar approach to C language of using printf and sprintf

    - printf => prints the formatted string
    - sprintf => returns the formatted string(will not print in stdoutput)
    - %v, %s => replaces %v with argument value EG: printf("Hi %v is manikandan", "this")
    - %d => interpolate for integer in decimal form
    - %f => interpolate for float
    - %.2f => interpolate to float with rounding 2.
TIP: Sprintf is similar to `\`` in js


## Structs

Structs are a collection of type(basically a js object literal ;)
EG:
    `type student Struct {
        name string
        age int
    }`
Structs can also be nested inside anothe Struct

    - Embedded Structs: These are similar to & in type in typescript.

## Interfaces

    Interfaces are similar to typescript except that the fact in Go, interfaces act as function signature
EG: `type message interface { getMessage() string }`
    Whomever structs implemented this interface will automatically inherits this interface

## slices

