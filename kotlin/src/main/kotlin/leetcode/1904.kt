package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class NumberOfRounds {
    fun numberOfRounds(startTime: String, finishTime: String): Int {
        fun timeToInt(s: String): Int {
            val (h, m) = s.split(":").map { it.toInt() }
            return h * 60 + m
        }
        var (s, f) = timeToInt(startTime) to timeToInt(finishTime)
        if (f < s) f += 24 * 60
        return maxOf(0, f / 15 - (s + 14) / 15)
    }

    @Test
    fun test1() {
        assertEquals(95, numberOfRounds("00:00", "23:59"))
        assertEquals(40, numberOfRounds("20:00", "06:00"))
        assertEquals(0, numberOfRounds("00:47", "00:57"))
    }
}