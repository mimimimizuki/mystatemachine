## my state machine

```

import "mystatemachine"

stateMachine := mystatemachine.New()

	painEvent := "skiing"
	happyEvent := "sleep"

	initstate := stateMachine.Init("happy") // set initialize state
	state := stateMachine.NewState("pain") // add another state

    // add edge (statemachine, from state, to state, events)

	mystatemachine.AddEdge(stateMachine, initstate, state, []string{painEvent})
	mystatemachine.AddEdge(stateMachine, state, initstate, []string{happyEvent})

	mystatemachine.AddEdge(stateMachine, initstate, initstate, []string{happyEvent})
	mystatemachine.AddEdge(stateMachine, state, state, []string{painEvent})

    events := []string{..., ..., ...} // set event

    mystatemachine.Compute(stateMachine, events, isPrint) // if isPrint is true, print each node

```
