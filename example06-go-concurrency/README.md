# Go concurrency

## Goroutine

A goroutine is a lightweight thread of execution.

## Channels

Channels are a golang primitive allow concurrency synchronization and communications between goroutines.

A few references :

- https://go101.org/article/channel.html
- https://docs.google.com/document/d/1yIAYmbvL3JxOKOjuCyon7JhW4cSv1wy5hC0ApeGMV9s/pub

## Question 01

To call a function synchronously:

```
print()
```

To call a function in a goroutine

```
go print()
```

Here the code will terminate before the goroutine is executed. In order to wait for the execution of the goroutine, a sleep instruction can be added.

```
import time
go print()
time.Sleep(time.Second)
```

Otherwise channels can be used.

> Channels are the pipes that connect concurrent goroutines. You can send values into channels from one goroutine and receive those values into another goroutine.

## Question 03

`go run -race` allows to detect race conditions.
