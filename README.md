Here is an example implementation of a bloom filter in Golang that includes an add and remove function, and is atomic aware:

To create a BloomFilter instance, call the NewBloomFilter function and pass in the desired size of the filter (the number of bits), as well as an array of hash functions to be used. For example:

```
filter := internals.NewBloomFilter(10000, []internals.HashFunc{internals.NewHashFunc()})
```

To add a key to the filter, call the Add method on the filter instance, passing in the key as a string. For example:

```
filter.Add("A")
filter.Add("B")
filter.Add("C")
```

To remove a key from the filter, call the Remove method on like below example:

```
filter.Remove("C")
```

To check if key exists:

```
filter.Check("C")
```
