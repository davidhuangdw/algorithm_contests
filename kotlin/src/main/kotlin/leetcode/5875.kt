package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class FinalValueAfterOperations5875 {
    fun finalValueAfterOperations(operations: Array<String>): Int {
        var x = 0
        for (o in operations) {
            x += if ('+' in o) 1 else -1
        }
        return x
    }


    @Test
    fun test1() {
        assertEquals(1, 1)
    }
}
