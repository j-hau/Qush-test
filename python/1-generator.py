"""
Define a generator which generates the positive integers up to and including a supplied value
which are divisible by another supplied positive integer and use it to calculate the sum of all
positive integers less than 102030 which are divisible by 3
"""

def generator(a, b):
    # Simple out of bounds constraint check
    if ((a >= 102030 or a <= 0) or (b >= 102030 or b <= 0)):
        return "One of the two inputs are out of bounds."

    """
    The first constraint in the problem is to use the first positive integer to generate a list of positive integers up to and incuding itself which are divisible by the second integer
    This produced list is then used to calculate the sum of all positive integers which are less than 102030 and divisible by 3.

    I would like to do this in a single step but due to specification of the problem I will produce a list first and then use this list for the sum.
    """
    divisibleIntegers = []
    # Looping for the amount of times entered in input1
    for x in range(1, a+1):
        # If the positive integer generated is divisible by the second input, add it to the list of divisible integers
        if x%b == 0 :
            divisibleIntegers.append(x)

    # Initialising the sum
    sum = 0    

    # Iterate over the each element in the divisible integer list.
    for divisibleInteger in divisibleIntegers:
        # If the divisible integer is divisible by 3 then add it to the sum
        if divisibleInteger % 3 == 0 :
            sum += divisibleInteger

    # Return the sum
    return sum

# Here is where we input the actual value.
# It also does some simple validation checks to check if it is actually an integer that is input.
input1 = input("Enter the number to generate positive integers up to and including for the overall sum: ")
input2 = input("Enter the number to check if the integers generated are divisible by this: ")
# Simple validation to make sure both inputs are integers
if (input1.isdigit() and input2.isdigit()) :
    print(generator(int(input1), int(input2)))
else:
    print("The input is not an integer.")