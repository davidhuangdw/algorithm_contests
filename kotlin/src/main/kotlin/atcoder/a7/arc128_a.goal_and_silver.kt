package atcoder.a7

import kotlin.math.ln

fun main() {
    fun inputInts() = readLine()!!.split(' ').map(String::toInt)
    val (N) = inputInts()
    val A = inputInts()
    val dp = List(N+1){ DoubleArray(2) }
    val dec = List(N+1){ IntArray(2) }
    dp[0][0] = .0
    dp[0][1] = -1000.0
    for((i,a) in A.withIndex()) {
        val v= ln(1.0*a)
        val k = i+1
        for(j in 0..1){
            val keep = dp[k-1][j]
            val exch = dp[k-1][1-j] + (if(j == 0) -v else v)
            dp[k][j] = maxOf(keep, exch)
            if(keep < exch){
                dec[k][j] = 1
            }
        }
    }
//    for(i in 0..N){
//        println(dp[i].joinToString(" "))
//    }

    val hist = mutableListOf<Int>()
    var j = 0
    for(i in N downTo  1){
        hist.add(dec[i][j])
        if(dec[i][j] == 1){
            j = j xor 1
        }
    }
    println(hist.reversed().joinToString(" "))
}
