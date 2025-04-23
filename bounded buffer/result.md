3) Sharing a variable

The result is not guareenteed to be correct because the operation on i is not safeguarded, meaning i can be accessed by both functions.

Problem with sleep: sleep blocks everything, time is being wasted doing nothing.

The operation on i is not atomic, incrementing (or decrementing) happens in three stages, read i, operate on i, replace old i with new i. So whilst both increment/decrement loops do run at 10000 and 10001 times respectively the values that they change is incorrect, leading to unexpected results.

The solution is to atomisize the critical part (block access to i when critical part is happenning). This is the case in C. In Go this is possible, but the motto is "Don't communicate by sharing memory; share memory by communicating". Therefore channels are used.

4) Sharing a variable, but properly

Difference between mutexes and semaphores:
- Mutex: Stands for MUTual EXecution. Only allows one thread to access the part that is locked. That one thread locks and unlocks the section it will operate on (mutex endorces ownership).
- Semaphores: Allows N threads to access the critical part(s). There is a permit/counting system, counting down from N to know when no other threads can access the operation whithout (at least) one releasing the permit. Therefore no ownership endorced.

