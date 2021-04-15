function countName(anArr, name){
  return anArr.filter($name => $name === name).length;
  };


  const MAINLIST = ["Bob","Ted","Amy","Alice","Bob","Ted","Amy","Ted","Amy","Fred"];  

  console.log(countName(MAINLIST, "Ted"));