package atcoder.dp

fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull{if(it != "") it.toInt() else null}

    val (N, W) = input_ints()

    val maxv = MutableList(W+1){0L}
    (1..N).forEach {
        val (w, v) = input_ints();
        (W downTo w).forEach {
            maxv[it] = maxOf(maxv[it], maxv[it - w] + v)
        }
    }
    println(maxv[W])
}
