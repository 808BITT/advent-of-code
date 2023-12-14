import time

CYCLES = 1000000000

def main():
    # timer = time.perf_counter()
    # print("Part 1:", part1("2023/day14/sample.txt"))
    # print("Time: {0:.3f}".format(time.perf_counter() - timer))
    timer = time.perf_counter()
    print("Part 2:", part2("2023/day14/input.txt"))
    print("Time: {0:.3f}".format(time.perf_counter() - timer))

def part2(file):
    memo = {}
    rock_map = parse_file(file)
    for i, line in enumerate(rock_map):
        rock_map[i] = list(line)

    start = 0
    end = 0
    
    # print("Initial Map - Height:", len(rock_map), "Width:", len(rock_map[0]))
    # show(rock_map)

    for i in range(CYCLES):
        print("Cycle:", i, "of", CYCLES, "Current Weight:", sum_weight(rock_map))
        memo[i] = to_rock_string(rock_map)

        show(rock_map)
        input()

        # test = from_rock_string(memo[i])
        # for i in range(len(test)-1):
        #     print(''.join(test[i]) + " " + ''.join(rock_map[i]))

        # NORTH TILT
        index = 1
        while True:
            if index == len(rock_map)-1:
                break

            rolling = False
            for j in range(len(rock_map)-1): # walk each row
                rock_map, delta = tilt_north(j+1, rock_map)
                # show(rock_map)
                # input()    
                if delta:
                    rolling = True
                    
            if not rolling:
                break

            index += 1

        print("NORTH TILT")
        show(rock_map)
        input()

        # WEST TILT
        index = 1
        while True:
            if index == len(rock_map[0]):
                break

            rolling = False
            for j in range(len(rock_map[0])-1): # walk each column
                rock_map, delta = tilt_west(j, rock_map)
                # show(rock_map)
                # input()    
                if delta:
                    rolling = True
                    
            if not rolling:
                break

            index += 1

        # print("WEST TILT DONE")
        show(rock_map)
        input()

        # SOUTH TILT
        index = 1
        while True:
            if index == len(rock_map)-1:
                break

            rolling = False
            for j in range(len(rock_map)-1): # walk each row
                rock_map, delta = tilt_south(j, rock_map)
                # show(rock_map)
                # input()    
                if delta:
                    rolling = True
                    
            if not rolling:
                break

            index += 1

        # print("SOUTH TILT DONE")
        show(rock_map)
        input()

        # EAST TILT
        index = 1
        while True:
            if index == len(rock_map)-1:
                break

            rolling = False
            for j in range(len(rock_map[0])-1): # walk each column
                rock_map, delta = tilt_east(j, rock_map)
                # show(rock_map)
                # input()    
                if delta:
                    rolling = True
                    
            if not rolling:
                break

            index += 1

        # print("EAST TILT DONE")
        show(rock_map)
        input()

        if to_rock_string(rock_map) in memo.values():
            # print("FOUND A LOOP AT CYCLE:", i+1, "WEIGHT:", sum_weight(rock_map))
            start = 0
            for j in range(i+1):
                if memo[j] == to_rock_string(rock_map):
                    start = j
                    break
            end = i + 1
            # print("Start:", start, "End:", end)
            remaining = CYCLES - end
            cycle_length = end - start
            # print("Cycle Length:", cycle_length, "Remaining:", remaining)
            offset = remaining % cycle_length
            # print("Offset:", offset)

            # print("Final Map:", start + offset)
            # print(memo[start + offset])
            cache = from_rock_string(memo[start + offset])
            cache = [list(line) for line in cache if line != ""]
                
            return sum_weight(cache)

        show(rock_map)
        input()

    return sum_weight(rock_map)

def to_rock_string(rock_map):
    string = ""
    for i, line in enumerate(rock_map):
        for char in line:
            string += char
        if i != len(rock_map)-1:
            string += "\n"
    return string

def from_rock_string(string):
    rock_map = []
    for line in string.split("\n"):
        rock_map.append(list(line))
    return rock_map

def part1(file):
    rock_map = parse_file(file)
    for i, line in enumerate(rock_map):
        rock_map[i] = list(line)
    
    print("Initial Map - Height:", len(rock_map), "Width:", len(rock_map[0]))
    show(rock_map)

    index = 1
    while True:
        if index == len(rock_map)-1:
            break

        rolling = False
        for i in range(len(rock_map)-1):
            rock_map, delta = tilt_north(i+1, rock_map)
            # show(rock_map)
            # input()    
            if delta:
                rolling = True


        if not rolling:
            break

        index += 1

    show(rock_map)

    return sum_weight(rock_map)

def sum_weight(rock_map):
    sum = 0
    for i, line in enumerate(rock_map):
        rocks = count_rocks(line)
        sum += rocks * (len(rock_map)-i)

    return sum

def count_rocks(line):
    count = 0
    for char in line:
        if char == "O":
            count += 1
    return count

def tilt_north(index, rock_map):
    delta = False
    for i in range(len(rock_map[0])): # walk each column
        if rock_map[index][i] == "O" and rock_map[index-1][i] == ".":
            rock_map[index][i] = "."
            rock_map[index-1][i] = "O"
            delta = True

    return rock_map, delta

def tilt_west(index, rock_map):
    delta = False
    # current_col = [line[index] for line in rock_map]
    # next_col = [line[index+1] for line in rock_map]
    # print("Index:", index)
    # print("Current Col:", current_col)
    # print("   Next Col:", next_col)
    # show(rock_map)
    # input()
    for i in range(len(rock_map)): # walk each row
        if rock_map[i][index+1] == "O" and rock_map[i][index] == ".":
            rock_map[i][index+1] = "."
            rock_map[i][index] = "O"
            delta = True

    return rock_map, delta

def tilt_south(index, rock_map):
    delta = False
    for i in range(len(rock_map[0])): # walk each column
        if rock_map[index][i] == "O" and rock_map[index+1][i] == ".":
            rock_map[index][i] = "."
            rock_map[index+1][i] = "O"
            delta = True

    return rock_map, delta

def tilt_east(index, rock_map): 
    delta = False
    for i in range(len(rock_map)): # walk each row
        if rock_map[i][index] == "O" and rock_map[i][index+1] == ".":
            rock_map[i][index] = "."
            rock_map[i][index+1] = "O"
            delta = True

    return rock_map, delta

def show(rock_map):
    for line in rock_map:
        print(''.join(line))
    print()

def parse_file(file):
    with open(file) as f:
        return [line.strip() for line in f.readlines()]
    
if __name__ == "__main__":
    main()