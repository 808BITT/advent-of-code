# two1nine -> 29
# eightwothree -> 83
# abcone2threexyz -> 13
# xtwone3four -> 24
# 4nineeightseven2 -> 42
# zoneight234 -> 14
# 7pqrstsixteen -> 76

NUMS = ["one", "two", "three", "four", "five", "six", "seven", "eight", "nine"]

def main():
    # open input file
    with open('input.txt', 'r') as file:
        lines = file.readlines()
    
    # test input
    # lines = [
    #     "two1nine",
    #     "eightwothree",
    #     "abcone2threexyz",
    #     "xtwone3four",
    #     "4nineeightseven2",
    #     "zoneight234",
    #     "7pqrstsixteen",
    # ]

    # initialize variables
    sum: int = 0
    

    # iterate through lines
    for line in lines:
        nums: list[str] = []
        word_buffer: str = ""
        # iterate through characters
        for char in line.strip():
            # if char is digit, add to nums
            if char.isdigit():
                nums.append(char)
                word_buffer = ""
            # if char is alpha, add to word_buffer
            elif char.isalpha():
                word_buffer += char
                # if word_buffer is in NUMS, add to nums and reset word_buffer to current char
                for n in NUMS:
                    if n in word_buffer:
                        nums.append(NUMS.index(n) + 1)
                        word_buffer = char
                        break
                        
        left = int(nums[0])
        if len(nums) == 1:
            right = int(nums[0])
        else:
            right = int(nums[-1])

        number = left*10 + right
        sum += number

        # print(f"{line.strip()} -> {number}")
    
    print(sum)
        



if __name__ == "__main__":
    main()