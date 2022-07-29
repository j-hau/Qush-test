"""
    Write a function which is passed an integer, n, and returns a list of n lists such that:
    f(0) returns []
    f(1) returns [[1]]
    f(2) returns [[1], [1,2]]
    f(3) returns [[1], [1,2], [1,2,3]]
    etc.
"""


from curses.ascii import isdigit


def createIntegerList(number):
    # We first want to induce a base case where f(0) returns [].  In this case we initialise the empty array of arrays to be returned empty if n=0.
    superSet = []

    if(number == 0):
        return superSet
    else:
        """
            We know one of the constraints of the problem is that f(1) returns [[1]], f(2) returns [[1], [1,2]].
            It can therefore be assumed that f(n) produces an array of arrays where there are n arrays contained inside of an array.
            The contents of the internal arrays are the numbers ascending to that index.
            e.g. the 4th array, will look like [1,2,3,4].  the nth array will look like [1,2,3,4 .... n] 

            We therefore want for the input number amount of times.  As f(1) is [[1]], this tells us we need to account for python indexing at 0 and start from 1.
            Furthermore, as python goes up to but not including the upper bound of a for loop, we need to loop until "number + 1".
        """
        for x in range(1, number+1):
            # Initialising the sub array to append to the superSet to be empty.
            subArray = []

            # Looping for the index of the sub array.  E.g. looking at the 3rd sub array in f(3) which is [1,2,3], we want to loop 3 times and add each increment to the sub array
            # This allows us to append it to the super set with ascending values up to and including that index.
            for i in range(x):
                # As python indexes from 0, we need to add 1 to each value we append.
                subArray.append(i+1)
            # Add this value to the super set.
            superSet.append(subArray)
    return superSet


# Here is where we input the actual value.
# It also does some simple validation checks to check if it is actually an integer that is input.
numberToLoop = input("Enter the number to create a list for: ")
if numberToLoop.isdigit():
    print(createIntegerList(int(numberToLoop)))
else:
    print("The input is not an integer.")