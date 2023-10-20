### [Readers Writers](https://en.wikipedia.org/wiki/Readers%E2%80%93writers_problem) 

> Readers and Writers problems deal with situations in which many concurrent threads of execution try to access the same shared resource at one time.

> Some threads may read and some may write, with the constraint that no thread may access the shared resource for either reading or writing while another 
> thread is in the act of writing to it. 

> In particular, we want to prevent more than one thread modifying the shared resource simultaneously and allow for two or more readers to access the 
> shared resource at the same time

![Picture](https://github.com/brpandey/readers_writers/blob/main/shared.png?raw=true)

```elixir
./readers_writers --help
Usage of ./readers_writers:
  -nr int
    	specify num readers, max is 7 (default 4)
  -nw int
    	specify num writers, max is 7 (default 1)
  -starve
    	indicates whether to classically starve writers
```

#### Implemented with classic writer starvation and also equal preference to both reader and writer
#### Tweak the options to see if there's a noticeable difference

```ocaml

$ ./readers_writers --starve=true  --nr=6 --nw=2
Specified options: -starve: true  -nr: 6  -nw: 2
Spawn order: [ R R W R R R W R ]

Reader (R) and Writer (W) interplay:
~~~ R1 Read purple -> 0
~~~ R2 Read obsidian -> 0
~~~ R5 Read hexagon -> 0
~~~ W1 Incr purple's value by 1 => 1
~~~ W0 Incr caramel's value by 1 => 1
~~~ R0 Read caramel -> 1
~~~ R4 Read atlantis -> 0
~~~ R1 Read purple -> 1
~~~ R3 Read graphite -> 0
~~~ R2 Read obsidian -> 0
~~~ R5 Read hexagon -> 0
~~~ W1 Incr purple's value by 1 => 2
~~~ R1 Read purple -> 2
~~~ W0 Incr caramel's value by 1 => 2
~~~ R0 Read caramel -> 2
~~~ R4 Read atlantis -> 0
~~~ R2 Read obsidian -> 0
~~~ R3 Read graphite -> 0
~~~ R5 Read hexagon -> 0
~~~ W1 Incr purple's value by 1 => 3
~~~ R1 Read purple -> 3
~~~ R2 Read obsidian -> 0
~~~ W0 Incr caramel's value by 1 => 3
~~~ R0 Read caramel -> 3
~~~ R4 Read atlantis -> 0
~~~ R5 Read hexagon -> 0
~~~ R3 Read graphite -> 0
~~~ R0 Read caramel -> 3
~~~ R4 Read atlantis -> 0
~~~ R3 Read graphite -> 0
```

