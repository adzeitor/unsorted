
const Graph = (nodes, edges) => {
  const neighbors = (node) => {
    let list = []
    edges.forEach((edge) => {
      if (edge[0] == node) {
        list.push(edge[1])
      }
      if (edge[1] == node) {
        list.push(edge[0])
      }
    })

    return list
  }

  const length = (u,v) => {
    const e = edges.find(edge =>
      (edge[0] == u && edge[1] == v) ||
      (edge[0] == v && edge[1] == u)
    )

    return e[2]
  }

  return {
    neighbors: neighbors,
    nodes: nodes,
    edges: edges,
    length: length,
  }
}


// Dijkstra's algorithm is an algorithm for finding the shortest paths between
// nodes in a graph
function dijkstra(graph, source) {
  // ditances from src to all nodes
  let dist = {}
  // previous node with shortest path
  let prev = {}
  let unvisited = []

  graph.nodes.forEach(x => {
    unvisited.push(x)
    dist[x] = Infinity
    prev[x] = undefined
  })
  dist[source] = 0

  while(unvisited.length > 0) {
    // index of node with minimal dist
    let u = unvisited[0]
    let index = 0
    // find node with minimal dist[x]
    unvisited.forEach((x,i) => {
      if (dist[x] < dist[u]) {
        u = x
        index = i
      }
    })

    // delete
    unvisited.splice(index, 1)

    let ns = graph.neighbors(u)

    ns.forEach(v => {
      let d = dist[u] + graph.length(v,u)
      if (d < dist[v]) {
        dist[v] = d
        prev[v] = u
      }
    })
  }

  return {
    prev: prev,
    dist : dist,
    path : (dest) => {
      let list = [dest]
      let x = dest
      while(prev[x] != source) {
        x = prev[x]
        list.unshift(x)
      }
      return list
    },
    cost : (dest) => {
      return dist[dest]
    }
  }
}


exports.Graph = Graph
exports.dijkstra = dijkstra
