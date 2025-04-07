# Go Memory Model Notes

## Happens-Before
- **Rule 1**: Package `sync` establishes happens-before via `Unlock` > `Lock`.
- **Rule 2**: Channel send happens before corresponding receive.

## Common Pitfalls
```go
var x int
go func() { x = 1 }()  // No guarantee main() sees this write
fmt.Println(x)         // May print 0 or 1