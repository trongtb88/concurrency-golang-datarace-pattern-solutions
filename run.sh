set -e
#listing 1
echo "----------------"
echo "Demo of loop index variable capture (Listing 1)"
echo "----------------"
go run -race loopIndexVariableCapture.go 2>&1 | tee loopIndexVariableCapture.log | grep -s "DATA RACE" && echo "PASS"

#listing 2
echo "----------------"
echo "Demo of error variable capture (Listing 2)"
echo "----------------"
go run -race errVariableCapture.go 2>&1 | tee errVariableCapture.log | grep -s "DATA RACE" && echo "PASS"

#listing 3
echo "----------------"
echo "Demo of named return variable capture (Listing 3)"
echo "----------------"
go run -race namedReturnVariableCapture.go 2>&1 | tee namedReturnVariableCapture.log | grep -s "DATA RACE" && echo "PASS"

#listing 4
echo "----------------"
echo "Demo of named return variable capture with deferred return (Listing 4)"
echo "----------------"
go run -race namedReturnVariableCaptureWithDefer.go 2>&1 | tee namedReturnVariableCaptureWithDefer.log | grep -s "DATA RACE" && echo "PASS"

#listing 5
echo "----------------"
echo "Demo of races in slices (Listing 5)"
echo "----------------"
go run -race racesInSlices.go 2>&1 | tee racesInSlices.log | grep -s "DATA RACE" && echo "PASS"

#listing 6
echo "----------------"
echo "Demo of races in concurrent access too hashmap (Listing 6)"
echo "----------------"
go run -race concurrentHashmapAccess.go 2>&1 | tee concurrentHashmapAccess.log | grep -s "DATA RACE" && echo "PASS"

#listing 7-8
echo "----------------"
echo "Demo of races due to confusion on whether method invocation is by value or pointer (Listing 7-8)"
echo "----------------"
go run -race methodInvocationByValueVsByPointer.go 2>&1 | tee methodInvocationByValueVsByPointer.log | grep -s "DATA RACE" && echo "PASS"

#listing 9
echo "----------------"
echo "Demo of mixing message passing with shared memory (Listing 9)"
echo "----------------"
go run -race mixingMessagePassingWithSharedmem.go 2>&1 | tee mixingMessagePassingWithSharedmem.log | grep -s "DATA RACE" && echo "PASS"



#listing 10
echo "----------------"
echo "Demo of races due to incorrect placement of WaitGroup.Add() (Listing 10)"
echo "----------------"
go run -race incorrectGroupAdd.go 2>&1 | tee incorrectGroupAdd.log | grep -s "DATA RACE" && echo "PASS"


#listing 10+
echo "----------------"
echo "Demo of races due to incorrect placement of WaitGroup.Done() (not a listing but mentioned)"
echo "----------------"
go run -race incorrectGroupDone.go 2>&1 | tee incorrectGroupDone.log | grep -s "DATA RACE" && echo "PASS"


#listing 11
echo "----------------"
echo "Demo of races due to mutating shared variable in a reader lock held in read-only mode (Listing 11)"
echo "----------------"
go run -race rlock.go 2>&1 | tee rlock.log | grep -s "DATA RACE" && echo "PASS"


#listing NA
echo "----------------"
echo "Demo of races due to running parallel tests in test suite/table (not a listing but mentioned)"
echo "----------------"
cd testTable
go test -race . 2>&1 | tee parallelTest.log | grep -s "DATA RACE" && echo "PASS"
cd ..




#mixingMessagePassingWithSharedmem.go
