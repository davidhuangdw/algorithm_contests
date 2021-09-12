package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class MaxProduct {
    fun maxProduct(s: String): Int {
        val n = s.length
        val M = 1 shl n
        val par = hashSetOf<Pair<Int, Int>>()
        for(b in 1 until M){
            val sub = mutableListOf<Char>()
            for(i in 0 until n){
                if((1 shl i) and b > 0)
                    sub.add(s[i])
            }
            var (l, r) = 0 to sub.size-1
            var ok = true
            while(l < r){
                if(sub[l] != sub[r]){
                    ok = false
                    break
                }
                l += 1
                r -= 1
            }
            if(ok) par.add(b to sub.size)
        }
        var ret = 0
        for((a, al) in par){
            for((b, bl) in par){
                if(a and b == 0){
                    ret = maxOf(ret, al * bl)
                }
            }
        }
        return ret
    }


    @Test
    fun test1() {
        assertEquals(1, 1)
    }
}
