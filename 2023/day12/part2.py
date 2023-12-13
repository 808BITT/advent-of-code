from re import split
import time

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
    split = line.split(' ')
    arrangement = split[0]
    values = [int(v) for v in split[1].split(',')]
    
    results = []

    # groups end with .
    # groups start with a #
    # groups are the length and in order of the values
    # groups are separated by a 

    buffer = arrangement[0]
    print(buffer)

    return len(results)


def countBroken(arrangement: str) -> int:
    count = 0
    for i in range(len(arrangement)):
        if arrangement[i] == '#' or arrangement[i] == '?':
            count += 1
    
    return 0

if __name__ == '__main__':
    main()