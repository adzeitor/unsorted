function isBalanced1(s) {
  const brackets = []
  const opens = {'(':')', '[':']', '{':'}'}
  const ends =  {')':'(', ']':'[', '}':'{'}

  for(i = 0; i < s.length; i++) {
    const v = s[i]
    // open bracket?
    if (opens[v]) {
       brackets.push(opens[v])
    } else if (ends[v]) {
      if (brackets.length == 0) return false;
      if (brackets.pop() != v) return false;
    }
  }
  if (brackets.length == 0) return true;
  return false
}

function isBalanced2(s) {
  const brackets = []
  const opens = ['(', '[', '{']
  const ends =  [')', '[', '{']

  for(i = 0; i < s.length; i++) {
    const v = s[i]
    let j = 0
    // open bracket?
    if ((j = opens.indexOf(v)) >= 0) {
       brackets.push(ends[j])
    } else if (ends.indexOf(v) >= 0) {
      if (brackets.length == 0) return false;
      if (brackets.pop() != v) return false;
    }
  }

  if (brackets.length == 0) return true;
  return false
}

// TODO
function isBalancedRecursive() {

}

exports.isBalanced1 = isBalanced1
exports.isBalanced2 = isBalanced2
