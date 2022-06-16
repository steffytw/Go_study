package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	/*
		Go provide a ‘math/rand’ package which has inbuilt support for generating pseudo-random numbers.
		This package defines methods which can be used to generate

		- A pseudo-random number within the range from 0 to n
		- A pseudo-random number without range specified. 
		  The range will depend upon the type of int i.e int64, int32, uint64, etc
		- Pseudo-Random Number Generator Functions with range.

		All functions take n as an argument and will panic if n<=0.


		Intn(n int) – It returns a non-negative pseudo-random number in [0,n)

		Int31n(n int32) – It returns a non-negative pseudo-random number in [0,n) but returns a int32

		Int63n(n int64) – It returns a non-negative pseudo-random number in [0,n) but returns a int64

		Pseudo Random Number Generator Functions without range.


		Int() – returns a non-negative pseudo-random int

		Int31() – returns a non-negative pseudo-random 31-bit integer as an int32

		Int63() – returns a non-negative pseudo-random 63-bit integer as an int64

		Uint32() – returns a pseudo-random 32-bit value as a uint32

		Uint64() – returns a pseudo-random 64-bit value as a uint64

		Pseudo Random Number Generator Functions with range for floats

		
		Float64() – returns, as a float64, a pseudo-random number in [0.0,1.0)

		Float32() – returns, as a float32, a pseudo-random number in [0.0,1.0)
	
	*/
	fmt.Println("Random Number:",rand.Intn(10)) // gives same output

	rand.Seed(time.Now().Unix()) 
	fmt.Println("Random Number:",rand.Intn(10)) // gives different output each time
	fmt.Printf("Random %T Number: %v  \n",rand.Int31n(10),rand.Int31n(10))
	fmt.Printf("Random %T Number: %v  \n",rand.Int63n(10),rand.Int63n(10))
	fmt.Printf("Random %T Number: %v  \n",rand.Uint32(),rand.Uint32())
	fmt.Printf("Random %T Number: %v  \n",rand.Uint64(),rand.Uint64())
	fmt.Printf("Random %T Number: %v  \n",rand.Float64(),rand.Float64())
	fmt.Printf("Random %T Number: %v  \n",rand.Float32(),rand.Float32())
	
}
