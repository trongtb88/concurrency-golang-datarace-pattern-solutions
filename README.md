# concurrency-golang-datarace-pattern-solutions
Concurrency in golang datarace pattern issues and solutions

# Commons issues at concurrency golang that causes data race
## You will see 2 code block at go file, 1 is ORIGINAL issue, 1 is solution to fix it

### List of commmon issue cause data race :  

1. concurrentHashmapAccess
2. errVariableCapture
3. incorrectGroupAdd
4. incorrectGroupDone
5. loopIndexVariableCapture
6. methodInvocationByValueVsByPointer
7. mixingMessagePassingWithSharedmem
8. namedReturnVariableCapture
9. namedReturnVariableCaptureWithDefer
10. racesInSlices
11. rlock
