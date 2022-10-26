console.log("Hello World!")

var name = "Carl";

function checkCarl(name) {
    if (name == "Carl") {
        return true
    }
    return false
}

console.log("Ist der Name " + name + " denn jetzt auf Carl gesetzt?")

console.log(checkCarl(name))

for (var i = 0; i < 10; i++) {
    console.log(i)
}

function checkPrime(number) {
    if (number < 2) return false;
    if (number == 2) return true;
    if (number % 2 == 0) return false;
    for (var divident = 2; divident < Math.floor(number / 2); divident++) if (number % divident == 0) return false;
    return true;
}

for (var i = 0; i < 100; i++) console.log(i, checkPrime(i))

function genFibonacci(index) {
    var n1 = 0, n2 = 1, next = 0;
    for (var limit = 1; limit <= index; limit++) {
        next = n1 + n2
        n1 = n2;
        n2 = next;
    }
    return next;
}

function recursiveFibonacci(index) {
    if (index <= 1) return index;
    return recursiveFibonacci(index-1)+recursiveFibonacci(index-2)
}

for (var i = 0; i < 40; i++) console.log(i, genFibonacci(i))
for (var i = 0; i < 40; i++) console.log(i, recursiveFibonacci(i))


function exercise(n) {
    if (n>10) {
        return exercise(n-2);
    } else {
        return (n-1);
    }
}

console.log(exercise(8),
exercise(10),
exercise(11),
exercise(15))

function genFac(number) {
    if (number<2) return 1;
    return genFac(number-1)*number;
}

for (var i = 0; i < 50; i++) console.log(genFac(i));

function genFacIter(limit) {
    number = 1;
    for (var i = 1; i <= limit; i++) {
        number = number * i;
    }
    return number;
}

for (var i = 0; i < 50; i++) console.log(genFacIter(i));

const multi = (a,b) => a*b;

console.log(
    multi(
        multi(
            multi(8,7),
            multi(6,5)
        ),
        multi(
            multi(4,3),
            multi(2,1)
        )
    )
)