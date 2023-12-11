with open("input.txt") as f:
    lines = f.readlines()

races = {}
times = lines[0].split(':')[1].strip().split(' ')
distances = lines[1].split(':')[1].strip().split(' ')

times = [int(x) for x in times if x != '']
distances = [int(x) for x in distances if x != '']

# print(times)
# print(distances)

ways_to_win = []
for i in range(len(times)):
    dis = distances[i]
    time = times[i]
    
    wins = 0
    for button_len in range(time+1):
        race_remaining = time - button_len
        travelled = race_remaining * button_len
        win = travelled > dis
        if travelled > dis:
            wins += 1
        # print(button_len, travelled)
        
    ways_to_win.append(wins)

total_ways_to_win = 1
for way in ways_to_win:
    total_ways_to_win *= way

print(total_ways_to_win)


