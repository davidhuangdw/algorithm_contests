package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class IsMatch {
    fun isMatch(s: String, p: String): Boolean {
//        println("$s - $p")
        val SKIP = ".*"
        val (n, m) = s.length to p.length
        val can = List(n + 1) { BooleanArray(m + 1) }
        can[n][m] = true

        fun eq(a: Char, b: Char) = a == '.' || b == '.' || a == b
        fun judge(i: Int, j: Int): Boolean {
//            println("${s.substring(i, n)} | ${p.substring(j, m)}")
            if (i == n && j == m) return true
            if(i+1 < n && s[i+1] == '*')
                return can[i+2][j] || (j < m && eq(s[i], p[j]) && can[i][j+1])
            if(j+1 < m && p[j+1] == '*')
                return can[i][j+2] || (i < n && eq(s[i], p[j]) && can[i+1][j])
            return (i < n && j < m && eq(s[i], p[j]) && can[i+1][j+1])
        }

        for (i in n downTo 0)                   // bug!: range from n, not n-1
            for (j in m downTo 0) {
                can[i][j] = judge(i, j)
//                println("${s.substring(i, n)} | ${p.substring(j, m)} | ${can[i][j]}")
            }
        return can[0][0]
    }

    @Test
    fun test1() {
        fun test(a: String, b: String, expected: Boolean) {
            assertEquals(expected, isMatch(a, b))
            assertEquals(expected, isMatch(b, a))
        }
        test("a", ".*..a*", false)
        test("aa", "a*", true)
        test("aab", "c*a*b", true)
        test("aa", "a", false)
        test("aa", ".*", true)
        test("aa", ".*", true)
        test("mis*is*p*.", "mississippi", false)
        test("mis*is*p*.", "mississppi", true)
    }
}
