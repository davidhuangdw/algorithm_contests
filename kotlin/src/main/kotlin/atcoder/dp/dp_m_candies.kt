package atcoder.dp
fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull{if(it != "") it.toInt() else null}
    fun input_longs() = split_input().mapNotNull{if(it != "") it.toLong() else null}

    val MOD = 1e9.toInt() + 7

    val (N, K) = input_ints()
    val DP = LongArray(K+1){0}
    DP[0] = 1

    input_ints().sorted().forEach { a ->
        val sum = DP.scan(0L){acc, x -> (acc + x) % MOD }.drop(1)
        for(k in 0..K){
            DP[k] = (MOD + sum[k] - (if(k-a-1 >= 0) sum[k-a-1] else 0)) % MOD
        }
    }
    println(DP[K])
}
