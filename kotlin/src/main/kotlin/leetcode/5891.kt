package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class MissingRolls5891 {
    fun missingRolls(rolls: IntArray, mean: Int, n: Int): IntArray {
        var s = mean * (rolls.size + n) - rolls.sum() - n
        if (s !in 0..5 * n)
            return intArrayOf()
        val r = IntArray(n) { 1 }
        for (i in 0 until n) {
            if (s == 0) break
            val t = minOf(s, 5)
            s -= t
            r[i] += t
        }
        return r
    }


    @Test
    fun test1() {
        assertEquals(1, 1)
    }
}
