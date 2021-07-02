package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class ProfitableSchemes{
    fun profitableSchemes(n: Int, minProfit: Int, group: IntArray, profit: IntArray): Int {
        val MOD = (1e9 + 7).toInt()
        val count = List(n+1){IntArray(minProfit+1){if(it == 0) 1 else 0} }

        for((gi, pi) in group.zip(profit)){
            for(g in n downTo gi)
                for(p in minProfit downTo 0){
                    count[g][p] = (count[g][p] + count[g-gi][maxOf(0,p-pi)]) % MOD
                }
        }
        return count[n][minProfit]
    }

    @Test
    fun test1(){
        assertEquals(2, profitableSchemes(5, 3, parseIntArray("[2,2]"), parseIntArray("[2,3]")))
        assertEquals(7, profitableSchemes(10, 5, parseIntArray("[2,3,5]"), parseIntArray("[6,7,8]")))
        assertEquals(7, profitableSchemes(10, 5, parseIntArray("[2,3,5]"), parseIntArray("[6,7,8]")))
    }
}