### [Readers Writers](https://en.wikipedia.org/wiki/Readers%E2%80%93writers_problem) 

> Readers and Writers problems deal with situations in which many concurrent threads of execution try to access the same shared resource at one time.

>> Some threads may read and some may write, with the constraint that no thread may access the shared resource for either reading or writing while another 
>> thread is in the act of writing to it. 

>> In particular, we want to prevent more than one thread modifying the shared resource simultaneously and allow for two or m>> ore readers to access the 
>> shared resource at the same time

#### Implemented classically with reader preference and writer starvation

With readers (3) and writers (2)
```ocaml
Reader (R) and Writer (W) interplay:
~~~ R1 Read purple -> 0
~~~ R0 Read caramel -> 0
~~~ R2 Read obsidian -> 0
~~~ W0 Incr caramel's value by 1 => 1
~~~ W1 Incr purple's value by 1 => 1
~~~ R0 Read caramel -> 1
~~~ R1 Read purple -> 1
~~~ W0 Incr caramel's value by 1 => 2
~~~ R2 Read obsidian -> 0
~~~ W1 Incr purple's value by 1 => 2
~~~ R0 Read caramel -> 2
~~~ W0 Incr caramel's value by 1 => 3
~~~ R1 Read purple -> 2
~~~ R2 Read obsidian -> 0
~~~ W1 Incr purple's value by 1 => 3
~~~ R0 Read caramel -> 3
~~~ R2 Read obsidian -> 0
~~~ R1 Read purple -> 3
```

Now with more readers (5) and writers (2) -- Notice how the writers appear to wait on the readers
```ocaml
Reader (R) and Writer (W) interplay:
~~~ R2 Read obsidian -> 0
~~~ R0 Read caramel -> 0
~~~ R1 Read purple -> 0
~~~ R4 Read atlantis -> 0
~~~ R3 Read graphite -> 0
~~~ W0 Incr caramel's value by 1 => 1 (Writer 0 waits for all readers to finish)
~~~ W1 Incr purple's value by 1 => 1 (Writer 1 appears to have queued behind W0)
~~~ R2 Read obsidian -> 0
~~~ R0 Read caramel -> 1
~~~ R1 Read purple -> 1
~~~ R3 Read graphite -> 0
~~~ R4 Read atlantis -> 0
~~~ R2 Read obsidian -> 0
~~~ W0 Incr caramel's value by 1 => 2
~~~ W1 Incr purple's value by 1 => 2
~~~ R1 Read purple -> 2
~~~ R0 Read caramel -> 2
~~~ R4 Read atlantis -> 0
~~~ R3 Read graphite -> 0
~~~ W0 Incr caramel's value by 1 => 3
~~~ R2 Read obsidian -> 0
~~~ W1 Incr purple's value by 1 => 3
~~~ R1 Read purple -> 3
~~~ R4 Read atlantis -> 0
~~~ R0 Read caramel -> 3
~~~ R3 Read graphite -> 0
```
