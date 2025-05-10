package BlogPostDescription

type BlogPostDescription struct {
	description string
}

func New(description string) (*BlogPostDescription, error) {
	return &BlogPostDescription{description}, nil
}

func (self BlogPostDescription) isEmpty() bool {
	return self.description == ""
}

func (self BlogPostDescription) Value() string {
	return self.description
}
