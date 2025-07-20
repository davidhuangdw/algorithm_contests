package atcoder.a7

fun main() {
    fun inputInts() = readLine()!!.split(' ').map(String::toInt)

    val (N, M) = inputInts()
    var cnt = IntArray(N)

    repeat(N){
        cnt[readLine()!!.count{it == '1'} % 2] += 1
    }
    println(1L * cnt[0] * cnt[1])
}
