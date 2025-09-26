/**
 * @param {number} n
 * @return {number}
 */
var subtractProductAndSum = function(n) {
    let str=n.toString().split("")
    let add=0
    let sub=1
    for(let i=0;i<str.length;i++){
        add+=Number(str[i])
        sub*=Number(str[i])
    }
    return sub-add
};