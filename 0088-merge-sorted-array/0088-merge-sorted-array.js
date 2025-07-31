/**
 * @param {number[]} nums1
 * @param {number} m
 * @param {number[]} nums2
 * @param {number} n
 * @return {void} Do not return anything, modify nums1 in-place instead.
 */
var merge = function (nums1, m, nums2, n) {
    var N1 = 0, N2 = 0, copyN1 = [...nums1];

    for (var i = 0; i < m + n; i++) {
        if (N2 == n) {
            nums1[i] = copyN1[N1];
            N1++;
        } else if (N1 == m) {
            nums1[i] = nums2[N2];
            N2++;
        } else if (copyN1[N1] <= nums2[N2]) {
            nums1[i] = copyN1[N1];
            N1++;
        } else {
            nums1[i] = nums2[N2];
            N2++;
        }
    }
};