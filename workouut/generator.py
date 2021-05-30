import random

def calculate(*args):
    workouts = []
    counter=0
    while True:
        for workout in args:
            string = ""
            for options in workout:
                random_index =  random.randint(0, len(workout[options])-1)
                string+=workout[options][random_index] +", "
            workouts.append(string)
        counter+=1
        if counter==30:
            break
    return workouts




legs = {"1":["Squats","Leg over head","smnthg else"],"2":["Cycling","Air cycles"],"3":["Running","Skipping"]}
arms = {"1":["dumbell 8 curls","wind circles","front and back stretch"],"2":["front and back with dumbels","push-ups","chin ups"],
"3":["Bell squats","Dead lifts"]}
cardio = {"1":["Planks","Situps","crunches"],"2":["Skipping","Jogging"],"3":["Swimming","Dancing"]}
x = calculate(legs,arms,cardio)
for stuff in x:
    print(stuff)