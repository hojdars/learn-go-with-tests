package maps

import "errors"

var (
	ErrElementNotFound  = errors.New("element not found")
	ErrWordExists       = errors.New("word already exists")
	ErrWordDoesNotExist = errors.New("word does not exist")
)

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	element, found := d[word]
	if !found {
		return "", ErrElementNotFound
	} else {
		return element, nil
	}
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrElementNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrElementNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(key string) {
	delete(d, key)
}
