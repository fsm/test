package test

// queueEmitter is a simple emitter that simply puts the Emitted items into a queue,
// that can be pulled out at a later time.
type queueEmitter struct {
	nodes []interface{}
}

// Dequeue pulls the first element out of the queue and returns it
func (q *queueEmitter) Dequeue() interface{} {
	if len(q.nodes) > 0 {
		result := q.nodes[0]
		q.nodes = q.nodes[1:]
		return result
	}
	return nil
}

// Flush clears out the queue and returns it all
func (q *queueEmitter) Flush() []interface{} {
	flushed := q.nodes
	q.nodes = make([]interface{}, 0)
	return flushed
}

// Emit enqueues to the queue
func (q *queueEmitter) Emit(i interface{}) error {
	q.nodes = append(q.nodes, i)
	return nil
}
