package atcoder.dp
fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull{if(it != "") it.toInt() else null}
    fun input_longs() = split_input().mapNotNull{if(it != "") it.toLong() else null}

    // idea trick: mapping
    val MOD = 1e9.toInt() + 7
    val (N, ) = input_ints()
    val s = split_input()[0]
    val dp = List(N+1){MutableList(N+1) {0L}}
    dp[1][1] = 1
    for(len in 2..N){
        if(s[len-2] == '<'){
            var pre_sum = 0L
            for(v in 1..len){
                dp[len][v] = pre_sum
                pre_sum = (pre_sum + dp[len-1][v]) % MOD
            }
        }else{
            var suf_sum = 0L
            for(v in len downTo 1){
                suf_sum = (suf_sum + dp[len-1][v]) % MOD
                dp[len][v] = suf_sum
            }
        }
    }
    println(dp[N].sum() % MOD)
}
