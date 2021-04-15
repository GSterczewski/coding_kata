// challege => https://www.codewars.com/kata/5500d54c2ebe0a8e8a0003fd

function gcd(num1,num2){
  const max = Math.max(num1,num2);
  const min = Math.min(num1,num2);
  
  const mod = max % min;
  if(mod === 0) {
    return min ;
  }
  return gcd(min, mod);
}
