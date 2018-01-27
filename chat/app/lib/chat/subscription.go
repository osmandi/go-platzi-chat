package chat

type Subscription struct {
	Archive []Event      // All events from the archive
	New     <-chan Event // New events coming  in
}

func (s Subscription) Cancel() {
	unsubscribe <- s.New
	drain(s.New)
}

func Subscribe() Subscription {
	resp := make(chan Subscription)
	subscribe <- resp
	return <-resp
}
