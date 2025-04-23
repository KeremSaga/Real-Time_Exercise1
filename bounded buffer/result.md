3) Sharing a variable

The result is not guareenteed to be correct because the operation on i is not safeguarded, meaning i can be accessed by both functions

4) Sharing a variable, but properly

Difference between mutexes and semaphores:
- Mutex: Stands for MUTual EXecution. Only allows one thread to access the part that is locked. That one thread locks and unlocks the section it will operate on (mutex endorces ownership).
- Semaphores: Allows N threads to access the critical part(s). There is a permit/counting system, counting down from N to know when no other threads can access the operation whithout (at least) one releasing the permit. Therefore no ownership endorced.