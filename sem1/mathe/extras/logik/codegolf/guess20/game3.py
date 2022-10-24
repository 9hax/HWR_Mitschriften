import random
tries=5
random_choice=random.randint(1,20)
while tries>0:
    guess=int(input(">"))  
    if guess==random_choice:print(":)");exit()
    elif guess<random_choice:print("h")
    else:print("l")
    tries-=1
print("f")