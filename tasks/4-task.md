Crate SumOf(elems ...any)float32

x


if each element is any type of number 
if elem is bool and true add 1 , false =0 
if elem is string, do parseInt and parsefloat add it


var(
    n1 uint8 =12
    n2 uint64 =12312
    n3 float32 =1232.23
    n4 float64 =12312.12312312
    n5 bool = true // add 1 
    n6 bool = false 
    n7 string="1232" // add 1232
    n8 string="12312.123"
    n9 any = 1123
    n10 any = 123.123
    n11 string= "1232o"
)
