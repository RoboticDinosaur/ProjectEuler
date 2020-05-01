# I do not understand this.  I don't know why its so different from the
# Go version.  I will need to investigate when I'm not tired.

maxValue = 600851475143

def primeFactor(n):
    i = 2
    factors = []
    while i * i <= n:
        if n % i:
            i += 1
        else:
            n //= i
            factors.append(i)

    if n > 1:
        factors.append(n)
    return factors

print(primeFactor(maxValue))