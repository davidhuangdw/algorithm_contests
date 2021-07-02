package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class NumOfWays{
    fun numOfWays(n: Int): Int {
        val MOD = (1e9+7).toInt()
        val line = mutableListOf<Int>()
        fun toColors(x: Int) = Triple(x % 3, (x/3) % 3, x/9)
        for(x in 0 until 27){
            val (a, b, c) = toColors(x)
            if(a != b && b !=c) line.add(x)
        }

        val w = List(2){IntArray(27)}
        for(x in line)
            w[1][x] = 1
        for(i in 2..n){
            val cur = w[i and 1]
            val pre = w[i and 1 xor 1]
            for(x in line){
                val (xa, xb, xc) = toColors(x)
                cur[x] = 0
                for(y in line){
                    val (ya, yb, yc) = toColors(y)
                    if(xa != ya && xb != yb && xc != yc)
                        cur[x] = (cur[x] + pre[y]) % MOD
                }
            }
        }
        return w[n and 1].reduce{acc, b -> (acc + b) % MOD}
    }

    @Test
    fun test1(){
        assertEquals(30228214, numOfWays(5000))
    }
}