package mystatemachine_test

import (
	"testing"
	"mystatemachine"
)

func TestMyStateMachine(t *testing.T) {
	stateMachine := mystatemachine.New()

	// painEvent := "skiing"
	// happyEvent := "sleep"

	testNoEvent := false
	testYesEvent := true

	initstate := stateMachine.Init("happy", false)
	state := stateMachine.NewState("pain", false)
	finalstate := stateMachine.NewState("dead", true)

	mystatemachine.AddEdge(stateMachine, initstate, state, []bool{testNoEvent})
	mystatemachine.AddEdge(stateMachine, state, initstate, []bool{testYesEvent})
	mystatemachine.AddEdge(stateMachine, state, finalstate, []bool{testNoEvent})
	
	mystatemachine.AddEdge(stateMachine, initstate, initstate, []bool{testYesEvent})
	mystatemachine.AddEdge(stateMachine, finalstate, state, []bool{testNoEvent})
	mystatemachine.AddEdge(stateMachine, finalstate, finalstate, []bool{testYesEvent})
	


	testcases := map[string] struct {
		Events []bool
		want  mystatemachine.State
	}{
		"compute simple state machine1" : {[]bool{testNoEvent, testYesEvent}, initstate},
		"compute simple state machine2" : {[]bool{testNoEvent, testNoEvent, testYesEvent}, finalstate},
		"compute simple state machine3 when reached final state" : {[]bool{testNoEvent, testNoEvent, testNoEvent}, finalstate},
	}

	for name , tc := range testcases {
		t.Run(name, func(t *testing.T) {
			finalState := mystatemachine.Compute(stateMachine, tc.Events, false)
			if finalState != tc.want {
				t.Errorf("reached another state %v",finalState.String())
			}
			stateMachine.CurrentState = initstate // init
		})
	}


}
