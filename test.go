package test

import (
	"github.com/fsm/fsm"
)

const Platform = "fsm.test"

// New creates and returns a new *TestingTraverser, which can be used
// to build automated tests for State Machines
func New(stateMachine fsm.StateMachine, store fsm.Store) *TestingTraverser {
	return &TestingTraverser{
		uuid:     uuid(),
		stateMap: fsm.GetStateMap(stateMachine),
		emitter: &queueEmitter{
			nodes: make([]interface{}, 0),
		},
		store: store,
	}
}
