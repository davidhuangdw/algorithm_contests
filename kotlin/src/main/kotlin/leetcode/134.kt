package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class CanCompleteCircuit134 {
    fun canCompleteCircuit(gas: IntArray, cost: IntArray): Int {
        val n = gas.size
        var sum = 0
        var min = Int.MAX_VALUE to -1
//        println("-".repeat(10))
        for (i in 0 until n) {
            val cur = gas[i] - cost[i]
            sum += cur
//            println("$i $cur $sum")
            if (sum < min.first)
                min = sum to (i + 1) % n
        }
        return if (sum >= 0) min.second else -1
    }

    @Test
    fun test1() {
        assertEquals(3, canCompleteCircuit(parseIntArray("[1,2,3,4,5]"), parseIntArray("[3,4,5,1,2]")))
        assertEquals(4, canCompleteCircuit(parseIntArray("[5,1,2,3,4]"), parseIntArray("[4,4,1,5,1]")))
    }
}
