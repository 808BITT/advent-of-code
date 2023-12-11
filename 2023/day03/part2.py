with open('input.txt') as f:
    lines = f.readlines()

# lines = [
#     "467..114..",
#     "...*......",
#     "..35..633.",
#     "......#...",
#     "617*......",
#     ".....+.58.",
#     "..592.....",
#     "......755.",
#     "...$.*....",
#     ".664.598..",
# ]

possible_gears = []
for y, line in enumerate(lines):
    for x, symbol in enumerate(line):
        if symbol == '*':
            possible_gears.append((x, y))

print(possible_gears)

# find numbers around each symbol


# find numbers around each symbol
sum = 0
for (x, y) in possible_gears:
    touch_coords = {}
    for dx in (-1, 0, 1):
        for dy in (-1, 0, 1):
            if dx == 0 and dy == 0:
                continue
            if lines[y + dy][x + dx] == '*':
                continue
            if lines[y + dy][x + dx].isdigit():
                touch_coords[(x + dx, y + dy)] = lines[y + dy][x + dx]
    to_delete = []
    for (x, y), number in touch_coords.items():
        if (x + 1, y) in touch_coords:
            to_delete.append((x, y))
            continue
    for item in to_delete:
        del touch_coords[item]

    nums = []
    for (x, y), number in touch_coords.items():
        new_num = number
        back = 1
        forward = 1
        start_found = False
        end_found = False
        while True:
            if not start_found and lines[y][x - back].isdigit():
                new_num = lines[y][x - back] + new_num
                back += 1
            else:
                start_found = True
            if not end_found and lines[y][x + forward].isdigit():
                new_num += lines[y][x + forward]
                forward += 1
            else:
                end_found = True
            if start_found and end_found:
                break
        
        nums.append(int(new_num))

    print(nums)

    if len(nums) == 2:
        sum += nums[0] * nums[1]

print(sum)

