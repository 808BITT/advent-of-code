from operator import index
from re import split
import time

from click import group

def main():
    timer = time.perf_counter()
    print("Part 2:", part2("sample.txt"))
    print("Time: {0:.3f}".format(time.perf_counter() - timer))


def part2(file):
    with open(file, 'r') as f:
        lines = f.readlines()

    for i in range(len(lines)):
        lines[i] = lines[i].strip()

    sum = 0
    for line in lines:
        sum += calculate(line)
    
    return sum

def calculate(line: str) -> int:
    print(line)
    split = line.split(' ')
    arrangement = split[0]
    values = [int(v) for v in split[1].split(',')]
    

    index = 0
    results = []
    for group in values:
        count = 0
        while True:
            if count == group:
                results.append(arrangement[index-group+1:index+1])
                break
            if arrangement[index] == '?' or arrangement[index] == '#':
                count += 1
            index += 1
        
            

        for r in results:
            print(r)


    return len(results)


if __name__ == '__main__':
    main()