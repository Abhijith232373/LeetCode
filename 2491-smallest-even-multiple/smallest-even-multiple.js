/**
 * @param {number} n
 * @return {number}
 */
var smallestEvenMultiple = function(n) {
 
  return  n % 2 === 0 ?  n :  n * 2
    // let mult=[]
    // for(let i=1;i<=2;i+2){
    //     mult=n*i
    //    let lastop=Math.min(mult)
    //    return  lastop
        // }
};
console.log(smallestEvenMultiple(5))