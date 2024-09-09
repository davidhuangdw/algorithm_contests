package atcoder.dp
fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull{if(it != "") it.toInt() else null}
    fun input_double() = split_input().mapNotNull{if(it != "") it.toDouble() else null}

    val (N, ) = input_ints()
    val DP = MutableList(N+1) { .0 }
    DP[0]=1.0

    input_double().withIndex().forEach { (j, p) ->
        (j+1 downTo 0).forEach {k ->
            DP[k] = DP[k] * (1-p) + if(k-1 >= 0) DP[k-1] * p else 0.0
        }
    }
    println(((N+1)/2..N).sumOf { DP[it] })
}
