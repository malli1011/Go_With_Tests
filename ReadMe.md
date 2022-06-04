**sync.Map** 

similar to concurrent hashmap in java. Can be used with multiple goRoutines. 

`type ProductDB struct { products sync.Map }`

It will be initialized with zero value and can be used without instantiation.

* Safe for concurrent use by multiuple goroutines

* Equivalent to a safe map[interface{}]interface{}
* The zero value is empty and ready for use
Incurs performance overhead and should only be used as nessary. 

*It has below methods:*

`func (m *Map) Load(key interface{})(value interface{}, ok bool)`
`func (m *Map) Store(key,value interface{})`
`func (m *Map) Range(f func(key,value interface{}))bool`

**sync.Mutex**

* The Mutex is initialised unlocked using `var m sync.Mutex`

* The zero value is empty and ready for use

`func (m *Mutex) Lock()`
`func (m *Mutex) Unlock()`


**Channel Syntax**

* The reserverdx keyword `chan` denotes a channel 
* Channel operator is the arrow operator `<-` 
* Channels are associated with data type
* The zero value of channels are `nil` `var ch chan T` is nil
* The syntax to declare a channel of type T is `ch := make(chan T)`
* Given the channel denoted by the variable ch, sending is done with `ch <- data`; the arrow points into the channel as the data travels into it. 
* Receiving is done with `data := <-ch`; the arrow points away from the channel as data travels out of it.

*Send and Receives are Blocking*

Code execution will stop until the send or receive is successfully completed. The definition depends on the setup of the channel. 

*Buffered Channels*

Predefined capacity channels which have the ability to store values for later processing. 
Sender can put the value into channel, if the channel has the space then it saves the value. 
Later a receiver can receive the value for processing, if the value is retrived successfully then 
the channel will remove the value. This mainly used for async communication. 

Bidirectional channel : `chan T`
Send only channel : `chan<- T`
Receive only channel: `<-chan T`

bidirectional can be casted to unidirectional, but the other way is not possible. 

*Clossing channel*

* Closing a channel signals no more values will be sent on it.
* The syntax is simple; `close(ch)`
* we can close bidrectional and send only channel, not receive only channels.
* Receivers immediately receive the zero value of the channel data type from a closed channel
* Senders panic when sending to a closed channel. 

checking for closed channels:

```
func doWork(ch chan string){
    data,ok:= <-ch
    if !ok{
        fmt.Println("channel is closed")
        return 
    }
    fmt.Println("channel is open:", data)
}
```
Range over a channel and print the values. It will complete when the channel closes. 
this is used when reading multiple values from a channel. 

```
for greeting := range ch{
    fmt.Println("Greeing received!",greeting)
}
```


**Select**

* Select us used to wait on multiple channel operations.
* It blocks until one of the operations on its channels is ready
* If multiple operations are ready, one of them is chosen at random

```
func doWork(ch1,ch2 chan string){
    select{
        case msg1:= <-ch1:
            fmt.Println("received from ch1: ", msg1)
        case msg2:= <-ch2:
            fmt.Println("received from ch2: ",msg2)
        default:
            fmt.Println("Nothing received")
    }
}
```
Branch: 02_06b ( for channels usecase)

sync.Once // calls a given code only once, even if its wrapped inside a for loop. 

**Worker pools**
Similar to java thread pool or executorservice. 