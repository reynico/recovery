package electrum

// Pool provides a shared pool of Clients that callers can acquire and release, limiting
// the amount of concurrent Clients in active use.
type Pool struct {
	nextClient chan *Client
}

// NewPool creates an initialized Pool with a `size` number of clients.
func NewPool(size int) *Pool {
	nextClient := make(chan *Client, size)

	for i := 0; i < size; i++ {
		nextClient <- NewClient()
	}

	return &Pool{nextClient}
}

// Acquire obtains an unused Client, blocking until one is released.
func (p *Pool) Acquire() <-chan *Client {
	return p.nextClient
}

// Release returns a Client to the pool, unblocking the next caller trying to `Acquire()`.
func (p *Pool) Release(client *Client) {
	p.nextClient <- client
}
