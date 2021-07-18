package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class FindRotateSteps514 {
    fun findRotateSteps(ring: String, key: String): Int {
        val n = ring.length
        val M = Int.MAX_VALUE
        var f = MutableList(n){M}
        f[0] = 0
        fun abs(x: Int) = if(x >= 0) x else -x
        for(ch in key){
            val nxt = MutableList(n){M}
            for(j in 0 until n)
                if(f[j] < M){
                    for(k in 0 until n)
                        if(ring[k] == ch)
                            nxt[k] = minOf(nxt[k], f[j] + minOf(abs(j-k), n-abs(j-k)) + 1)
                }
            f = nxt
        }
        return f.minOrNull()!!
    }


    @Test
    fun test1() {
        assertEquals(4, findRotateSteps("godding", "gd"))
    }
}
