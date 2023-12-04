import time
start_timer = time.time()
with open ('input.txt') as f:
    lines = f.readlines()

# lines = [
#     "Card   1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
#     "Card   2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
#     "Card   3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
#     "Card   4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
#     "Card   5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
#     "Card   6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11"
# ]

cards = {}
for l in lines:
    card, nums = l.split(': ')
    card = card.split(' ')[-1]
    cards[card] = 1

# print(cards)

for card in cards:
    # pause = input()

    idx = int(card)-1
    _, nums = lines[idx].split(': ')
    winning, possible = nums.split(' | ')
    possible = possible.split(' ')
    winning = winning.split(' ')
    winning = [i for i in winning if i != '']
    possible = [i for i in possible if i != '']
    # remove newline from last item
    possible[-1] = possible[-1][:-1]
    # print (winning, possible)

    matches = 0
    for i in range(len(winning)):
        for j in range(len(possible)):
            if winning[i] == possible[j]:
                matches += 1
                # print(f'{winning[i]} matches {possible[j]}')
    # print(card, matches)

    if matches > 0:
        # print(f'Card {card}: {matches} matches, adding {cards[card]} to {matches} cards')
        for i in range(matches):
            # print(f'adding {cards[card]} to {str(int(card)+i+1)}')
            cards[str(int(card)+i+1)] += cards[card]
    

    # print(f'After Round {card}: {cards}')

sum = 0
for card in cards:
    sum += cards[card]

print(sum)
print(f"{(time.time() - start_timer) * 1000} ms")