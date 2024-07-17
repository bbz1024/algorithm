package main

func main() {

}
func canFinish(numCourses int, prerequisites [][]int) bool {
	var (
		edges  = make([][]int, numCourses)
		indeg  = make([]int, numCourses) // 入边，即依赖的
		result []int
	)
	for _, info := range prerequisites {
		edges[info[1]] = append(edges[info[1]], info[0]) //当前顶点相邻的边，1->2 1->3
		indeg[info[0]]++                                 // 相邻边的顶点入边数量+1
	}
	var q []int // 队列，逐个出队无入边的顶点
	// 1. 记录无入边的顶点
	for i := 0; i < numCourses; i++ {
		if indeg[i] == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		result = append(result, u)
		for _, v := range edges[u] {
			//遍历该顶点的所有邻接顶点，将其入度减1
			indeg[v]--
			//如果某个邻接顶点的入度减为0，则将其加入队列
			if indeg[v] == 0 {
				q = append(q, v)
			}
		}
	}
	//判断是否存在环路
	return len(result) == numCourses
}
