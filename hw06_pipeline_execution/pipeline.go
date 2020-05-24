package hw06_pipeline_execution //nolint:golint,stylecheck

type (
	I   = interface{}
	In  = <-chan I
	Out = In
	Bi  = chan I
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	for _, stage := range stages {
		out := stage(in)
		rez := make(Bi)
		go func(out, done In, rez Bi) {
			for i := range out {
				select {
				case <-done:
					close(rez)
					return
				default:
					rez <- i
				}
			}
			close(rez)
		}(out, done, rez)
		in = rez
	}

	return in
}
