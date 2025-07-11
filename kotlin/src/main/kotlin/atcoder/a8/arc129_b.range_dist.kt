package atcoder.a8

fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().map { it.toInt() }

    val (N,) = input_ints()
    var min_r = 2e9.toInt()
    var max_l = 0
    repeat(N){
        val (l, r) = input_ints()
        max_l = maxOf(max_l, l)
        min_r = minOf(min_r, r)
        println(maxOf(0, (max_l - min_r+1)/2))
    }
}
