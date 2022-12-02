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
<summary>Queue Pub/Sub example</summary>

Queue acts as a non-blocking message queue backing with unbuffered channel.
Publish/Subscribe functions are not blocks code execution.

```go
q := NewQueue[int]()
q.Publish(context.Background(), 1)
q.Publish(context.Background(), 2)

q.Subscribe(context.Background(), func(data int) {
	fmt.Println("data readed ", data)
})

<-make(chan struct{})
```

You can also give timeout to both message publish and subscribe functions:

```go
q := NewQueue[int]()
ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
defer cancel()
q.Publish(context.Background(), 1)
q.Publish(ctx, 2)

q.Subscribe(ctx, func(data int) {
	fmt.Println("data readed ", data)
})

<-make(chan struct{})
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

#### String to byte array zero allocation

```go
str := String([]byte("test"))
```
#### Byte to string

```go
byteArr := Byte("test")
```

</details>

<details>
<summary>Time utils</summary>

#### Function execution elapsed time utility

```go
elapsedTime := ElapsedTime(func() {
	time.Sleep(100 * time.Millisecond)
})
```
	

</details>
