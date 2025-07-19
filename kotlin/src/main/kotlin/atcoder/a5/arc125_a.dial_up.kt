package atcoder.a5

fun main() {
    fun inputInts() = readLine()!!.split(' ').map(String::toInt)

    val (N, M) = inputInts()
    val (S, T) = List(2) { inputInts() }
    val s = S.toHashSet()
    val t = T.toHashSet()
    if (s.size == 1 && (s.size < t.size || S[0] != T[0])) {
        println(-1)
        return
    }
    var cnt = M
    val k = (0 until M).find { T[it] != S[0] } ?: M
    if (k < M) {
        cnt += minOf((0 until N).find { S[it] == T[k] }!!,
            (1..N).find { S[N - it] == T[k] }!!
        )
    }
    cnt += (k + 1 until M).count { T[it] != T[it - 1] }
    println(cnt)
}
