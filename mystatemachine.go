package mystatemachine

import (
	"fmt"
	"strconv"
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/multi"
	"golang.org/x/exp/constraints"
)

type State struct {
	Id  int64
	Val interface{}
}

type comparable interface {
	constraints.Ordered | bool
}

type Edge[T comparable] struct {
	Id        int64
	Entry, Exit  graph.Node
	Rules     []T
}

func (s State) ID() int64{
	return s.Id
}

func (e Edge[T]) From() graph.Node {
	return e.Entry
}

func (e Edge[T]) To() graph.Node {
	return e.Exit
}

func (e Edge[T]) ID() int64 {
	return e.Id
}

func (e Edge[T]) ReversedLine() graph.Line {
	return Edge[T]{Entry: e.Exit, Exit: e.Entry}
}

var NodeCounter = 0
var EdgeCounter = 1

type StateMachine struct {
	CurrentState State
	g            *multi.DirectedGraph
}

func New() *StateMachine {
	s := &StateMachine{}
	s.g = multi.NewDirectedGraph()

	return s
}

func (s *StateMachine) Init(initStateValue interface{}) State {
	s.CurrentState = State{Id: int64(NodeCounter), Val: initStateValue}
	s.g.AddNode(s.CurrentState)
	NodeCounter++
	return s.CurrentState
}

func (s *StateMachine) NewState(stateValue interface{}) State {
	state := State{Id: int64(NodeCounter), Val: stateValue}
	
	s.g.AddNode(state)
	NodeCounter++
	return state
}

func AddEdge[T comparable](s *StateMachine, s1, s2 State, rule []T) {
	s.g.SetLine(Edge[T]{Entry: s1, Exit: s2, Id: int64(EdgeCounter), Rules: rule})
	EdgeCounter++
}

func (n State) String() string{
	switch n.Val.(type) {
	case int:
		return strconv.Itoa(n.Val.(int))
	case float32:
		return fmt.Sprintf("%f", n.Val.(float32))
	case float64:
		return fmt.Sprintf("%f", n.Val.(float64))
	case bool:
		return strconv.FormatBool(n.Val.(bool))
	case string:
		return n.Val.(string)
	default:
		return ""
	}
}

func FireEvent[T comparable](s *StateMachine, e T) error {
	currentNode := s.CurrentState

	it := s.g.From(currentNode.Id)
	for it.Next() {
		n := s.g.Node(it.Node().ID()).(State)
		edge := graph.LinesOf(s.g.Lines(currentNode.Id, n.Id))[0].(Edge[T]) // There can be one defined path between two distinct states
		for _, val := range edge.Rules {
			if val == e {
				s.CurrentState = n
				return nil
			}
		}
	}
	return nil
}


func Compute[T comparable](s *StateMachine, events []T, printState bool) State {
	for _, e := range events {
		FireEvent(s, e)
		if printState {
			fmt.Printf("%s\n", s.CurrentState.String())
		}
	}
	return s.CurrentState
}
