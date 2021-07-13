package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class ReorganizeString767 {
    fun reorganizeString(s: String): String {
        val n = s.length
        val count = hashMapOf<Char, Int>()
        for(ch in s)
            count[ch] = count.getOrDefault(ch, 0) + 1
        val cs = count.keys.sortedByDescending { count[it] }
        if(count[cs[0]]!! > (n+1)/2) return ""
        val res = MutableList(n){' '}
        var i = 0
        for(ch in cs){
            repeat(count[ch]!!){
                res[i] = ch
                i += 2
                if(i >= n) i = 1
            }
        }
        return res.joinToString("")
    }


    @Test
    fun test1() {
        assertEquals("aba", reorganizeString("aab"))
        assertEquals("", reorganizeString("aaab"))
    }
}
