package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class MergeTriplets {
    fun mergeTriplets(triplets: Array<IntArray>, target: IntArray): Boolean {
        var cur = mutableListOf(0, 0, 0)
        for(arr in triplets){
            var ok = true
            for(i in 0 until 3)
                if(target[i] < arr[i])
                    ok = false
            if(!ok) continue
            for(i in 0 until 3)
                cur[i] = maxOf(cur[i], arr[i])
        }
        return cur == target.toList()
    }


    @Test
    fun test1() {
        assertEquals(true, mergeTriplets(parse2dIntArray("[[2,5,3],[2,3,4],[1,2,5],[5,2,3]]"), parseIntArray("[5,5,5]")))
        assertEquals(false, mergeTriplets(parse2dIntArray("[[3,4,5],[4,5,6]]"), parseIntArray("[3,2,5]")))
    }
}
