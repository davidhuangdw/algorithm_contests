package google_kickstart.`2017`.D

fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull{if(it != "") it.toInt() else null}
    fun input_longs() = split_input().mapNotNull{if(it != "") it.toLong() else null}

    fun begin(fr: Int, s: Int, f: Int): Int {
        val k = maxOf(0, (fr-s+f-1)/f)
        return s + k*f
    }
    val (T, ) = input_ints()
    (1..T).forEach {
        val (NN, TS, TF) = input_ints()
        val N = NN-1
        val V = List(N) { input_ints() }
        var mx = -1
        fun small(){
            for(m in 0 until (1 shl N)){
                var cnt = 0
                var t = 0
                for(i in 0 until N){
                    if((1 shl i) and m > 0){
                        cnt ++
                        t += TS
                    }
                    val (s, f, d) = V[i]
                    t = begin(t, s, f) + d
                    if(t > TF){
                        cnt = -1
                        break
                    }
                }
                mx = maxOf(mx, cnt)
            }
        }
        fun large(){
            // idea: dp: minTime for i-th bus with k sees: f(i,k) = minOf(calc(i, f(i-1, k)), calc(i, f(i-1, k-1)))
            val DP = List(N+1){ MutableList(N+1){TF+1} }
            DP[0][0] = 0
            for(i in 1..N){
                val (s, f, d) = V[i-1]
                for(k in 0..i){
                    DP[i][k] = d + minOf(
                        begin(DP[i-1][k], s, f),
                        if(k-1 >= 0) begin(DP[i-1][k-1] + TS, s, f) else (TF+1),
                        )
                }
            }
            mx = (0..N).filter{DP[N][it] <= TF}.maxOrNull() ?: -1
        }
//        small()
        large()
        println("Case #${it}: ${if(mx>=0) mx else "IMPOSSIBLE"}")
    }
}
