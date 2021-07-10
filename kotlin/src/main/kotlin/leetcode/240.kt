package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class SearchMatrix240 {
    fun searchMatrix(matrix: Array<IntArray>, target: Int): Boolean {
        var r = matrix[0].size-1
        for(row in matrix){
            while(r >= 0 && row[r] > target)
                r -= 1
            if(r >= 0 && row[r] == target) return true
        }
        return false
    }

    @Test
    fun test1() {
        assertEquals(true, searchMatrix(parse2dIntArray("[[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]]"), 5))
        assertEquals(false, searchMatrix(parse2dIntArray("[[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]]"), 20))
    }
}
