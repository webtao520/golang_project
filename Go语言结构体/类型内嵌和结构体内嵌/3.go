var c Color
c.BasicColor.R = 1
c.BasicColor.G = 1
c.BasicColor.B = 0
一个结构体只能嵌入一个同类型的成员，无须担心结构体重名和错误赋值的情况，编译器在发现可能的赋值歧义时会报错。