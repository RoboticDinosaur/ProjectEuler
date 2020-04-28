inputValue = 1000

multiples = []
total = 0

number = 0

while number < inputValue:
    if number % 3 == 0 or number % 5 == 0:
        multiples.append(number)
    number += 1

for number in multiples:
    total += number

print('The answer is: ', total)