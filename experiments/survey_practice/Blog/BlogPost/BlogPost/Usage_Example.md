### Accessing a BlogPost WITH validation

```go
var post, _ = BlogPost.Create(BlogPost.BlogPostProps{
	Title:       "How to measure productivity?",
	Description: "Find out how to measure productivity...",
})

if post.HasDescription() {
	fmt.Println("Here is the description:", post.UnwrapDescription())
}
```


### Accessing a BlogPost WITHOUT validation

```go
var post, _ = BlogPost.Create(BlogPost.BlogPostProps{
	Title: "How to measure productivity?",
	Description: "Find out how to measure productivity...",
})

fmt.Println("Here is the description:", post.UnwrapDescription())
```
