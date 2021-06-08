// challenge => https://www.codewars.com/kata/513e08acc600c94f01000001

function roundToValidRGB(n){
  if(n > 255) return 255;
  if(n < 0) return 0;
  return n;
}

function toHex(n){
  if(n < 16) return `0${n.toString(16)}`;
  return n.toString(16);
}
// this is very easy way, using built-in number to string conversion with base 16;
function rgb(r, g, b){

  return `${toHex(roundToValidRGB(r))}${toHex(roundToValidRGB(g))}${toHex(roundToValidRGB(b))}`.toUpperCase();  
}

console.log(rgb(10,10,15));