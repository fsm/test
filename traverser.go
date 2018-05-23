package test

import (
	"github.com/fsm/fsm"
)

// TestingTraverser is a fsm.Traverser, with an additional three methods:
// Send(), GetReceived(), and GetAllReceived()
type TestingTraverser struct {
	uuid     string
	platform string
	stateMap fsm.StateMap
	emitter  *queueEmitter
	store    fsm.Store
}

// ===== UUID =====

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

// ===== Platform =====

// Platform returns the Traversers Platform
func (t *TestingTraverser) Platform() string {
	return t.platform
}

// SetPlatform sets the Traversers Platform
func (t *TestingTraverser) SetPlatform(platform string) {
	t.platform = platform
}

// ===== State =====

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

// ===== Other Data =====

// Upsert updates (or creates) a variable for this traverser.
// Be sure to call this only after you have called at least one `Send()` to this Traverser.
func (t *TestingTraverser) Upsert(key string, value interface{}) {
	// TODO
	traverser, err := t.store.FetchTraverser(t.uuid)
	if err != nil {
		panic("Failed to access *TestingTraverser's underlying traverser.  It is likely you have forgotten to call Send() prior to calling Upsert().")
	}
	traverser.Upsert(key, value)
}

// Fetch returns a variable set to this traverser.
// Be sure to call this only after you have called at least one `Send()` to this Traverser.
func (t *TestingTraverser) Fetch(key string) interface{} {
	traverser, err := t.store.FetchTraverser(t.uuid)
	if err != nil {
		panic("Failed to access *TestingTraverser's underlying traverser.  It is likely you have forgotten to call Send() prior to calling Fetch().")
	}
	return traverser.Fetch(key)
}

// Delete deletes a variable set to this traverser.
// Be sure to call this only after you have called at least one `Send()` to this Traverser.
func (t *TestingTraverser) Delete(key string) {
	traverser, err := t.store.FetchTraverser(t.uuid)
	if err != nil {
		panic("Failed to access *TestingTraverser's underlying traverser.  It is likely you have forgotten to call Send() prior to calling Delete().")
	}
	traverser.Delete(key)
}

// ===== Other Functions =====

// Send emulates a user sending a message to a chat-bot.
func (t *TestingTraverser) Send(input string) {
	fsm.Step(Platform, t.uuid, input, fsm.TextInputTransformer, t.store, t.emitter, t.stateMap)
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
