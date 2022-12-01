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

#### Run posix command and get output as byte array

```go
out := Exec("echo", "test")
```

#### Run posix command with pipes

```go
strReader := strings.NewReader("hello world")

outWriter := bytes.NewBuffer(nil)
errWriter := bytes.NewBuffer(nil)

ExecPipe(strReader, outWriter, errWriter, "echo", "test")
outputStr := outWriter.String()
```


</details>

<details>
<summary>Zero alloc string byte conversion</summary>

</details>

<details>
<summary>Time utils</summary>

</details>
