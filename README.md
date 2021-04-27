# lock

A small Go library to handle scoped locks in a single line.

sync.RWMutex examples:

```defer lock.Read(&rwmutex).Unlock()```

```defer lock.Write(&rwmutex).Unlock()```

sync.Mutex examples:

```defer lock.Mutex(&mutex).Unlock()```
