# Mergesort adversarial list creation

Inspired by [this example](https://www.baeldung.com/cs/merge-sort-time-complexity),
the program creates a linked list with integer data values that
requires (almost?) every possible test.

```
$ go build $PWD
$ ./worstcase 8
1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7 -> 8 ->
5 -> 1 -> 7 -> 3 -> 6 -> 2 -> 8 -> 4 ->
enter recursiveMergeSort: 5 -> 1 -> 7 -> 3 -> 6 -> 2 -> 8 -> 4 ->
enter recursiveMergeSort: 5 -> 1 -> 7 -> 3 ->
enter recursiveMergeSort: 5 -> 1 ->

merging lists
    left: 5 -> 
    right: 1 ->
merged list: 1 ->
left list : 5 ->
...
merging lists 
    left: 1 -> 3 -> 5 -> 7 ->
    right: 2 -> 4 -> 6 -> 8 ->
merged list: 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7 ->
right list : 8 ->
1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7 -> 8 ->
List of length 8 is sorted
```

Web examples, like the inspiration for this repo,
tend to only show the easy case or one of the easy cases.
The inspiration code uses a list of length 8, 2<sup>3</sup>.
How do you create an adversarial list of length 9?

## Algorithm

It seems easiest to do this recursively, starting with an in-order sorted list.

```
func createWorstCase(list) {
    if len(list) == 1 { return list }

    left, right := alternatingSplit(list)

    if len(left) == 1 && len(right) == 1 {
        return left + right concatenated in descending order
    }

    left = createWorstCase(left)
    right = createWorstCase(right)

    concatenate left + right

    return left
}
```

`func alternatingSplit` distructively splits an input list into
two output lists by alternately appending the head node of the
(remaining) input list to each of two output lists.

There's no single "worst" input list of any given length,
although a lot of them are only trivially different.
One way to create a worst case list different than what the
above function creates would be to concatenation the lists
the other way, concatenate `left` on the end of `right`,
then return `right`.
