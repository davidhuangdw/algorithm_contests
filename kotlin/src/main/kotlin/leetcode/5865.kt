package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class FirstDayBeenInAllRooms5865 {
    fun firstDayBeenInAllRooms(nextVisit: IntArray): Int {
        val MOD = 1000_000_007
        val n = nextVisit.size
        val f = MutableList(n) { 0L }
        val s = MutableList(n + 1) { 0L }
        for ((i, x) in nextVisit.withIndex()) {
            f[i] = (s[i] - s[x] + 2 + MOD) % MOD
            s[i + 1] = (s[i] + f[i]) % MOD
        }
        println(f)
        println(s)
        return s[n - 1].toInt()
    }


    @Test
    fun test1() {
        assertEquals(2, firstDayBeenInAllRooms(parseIntArray("[0,0]")))
    }
}
