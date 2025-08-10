var plusOne = function(digits) {
    let joined = digits.join("");  
    let num = BigInt(joined) + 1n; 
    return num.toString().split("").map(Number);
};

console.log(plusOne([1,2,3]));
