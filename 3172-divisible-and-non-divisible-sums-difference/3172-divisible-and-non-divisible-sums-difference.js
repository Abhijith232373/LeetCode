/**
 * @param {number} n
 * @param {number} m
 * @return {number}
 */
var differenceOfSums = function(n, m) {
    let div=0
    let not=0
    for(let i=1;i<=n;i++){
        if(i%m===0){
            div+=i
        }else{
            not+=i
        }
    }
    return not-div
};