/**
 * problem -> https://leetcode.com/problems/best-time-to-buy-and-sell-stock/;
 * @param {number[]} prices
 * @return {number}
 */
const maxProfit = prices => {
    
  const length = prices.length; 
  let min = prices[0];
  let profit = 0;
  
  for(let i = 1; i < length; i++){
     
      if( prices[i] < min){
          min = prices[i];
      }
      else {
          profit = Math.max(profit,prices[i] - min);
      }
      }
  return profit;
  };