package google_kickstart.`2022`.H
fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull{if(it != "") it.toInt() else null}
    fun input_longs() = split_input().mapNotNull{if(it != "") it.toLong() else null}

    val (T, ) = input_ints()
    (1..T).forEach {
        val (L,)  = input_ints()
        val dp = List(L+1){MutableList(2){1e9.toLong()} }
        dp[0][0] = 0L
        for(i in 0 until L){
            dp[i][1] = minOf(dp[i][1], dp[i][0] + 4)
            dp[i + 1][0] = minOf(dp[i+1][0], minOf(dp[i][0], dp[i][1]) + 1)
            if(i < 3) continue
            for(k in 2..L/i){
                dp[k*i][0]  = minOf(dp[k*i][0], dp[i][1] + 2*(k-1))
            }
        }
        println("Case #${it}: ${dp[L].minOrNull()}")
    }
}
