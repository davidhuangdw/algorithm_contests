package google_code_jam.`2018`.`1c`
fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull{if(it != "") it.toInt() else null}
    fun input_longs() = split_input().mapNotNull{if(it != "") it.toLong() else null}

    val (T, ) = input_ints()
    (1..T).forEach {
        val (N, ) = input_ints()
        val W = input_longs()
        fun set1(): Int{
            val S = W.sum().toInt()
            var dp = List(N+1){MutableList(S+1){0}}
            for((_i, w) in W.map{it.toInt()}.withIndex()){
                val i = _i + 1
                for(v in 0..S){
                    var max = dp[i-1][v]
                    if(v-1 >= 0)
                        max = maxOf(max, dp[i][v-1])
                    if(v in w..7*w)
                        max = maxOf(max, dp[i-1][v-w] + 1)
                    dp[i][v] = max
                }
            }
            return dp[N][S]
        }

        fun set2(): Int{
            val dp = MutableList(N+1){i -> if(i ==0) 0L else Long.MAX_VALUE}
            var bound = 0
            // trick: the w increase very fast, so bound << N
            for(w in W){
                val sum = w * 6
                for(i in bound+1 downTo 1){
                    if(dp[i-1] <= sum)
                        dp[i] = minOf(dp[i], dp[i-1]+w)
                }
                if(dp[bound+1] < Long.MAX_VALUE)
                    bound ++
            }
            return bound
        }
        val r = set2()
        println("Case #${it}: $r")
    }
}
