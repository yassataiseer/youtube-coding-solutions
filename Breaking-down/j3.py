def solve(data,location):
    answer = []
    prev_data = data[0:location]
    next_data = data[location:len(data)]
    for i in range(len(prev_data)):
        answer.append(sum(prev_data[i:len(prev_data)]))
    for i in range(len(next_data)):
        answer.append(sum(next_data[0:i+1]))
    return answer
def shifts(locations):
    answer = []
    for i in range(len(locations)):
        #[0,2,10,15,20]
        if i == 0:
            answer.append(solve(locations,i))
        else:
            last = locations.pop(i)
            locations.insert(i-1,last)
            answer.append(solve(locations,i))
    return answer

locations = input()
locations = locations.split(" ")
locations = list(map(int,locations))
locations.insert(0,0)
answer = shifts(locations)
for things in answer:
    output = ""
    for letter in things:
        output+=str(letter)+" "
    print(output.strip())


