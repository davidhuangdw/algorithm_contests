package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class IsMatch44 {
    fun isMatch(s: String, p: String): Boolean {
        val m = p.length
        val match = MutableList(m+1){false}
        match[0] = true
        for((i, ch) in p.withIndex()){
            if(ch != '*') break
            match[i+1] = true
        }
        for (sc in s) {
            var pre = match[0]
            match[0] = false
            for ((i, pc) in p.withIndex()) {
                val j = i + 1
                val before = match[j]
                match[j] = if (pc == '*') {
                    match[i] || match[j]
                } else if (pc == '?' || pc == sc) {
                    pre
                } else false
                pre = before
            }
        }
        return match[m]
    }

    @Test
    fun test1() {
        assertEquals(false, isMatch("aa", "a"))
        assertEquals(true, isMatch("aa", "*"))
        assertEquals(true, isMatch("cb", "?b"))
        assertEquals(false, isMatch("cb", "?a"))
        assertEquals(true, isMatch("adceb", "*a*b"))
        assertEquals(false, isMatch("acdcb", "a*c?b"))
    }
}
