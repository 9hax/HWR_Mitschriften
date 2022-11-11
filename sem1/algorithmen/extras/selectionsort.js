randomArray = [46, 21,62,59,53,87,64,6,99,88,24,84,9,64,93,73,30,7,72]

notSoRandomArray = []
for(let i = 255; i >= 0; i--) {
    notSoRandomArray.push(i)
}

swaps = 0
function selectionSort(disarray) {
    let n = disarray.length;
        
    for(let i = 0; i < n; i++) {
        let min = i;
        for(let j = i+1; j < n; j++){
            if(disarray[j] < disarray[min]) {
                min=j; 
            }
         }
         if (min != i) {
             let tmp = disarray[i]; 
             disarray[i] = disarray[min];
             disarray[min] = tmp;
             swaps = swaps + 1
        }
    }
    return disarray;
}

function outputLongArray(array) {
    output = "["
    for ( let i = 0; i < array.length; i++ ) {
        output = output + array[i] + ", "
    }
    console.log(output+"]")
}

//outputLongArray(selectionSort(notSoRandomArray))

console.log(JSON.stringify(selectionSort(notSoRandomArray), null, 2))