n = int(input())
numbers = list(map(int, input().split()))

for x in numbers:
    if x < 10:
        continue
    
    last_digit = x % 10
    temp = x
    good = True
    
    while temp > 0:
        if temp % 10 != last_digit:
            good = False
            break
        temp = temp // 10
    
    if good:
        print(x, end=" ")
print()


