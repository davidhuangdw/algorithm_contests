package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class NumRabbits{
    fun numRabbits(answers: IntArray): Int {
        val count = hashMapOf<Int, Int>()
        for(v in answers)
            count[v] = count.getOrDefault(v, 0) + 1

        var res = 0
        for((v, cnt) in count){
            val grp = v + 1
            res += (cnt+grp-1)/grp * grp
        }
        return res
    }

    @Test
    fun test1(){
        assertEquals(10, numRabbits(parseIntArray("[2,2,2,2,3]")))
    }
}