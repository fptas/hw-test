package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	monitor := func(in In, done In, stage Stage) Out {
		myOut := make(Bi)
		stageOut := stage(myOut)
		go func() { 
/* горутина слушает канал предыдущего стейджа и передает следующему
если done закрывается, то закрывает свой выходной канал,
что приведет впоследствии к прекращению работы последующего стейджа*/
			defer close(myOut)
			for {
				select {
				case <-done:
					return
				case i, ok := <-in:
					if ok {
						myOut <- i
					} else {
						return
					}
				}
			}
		}()
		return stageOut
	}

	in2 := in
	for i := range stages {
		in2 = monitor(in2, done, stages[i])
	}

	return in2
}
