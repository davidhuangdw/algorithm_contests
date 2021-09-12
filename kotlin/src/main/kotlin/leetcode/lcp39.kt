package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class MinimumSwitchingTimesLcp39 {
    fun minimumSwitchingTimes(source: Array<IntArray>, target: Array<IntArray>): Int {
        val s = hashMapOf<Int, Int>()
        val t = hashMapOf<Int, Int>()
        for (a in source)
            for (v in a)
                s[v] = s.getOrDefault(v, 0) + 1
        for (a in target)
            for (v in a)
                t[v] = t.getOrDefault(v, 0) + 1
        var sum = 0
        for ((v, c) in t)
            sum += maxOf(0, c - s.getOrDefault(v, 0))
        return sum
    }


    @Test
    fun test1() {
        assertEquals(1, 1)
    }
}
