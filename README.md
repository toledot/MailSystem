# MailSystem
MailSystem:


The general input order is expected to be:
number of cities
number of operations

*   The system provides the option to skip on arguments, for example:
    Number of Cities:
    2       // number of cities
    A       // city
    1       // number of branches
    0 1 3   // branche
    skip    // will skip on the next city
    1       // number of operations

*   To exit, type exit

*   The City Name can be any string, but it must be unique
*   Cities, Operations, and Branches must be integers greater than or equal to zero
*   Weights are float32 variables

In Case of invaild input, Invaild Input will be printed to the screen and the system wait for vaild input

Operations:

1. Getting the city name and printing her content
2. Transferring packages from src to dst. RunTime O(number of Packages in source branch)
3. The city with the most packages is printed. RunTime O(Number of Cities)

* Unit test is MailSystem_test.go file
