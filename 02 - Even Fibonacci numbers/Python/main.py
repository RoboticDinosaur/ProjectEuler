maxValue = 4000000

currentValue = 1
previousValue = 2

array = []

number = 0
result = 0

while previousValue < maxValue:
    print('Previous Value:', previousValue)
    temp = currentValue + previousValue
    currentValue = previousValue
    previousValue = temp

    if currentValue % 2 == 0:
        array.append(currentValue)

for number in array:
    result += number

print(result)