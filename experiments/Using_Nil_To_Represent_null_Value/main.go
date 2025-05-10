package main

import "fmt"

type BlogPost struct {
	title       string
	description *string
}

func (self BlogPost) HasDescription() bool {
	return self.description != nil
}

func main() {
	{
		// Creating a BlogPost WITHOUT description
		var post = BlogPost{
			title: "How to measure productivity?",
		}

		fmt.Println("Does the post has a description?", post.HasDescription())
	}

	{
		// Creating a BlogPost WITH description
		var description = "Find out how to measure productivity..."

		var post = BlogPost{
			title:       "How to measure productivity?",
			description: &description,
		}

		fmt.Println("Does the post has a description?", post.HasDescription())
	}
}
