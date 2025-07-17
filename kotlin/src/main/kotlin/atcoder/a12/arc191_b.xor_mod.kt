package atcoder.a12

fun main() {
    fun inputInts() = readLine()!!.split(' ').map(String::toInt)

    //    for(n in 1..10){
//        println("-".repeat(10) + Integer.toBinaryString(n))
//        for(x in 0 until n*2){
//            if(x %n == x xor n){
//                println(Integer.toBinaryString(x))
//            }
//        }
//    }
    fun sol(n: Int, k: Int): Int {
        val s = Integer.toBinaryString(n)
        val z = s.count { it == '0' }
        if (k > 1 shl z) {
            return -1
        }
        var rank = k-1
        var res = n
        var bit = 1
        while(bit <= n){
            if(n and bit == 0){
                if(rank and 1 > 0){
                    res += bit
                }
                rank /= 2
            }
            bit = bit shl 1
        }
        return res
    }
    val (T) = inputInts()
    repeat(T) {
        val (N, K) = inputInts()
        println(sol(N, K))

    }
}
