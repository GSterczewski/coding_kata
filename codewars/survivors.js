// challenge => https://www.codewars.com/kata/60a516d70759d1000d532029/

const testMomentum = [3,2,1,0];
const testPowerUps = [[1,0,0],[0,2,0,0],[0,9],[8,8]];

function survivors(listOfMomentum, listOfPowerups) {
  const result = [];
  listOfMomentum.forEach((momentum,index)=>{
    if(momentum <= 0){
      return;
    }
    let powerUps = listOfPowerups[index];
    let distance = powerUps.length;
     if(momentum > distance){
       result.push(index);
       return;
     }
     let position = 0;
     while(position < distance) {
       momentum += powerUps[position] - 1;
       position++;
       if(momentum <= 0) break; 
     }

     if(momentum > 0){
       result.push(index);
     }
  });
  return result;
}


console.log(survivors(testMomentum,testPowerUps));