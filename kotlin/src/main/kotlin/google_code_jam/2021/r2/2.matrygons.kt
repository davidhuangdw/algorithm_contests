package google_code_jam.`2021`.r2

fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().map{it.toInt()}
    fun input_longs() = split_input().map{it.toLong()}

    val MA = 1000000
    val dp = IntArray(MA+1){1}
    dp[0] = 0
    dp[1] = 0
    for(i in 3..MA){
        var x = 2
        while(x*x <= i) {
            if (i % x == 0) {
                val j = i/x
                dp[i] = maxOf(dp[i], 1 + dp[j - 1])
                dp[i] = maxOf(dp[i], 1 + dp[x - 1])
            }
            x += 1
        }
    }
//    println(dp.take(42).mapIndexed{i,v -> "$i $v"})

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val N = readLine()!!.toInt()
        var x = 3
        var res = 1
        while(x*x <= N) {
            if (N % x == 0) {
                val j = N/x
                res = maxOf(res, 1 + dp[j - 1])
                res = maxOf(res, 1 + dp[x - 1])
            }
            x += 1
        }

        println("Case #${it}: ${res}")
    }
}
