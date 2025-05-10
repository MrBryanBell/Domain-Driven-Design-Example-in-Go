package BlogPost

import BlogPostDescription "mymodule/experiments/survey_practice/Blog/BlogPost/BlogPost__Description"

type BlogPost struct {
	title       string
	description *BlogPostDescription.BlogPostDescription
}

type BlogPostProps struct {
	Title       string
	Description string
}

func Create(props BlogPostProps) (*BlogPost, error) {
	var description, error = BlogPostDescription.New(props.Description)
	if error != nil {
		return nil, error
	}

	return &BlogPost{
		title:       props.Title,
		description: description,
	}, nil
}

func (self BlogPost) HasDescription() bool {
	return self.description != nil
}

func (self BlogPost) UnwrapDescription() string {
	return self.description.Value()
}
