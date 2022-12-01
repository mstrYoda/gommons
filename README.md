# gommons

Swiss army knife for Golang developers

## Features

- [X] Async tasks
- [X] Async tasks with results
- [X] Command exec utils
- [X] Zero alloc string-byte conversion
- [X] Time utils
- [ ] Array utils

<details>
<summary>Async</summary>

#### Async tasks

```go
New().Task(
    func () {
        a = 1
        fmt.Println("1")
    }, func () {
        b = 1
        fmt.Println("2")
    }).Await()
```

#### Async tasks with results

```go
results := NewAsyncWorkWithResult[int]().TaskWithResult(
    func() int {
        return 5
    }, func() int {
        return 11
    }).AwaitResult()
```


</details>

<details>
<summary>Command exec</summary>

</details>

<details>
<summary>Zero alloc</summary>

</details>

<details>
<summary>Time utils</summary>

</details>
