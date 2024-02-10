module module_example/hello

go 1.22.0

replace module_example/welcome => ../welcome

require module_example/welcome v0.0.0-00010101000000-000000000000
