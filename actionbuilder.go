package nettrigger

// ActionBuilder builds actions from specifications.
type ActionBuilder func(ActionSpec) (Action, error)
