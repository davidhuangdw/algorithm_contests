package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class AddDigits{
    fun addDigits(num: Int): Int {
        if(num == 0) return 0
        val r = num%9
        return if(r==0) 9 else r
    }

    @Test
    fun test1(){
//        for(i in 0..110)
//            println("$i -> ${addDigits(i)}")
        assertEquals(addDigits(38), 2)
    }
}