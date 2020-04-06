# 06 - Sum square difference

## The Problem
The sum of the squares of the first ten natural numbers is,
`1^2 + 2^2 + ... + 10^2 = 385`

The square of the sum of the first ten natural numbers is,
`(1+2+...+10)^2 = 55^ = 3025`

Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is, 
`3025 − 385 = 2640`

Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.

### Natural Numbers
In mathematics, the natural numbers are those used for counting (as in "there are six coins on the table") and ordering (as in "this is the third largest city in the country"). In common mathematical terminology, words colloquially used for counting are "cardinal numbers" and words connected to ordering represent "ordinal numbers". The natural numbers can, at times, appear as a convenient set of codes (labels or "names"); that is, as what linguists call nominal numbers, forgoing many or all of the properties of being a number in a mathematical sense. 

### The Solution
This problem needs to be solved with a couple of functions. 

The first will be the SumofSqaure() which will count from keep a running total while counting from 1 - 101

The second is a very simple function to calculate the SquareofSum() which will add all of the numbers together and then square.

The final function will return the result to fmt.Println()