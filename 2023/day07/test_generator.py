import random


with open('test.txt', 'w') as file:
    for i in range(10):
        file.write(f"{''.join(random.choices('23456789TJQKA', k=5))} {random.randint(1, 100)}\n")