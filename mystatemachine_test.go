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

	initstate := stateMachine.Init("happy")
	state := stateMachine.NewState("pain")

	mystatemachine.AddEdge(stateMachine, initstate, state, []bool{testNoEvent})
	mystatemachine.AddEdge(stateMachine, state, initstate, []bool{testYesEvent})
	
	mystatemachine.AddEdge(stateMachine, initstate, initstate, []bool{testYesEvent})
	mystatemachine.AddEdge(stateMachine, state, state, []bool{testNoEvent})

	testcases := map[string] struct {
		Events []bool
		want  mystatemachine.State
	}{
		"compute simple state machine" : {[]bool{testNoEvent, testNoEvent}, state},
		"compute simple state machine2" : {[]bool{testNoEvent, testYesEvent}, initstate},
	}

	for name , tc := range testcases {
		t.Run(name, func(t *testing.T) {
			finalState := mystatemachine.Compute(stateMachine, tc.Events, false)
			if finalState != tc.want {
				t.Error("reached another state")
			}
		})
	}


}

