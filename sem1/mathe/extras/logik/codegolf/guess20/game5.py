import random
tries = 5
random_choice = random.randint(1, 20)
while tries > 0:
    guess = int(input("Guess the number between 1 and 20! >"))
    if guess == random_choice:
        print("Good job!")
        exit()
    elif guess < random_choice:
        print("The number is higher.")
    else:
        print("The number is lower.")
    tries = tries - 1
print("You failed.")
