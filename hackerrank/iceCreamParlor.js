/**
 * problem https://www.hackerrank.com/challenges/ctci-ice-cream-parlor/problem
 * @param {number[]} cost 
 * @param {number} money 
 */
function whatFlavors(cost, money) {

  const costMap = {};
  for(let i = 0; i < cost.length; i++){
      
      if(costMap[money- cost[i]] >= 0){
          console.log(`${costMap[money-cost[i]] + 1} ${i + 1}`)
          break;
      } else{
          costMap[cost[i]] = i;
      }
      
  }
  
}