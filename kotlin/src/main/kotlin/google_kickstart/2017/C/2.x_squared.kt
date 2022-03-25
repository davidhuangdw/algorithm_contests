package google_kickstart.`2017`.C

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (N) = input_ints()
        val sq = (1..N).map { split_input()[0] }
        val cnt = hashMapOf<Pair<Int, Int>, Int>()
        val col = MutableList(N) { 0 }
        fun valid(): Boolean {
            var one = 0
            for (line in sq) {
                var xs = mutableListOf<Int>()
                for ((i, ch) in line.withIndex()) {
                    if (ch == 'X') {
                        xs.add(i)
                        col[i]++
                    }
                }
                when (xs.size) {
                    1 -> {
                        one++
                        if (one > 1) return false
                    }
                    2 -> {
                        col
                        val p = xs[0] to xs[1]
                        cnt[p] = cnt.getOrDefault(p, 0) + 1
                        if (cnt[p]!! > 2) return false
                    }
                    else -> return false
                }
            }
            for ((_, v) in cnt) {
                if (v != 2) return false
            }
            for (v in col) {      // bug: forget to check column_count
                if (v !in 1..2)
                    return false
            }
            return one == 1
        }

        val res = if (valid()) "POSSIBLE" else "IMPOSSIBLE"
        println("Case #${it}: $res")
    }
}
