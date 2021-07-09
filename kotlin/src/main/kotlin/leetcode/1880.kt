package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class IsSumEqual {
    fun isSumEqual(firstWord: String, secondWord: String, targetWord: String): Boolean {
        fun toNumber(s: String): Int {
            var x = 0
            for (ch in s)
                x = x * 10 + (ch - 'a')
            return x
        }
        return toNumber(firstWord) + toNumber(secondWord) == toNumber(targetWord)
    }


    @Test
    fun test1() {
        assertEquals(true, isSumEqual("acb", "cba", "cdb"))
        assertEquals(false, isSumEqual("aaa", "a", "aab"))
        assertEquals(true, isSumEqual("aaa", "a", "aaaa"))
    }
}
