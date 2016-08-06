const test = require('tape')
const brackets = require('.')

const table = [
  {in:"{({[[{}]]}){}{}[]}", res:true},
  {in:"abcdef", res:true},
  {in:"][", res:false},
  {in:"[", res:false},
  {in:"", res:true},
  {in:"}", res:false},
  {in:"[)", res: false},
  {in:"(}", res: false},
  {in:"{]", res: false},
]


test('iterative version 1', function (t) {
  t.plan(table.length)
  table.forEach((x) => t.equal(brackets.isBalanced1(x.in), x.res, x.in))
})

test('iterative version 2', function (t) {
  t.plan(table.length)
  table.forEach((x) => t.equal(brackets.isBalanced2(x.in), x.res, x.in))
})
