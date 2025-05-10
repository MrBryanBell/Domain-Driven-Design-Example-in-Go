package Survey

import "errors"

type Survey struct {
	title string
}

const TITLE_MINIMUM_CHARACTERS int = 12
const TITLE_MAXIMUM_CHARACTERS int = 200

func Create(title string) (*Survey, error) {

	if len(title) < TITLE_MINIMUM_CHARACTERS {
		return nil, errors.New("Survey Title should have at least 5 characters")
	}

	if len(title) > TITLE_MAXIMUM_CHARACTERS {
		return nil, errors.New("Survey Title should have less than 200 characters")
	}

	return &Survey{
		title,
	}, nil
}
