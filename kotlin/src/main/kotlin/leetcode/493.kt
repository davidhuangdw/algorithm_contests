package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class ReversePairs493 {
    fun reversePairs(nums: IntArray): Int {
        val ordered =nums.toSet().sorted()
        val v_to_i = hashMapOf<Int, Int>()
        for((i, v) in ordered.withIndex())
            v_to_i[v] = i+1
        val max = ordered[ordered.size-1]

        val n = ordered.size
        val tree = MutableList(n+1){0}
        fun add(ii: Int){
            var i = ii
            while(i <= n){
                tree[i] += 1
                i += i and -i
            }
        }
        fun query(ii: Int): Int {
            var i = ii
            var sum = 0
            while(i > 0){
                sum += tree[i]
                i -= i and -i
            }
            return sum
        }
        fun lowerPos(v: Long): Int {
            var (l, r) = 0 to n-1
            while(l <= r){
                val md = (l+r) ushr 1
                if(ordered[md] <= v) l = md+1
                else r = md - 1
            }
            return l
        }

        var cnt = 0
        for((i, v) in nums.withIndex()){
            cnt += i - query(lowerPos(v * 2L))
            add(v_to_i[v]!!)
        }
        return cnt
    }

    fun reversePairs1(nums: IntArray): Int {    // by merge
        var cnt = 0
        fun merge(l: Int, r: Int){      // [l, r)
            if(r-l <= 1) return
            val md = (l+r)/2
            merge(l, md)
            merge(md, r)
            // count
            var j = l
            for(i in md until r){
                val rv = nums[i]*2L     // bug: *2 will overflow
                while(j < md && nums[j] <= rv)
                    j += 1
                cnt += md - j
            }

            // merge
            val left = (l until md).map{nums[it]}
            j = md
            var i = l
            for(v in left){
                while(j < r && nums[j] < v){
                    nums[i] = nums[j]
                    i += 1
                    j += 1
                }
                nums[i] = v
                i += 1
            }
        }
        merge(0, nums.size)
        return cnt
    }

    @Test
    fun test1() {
        assertEquals(2, reversePairs(parseIntArray("[1,3,2,3,1]")))
        assertEquals(3, reversePairs(parseIntArray("[2,4,3,5,1]")))
    }
}
