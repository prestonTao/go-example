/**
 * KMP Algorithm
 * @author Jesse Wong (@straybugs)
 */

function kmp(sstr, lstr) {

  'use strict';

  if (typeof(sstr) !== 'string' || typeof(lstr) !== 'string') {
    return '';
  }
 
  // sstr should be shorter
  if (sstr.length > lstr.length) {
    sstr = [lstr, lstr = sstr][0];
  }

  var slen = sstr.length
  , llen = lstr.length
  , slenm1 = slen - 1
  , next = []
  , i = 0
  , j = -1
  ;
  next[0] = -1;
  
  // next
  while (i < slenm1) {
    if (j === -1 || sstr.charAt(i) === sstr.charAt(j)) {
      if (sstr.charAt(++i) !== sstr.charAt(++j)) {
        next[i] = j;
      } else {
        next[i] = next[j];
      }
    } else {
      j = next[j];
    }
  }

  // kmp
  i = j = 0;
  while (i < llen && j < slen) {
    if (j === -1 || lstr.charAt(i) === sstr.charAt(j)) {
      i += 1; 
      j += 1;
    } else {
      j = next[j];
    }
  }
  if (j >= slen) {
    return (i - slen);
  } else {
    return 0;
  }
}