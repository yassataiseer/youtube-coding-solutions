def solve(data):
    previous_step = ""
    for numbers in data:
        if int(numbers)==99999:
            pass
        else:
            summ = int(numbers[0])+int(numbers[1])
            printed_response = ""
            if summ==0:
                pass
            elif summ%2==0:
                previous_step = "right"
            else:
                previous_step = "left"
            printed_response = previous_step+" "+numbers[2:5]
            print(printed_response)

data = []
while True:
    instructions = input()
    data.append(instructions)
    if instructions=="99999": 
        break
solve(data)        

