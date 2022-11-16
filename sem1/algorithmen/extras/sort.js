notSoRandomArray = []
for(let i = 100000; i >= 0; i--) {
    notSoRandomArray.push(i)
}
function shuffle(array) {
    let currentIndex = array.length,  randomIndex;
  
    // While there remain elements to shuffle.
    while (currentIndex != 0) {
  
      // Pick a remaining element.
      randomIndex = Math.floor(Math.random() * currentIndex);
      currentIndex--;
  
      // And swap it with the current element.
      [array[currentIndex], array[randomIndex]] = [
        array[randomIndex], array[currentIndex]];
    }
  
    return array;
  }
randomArray = shuffle(notSoRandomArray)


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

function insertionSort(disarray) {
    for (let i = 1; i < disarray.length; i++) {
        let current = disarray[i];
        let previousIndex = i-1; 
        while ((previousIndex > -1) && (current < disarray[previousIndex])) {
            disarray[previousIndex+1] = disarray[previousIndex];
            previousIndex--;
        }
        disarray[previousIndex+1] = current;
    }
    return disarray
}

function bubbleSort(disarray) {
    hasSwapped = true
    while (hasSwapped) {
        hasSwapped = false
        for (let i = 0; i < disarray.length - 1; i++) {
            if (disarray[i] > disarray[i+1]) {
                [disarray[i], disarray[i+1]] = [disarray[i+1], disarray[i]]
                hasSwapped = true;
            }
        }
    }
    return disarray;
}

/* The following code is the quicksort Implementation. */
function partition(items, left, right) {
    var pivot = items[Math.floor((right + left) / 2)]
    while (left <= right) {
        while (items[left] < pivot) {
            left++;
        }
        while (items[right] > pivot) {
            right--;
        }
        if (left <= right) {
            [items[left], items[right]] = [items[right], items[left]]
            left++;
            right--;
        }
    }
    return left;
}

function quickSort(items, left, right) {
    if (right < 0) {
        right = items.length-1
    }
    var index;
    if (items.length > 1) {
        index = partition(items, left, right);
        if (left < index - 1) {
            quickSort(items, left, index - 1);
        }
        if (index < right) {
            quickSort(items, index, right);
        }
    }
    return items;
}

/* End of quicksort */

/* Mergesort */

function mergeLists(lista, listb) {
    let biglist= []
    while (lista.length && listb.length) {
        if(lista[0] < listb[0]) {
            biglist.push(lista.shift())
        } else {
            biglist.push(listb.shift())
        }
    }
    return [ ...biglist, ...lista, ...listb]
}

function mergeSort(array) {
    if(array.length < 2){
      return array 
    }
    const left = array.splice(0, array.length/2)
    return mergeLists(mergeSort(left),mergeSort(array))
}

/* End of Mergesort */

function outputLongArray(array) {
    output = "["
    for ( let i = 0; i < array.length; i++ ) {
        output = output + array[i] + ", "
    }
    console.log(output+"]")
}

function visualize(array) {
    array.forEach(element => {
        console.log(String(element).padEnd(4," ")+"".repeat(element))
    });
}

visualize(randomArray)
console.log("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
var selSortPerf = performance.now()
console.log("selectionSort");visualize(selectionSort(randomArray))
selSortPerf = performance.now()-selSortPerf
console.log("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
var insSortPerf = performance.now()
console.log("insertionSort");visualize(insertionSort(randomArray))
insSortPerf = performance.now()-insSortPerf
console.log("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
var bubSortPerf = performance.now()
console.log("bubbleSort");visualize(bubbleSort(randomArray))
bubSortPerf = performance.now() - bubSortPerf
console.log("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
var quiSortPerf = performance.now()
console.log("quickSort");visualize(quickSort(randomArray, 0, -1))
quiSortPerf = performance.now() - quiSortPerf
console.log("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
var merSortPerf = performance.now()
console.log("mergeSort");visualize(mergeSort(randomArray))
merSortPerf = performance.now()- merSortPerf


console.log(`SelectionSort took ${selSortPerf}ms.`)
console.log(`InsertionSort took ${insSortPerf}ms.`)
console.log(`BubbleSort took ${bubSortPerf}ms.`)
console.log(`QuickSort took ${quiSortPerf}ms.`)
console.log(`MergeSort took ${merSortPerf}ms.`)



//console.table(JSON.stringify(selectionSort(randomArray), null, 2))
//console.table(JSON.stringify(insertionSort(randomArray), null, 2))
