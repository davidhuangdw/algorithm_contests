package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class MinTimeToType5834 {
    fun minTimeToType(word: String): Int {
        val n = word.length
        val dp = List(n + 1) { MutableList(26) { 0 } }
        for (i in n - 1 downTo 0) {
            val tar = word[i] - 'a'
            for (j in 0 until 26) {
                val dis = minOf((tar - j + 26) % 26, (j - tar + 26) % 26)
                dp[i][j] = dis + 1 + dp[i + 1][tar]
            }
        }
        return dp[0][0]
    }


    @Test
    fun test1() {
        assertEquals(5, minTimeToType("abc"))
        assertEquals(7, minTimeToType("bza"))
        assertEquals(34, minTimeToType("zjpc"))
    }
}
