package atcoder.dp
fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull{if(it != "") it.toInt() else null}
    fun input_longs() = split_input().mapNotNull{if(it != "") it.toLong() else null}

    val (N, ) = input_ints()
    val A = (1..N).map { val (w, s, v) = input_ints(); listOf(w, w+s, v) }.sortedBy { it[1] }
//    println(A)
    val S = A.last()[1] + 1
    val DP = List(N){ MutableList(S){-1L} }
    fun dp(i: Int, lim: Int): Long {
        if(i < 0 || lim <= 0) return 0
        if(DP[i][lim] < 0){
            val (w, s, v) = A[i]
            var mx = dp(i-1, lim)
            if(w <= lim){
                mx = maxOf(mx, dp(i-1, minOf(lim, s)-w) + v)
            }
            DP[i][lim] = mx
//            println("----- $i $lim: $mx")
        }
        return DP[i][lim]
    }
    println((0 until N).maxOf{ dp(it, A[it][1]) })
}
