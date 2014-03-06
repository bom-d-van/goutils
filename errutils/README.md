package errutils is designed to offer a good way to profile errors with a
complete running stacks of the program where the error breaks out.

Suppose you have there functions are there compose in a way like this:

```
func func1() (err error) {
	err = DoSomethingAndGoWrong()
	return
}

func func2() (err error) {
	err = func1()
	return
}

func func3() (err error) {
	err = func2()
	return
}
```
