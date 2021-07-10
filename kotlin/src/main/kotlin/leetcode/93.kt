package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class RestoreIpAddresses {
    fun restoreIpAddresses(s: String): List<String> {
        val n = s.length
        if(n !in 4..12) return listOf()

        fun valid(pos: Pair<Int, Int>): Boolean {
            val (st, ed) = pos
            var sum = 0
            if(ed - st > 1 && s[st] == '0') return false
            for(i in st until ed)
                sum = sum*10 + (s[i]-'0')
            return sum in 0..255
        }

        val res = mutableListOf<String>()
        for(a in 1..n-3)
            for(b in a+1..n-2)
                for(c in b+1 until n){
                    val pos = listOf(0, a, b, c, n)
                    val sides = pos.zipWithNext()
                    if(sides.all{valid(it)}){
                        res.add(sides.map{s.substring(it.first until it.second)}.joinToString("."))
                    }
                }
        return res
    }


    @Test
    fun test1() {
        assertEquals(setOf("255.255.11.135","255.255.111.35"), restoreIpAddresses("25525511135").toSet())
        assertEquals(setOf("0.0.0.0"), restoreIpAddresses("0000").toSet())
        assertEquals(setOf("1.1.1.1"), restoreIpAddresses("1111").toSet())
        assertEquals(setOf("0.10.0.10","0.100.1.0"), restoreIpAddresses("010010").toSet())
        assertEquals(setOf("1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"), restoreIpAddresses("101023").toSet())
    }
}
