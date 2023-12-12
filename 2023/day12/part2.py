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
    
    return 0

if __name__ == '__main__':
    main()