/**
 * @param {string[]} words
 * @param {character} separator
 * @return {string[]}
 */
var splitWordsBySeparator = function(words, separator) {
    let result=[]
    for(let w of words){
        result.push(...w.split(separator))
    }
    return result.filter(str=>str!=="")
};