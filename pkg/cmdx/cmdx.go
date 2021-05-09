package cmdx

type Task struct {
	Type    string
	Payload Payload
}

func NewTask(t string, payload map[string]interface{}) *Task {
	return &Task{
		Type:    t,
		Payload: Payload{payload},
	}
}
