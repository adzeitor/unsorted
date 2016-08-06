// recursive version of merge sort
function sort(a) {
  if (a.length <= 1) return a

  if (a.length == 2) {
    if (a[0] > a[1]) return [a[1],a[0]]
    return a
  }

  const center = Math.floor(a.length/2)
  return merge(
    sort(a.slice(0, center)),
    sort(a.slice(center)))
}

// merge two sorted arrays
function merge(a,b) {
  let items = []
  let i = 0
  let j = 0

  while(i < a.length && j < b.length) {
    if (a[i] <= b[j]) {
      items.push(a[i])
      i++
    } else {
      items.push(b[j])
      j++
    }
  }

  if (i >= a.length) return items.concat(b.slice(j))
  if (j >= b.length) return items.concat(a.slice(i))
  return items
}

exports.sort = sort
exports.merge = merge
