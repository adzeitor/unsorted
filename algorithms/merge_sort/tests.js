const test = require('tape')
const merge = require('.')

const table = [
  {in:[1,2,3,4,5], out:[1,2,3,4,5]},
  {in:[5,4,3,2,1], out:[1,2,3,4,5]},
  {in:[], out:[]},
  {in:[1,2], out:[1,2]},
  {in:[6,5], out:[5,6]},
  {in:[1], out:[1]},
  {in:[3,1,2], out:[1,2,3]},
]

const mergeTable = [
  {a:[1,2,3], b:[4,5], out:[1,2,3,4,5]},
  {a:[], b:[1], out:[1]},
  {a:[1], b:[], out:[1]},
  {a:[1], b:[0], out:[0,1]},
  {a:[1,3,5,7], b:[2,4,6], out:[1,2,3,4,5,6,7]},
]

test('merge arrays test', function (t) {
  t.plan(mergeTable.length)
  mergeTable.forEach((x) => t.deepEqual(merge.merge(x.a,x.b), x.out, JSON.stringify([x.a,x.b])))
})

test('recursive mergesort test', function (t) {
  t.plan(table.length)
  table.forEach((x) => t.deepEqual(merge.sort(x.in), x.out, JSON.stringify(x.in)))
})
