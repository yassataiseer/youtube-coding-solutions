import time
from functools import cache
import sys
start_time = time.time()
@cache
def fibonacci(number):
    if number<=2:
        return 1
    else:
        return fibonacci(number-1)+fibonacci(number-2)
print(fibonacci(50))
print("--- %s seconds ---" % (time.time() - start_time))