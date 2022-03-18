## my state machine

```

import "mystatemachine"

stateMachine := mystatemachine.New()

	painRule := "skiing" // <= the type T
	happyRule := "sleep"

	initstate := stateMachine.Init("happy", false) // set initialize state, bool is whether the final node or not
	state := stateMachine.NewState("pain", false) // add another state

    // add edge (statemachine, from state, to state, events)

	mystatemachine.AddEdge(stateMachine, initstate, state, []string{painRule})
	mystatemachine.AddEdge(stateMachine, state, initstate, []string{happyRule})

	mystatemachine.AddEdge(stateMachine, initstate, initstate, []string{happyRule})
	mystatemachine.AddEdge(stateMachine, state, state, []string{painRule})

    rules := []string{..., ..., ...} // set event

    mystatemachine.Compute(stateMachine, rules, isPrint) // if isPrint is true, print each node

```
