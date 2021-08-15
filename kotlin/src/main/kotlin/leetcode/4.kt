package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class FindMedianSortedArrays {
    fun findMedianSortedArrays(nums1: IntArray, nums2: IntArray): Double {
        val (a, b) = if (nums1.size <= nums2.size) nums1 to nums2 else nums2 to nums1
        val (n, m) = a.size to b.size
        fun kth(kk: Int): Int {      // 1 based
            var k = kk
            var (i, j) = 0 to 0
            if (n < k) {
                j = k - n - 1
                k -= k - n - 1
            }

            while (k > 1) {
                val (x, y) = k / 2 to (k + 1) / 2
                if (a[i + x - 1] <= b[j + y - 1]) {
                    i += x
                    k -= x
                } else {
                    j += y
                    k -= y
                }
            }
            return minOf(if (i < n) a[i] else Int.MAX_VALUE, if (j < m) b[j] else Int.MAX_VALUE)
        }

        val sz = n + m
        return when {
            sz == 0 -> 0.0
            sz % 2 == 1 -> kth((sz + 1) / 2) * 1.0
            else -> (kth((sz + 1) / 2) + kth((sz + 2) / 2)) / 2.0
        }
    }


    @Test
    fun test1() {
        assertEquals(2.0, findMedianSortedArrays(parseIntArray("[1,3]"), parseIntArray("[2]")))
        assertEquals(2.5, findMedianSortedArrays(parseIntArray("[1,2]"), parseIntArray("[3,4]")))
        assertEquals(0.0, findMedianSortedArrays(parseIntArray("[0,0]"), parseIntArray("[0, 0]")))
        assertEquals(1.0, findMedianSortedArrays(parseIntArray("[1]"), parseIntArray("[]")))
        assertEquals(2.0, findMedianSortedArrays(parseIntArray("[]"), parseIntArray("[2]")))
    }
}
