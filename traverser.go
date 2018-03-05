package test

import (
	"github.com/fsm/fsm"
	targetutil "github.com/fsm/target-util"
)

// TestingTraverser is a fsm.Traverser, with an additional three methods:
// Send(), GetReceived(), and GetAllReceived()
type TestingTraverser struct {
	uuid     string
	stateMap fsm.StateMap
	emitter  *queueEmitter
	store    fsm.Store
}

// UUID returns the Traversers UUID
func (t *TestingTraverser) UUID() string {
	return t.uuid
}

// SetUUID sets a new UUID to the traverser.  It's not advised to set this yourself.
// Be sure to call this only after you have called at least one `Send()` to this Traverser.
func (t *TestingTraverser) SetUUID(uuid string) {
	traverser, err := t.store.FetchTraverser(t.uuid)
	if err != nil {
		panic("Failed to access *TestingTraverser's underlying traverser.  It is likely you have forgotten to call Send() prior to calling SetUUID().")
	}
	t.uuid = uuid
	traverser.SetUUID(uuid)
}

// CurrentState returns the current active state the Traverser is in.
// Be sure to call this only after you have called at least one `Send()` to this Traverser.
func (t *TestingTraverser) CurrentState() string {
	traverser, err := t.store.FetchTraverser(t.uuid)
	if err != nil {
		panic("Failed to access *TestingTraverser's underlying traverser.  It is likely you have forgotten to call Send() prior to calling CurrentState().")
	}
	return traverser.CurrentState()
}

// SetCurrentState sets the current state.  This can be useful for jumping around
// specific parts of your bot flow.
// Be sure to call this only after you have called at least one `Send()` to this Traverser.
func (t *TestingTraverser) SetCurrentState(state string) {
	traverser, err := t.store.FetchTraverser(t.uuid)
	if err != nil {
		panic("Failed to access *TestingTraverser's underlying traverser.  It is likely you have forgotten to call Send() prior to calling SetCurrentState().")
	}
	traverser.SetCurrentState(state)
}

// Upsert updates (or creates) a variable for this traverser.
// Be sure to call this only after you have called at least one `Send()` to this Traverser.
func (t *TestingTraverser) Upsert(key string, value interface{}) error {
	// TODO
	traverser, err := t.store.FetchTraverser(t.uuid)
	if err != nil {
		panic("Failed to access *TestingTraverser's underlying traverser.  It is likely you have forgotten to call Send() prior to calling Upsert().")
	}
	return traverser.Upsert(key, value)
}

// Fetch returns a variable set to this traverser.
// Be sure to call this only after you have called at least one `Send()` to this Traverser.
func (t *TestingTraverser) Fetch(key string) (interface{}, error) {
	traverser, err := t.store.FetchTraverser(t.uuid)
	if err != nil {
		panic("Failed to access *TestingTraverser's underlying traverser.  It is likely you have forgotten to call Send() prior to calling Fetch().")
	}
	return traverser.Fetch(key)
}

// Delete deletes a variable set to this traverser.
// Be sure to call this only after you have called at least one `Send()` to this Traverser.
func (t *TestingTraverser) Delete(key string) error {
	traverser, err := t.store.FetchTraverser(t.uuid)
	if err != nil {
		panic("Failed to access *TestingTraverser's underlying traverser.  It is likely you have forgotten to call Send() prior to calling Delete().")
	}
	return traverser.Delete(key)
}

// Send emulates a user sending a message to a chat-bot.
func (t *TestingTraverser) Send(input string) {
	targetutil.Step(t.uuid, input, t.store, t.emitter, t.stateMap)
}

// GetReceived returns the first message that was sent to the traverser that hasn't been received yet.
// Think of this as a queue and GetReceived() is simply dequeue().
func (t *TestingTraverser) GetReceived() interface{} {
	return t.emitter.Dequeue()
}

// GetAllReceived returns an array of all of the messages sent to this traverser.
// Think of this as a queue, and GetAllReceived() is simply flush().
func (t *TestingTraverser) GetAllReceived() []interface{} {
	return t.emitter.Flush()
}
