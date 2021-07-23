package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class MinSwapsCouples765 {
    fun minSwapsCouples(row: IntArray): Int {
        val n = row.size
        val pos = MutableList(n) { 0 }
        for ((i, v) in row.withIndex())
            pos[v] = i

        var cnt = 0
        for (i in 0 until n step 2) {
            if (row[i + 1] != row[i] xor 1) {
                val j = pos[row[i] xor 1]
                row[j] = row[i + 1]
                pos[row[i + 1]] = j
                cnt += 1
            }
        }
        return cnt
    }

    fun minSwapsCouples1(row: IntArray): Int {
        val n = row.size
        val par = MutableList(n) { it }
        val roots = (0 until n).toHashSet()

        fun find(x: Int): Int {
            if (par[x] != x)
                par[x] = find(par[x])
            return par[x]
        }

        fun union(a: Int, b: Int) {
            val (ra, rb) = find(a) to find(b)
            if (ra != rb) {
                par[rb] = ra
                roots.remove(rb)
            }
        }
        for (i in 0 until n step 2) {
            union(i, i + 1)
            union(row[i], row[i + 1])
        }
        return n / 2 - roots.size
    }


    @Test
    fun test1() {
        assertEquals(3, minSwapsCouples(parseIntArray("[0,2,4,6,7,1,3,5]")))
        assertEquals(1, minSwapsCouples(parseIntArray("[0, 2, 1, 3]")))
        assertEquals(0, minSwapsCouples(parseIntArray("[3, 2, 0, 1]")))
    }
}
