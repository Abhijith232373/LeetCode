/**
 * @param {number[]} nums
 * @param {number} n
 * @return {number[]}
 */
var shuffle = function(nums, n) {
let newarr=[]
for(let k=0;k<n;k++){
    newarr.push(nums[k],nums[k+n])
}
return newarr
};
shuffle([3,4,7,2,5,1],3)