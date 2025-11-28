package metrics


type TodoMetrics struct {

}

func(t *TodoMetrics)Inc(ket string,value int){

}

func NewTodoMetrics() *TodoMetrics {
	return &TodoMetrics{}
}