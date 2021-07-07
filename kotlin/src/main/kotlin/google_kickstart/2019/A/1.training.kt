package google_kickstart.`2019`.A

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (N, P) = input_ints()
        val pre_sum = input_ints().sorted().scan(0, Int::plus)

        var min = Int.MAX_VALUE
        for(i in P..N){
            min = minOf(min , (pre_sum[i] - pre_sum[i-1]) * P - (pre_sum[i] - pre_sum[i-P]))
        }
        println("Case #${it}: $min")
    }
}
