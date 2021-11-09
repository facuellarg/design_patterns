# Factory
we implement the factory design pattern and for extra learning was added a function to add more types to our factory, in this way the client can add his owns type as long as its type implements the `Deliverer` interface.
```GO
type Deliverer interface {
	Deliver() error
}
```