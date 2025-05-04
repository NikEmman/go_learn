package main

import (
	"fmt"
	"sync"
	"time"
)

// Go Routines and Mutexes

// this is a wait group. Go's async execution. We add a wait counter in the loop below
var wg = sync.WaitGroup{}
var dbData = []string{"id1","id2","id3","id4","id5"}
var results = []string{}
//since cocurrent execution may write at the same memory and corrupt the data, we might not get all of the correctly.
//to avoid that we use Mutex ( mutual exclusion)
//var m = sync.Mutex{}
// to allow other Go Routines to access the results values while a Lock is in place, we use the Read Write Mutex
var m = sync.RWMutex{}


func main5(){
	t0:=time.Now()
	for i:=0; i<len(dbData);i++{
		wg.Add(1)
		// we use the go keyword infront of the function we want to run concurrently
		go dbCall(i)

	}
	// we wait till all the tasks in the wait group are completed and then execute the rest of the code
	wg.Wait()
	
	fmt.Printf("\nTotal execution time: %v",time.Since(t0))
	fmt.Printf("\n the results are %v", results)
}
func dbCall(i int){
	//simulate DB call delay
	var delay float32 = 2000
	time.Sleep(time.Duration(delay)*time.Millisecond)
	save(dbData[i])
	log()

	//we call Done() to remove a counter  from the wait group
	wg.Done()

}

func save (result string){
	// to safeguard the memory/data corruption we lock the memory before we write in the results Slice
	// execution will wait on Lock to see of a lock is present by another Go Routine ( async execution) and continue when it can and set the lock itself
	//one drawback of Lock is that we prevent other Go Routines rom accessing results variable for the lock duration
	// Lock will wait until all other Locks or Readlocks are cleared, to start writing on results
	m.Lock()
	results = append(results, result)
	m.Unlock()
}
func log(){
	// read locks are softer versions of Lock, will allow other Go Routines to read the data at the same time but will prevent writing until they are done reading
	m.RLock()
	fmt.Printf("\nThe current results are :%v" ,results)
	m.RUnlock()

}

// without the async the db calls execute one after another, thus needing more total time (ie 8 sec) but with the
//async method they take 1.8 sec, or w/e the slowest call is 
// using RWMutex will allow multiple go routines from reading a variable but will stop when a full lock is in place, where data is potentially added to the variable