from itertools import count
import time
from typing import List, Dict, Tuple

class Direction:
    def __init__(self, name: str, dx: int, dy: int):
        self.name = name
        self.dx = dx
        self.dy = dy

class Pipe:
    def __init__(self, connections: List[Direction]):
        self.connections = connections

class Position:
    def __init__(self, x: int, y: int):
        self.x = x
        self.y = y

    def __eq__(self, other):
        return self.x == other.x and self.y == other.y

def part2(filename: str) -> int:
    north = Direction("North", 0, -1)
    south = Direction("South", 0, 1)
    east = Direction("East", 1, 0)
    west = Direction("West", -1, 0)

    lookup = {
        "|": Pipe([north, south]),
        "-": Pipe([east, west]),
        "L": Pipe([north, east]),
        "J": Pipe([north, west]),
        "7": Pipe([south, west]),
        "F": Pipe([south, east]),
        "S": Pipe([south, east, north, west])
    }

    lines = read_file(filename)
    pipe_map, start_x, start_y = generate_map(lines)

    current_x = start_x
    current_y = start_y


    #try to open visited.txt, if it doesn't exist, create it
    visited = []
    try:
        with open("visited.txt", "r") as file:
            visited = [Position(int(line.split()[0]), int(line.split()[1])) for line in file.readlines()]
    except FileNotFoundError:
        back_to_start = False
        visited = [Position(current_x, current_y)]
        steps = 0
        while not back_to_start:
            current = pipe_map[current_y][current_x]
            for direction in lookup[current].connections:
                new = pipe_map[current_y+direction.dy][current_x+direction.dx]

                pipe_found = False
                for p in lookup:
                    if new == p and Position(current_x + direction.dx, current_y + direction.dy) not in visited:
                        visited.append(Position(current_x + direction.dx, current_y + direction.dy))
                        pipe_found = True
                        current_x += direction.dx
                        current_y += direction.dy
                        break
                    elif new == "S" and steps > 1:
                        pipe_found = True
                        back_to_start = True
                        break
                if pipe_found:
                    steps += 1
                    break
                
        with open("visited.txt", "w") as file:
            for v in visited:
                file.write(f"{v.x} {v.y}\n")

    world_width = len(pipe_map[0])
    world_height = len(pipe_map)

    # print(f"World width: {world_width}, World height: {world_height}")

    # grid = [
    #     "...........",
    #     ".S-------7.",
    #     ".|F-----7|.",
    #     ".||.....||.",
    #     ".||.....||.",
    #     ".|L-7.F-J|.",
    #     ".|..|.|..|.",
    #     ".L--J.L--J.",
    #     "..........."
    # ]

    visited = [(v.x, v.y) for v in visited]
    print("Visited:", visited)
    count, enclosed = count_enclosed_areas(pipe_map, visited)
    print("Enclosed areas:", count)

    # draw the visited path using turtle
    import turtle
    # set zoom level
    turtle.setup(1200, 1200)
    turtle.setworldcoordinates(-1, -1, world_width, world_height)  # Set the origin to bottom-left
    turtle.speed(10)
    turtle.penup()
    
    # draw the grid of dots
    for y, row in enumerate(pipe_map):
        for x, char in enumerate(row):
            turtle.goto(x, world_height - y - 1)  # Invert y-coordinate
            turtle.dot(3, "black")

    # paint the pipe
    turtle.goto(start_x, world_height - start_y - 1)  # Invert start y-coordinate
    turtle.dot(5, "green")
    turtle.pendown()
    turtle.color("green")
    for v in visited:
        turtle.goto(v[0], world_height - v[1] - 1)  # Invert y-coordinate for each visited position
    turtle.goto(start_x, world_height - start_y - 1)  # Invert start y-coordinate
    turtle.penup()

    # put a dot at each enclosed area
    for y, row in enumerate(pipe_map):
        for x, char in enumerate(row):
            if (x, y) not in visited and (x, y) != (start_x, start_y) and (x, y) not in enclosed:
                turtle.goto(x, world_height - y - 1)  # Invert y-coordinate
                turtle.dot(5, "red")

    # put a dot at each enclosed area
    for y, row in enumerate(pipe_map):
        for x, char in enumerate(row):
            if (x, y) not in visited and (x, y) != (start_x, start_y) and (x, y) in enclosed:
                turtle.goto(x, world_height - y - 1)  # Invert y-coordinate
                turtle.dot(5, "blue")
    turtle.done()

    return count

def flood_fill(grid, x, y, visited, loop_parts):
    if x < 0 or x >= len(grid[0]) or y < 0 or y >= len(grid) or (x, y) in visited or (x, y) in loop_parts or grid[y][x] == "S":
        return
    if (x, y) in loop_parts:
        # Check for squeezable gaps
        if not is_squeezable_gap(grid, x, y, loop_parts):
            return
    visited.add((x, y))

    directions = [(-1, 0), (1, 0), (0, -1), (0, 1)]  # Up, Down, Left, Right
    for dx, dy in directions:
        flood_fill(grid, x + dx, y + dy, visited, loop_parts)

def is_squeezable_gap(grid, x, y, loop_parts):
    directions = [(-1, 0), (1, 0), (0, -1), (0, 1)]  # Up, Down, Left, Right
    for dx, dy in directions:
        nx, ny = x + dx, y + dy
        if 0 <= nx < len(grid[0]) and 0 <= ny < len(grid) and (nx, ny) not in loop_parts:
            return True
    return False

def count_enclosed_areas(grid, loop_parts):
    visited = set()
    enclosed = []
    for y in range(len(grid)):
        for x in range(len(grid[0])):
            # Start flood fill from the borders
            if x == 0 or y == 0 or x == len(grid[0]) - 1 or y == len(grid) - 1:
                flood_fill(grid, x, y, visited, loop_parts)

    # Count non-loop cells not visited by the flood fill
    count = 0
    for y, row in enumerate(grid):
        for x, cell in enumerate(row):
            if (x, y) not in visited and (x, y) not in loop_parts:
                count += 1
                enclosed.append((x, y))

    return count, enclosed

def read_file(filename: str) -> List[str]:
    with open(filename, 'r') as file:
        lines = [line.strip() for line in file.readlines()]
    return lines

def generate_map(lines: List[str]) -> Tuple[List[List[str]], int, int]:
    grid = []
    x, y = 0, 0
    for l, line in enumerate(lines):
        row = []
        for c, char in enumerate(line):
            row.append(char)
            if char == "S":
                x = c
                y = l
        grid.append(row)
    return grid, x, y


if __name__ == "__main__":
    timer = time.time()
    print("Part 2 Solution:", part2("sample.txt"))
    print("Time taken:", time.time() - timer)