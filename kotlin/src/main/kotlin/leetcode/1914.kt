package leetcode

class RotateGrid {
    fun rotateGrid(grid: Array<IntArray>, k: Int): Array<IntArray> {
        val (n, m) = grid.size to grid[0].size
        val res = Array(n) { IntArray(m) }

        var (l, r, u, d) = listOf(0, m - 1, 0, n - 1)
        while (l < r && u < d) {
            val cycle_index = mutableListOf<Pair<Int, Int>>()
            for (j in l..r) cycle_index.add(u to j)
            for (i in u + 1..d) cycle_index.add(i to r)
            for (j in r - 1 downTo l) cycle_index.add(d to j)
            for (i in d - 1 downTo u + 1) cycle_index.add(i to l)

            val len = cycle_index.size
            val off = k % len
            for ((d, ij) in cycle_index.withIndex()) {
                val (i, j) = ij
                val (ti, tj) = cycle_index[(off + d) % len]
                res[i][j] = grid[ti][tj]
            }

            l += 1
            r -= 1
            u += 1
            d -= 1
        }
        return res
    }

}