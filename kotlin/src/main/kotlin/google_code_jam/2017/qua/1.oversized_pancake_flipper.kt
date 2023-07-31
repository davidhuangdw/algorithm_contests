package google_code_jam.`2017`.qua

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val (T) = input_ints()
    (1..T).forEach {
        val (S, _K) = split_input()
        val K = _K.toInt()
        fun set(): String {
            val state = S.map { if (it == '+') 0 else 1 }.toMutableList()
//            var cnt = 0
            val n = state.size
            val sum = MutableList(n+1){0}
            for (i in 0 until n) {
                val s = sum[i] - if(i-K >= 0)sum[i-K+1] else 0
                sum[i+1] = sum[i]
                if ((s + state[i]) % 2 != 0) {
                    if (i > n - K)
                        return "IMPOSSIBLE"
                    sum[i+1] ++
//                    cnt++
//                    for (j in i until i + K)
//                        state[j] = 1 - state[j]
                }
            }
            return sum[n].toString()
        }

        val r = set()
        println("Case #${it}: $r")
    }
}
