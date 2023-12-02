def main():
    # open input file
    with open('input.txt', 'r') as file:
        lines = file.readlines()
    
    # test input
    # lines = [
    #     "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
    #     "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
    #     "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
    #     "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
    #     "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
    # ]

    sum: int = 0
    for line in lines:
        min_red = 0
        min_blue = 0
        min_green = 0

        game_name = line.split(":")[0].strip()
        game_id = game_name.split(" ")[1]
        games = line.split(":")[1].strip()
        games = games.split(";")
        games = [g.strip() for g in games]
        # print(game_id, games)

        for game in games:
            cubes = game.split(",")
            cubes = [c.strip() for c in cubes]
            # print(game_id, cubes)

            for cube in cubes:
                cube = cube.split(" ")
                cube_color = cube[1]
                cube_count = int(cube[0])
                # print(game_id, cube_color, cube_count)

                if cube_color == "red" and cube_count > min_red:
                    min_red = cube_count
                elif cube_color == "blue" and cube_count > min_blue:
                    min_blue = cube_count
                elif cube_color == "green" and cube_count > min_green:
                    min_green = cube_count

        power = min_red * min_blue * min_green
        print(game_id, "power:", power)

        sum += power
        
    print("Sum:", sum)

if __name__ == "__main__":
    main()