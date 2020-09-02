package main

import "fmt"

// Typical recursive fibonacci
func FibonacciRecursion(n int) int {
    if n <= 1 {
        return n
    }
    return FibonacciRecursion(n-1) + FibonacciRecursion(n-2)
}

// Accumulate Fibonacci numbers using a 'tail'
func FibonacciTailRecursion(number int, p int, c int, tail []int) [] int{

    // Special initial condition
    if p == 0 && c == 1 {
        tail = append(tail, 0)

            tail = append(tail, 1)
            number--
    }

    if number == 0 {
        return tail
    }

    tail = append(tail, p+c)
    return FibonacciTailRecursion(number - 1, c, p+c, tail)
}

// Implementation of fibonacci in a normal loop with accumulator
func FibonacciLoop(n int) [] int {
    accum := make([]int,0)
    if n < 2 {
        accum = accum[0:2]
    }

    accum = append(accum, 0, 1)

    for i := 2; i <= n; i++ {
        accum = append(accum, accum[i-1] + accum[i-2])
    }
    return accum
}

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
    // These variables live in a special scope
    // Aka lexical scoping
    n1, n := -1, 1

    return func() int {
        n1, n = n, n1+n
        return n
    }
}

func main() {
    fibLimit := 15

    fmt.Println(FibonacciRecursion(fibLimit))
    fmt.Println(FibonacciTailRecursion(fibLimit, 0, 1, make([]int,0)))
    fmt.Println(FibonacciLoop(fibLimit))

    accum := make([]int, 0)
    f := fibonacci()

    for i := 0; i <= fibLimit; i++ {
        accum = append(accum, f())
    }

    fmt.Println(accum)

}

