Exercise 1 - Theory questions
-----------------------------

### Concepts

What is the difference between *concurrency* and *parallelism*?
> *Both concurrency and parallelism is about attempting to do several operations simultaniously. Concurrancy is done on one CPU, where the tasks are swapped very fast. So they take turns (not simultanious technically). Parallelism is when more than one CPU is used (this one is simultanious in every sense).*

What is the difference between a *race condition* and a *data race*? 
> *Data race is a type of race condition, but there are race conditions which aren't data races. So data races are a subcatogory of race conditions. Race conditions are when the (sub)programs correct/desired output depends on the timing of inputs. Data race is when a shared memory block is accessed by two or more threads where one thread alters the data before the other(s) can read or write to it.* 
 
*Very* roughly - what does a *scheduler* do, and how does it do it?
> *A schedular decides which thread to run next.* 


### Engineering

Why would we use multiple threads? What kinds of problems do threads solve?
> *Your answer here*

Some languages support "fibers" (sometimes called "green threads") or "coroutines"? What are they, and why would we rather use them over threads?
> *Your answer here*

Does creating concurrent programs make the programmer's life easier? Harder? Maybe both?
> *Your answer here*

What do you think is best - *shared variables* or *message passing*?
> *Both have their pros and cons. In an optimal scenario shared variables is better/best since if the concurrancy issues are handled the operation gets done faster. But this bit, concurrancy struggles, is the con of shared variables, thus making message passing better if the system doesn't require to be "real-time"/fast. Fast is a relative term, but so be it..*


