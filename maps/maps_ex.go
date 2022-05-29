package maps

var dictionary = make(map[string]string)

type Dictionary map[string]string

const (
	ErrNotFound   = DictionaryErr("could not find the word you were looking for")
	ErrWordExists = DictionaryErr("cannot add word because it already exists")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	/*
		we are using an interesting property of the map lookup. It can return 2 values.
		The second value is a boolean which indicates if the key was found successfully.
	*/
	s, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return s, nil
}

func Init() {
	dictionary["test"] = "default test value"
}

func Search(dict map[string]string, word string) string {
	if dict == nil {
		return dictionary[word]
	}
	return dict[word]
}

/*
An interesting property of maps is that you can modify them without passing as an address to it (e.g &myMap)
This may make them feel like a "reference type", but as Dave Cheney describes they are not.

"A map value is a pointer to a runtime.hmap structure."

So when you pass a map to a function/method, you are indeed copying it,
but just the pointer part, not the underlying data structure that contains the data.
A gotcha with maps is that they can be a nil value.
A nil map behaves like an empty map when reading,
but attempts to write to a nil map will cause a runtime panic.

Therefore, you should never initialize an empty map variable:
`var m map[string]string`

Instead, you can initialize an empty map like we were doing above,
 or use the make keyword to create a map for you:

 var dictionary = map[string]string{}

// OR

var dictionary = make(map[string]string)

Both approaches create an empty hash map and point dictionary at it.
Which ensures that you will never get a runtime panic.
*/

func (d Dictionary) Add(word string, def string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = def
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, def string) error {
	_, err := d.Search(word)
	if err == nil {
		d[word] = def
		return nil
	}
	return ErrNotFound
}
