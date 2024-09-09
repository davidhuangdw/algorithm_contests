package atcoder.dp

fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull{if(it != "") it.toInt() else null}
    fun input_longs() = split_input().mapNotNull{if(it != "") it.toLong() else null}

    val (N, ) = input_ints()
    val cost = (1..N).map { input_ints() }

    val cache = mutableMapOf<Pair<Int, Int>, Int>()
    fun dp(pre: Int, i: Int): Int{
        if(i >= N) return 0;
        if(pre to i !in cache) {
            cache[pre to i] = (0..2).filter { it != pre }.maxOf { cost[i][it] + dp(it, i+1) }
        }
        return cache[pre to i]!!
    }
    println((0..2).maxOf{ cost[0][it] + dp(it, 1)})
}
