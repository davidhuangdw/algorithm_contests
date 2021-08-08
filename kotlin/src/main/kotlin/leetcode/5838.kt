package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class IsPrefixString5838 {
    fun isPrefixString(s: String, words: Array<String>): Boolean {
        val n = s.length
        val m = words.size
        var i = 0
        var j = 0
        while (i < n) {
            if (j < m) {
                for (ch in words[j]) {
                    if (i < n && s[i] == ch)
                        i += 1
                    else
                        return false
                }
                j += 1
            } else
                return false
        }
        return i == n
    }


    @Test
    fun test1() {
        assertEquals(1, 1)
    }
}
