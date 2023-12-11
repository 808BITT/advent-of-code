
RED_CUBES = 12
BLUE_CUBES = 14
GREEN_CUBES = 13


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
        valid_game = True
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

                if cube_color == "red" and cube_count > RED_CUBES:
                    valid_game = False
                    break
                elif cube_color == "blue" and cube_count > BLUE_CUBES:
                    valid_game = False
                    break
                elif cube_color == "green" and cube_count > GREEN_CUBES:
                    valid_game = False
                    break
            
            if not valid_game:
                break
        
        if valid_game:
            print(game_id, "is valid")
            sum += int(game_id)
        
    print("Sum:", sum)

if __name__ == "__main__":
    main()