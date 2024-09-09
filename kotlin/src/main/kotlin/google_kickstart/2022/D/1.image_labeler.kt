package google_kickstart.`2022`.D

fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull{if(it != "") it.toInt() else null}
    fun input_longs() = split_input().mapNotNull{if(it != "") it.toLong() else null}

    fun med(sorted: List<Long>): Double {
        val n = sorted.size
        if(n % 2 == 1){return 1.0 * sorted[n/2]}
        return (1.0 * sorted[n/2-1] + sorted[n/2])/2
    }

    val (T, ) = input_ints()
    (1..T).forEach {
        val (N, M) = input_ints()
        val arr = input_longs().sorted()
        val sum = arr.subList(N-M+1, N).sum() + med(arr.subList(0, N-M+1));
        println("Case #${it}: ${String.format("%.2f", sum)}")
    }
}
