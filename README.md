# ParallelCap

ParallelCap is a Go package that provides a structure for managing parallel execution with a specified capacity. It
leverages channels and wait groups to handle concurrent tasks, allowing you to add tasks up to the specified capacity,
wait for them to complete, and signal when they are done.

## Installation

You can install ParallelCap by running:

```bash
go get github.com/thewisebit/parallelcap
```

## Usage

### Creating a ParallelCap

You can create a new ParallelCap with the given capacity:

```go
p := parallelcap.NewParallelCap(5)
```

If the capacity is less than or equal to 0, the capacity will be set to the maximum integer value.

### Adding Tasks

You can add a task to the ParallelCap:

```go
p.Add()
```

Or you can add a task with context:

```go
err := p.AddWithContext(ctx)
```

If the context is done, or the ParallelCap is full, `AddWithContext` will return an error.

### Signaling Completion

You can signal that a task is done by calling:

```go
p.Done()
```

### Waiting for Completion

You can wait for all tasks to complete by calling:

```go
p.Wait()
```

## Example

```go
package main

import "github.com/thewisebit/parallelcap"

func main() {
	p := parallelcap.NewParallelCap(5)

	for i := 0; i < 10; i++ {
		p.Add()
		go func() {
			// Your concurrent task here
			p.Done()
		}()
	}

	p.Wait() // Wait for all tasks to complete
}
```
