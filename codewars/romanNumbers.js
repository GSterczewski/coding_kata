// challenge => https://www.codewars.com/kata/51b6249c4612257ac0000005/train/javascript


function solution(roman){
  const numeralsMap = {
    "M":1000,
    "D": 500,
    "C": 100,
    "L":50,
    "X":10,
    "V": 5,
    "I":1
  };
  
  const result = [];
  const numerals = roman.split("");
  let index = 0;
  while(index < numerals.length){
    let currentNumeral = numerals[index];
    let nextNumeral = numerals[index + 1];
    if(currentNumeral === "I"){
      if(nextNumeral === "X" || nextNumeral === "V"){
        result.push(numeralsMap[nextNumeral] - 1);
        index += 2;
      } else{
        result.push(1);
        index += 1;
      }
    } else {
      result.push(numeralsMap[currentNumeral]);
      index += 1;
    }
  }

  return result.reduce((sum,next) => sum + next, 0);   
}

const testCase1 = "XXI";
const testCase2 = "I";
const testCase3 = "MMVIII";
const testCase4 = "IV";
const testCase5 = "VI";
const testCase6 = "MDCLXVI";

console.log(solution(testCase1));
console.log(solution(testCase2));
console.log(solution(testCase3));
console.log(solution(testCase4));
console.log(solution(testCase5));
console.log(solution(testCase6));