NUMS = ["one", "two", "three", "four", "five", "six", "seven", "eight", "nine"]

def main():
    with open('input.txt', 'r') as file:
        lines = file.readlines()

    # lines = [
    #     "two1nine",
    #     "eightwothree",
    #     "abcone2threexyz",
    #     "xtwone3four",
    #     "4nineeightseven2",
    #     "zoneight234",
    #     "7pqrstsixteen",
    # ]

    sum: int = 0
    
    for line in lines:
        nums: list[str] = []
        word_buffer: str = ""
        for char in line.strip():
            if char.isdigit():
                nums.append(char)
                word_buffer = ""
            elif char.isalpha():
                word_buffer += char
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