
from itertools import count
import time

def main():
    print('-'*20+'\n')
    timer = time.perf_counter()
    print("Part 1:", part1("input.txt"))
    print("Time: {0:.3f}".format(time.perf_counter() - timer))
    # timer = time.perf_counter()
    # print("Part 2:", part2("sample.txt"))
    # print("Time: {0:.3f}".format(time.perf_counter() - timer))



def part1(file):
    pictures = parse_file(file)

    count = 0
    sum = 0
    for picture in pictures:
        show(picture)
        vertical, horizontal = find_symmetry(picture)
        if vertical is not None:
            print("VERTICAL @", vertical, '\n')
            print('-'*20+'\n')
            sum += vertical+1
            count += 1
        if horizontal is not None:
            print("HORIZONTAL @", horizontal, '\n')
            print('-'*20+'\n')
            sum += (horizontal+1) * 100
            count += 1
        print(f'Found {count}/{len(pictures)} symmetries - Sum: {sum}\n')
        input()

    print(f'Found {count}/{len(pictures)} symmetries\n')

    return sum

def show(picture):
    for line in picture:
        print(line)
    print()

def find_symmetry(picture):
    # print("Checking for vertical symmetry:")
    possible_verticals = []
    for i in range(len(picture[0])-1):
        column1 = [line[i] for line in picture]
        column2 = [line[i+1] for line in picture]
        # print(column1)
        # print(column2)
        if column1 == column2:
            possible_verticals.append(i)
            
    
    # print("Checking for horizontal symmetry:")
    possible_horizontals = []
    for i in range(len(picture)-1):
        line1 = picture[i]
        line2 = picture[i+1]
        # print(line1)
        # print(line2)
        if line1 == line2:
            possible_horizontals.append(i)
    
    possible_verticals.sort(reverse=True)
    possible_horizontals.sort(reverse=True)

    print("Possible verticals:", possible_verticals)
    print("Possible horizontals:", possible_horizontals)

    for possible in possible_verticals:
        print("\nChecking vertical at", possible)
        if verify_vertical(picture, possible):
            return possible, None

    for possible in possible_horizontals:
        print("\nChecking horizontal at", possible)
        if verify_horizontal(picture, possible):
            return None, possible

    print("No symmetry found")
    return None, None

def verify_vertical(picture, index):
    picture_width = len(picture[0])
    print("Picture width:", picture_width)
    cols_to_verify = min(index+1, picture_width-index-1)
    print("Cols to verify:", cols_to_verify)
    for i in range(cols_to_verify):
        column1 = [line[index-i] for line in picture]
        column2 = [line[index+1+i] for line in picture]
        print(column1)
        print(column2)
        if column1 != column2:
            print('-'*20)
            return False
    return True

def verify_horizontal(picture, index):
    try:
        picture_height = len(picture)
        print("Picture height:", picture_height)
        lines_to_verify = min(index+1, picture_height-index-1)
        print("Lines to verify:", lines_to_verify)
        for i in range(lines_to_verify):
            line1 = picture[index-i]
            line2 = picture[index+1+i]
            print(line1)
            print(line2)
            if line1 != line2:
                print('-'*20)
                return False
        return True
    except Exception as e:
        print(e)
        return False


def parse_file(file):
    with open(file, 'r') as f:
        lines = f.readlines()

    pictures = []
    picture = []
    for line in lines:
        if line == '\n':
            pictures.append(picture)
            picture = []
        else:
            picture.append(line.strip())
    pictures.append(picture)

    return pictures
    
    

if __name__ == '__main__':
    main()
