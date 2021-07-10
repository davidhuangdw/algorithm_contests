package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class SuperEggDrop887 {
    fun superEggDrop(k: Int, n: Int): Int {
        val max_floors = MutableList(n + 1) { 0 }
        var res = 0
        repeat(k) {
            var pre = 0
            for (t in 1..n) {
                max_floors[t] = (1 + max_floors[t - 1] + pre).also {
                    pre = max_floors[t]
                }
                if(max_floors[t] >= n){
                    res = t
                    break
                }
            }
        }
        return res
    }

    fun superEggDrop1(k: Int, n: Int): Int {
        val min = List(2) { MutableList(n + 1) { it } }
        for (i in 2..k) {
            var cur = min[i and 1]
            var last = min[i and 1 xor 1]
            var t = 1
            for (x in 2..n) {
                while (t < x && cur[t] < last[x - 1 - t]) t += 1
                cur[x] = 1 + cur[t]
            }
        }
        return min[k and 1][n]
    }


    @Test
    fun test1() {
        assertEquals(2, superEggDrop(1, 2))
        assertEquals(3, superEggDrop(2, 6))
        assertEquals(4, superEggDrop(3, 14))
    }
}
