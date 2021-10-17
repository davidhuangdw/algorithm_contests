fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }


    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (N, K) = input_ints()
        val B = input_ints()
        val pre = B.runningFold(0L, Long::plus)
        val rmin = hashMapOf<Int, Int>()
        rmin[0] = 0
        var min_cost = N * 2
        for (i in N downTo 1) {
            // query
            for (j in i - 1 downTo 0) {
                val v = pre[i] - pre[j]
                val c = i - j
                if (K - v < 0) continue
                val x = (K - v).toInt()
                if (x in rmin)
                    min_cost = minOf(min_cost, rmin[x]!! + c)
            }
            // update
            for (j in i..N) {
                val v = pre[j] - pre[i - 1]
                val c = j - i + 1
                if (v > K) continue
                val x = v.toInt()
                rmin[x] = minOf(rmin.getOrDefault(x, Int.MAX_VALUE), c)
            }
        }

        if (min_cost > N)
            min_cost = -1
        println("Case #${it}: $min_cost")
    }
}
