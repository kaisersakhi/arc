# arc

Data structures in Go.

## Linked List

Doubly linked list with generics. Supports custom comparators via `NewLinkedList` or default ordering via `NewComparableLinkedList[T cmp.Ordered]`.

### Benchmarks

Run on Apple M2 (`go test ./linked_list/... -bench=. -benchmem`).

| Benchmark | N | Time/op | Allocs/op | Bytes/op |
|---|---|---|---|---|
| Append | 1000 | 14.5 µs | 1001 | 24016 B |
| Prepend | 1000 | 15.2 µs | 1001 | 24016 B |
| PopStart | 1000 | 17.1 µs | 1002 | 24048 B |
| PopEnd | 1000 | 16.8 µs | 1002 | 24048 B |
| GetAt | 1000 (mid) | 401 ns | 0 | 0 B |
| IsPresent | 1000 (mid) | 971 ns | 0 | 0 B |
| Sort | 100 | 4.7 µs | 200 | 4808 B |
| Sort | 1000 | 54.7 µs | 2000 | 48008 B |
| Sort | 10000 | 668 µs | 20000 | 480009 B |

Sort allocates `2N` nodes per call: `N` for the copy returned, `N-1` dummy nodes inside `merge`.
