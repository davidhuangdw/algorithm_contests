package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class MakeFancyString5193 {
    fun makeFancyString(s: String): String {
        val n = s.length
        var i = 0
        return buildString {
            while (i < n) {
                val ch = s[i]
                var cnt = 0
                var j = i
                while (j < n && s[j] == ch) {
                    cnt += 1
                    if (cnt <= 2)
                        append(ch)
                    j += 1
                }
                i = j
            }
        }
    }


    @Test
    fun test1() {
        assertEquals(1, 1)
    }
}
