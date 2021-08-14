package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class LastRemaining62 {
    fun lastRemaining(n: Int, m: Int): Int {
        if (m == 1) return n - 1
        fun f(l: Int): Int {
            if (l == 1) return 0
            if (l < m) return (f(l - 1) + m) % l
            val s = l / m
            var i = f(l - s) + s * m
            return if (i < l) i else i % l + i % l / (m - 1)
        }
//        for(l in 1 until 10)
//            println("$m, $l: ${f(l)} : ${lastRemaining2(l, m)}")
        return f(n)
    }

    fun lastRemaining2(n: Int, m: Int): Int {
        var i = 0
        for (l in 2..n)
            i = (i + m) % l
        return i
    }


    @Test
    fun test1() {
        assertEquals(3, lastRemaining(5, 3))
        assertEquals(2, lastRemaining(10, 17))
    }
}
