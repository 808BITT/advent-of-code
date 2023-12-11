import time
start_timer = time.time()
with open ("input.txt", "r") as f:
    data = f.readlines()

sum = 0
for line in data:
    left, right = None, None
    for i in range(len(line)):
        if line[i].isdigit():
            if left is None:
                left = line[i]
                right = line[i]
            else:
                right = line[i]

    sum += int(left)*10 + int(right)
print(sum)
print(f"{(time.time() - start_timer) * 1000} ms")