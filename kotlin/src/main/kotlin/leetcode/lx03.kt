package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class GetKthNumLx03 {
    fun getKthNum(k: Int): Int {
        val pow = (1..10).scan(1L) { acc, _ -> acc * 10 }
        fun countLen(x: Int): Int {
            var n = 0
            while (pow[n + 1] <= x)
                n += 1
            val dp = List(n + 1) { List(n + 1) { MutableList(2) { 1L } } }
            for (i in 0..n) {
                val suf = x % pow[i + 1]
                val dig = (suf / pow[i]).toInt()
                for (l in 0..(n - i))
                    for (lt in 0..1) {
                        val eq = lt == 0
                        var cur =
                            0L // l * if (eq) suf+1 else pow[i+1]   // bug: 0, not (suf+1 or pow[i+1]), only compute when i<0
                        for (d in 0..if (eq) dig else 9)
                            cur += if (i - 1 >= 0) dp[i - 1][if (l == 0 && d == 0) 0 else l + 1][if (eq && d == dig) 0 else 1] else l + 1L         // bug: l+1, not l
                        dp[i][l][lt] = cur
                    }
            }
            return dp[n][0][0].toInt()
        }

        countLen(100)
//        for(i in 1..101){
//            println("$i ${countLen(i)}")
//        }
        val zeroK = k + 1
        var (l, r) = 1 to k
        while (l <= r) {
            val md = (l + r) / 2
            val len = countLen(md)
//            println("$md $len")
            if (len < zeroK) l = md + 1
            else r = md - 1
        }

        var x = l
//        println(x)
//        println(countLen(x))
//        println("-".repeat(10))
        return x / pow[countLen(x) - zeroK].toInt() % 10
    }


    @Test
    fun test1() {
        assertEquals(6, getKthNum(6))
        assertEquals(1, getKthNum(20))
        assertEquals(4, getKthNum(99))
        assertEquals(1, getKthNum(1000_000))
        assertEquals(7, getKthNum(9898989))
        assertEquals(6, getKthNum(123_123_123))
    }
}
