package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class FractionToDecimal166 {
    fun fractionToDecimal(numerator: Int, denominator: Int): String {
        return buildString {
            var (x, y) = numerator.toLong() to denominator.toLong()
            if (y < 0) {
                x = -x
                y = -y
            }
            if (x < 0) {
                x = -x
                append("-")
            }
            append(x / y)
            var cur = x % y
            if (cur == 0L) return@buildString

            append(".")
            cur *= 10
            val pos = hashMapOf<Long, Int>()
            val past = mutableListOf<Long>()
            while (cur > 0 && cur !in pos) {
                pos[cur] = past.size
                past.add(cur / y)
                cur = (cur % y)
                cur *= 10
            }
            if (cur == 0L) append(past.joinToString(""))
            else {
                val st = pos[cur]!!
                append(past.subList(0, st).joinToString(""))
                append("(")
                append(past.subList(st, past.size).joinToString(""))
                append(")")
            }
        }
    }


    @Test
    fun test1() {
        assertEquals("-6.25", fractionToDecimal(-50, 8))
        assertEquals("0.5", fractionToDecimal(1, 2))
        assertEquals("2", fractionToDecimal(2, 1))
        assertEquals("0.(6)", fractionToDecimal(2, 3))
        assertEquals("0.(012)", fractionToDecimal(4, 333))
        assertEquals("0.2", fractionToDecimal(1, 5))
    }
}
