# Open file: AOC-2021/day1.txt
with open('day1.txt') as f:
    data = f.read().splitlines()
    data = map(int, data)

    last_n = None
    increases = 0
    for n in data:
        if not last_n:
            last_n = n
            continue

        if n > last_n:
            increases += 1

        last_n = n
    
    print(increases)