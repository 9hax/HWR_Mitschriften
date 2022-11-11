gleise = []

for (var i = 0; i < 3; i++) {
    gleise[i] = []
}

gleise[0].push("ZugA")
gleise[0].push("WagenA1")
gleise[0].push("WagenA2")
gleise[1].push("ZugB")
gleise[1].push("WagenB1")
gleise[1].push("WagenB2")

console.log(gleise)

function pop(gleis) {
    return gleise[gleis].pop()
}

function push(gleis, zug) {
    gleise[gleis].push(zug)
}

push(2,pop(1))
push(2,pop(1))
push(2,pop(1))

push(1,pop(0))
push(1,pop(0))
push(1,pop(0))

push(0,pop(2))
push(0,pop(2))
push(0,pop(2))

console.log(gleise)