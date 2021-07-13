package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class LongestPalindromeSubseq516 {
    fun longestPalindromeSubseq(s: String): Int {
        val n = s.length
        val f = List(n){MutableList(n){0} }
        for(i in n-1 downTo 0){
            f[i][i] = 1
            for(j in i+1 until n){
                f[i][j] = maxOf(f[i][j-1], f[i+1][j], if(s[i] == s[j]) f[i+1][j-1] + 2 else 0)
            }
        }
        return f[0][n-1]
    }


    @Test
    fun test1() {
        assertEquals(4, longestPalindromeSubseq("bbbab"))
    }
}
