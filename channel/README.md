# Channels
Channels are a typed conduit through which you can send and receive values with the channel operator, `<-`.

`
ch <- v    // Send v to channel ch.
`

`
v := <-ch  // Receive from ch, and assign value to v.
`

Like maps and slices, channels must be created before use:

`
ch := make(chan int)
`

By default, sends and receives block until the other side is ready. This allows goroutines to synchronize without explicit locks or condition variables.

# Buffered Channels
Channels can be buffered. Provide the buffer length as the second argument to make to initialize a buffered channel:

`
ch := make(chan int, 100)
`

Sends to a buffered channel block only when the buffer is full. Receives block when the buffer is empty.

Modify the example to overfill the buffer and see what happens.

# Range and Close
A sender can close a channel to indicate that no more values will be sent. Receivers can test whether a channel has been closed by assigning a second parameter to the receive expression: after

`
v, ok := <-ch
`

ok is false if there are no more values to receive and the channel is closed.

The loop 

```
    c := make(chan int, 10)
    for i := range c {
        fmt.Println(i)
    }
```

receives values from the channel repeatedly until it is closed.

**Note**: Only the sender should close a channel, never the receiver. Sending on a closed channel will cause a panic.

**Another note**: Channels aren't like files; you don't usually need to close them. Closing is only necessary when the receiver must be told there are no more values coming, such as to terminate a range loop.

# Select
The `select` statement lets a goroutine wait on multiple communication operations.

A select blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready.

# Default Selection
The default case in a select is run if no other case is ready.

Use a default case to try a send or receive without blocking:
```
select {
case i := <-c:
    // use i
default:
    // receiving from c would block
}

```
