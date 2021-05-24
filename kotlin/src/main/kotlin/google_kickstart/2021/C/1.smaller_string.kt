package google_kickstart.`2021`.C

fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().map{it.toInt()}
    fun input_longs() = split_input().map{it.toLong()}
    val MOD = 1000_000_007L
    fun pow(a: Int, kk: Int): Long{
        var (x, mul, k) = Triple(1L, a.toLong(), kk)
        while(k > 0){
            if(k and 1 != 0) x = x*mul % MOD
            k = k shr 1
            mul = mul * mul % MOD
        }
        return x
    }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (N, K) = input_ints()
        val S = readLine()!!
        var res = 0L
        val md_len = (N+1)/2
        val half = N/2
        for(eq_len in 0..minOf(half, md_len-1)){
            res += (S[eq_len] - 'a')*pow(K, md_len-eq_len-1) % MOD
        }
        for(i in 0 until half)
            if(S[half-1-i] != S[md_len + i]){
                if(S[half-1-i] < S[md_len + i]) res +=1
                break
            }
        println("Case #${it}: ${res % MOD}")
    }
}
