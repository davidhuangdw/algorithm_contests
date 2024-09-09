package atcoder.dp

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }

    val (N, W) = input_ints()

    val V = N*1000;
    val minw = MutableList(V+1){W+1}
    minw[0] = 0
    (1..N).forEach {
        val (w, v) = input_ints();
        (V downTo 1).forEach {
            minw[it] = minOf(minw[it], if(it-v >= 0)minw[it - v] + w else W+1, if(it+1 <=V) minw[it+1] else W+1, W+1)
        }
    }
    println((V downTo 0).find{minw[it] <= W})
}
