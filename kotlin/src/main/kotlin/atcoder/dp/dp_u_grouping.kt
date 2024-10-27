package atcoder.dp

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val (N,) = input_ints()
    val U = (1 shl N)
    val A = (0 until N).map { input_ints() }
    val S = MutableList(U) { 0L }
    for (v in 1 until U) {
        val e = (0 until N).filter { (1 shl it) and v > 0 }
        var sum = 0L
        for (i in 0 until e.size)
            for (j in i + 1 until e.size) {
                sum += A[e[i]][e[j]]
            }
        S[v] = sum
    }

    val dp = MutableList(U) { 0L }
    for (v in 1 until U) {
        var mx = 0L
//        val e = (0 until N).filter { (1 shl it) and v > 0 }
//        for (x in 1 until (1 shl e.size)) {
//            val xv = (0 until e.size).filter { (1 shl it) and x > 0 }.sumOf { 1 shl e[it] }
//            mx = maxOf(mx, S[xv] + dp[v - xv])
//        }
        var sub = v
        while (sub > 0){
            mx = maxOf(mx, S[sub] + dp[v - sub])
            sub = (sub - 1) and v // trick: search directly all subsets
        }
        dp[v] = mx
    }
    println(dp[U - 1])
}
