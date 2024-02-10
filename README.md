# Go Language Beginner readme

Go is a statically typed language and also its a compiled language unlike JS, Python(dynamically typed). It also supports
concurrency model, It also supports type system out of the box(like typescript for javascript). Since it is compiled language
it turns our code into machine code quickly which makes the application faster on the run time. It also has
built in garbage collection.

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

