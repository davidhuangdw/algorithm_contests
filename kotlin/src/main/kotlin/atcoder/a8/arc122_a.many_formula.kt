package atcoder.a8

fun main() {
    fun inputInts() = readLine()!!.split(' ').map(String::toInt)

    val (N) = inputInts()
    val A = inputInts()
    val MOD = (1e9+7).toInt()

    val cnt = LongArray(N+1)
    cnt[0] = 1
    cnt[1] = 1
    for(i in 2..N){
        cnt[i] = (cnt[i-1] + cnt[i-2]) % MOD
    }

    val dp = LongArray(N+1)
    for((i, v) in A.withIndex()){
        dp[i+1] = (v*cnt[i] % MOD + dp[i] + (if(i-1>=0) (A[i-1]-v)*cnt[i-1] % MOD+dp[i-1] else 0)) % MOD
    }
    println(dp[N])
}
