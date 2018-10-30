package nettrigger

import (
	"context"
	"fmt"
)

// Action performs some action based on the environment and providers.
type Action func(context.Context, Environment, Providers) error

// BuildAction uses builders to construct an action from the specification.
func BuildAction(spec ActionSpec, builders ...ActionBuilder) (Action, error) {
	for _, builder := range builders {
		action, err := builder(spec)
		if err != nil {
			return action, err
		}
		if action != nil {
			return action, nil
		}
	}
	return nil, fmt.Errorf("unknown action type \"%s\"", spec.Type)
}

// BuildActions converts the given specifications into an action list.
func BuildActions(specs []ActionSpec, builders ...ActionBuilder) ([]Action, error) {
	var actions []Action
	for i, a := range specs {
		action, err := BuildAction(a, builders...)
		if err != nil {
			return nil, fmt.Errorf("invalid action %d: %v", i, err)
		}
		actions = append(actions, action)
	}
	return actions, nil
}
