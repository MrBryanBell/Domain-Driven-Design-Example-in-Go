# Simulating Null Value using Nil

En la mayoría de lenguajes de programación podemos representar la ausencia de valor con algún primitivo:
- Rust (Option<T>)
- Javascript (null)
- Python (None)

Sin embargo, en Golang no existe un primitivo para representar la ausencia de valor. 
A common way to represent an optional value in Go is to use a pointer. If the pointer is nil, it indicates the value is absent.

```go
type BlogPost struct {
	title       string
	description *string
}

func (self BlogPost) HasDescription() bool {
	return self.description != nil
}
```

Gracias a esta decisión de diseño, podemos crear un BlogPost sin descripción. Ten cuidado al intentar acceder al valor de un puntero, si el valor almacenado es nil, se producirá un error.

```go
func main() {
	
	// Creating a BlogPost WITH description
	var description = "Find out how to measure productivity..."

	var post = BlogPost{
		title:       "How to measure productivity?",
		description: &description,
	}

	fmt.Println("Does the post has a description?", post.HasDescription())
	
}
```

```go
func main() {

	// Creating a BlogPost WITH description
	var description = "Find out how to measure productivity..."

	var post = BlogPost{
		title:       "How to measure productivity?",
		description: &description,
	}

	fmt.Println("Does the post has a description?", post.HasDescription())
	
}
```