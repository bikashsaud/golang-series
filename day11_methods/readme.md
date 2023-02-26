> METHODS

> In Go, a method is a function associated with a particular type. 

>   Methods are defined using the func keyword, followed by the receiver type and a method name, a set of parentheses containing zero or more parameters, and a block of code enclosed in curly braces.

```


type Rectangle struct {
    width  float64
    height float64
}

func (r Rectangle) area() float64 {
    return r.width * r.height
}

func main() {
    rect := Rectangle{width: 3, height: 4}
    fmt.Println("area:", rect.area())
}
```



