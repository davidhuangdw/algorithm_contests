package atcoder.dp
fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull{if(it != "") it.toInt() else null}
    fun input_longs() = split_input().mapNotNull{if(it != "") it.toLong() else null}

    val MOD = 1e9.toInt()+7;

    val (H, W) = input_ints()
    val mp = (1..H).map{ readLine()!!.map{ch -> if(ch=='.') true else false}}
    val dp = MutableList(W){0}
    dp[0] = 1
    for(i in 0 until H)
        for(j in 0 until W){
            dp[j] = if(mp[i][j]){
                (dp[j] + if(j-1 >= 0)dp[j-1] else 0) % MOD
            }else{0}
        }

    println(dp[W-1])
}
