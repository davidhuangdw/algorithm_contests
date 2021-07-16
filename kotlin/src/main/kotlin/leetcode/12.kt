package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class IntToRoman12 {
    fun intToRoman(num: Int): String {
        val keys = listOf(
            1000 to "M-",
            100 to "CD",
            10 to "XL",
            1 to "IV",
        )
        var cur = num
        return buildString {
            var pre = ""
            for ((base, key) in keys) {
                if (cur >= base) {
                    val d = cur / base
                    if (d == 9) {
                        append(key[0], pre[0])
                    } else if (d >= 5) {
                        append(key[1])
                        repeat(d - 5) { append(key[0]) }
                    } else if (d == 4) {
                        append(key[0])
                        append(key[1])
                    } else repeat(d) { append(key[0]) }
                    cur %= base
                }
                pre = key
            }
        }
    }


    @Test
    fun test1() {
        assertEquals("MCMXCIV", intToRoman(1994))
        assertEquals("LVIII", intToRoman(58))
        assertEquals("III", intToRoman(3))
    }
}
