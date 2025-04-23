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
> *If you want to accomplish multiple tasks at once, you need to use threads. Threads mimic how a human mind thinks better. You essantially divide the task into independent task that at certain point interact or trigger something. Like a timer. Timer is something that is independent from the main task, but when it goes off, the main task needs to be distrupted.*

Some languages support "fibers" (sometimes called "green threads") or "coroutines"? What are they, and why would we rather use them over threads?
> *Fibers are lightweight threads managed by the language runtime, not the OS. They are faster to create and use less memory, in other words they give less overhead than threads. Therefore they should be used whenever "possible". They aren't possible to use, meaning they won't work when the code has blocking functions like sleep() or if you want to run code on multiple cores.*

Does creating concurrent programs make the programmer's life easier? Harder? Maybe both?
> *Both. Easier in the long run and overall easier, eitherwise there would be no points in using consurrancy. The hardness comes in the form of added complexity and potential issues like race conditions, deadlocks and less straight forward debugging.*

What do you think is best - *shared variables* or *message passing*?
> *Both have their pros and cons. In an optimal scenario shared variables is better/best since if the concurrancy issues are handled the operation gets done faster. But this bit, concurrancy struggles, is the con of shared variables, thus making message passing better if the system doesn't require to be hard real-time fast, meaning the time constraints for the system response allows for massage passing.*


