const test = require('tape')
const dijkstra = require('.')

test('simple cases', function (t) {
  const g1 = dijkstra.Graph([1,2,3,4,5], [[1,3,10]])
  const d1 = dijkstra.dijkstra(g1, 1)

  t.plan(3)
  t.equal(d1.dist['2'], Infinity)
  t.equal(d1.dist['3'], 10)
  t.equal(d1.dist['1'], 0)
})

test('loop test', function (t) {
  const g1 = dijkstra.Graph([1,2,3], [[1,2,5], [2,3,5], [3,1,5]])
  const d1 = dijkstra.dijkstra(g1, 1)

  t.plan(2)
  t.equal(d1.dist['2'], 5)
  t.equal(d1.dist['3'], 5)
})

test('wikipedia example test', function (t) {
  const g1 = dijkstra.Graph([1,2,3,4,5,6], [
    [1,2,7],
    [1,3,9],
    [1,6,14],
    [2,3,10],
    [2,4,15],
    [3,6,2],
    [3,4,11],
    [6,5,9],
    [5,4,6]
  ]);
  const d = dijkstra.dijkstra(g1, 1)


  t.plan(2)
  t.deepEqual(d.dist, { '1': 0, '2': 7, '3': 9, '4': 20, '5': 20, '6': 11 })
  t.deepEqual(d.path(5), [3, 6, 5])
})
