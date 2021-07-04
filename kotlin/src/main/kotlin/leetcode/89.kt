package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class GrayCode {
    fun grayCode(n: Int): List<Int> {
        var cur = mutableListOf(0)
        for(i in 0 until n){
            val (head, m) = 1 shl i to cur.size
            for(j in m-1 downTo 0)
                cur.add(head or cur[j])
        }
        return cur
    }

    fun grayCode1(n: Int): List<Int> {
        var x = 0
        val done = hashSetOf(0)
        val res = mutableListOf(0)
        do{
            var found = false
            for(k in 0 until n)
                if(x xor (1 shl k) !in done){
                    x = x xor (1 shl k)
                    found = true
                    done.add(x)
                    res.add(x)
                    break
                }
        }while(found)
        return res
    }

    @Test
    fun test1() {
        assertEquals(listOf(0, 1, 3, 2), grayCode(2))
        assertEquals(listOf(0), grayCode(0))
        for(n in 3..10)
            assertEquals(1 shl n, grayCode(n).size)
    }
}
