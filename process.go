package process

//Context process context
type Context interface{}

//Step function step
type Step func(c Context) (interface{}, error)

//Process represents the process struct
type Process struct {
	steps      []Step
	err        error
	actualStep int
	ctx        Context
	result     interface{}
}

//New  new process
func New() *Process {
	return &Process{}
}

//AddStep add step to process
func (p *Process) AddStep(s Step) *Process {
	p.steps = append(p.steps, s)
	return p
}

//Start a new process execution
func (p *Process) Start(ctx Context) *Process {
	p.actualStep = 0
	p.err = nil
	p.result = nil
	p.ctx = ctx
	return p
}

//Next Is there a new step
func (p *Process) Next() bool {
	return p.actualStep < len(p.steps)
}

//RunStep run step
func (p *Process) RunStep() *Process {
	if p.Next() && p.Error() == nil {
		p.result, p.err = p.steps[p.actualStep](p.ctx)
	}
	p.actualStep++
	return p
}

//RunAll run all step
func (p *Process) RunAll() *Process {
	p.Start(p.ctx)
	for p.Next() {
		p.RunStep()
	}
	return p
}

//Result return value
func (p *Process) Result() interface{} {
	return p.result
}

//Error run all step
func (p *Process) Error() error {
	return p.err
}
