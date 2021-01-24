// PASSING 100% TEST CASES

function processData(input) {
    
    // handle input string and parse it into 2d array of numbers without first index(unnecessary)
    const parsedInput = input.split("\n").slice(1).map(s => s.split(" ").map(s => Number(s)) )
    
    const stack = []
    const getLastElement = arr => arr[arr.length - 1]
    const getMaxElement = (arr,value) => Math.max(getLastElement(arr), value )
    parsedInput.forEach(line => {
        const queryType = line[0]
    
        switch(queryType){
            case 1: stack.length > 0 ? stack.push(getMaxElement(stack,line[1])) : stack.push(line[1]);break;
            case 2 : stack.pop(); break;
            case 3 : console.log(getLastElement(stack)); break;
        }
        
    })
    
}