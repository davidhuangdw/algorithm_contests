//package leetcode
//
//import org.junit.jupiter.api.Assertions.assertEquals
//import org.junit.jupiter.api.Test
//
//class FindKthNumber440 {
//    fun findKthNumber(n: Int, k: Int): Int {
//        var x = n/10
//        var pow = 1
//        while(x > 0){
//            x /= 10
//            pow *= 10
//        }
//
//        var res  = 0
//        var left = k
//        var allow_zero = 0
//        while(left > 0){
//            val k = (left+pow-1)/pow
//            res = res*10 + (k - allow_zero)
//            left -= (k-1) * pow + 1
//            pow /= 10
//            allow_zero = 1
//        }
//        return res
//    }
//
//
//    @Test
//    fun test1() {
//        assertEquals(10, findKthNumber(13, 2))
//        assertEquals(35, findKthNumber(100, 30))
//    }
//}
