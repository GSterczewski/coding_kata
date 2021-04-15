// challenge => https://www.codewars.com/kata/54bf1c2cd5b56cc47f0007a1

function duplicateCount(text){
  text = text.toLowerCase();
  const counts = {};
  text.split("").forEach(letter => {
    counts[letter] ? counts[letter]++ : counts[letter] = 1;
    
  })
  return Object.values(counts).filter(val => val > 1).length;
}

