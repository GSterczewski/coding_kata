

function gcd(num1,num2){
  const max = Math.max(num1,num2);
  const min = Math.min(num1,num2);
  
  const mod = max % min;
  if(mod === 0) {
    return min ;
  }
  return gcd(min, mod);
}

console.log(gcd(10,3));
console.log(gcd(100,15));
console.log(gcd(4,8));
console.log(gcd(1333344,24348));