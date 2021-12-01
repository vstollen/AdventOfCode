# Open file: AOC-2021/day1.txt
with open('AOC-2021/day1.txt') as f:
    data = f.read().splitlines()
    data = list(map(int, data))

    increases = 0
    for i in range(3, len(data)):
        last_sum = sum(data[i-3:i])
        current_sum = sum(data[i-2:i+1])

        if current_sum > last_sum:
            increases += 1
    
    print(increases)