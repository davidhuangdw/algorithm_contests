package leetcode

import org.junit.jupiter.api.Assertions.assertArrayEquals
import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

// https://leetcode-cn.com/problems/max-submatrix-lcci/
class GetMaxMatrix {
    fun getMaxMatrix(matrix: Array<IntArray>): IntArray {
        val (n, m) = matrix.size to matrix[0].size
        val pre_sum = List(n + 1) { MutableList(m + 1) { 0L } }
        for (i in 1..n)
            for (j in 1..m)
                pre_sum[i][j] = pre_sum[i - 1][j] + pre_sum[i][j - 1] - pre_sum[i - 1][j - 1] + matrix[i - 1][j - 1]

        var res = Long.MIN_VALUE to intArrayOf(-1, -1, -1, -1)
        for (u in 0 until n)
            for (d in u + 1..n) {
                var lmin = 0L to 0
                for (j in 1..m) {
                    val cur = pre_sum[d][j] - pre_sum[u][j]
                    if (cur - lmin.first > res.first) {
                        res = cur - lmin.first to intArrayOf(u, lmin.second, d - 1, j - 1)
                    }
                    if (cur < lmin.first)
                        lmin = cur to j
                }
            }

//        println(res.first to res.second.toList())
        return res.second
    }


    @Test
    fun test1() {
        assertEquals(intArrayOf(0, 1, 0, 1).toList(), getMaxMatrix(parse2dIntArray("[[-1,0],[0, -1]]")).toList())
        assertEquals(
            intArrayOf(1, 0, 1, 2).toList(),
            getMaxMatrix(parse2dIntArray("[[-5, 2, -2],[4, -2, 3]]")).toList()
        )
    }
}
