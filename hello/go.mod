module hello

go 1.18

//这里用来引用相对路径模块
replace greetings => ../greetings  

require greetings v0.0.0-00010101000000-000000000000
