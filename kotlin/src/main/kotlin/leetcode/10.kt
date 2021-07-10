package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class IsMatch10 {
    fun isMatch(s: String, p: String): Boolean {
        val (n, m) = s.length to p.length
        val can = MutableList(n+1){it == 0}
        var i = 0;
        fun eq(a: Char, b: Char) = a == '.' || a==b
        while(i < m){
            val star = i+1 < m && p[i+1] == '*'
            val pc = p[i]
            var pre = can[0]
            can[0] = can[0] && star
            for(j in 1..n){
                val sc = s[j-1]
                val cur = can[j]
                can[j] = if(star){
                    if(eq(pc, sc)) can[j-1] || can[j]  // bug: for "a" "aa*", need leave "a" to previous one
                    else can[j]
                }else eq(pc, sc) && pre
                pre = cur
            }
            i += if(star) 2 else 1
//            println("-".repeat(10))
//            println(s)
//            println(p.substring(0, i))
//            println(can)
        }
        return can[n]
    }


    @Test
    fun test1() {
        assertEquals(true, isMatch("aasdfasdfasdfasdfas", "aasdf.*asdf.*asdf.*asdf.*s"))
        assertEquals(true, isMatch("aaa", "ab*a*c*a"))
        assertEquals(false, isMatch("aa", "a"))
        assertEquals(true, isMatch("aa", "a*"))
        assertEquals(true, isMatch("ab", ".*"))
        assertEquals(true, isMatch("mississippi", "mis*is*ip*.i"))
        assertEquals(false, isMatch("mississippi", "mis*is*p*."))
    }
}
