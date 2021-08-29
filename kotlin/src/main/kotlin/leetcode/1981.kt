package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class MinimizeTheDifference1981 {
    fun minimizeTheDifference(mat: Array<IntArray>, target: Int): Int {
        val L = target + 70
        var pre = hashSetOf<Int>(0)
        for (row in mat) {
            val cur = hashSetOf<Int>()
            for (v in pre) {
                for (x in row)
                    if (v + x < L)
                        cur.add(v + x)
            }
            pre = cur
        }
        var min = Int.MAX_VALUE
        for (x in pre)
            min = minOf(min, if (target >= x) target - x else (x - target))
        return min
    }


    @Test
    fun test1() {
        assertEquals(1, 1)
    }
}
