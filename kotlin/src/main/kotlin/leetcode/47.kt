package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class PermuteUnique47 {
    fun permuteUnique(nums: IntArray): List<List<Int>> {
        val n = nums.size
        val all = mutableListOf<List<Int>>()
        fun dfs(i: Int){
            if(i == n){
                all.add(nums.toList())
            }else{
                val done = hashSetOf<Int>()
                for(j in i until n){
                    val v = nums[j]
                    if(v !in done){
                        done.add(v)
                        nums[i] = nums[j].also{nums[j] = nums[i]}
                        dfs(i+1)
                        nums[i] = nums[j].also{nums[j] = nums[i]}
                    }
                }
            }
        }
        dfs(0)
        return all
    }


    @Test
    fun test1() {
        assertEquals(parse2dIntArray("[[1,1,2,2],[1,2,1,2],[1,2,2,1],[2,1,1,2],[2,1,2,1],[2,2,1,1]]").map{it.toList()}, permuteUnique(parseIntArray("[2,2,1,1]")))
        assertEquals(parse2dIntArray("[[1,1,2],[1,2,1],[2,1,1]]").map{it.toList()}, permuteUnique(parseIntArray("[1,1,2]")))
        assertEquals(parse2dIntArray("[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,2,1],[3,1,2]]").map{it.toList()}, permuteUnique(parseIntArray("[1,2,3]")))
    }
}
