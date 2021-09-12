package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class MaxmiumScoreLcp40 {
    fun maxmiumScore(cards: IntArray, cnt: Int): Int {
        cards.sort()
        val n = cards.size
        var (le, lo) = -1 to -1
        var (re, ro) = -1 to -1
        var sum = 0
        var odd = 0
        for ((i, v) in cards.withIndex())
            if (i < n - cnt) {
                if (v % 2 == 0)
                    le = v
                else
                    lo = v
            } else {
                sum += v
                if (v % 2 == 1)
                    odd = odd xor 1
                if (v % 2 == 0 && re == -1)
                    re = v
                else if (v % 2 == 1 && ro == -1)
                    ro = v
            }
        if (odd == 0)
            return sum
        var ret = 0
        if (re >= 0 && lo >= 0)
            ret = maxOf(ret, sum - re + lo)
        if (ro >= 0 && le >= 0)
            ret = maxOf(ret, sum - ro + le)
        return ret
    }


    @Test
    fun test1() {
        assertEquals(1, 1)
    }
}
