with open('input.txt') as f:
    lines = f.readlines()

symbols = []
for line in lines:
    for char in line:
        if not char.isdigit() and char not in ('.', '\n') and char not in symbols:
            symbols.append(char)

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


symbol_coords = {}
for y, line in enumerate(lines):
    for x, symbol in enumerate(line):
        if symbol in symbols:
            symbol_coords[(x, y)] = symbol



# print(symbol_coords)


touch_coords = {}

# find numbers around each symbol
for (x, y), symbol in symbol_coords.items():
    for dx in (-1, 0, 1):
        for dy in (-1, 0, 1):
            if dx == 0 and dy == 0:
                continue
            if (x + dx, y + dy) in symbol_coords:
                continue
            if lines[y + dy][x + dx] in symbols:
                continue
            if lines[y + dy][x + dx].isdigit():
                touch_coords[(x + dx, y + dy)] = lines[y + dy][x + dx]

to_delete = []
for (x, y), number in touch_coords.items():
    # remove items if the y are the same and the x is one apart
    if (x + 1, y) in touch_coords:
        to_delete.append((x, y))
        continue

for item in to_delete:
    del touch_coords[item]

print(touch_coords)

sum = 0
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
    
    touch_coords[(x, y)] = new_num
    sum += int(new_num)


print(touch_coords)
print(sum)


