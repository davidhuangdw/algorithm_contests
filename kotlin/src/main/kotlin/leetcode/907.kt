package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test
import java.util.*

class SumSubarrayMins {
    fun sumSubarrayMins(arr: IntArray): Int {
        val MOD = (1e9 + 7).toLong()
        val n = arr.size
        val right_lt = IntArray(n)
        val st = mutableListOf<Int>()
        for (i in n - 1 downTo 0) {
            while (st.isNotEmpty() && arr[st.last()] >= arr[i])
                st.removeAt(st.size - 1)
            right_lt[i] = if (st.isNotEmpty()) st.last() else n
            st.add(i)
        }

        var sum = 0L
        st.clear()
        for (i in 0 until n) {
            while (st.isNotEmpty() && arr[st.last()] > arr[i])
                st.removeAt(st.size - 1)
            val left_le = if (st.isNotEmpty()) st.last() else -1
            sum = ((i - left_le) * (right_lt[i] - i) % MOD * arr[i] % MOD + sum) % MOD
            st.add(i)
        }

        return sum.toInt()
    }


    @Test
    fun test1() {
        assertEquals(7, sumSubarrayMins(parseIntArray("[1,2,1]")))
        assertEquals(17, sumSubarrayMins(parseIntArray("[3,1,2,4]")))
        assertEquals(444, sumSubarrayMins(parseIntArray("[11,81,94,43,3]")))
    }
}
